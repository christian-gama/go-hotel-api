PWD = $(shell pwd)
GENERATED_DIR ?= $(PWD)/.generated
CACHE_DIR ?= $(PWD)/.cache
COVERAGE_DIR ?= $(GENERATED_DIR)/coverage
BUILD_DIR ?= $(GENERATED_DIR)/build
APP_NAME ?= go_booking

-include $(ENV_FILE)

init: cmd-exists-git cmd-exists-node
	@git config core.hooksPath .githooks
	@chmod +x .githooks/*
	@echo "Initialized git hooks"


run: cmd-exists-go
	@make build
	@$(BUILD_DIR)/$(APP_NAME) -env $(ENV_FILE)


build: cmd-exists-go
	@CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/api/*.go
	@chmod +x $(BUILD_DIR)/$(APP_NAME)
	@echo "Build was generated at $(BUILD_DIR)/$(APP_NAME)"


test_integration: cmd-exists-go
	@TEST_MODE=unit gotestsum --format pkgname -- $(TEST_FLAGS) -p 1 ./...


test: cmd-exists-go cmd-exists-gotestsum
	@TEST_MODE=unit gotestsum --format pkgname -- $(TEST_FLAGS) ./...


cover: cmd-exists-go
	@if [ ! -d "$(COVERAGE_DIR)" ]; then \
		mkdir -p $(COVERAGE_DIR); \
	fi
	@make test TEST_FLAGS=-coverprofile=$(COVERAGE_DIR)/coverage.out 


cover-html: cmd-exists-go
	@go tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html


lint: cmd-exists-docker
	@sh -c "./scripts/linter.sh"


migrate-%: cmd-exists-docker
	@if [ -z "$(*)" ]; then \
		echo "Error: expected [up|down|force]"; \
		exit 1; \
	fi;
	@ENV_FILE=.env.dev MIGRATION=$(*) make migrate


migrate: cmd-exists-docker
	@mkdir -p $(PWD)/$(MIG_DIR)
	@docker run \
		-v $(PWD)/$(MIG_DIR):/migrations \
		--rm \
		--network host \
		migrate/migrate \
		-path=/migrations/ \
		-database $(DB_SGBD)://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_EXPOSED_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE) $(MIGRATION) $(VERSION)
	
mock: cmd-exists-docker
	docker run -v "$(PWD)":/src -w /src vektra/mockery --all --exported --dir ./internal


clear:
	@rm -rf $(GENERATED_DIR)
	@rm -rf $(CACHE_DIR)


docker: cmd-exists-docker
	@docker compose --env-file "$(ENV_FILE)" run \
		--name $(APP_NAME) \
		-p $(APP_PORT):$(APP_PORT) \
		--rm \
		-e ENV_FILE=$(ENV_FILE) \
		api \
		$(ENTRY_POINT)


docker_dev: cmd-exists-docker
	@make docker ENV_FILE=.env.dev ENTRY_POINT="make run" 


docker_prod: cmd-exists-docker
	@make docker ENV_FILE=.env.prod ENTRY_POINT="make run" 


docker_test: cmd-exists-docker
	@make docker ENV_FILE=.env.test ENTRY_POINT="make test"


docker_test_integration: cmd-exists-docker
	@make docker ENV_FILE=.env.test ENTRY_POINT="make test_integration" 


docker_cover: cmd-exists-docker
	@make docker ENV_FILE=.env.test ENTRY_POINT="make cover" 


docker_cover_html: cmd-exists-docker
	@make docker ENV_FILE=.env.test ENTRY_POINT="make cover-html"


docker_down: cmd-exists-docker
	@docker compose --env-file ".env.dev" down 


docker_kill: cmd-exists-docker
	@docker kill $(APP_NAME) 


docker_rebuild:
	@docker compose --env-file ".env.dev" build


cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)
