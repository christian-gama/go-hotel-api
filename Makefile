PWD = $(shell pwd)
GENERATED_DIR ?= $(PWD)/.generated
CACHE_DIR ?= $(PWD)/.cache
COVERAGE_DIR ?= $(GENERATED_DIR)/coverage
BUILD_DIR ?= $(GENERATED_DIR)/build
MAKE = make --no-print-directory

# WORKDIR is used to set the working directory for Dockerfile builds.
export WORKDIR=/go/src/github.com/christian-gama/go-booking-api
export APP_NAME = $(shell basename $(PWD))

-include $(ENV_FILE)

init: cmd-exists-git
	@git config core.hooksPath .githooks
	@chmod +x .githooks/*
	@echo "Initialized git hooks" 
	@$(MAKE) remake


run: cmd-exists-go
	@$(MAKE) build
	@$(BUILD_DIR)/$(APP_NAME) -env $(ENV_FILE)


build: cmd-exists-go
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/api/*.go
	@chmod +x $(BUILD_DIR)/$(APP_NAME)
	@echo "Build was generated at $(BUILD_DIR)/$(APP_NAME)"


test: cmd-exists-go cmd-exists-gotestsum
	@TEST_MODE=unit APP_NAME=$(APP_NAME) gotestsum --format pkgname -- $(TEST_FLAGS) ./...


test-integration: cmd-exists-go
	@TEST_MODE=integration APP_NAME=$(APP_NAME) gotestsum --format pkgname -- $(TEST_FLAGS) -p 1 ./...


test-verbose: cmd-exists-go cmd-exists-gotestsum
	@TEST_MODE=unit APP_NAME=$(APP_NAME) gotestsum --format standard-verbose -- $(TEST_FLAGS) ./...


cover: cmd-exists-go
	@if [ ! -d "$(COVERAGE_DIR)" ]; then \
		mkdir -p $(COVERAGE_DIR); \
	fi
	@TEST_MODE=both APP_NAME=$(APP_NAME) gotestsum --format pkgname -- -coverprofile=$(COVERAGE_DIR)/coverage.out -p 1 ./...
	@$(MAKE) cover-html


cover-html: cmd-exists-go
	@go tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html


cover-open:
	@google-chrome $(COVERAGE_DIR)/coverage.html


lint: cmd-exists-docker
	@sh -c "./scripts/linter.sh"


migrate-%: cmd-exists-docker
	@if [ -z "$(*)" ]; then \
		echo "Error: expected [up|down|force]"; \
		exit 1; \
	fi;
	@ENV_FILE=.env.dev MIGRATION=$(*) make migrate


migrate: cmd-exists-docker
	@docker compose --env-file $(ENV_FILE) up -d psql
	@until docker exec go_booking_psql pg_isready ; do sleep 4 ; done

	@mkdir -p $(PWD)/$(MIG_DIR)
	@docker run -it \
		-v $(PWD)/$(MIG_DIR):/migrations \
		--rm \
		--network host \
		migrate/migrate \
		-path=/migrations/ \
		-database $(DB_SGBD)://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_EXPOSED_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE) $(MIGRATION) $(VERSION)

	
mock: cmd-exists-docker
	@docker run \
		-v "$(PWD)":/src \
		-w /src vektra/mockery --all --case underscore --exported --dir ./internal --quiet


clear:
	@docker compose --env-file .env.dev down
	@rm -rf $(GENERATED_DIR)
	@rm -rf $(CACHE_DIR)
	@rm -rf $(PWD)/mocks


remake:
	@$(MAKE) clear
	@WORKDIR=$(WORKDIR) docker compose --env-file .env.dev up api -d
	@$(MAKE) migrate-up

	@$(MAKE) docker ENV_FILE=.env.dev ENTRY_POINT="make build"
	@$(MAKE) mock


docker: cmd-exists-docker
	@WORKDIR=$(WORKDIR) APP_NAME=$(APP_NAME) docker compose --env-file "$(ENV_FILE)" run \
		--name $(APP_NAME) \
		-p $(APP_PORT):$(APP_PORT) \
		--rm \
		-e ENV_FILE=$(ENV_FILE) \
		api \
		$(ENTRY_POINT)


docker-dev: cmd-exists-docker
	@$(MAKE) docker ENV_FILE=.env.dev ENTRY_POINT="make run" 


docker-prod: cmd-exists-docker
	@$(MAKE) docker ENV_FILE=.env.prod ENTRY_POINT="make run" 


docker-test: cmd-exists-docker
	@$(MAKE) docker ENV_FILE=.env.test ENTRY_POINT="make test"


docker-test-integration: cmd-exists-docker
	@$(MAKE) docker ENV_FILE=.env.test ENTRY_POINT="make test-integration" 


docker-cover: cmd-exists-docker
	@$(MAKE) docker ENV_FILE=.env.test ENTRY_POINT="make cover" 


docker-cover-html: cmd-exists-docker
	@$(MAKE) docker ENV_FILE=.env.test ENTRY_POINT="make cover-html"


docker-down: cmd-exists-docker
	@docker compose --env-file ".env.dev" down 


docker-rebuild:
	@docker compose --env-file ".env.dev" build


cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)
