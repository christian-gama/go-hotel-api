GENERATED_DIR ?=./.generated
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
	CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/api/*.go
	@chmod +x $(BUILD_DIR)/$(APP_NAME)
	@echo "Build was generated at $(BUILD_DIR)/$(APP_NAME)"


test_integration: cmd-exists-go
	@TEST_MODE=integration go test -p 1 $(TEST_FLAGS) ./...


test: cmd-exists-go
	@TEST_MODE=unit go test $(TEST_FLAGS) ./...


cover: cmd-exists-go
	@if [ ! -d "$(COVERAGE_DIR)" ]; then \
		mkdir -p $(COVERAGE_DIR); \
	fi
	@make test TEST_FLAGS=-coverprofile=$(COVERAGE_DIR)/coverage.out 


cover-html: cmd-exists-go
	go tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html


lint: cmd-exists-golangci-lint
	golangci-lint run --config .golangci.yml 


migrate-create-%: cmd-exists-migrate
	@if [ -z "$(*)" ]; then \
		echo "Error: SEQ must be specified. migrate-create-SEQ"; \
		exit 1; \
	fi; \
	migrate create -ext $(MIG_EXT) -dir $(MIG_DIR) -seq -digits 4 $(*)


migrate-%: cmd-exists-migrate
	@if [ -z "$(*)" ]; then \
		echo "Error: expected [up|down|drop|force]"; \
		exit 1; \
	fi
	migrate -source file://$(MIG_DIR) -database $(DB_SGBD)://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE) $(*) $(VERSION)


clear:
	@rm -rf $(GENERATED_DIR)


docker: cmd-exists-docker
	@docker-compose --env-file "$(ENV_FILE)" run --name $(APP_NAME) -p $(APP_PORT):$(APP_PORT) --rm -e ENV_FILE=$(ENV_FILE) api $(ENTRY_POINT)


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
	@docker-compose --env-file ".env.dev" down 


docker_kill: cmd-exists-docker
	@docker kill $(APP_NAME) 


docker_rebuild:
	@docker-compose --env-file ".env.dev" build


cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)
