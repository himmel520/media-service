include .env
export

TOOLS_PATH=bin/tools
migrate=$(TOOLS_PATH)/migrate
ogen=$(TOOLS_PATH)/ogen

$(migrate):
	GOBIN=`pwd`/$(TOOLS_PATH) go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

$(ogen):
	GOBIN=`pwd`/$(TOOLS_PATH) go install github.com/ogen-go/ogen/cmd/ogen@v1.6.0

setup: $(migrate) $(ogen)

.PHONY: generate.api
generate.api: $(ogen)
	$(ogen) --loglevel error --clean --config .ogen.yml --target ./api/oas ./api/openapi.yml
	docker run --rm -v `pwd`:/spec redocly/cli bundle ./api/openapi.yml > ./api/bundle.yml

.PHONY: run
run:
	go run cmd/main.go

# COMPOSE
.PHONY: compose.up
compose.up:
	docker compose -f deployments/dev/compose.yml -p media up --build --no-log-prefix --attach media

.PHONY: compose.down
compose.down:
	docker compose down

# MIGRATION
.PHONY: migrate.create
migrate.create: $(migrate)
	$(migrate) create -ext sql -dir migrations $(name) 

.PHONY: migrate.up
migrate.up:
	docker run --rm \
		--network media_default \
		-v `pwd`/migrations:/migrations \
		migrate/migrate \
		-path=/migrations -database $(POSTGRES_CONN) up

.PHONY: migrate.down
migrate.down:
	docker run --rm \
		--network media_default \
		-v `pwd`/migrations:/migrations \
		migrate/migrate \
		-path=/migrations -database $(POSTGRES_CONN) down -all

# TEST
cover-html: ### run test with coverage and open html report
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out
.PHONY: coverage-html

cover: ### run test with coverage
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm coverage.out
.PHONY: coverage

# ADDITION
.PHONY: deep
deep:
	dep-tree entropy cmd/main.go