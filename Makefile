# COLORS
ccgreen=$(shell printf "\033[32m")
ccred=$(shell printf "\033[0;31m")
ccyellow=$(shell printf "\033[0;33m")
ccend=$(shell printf "\033[0m")

# Environment variables for project
-include $(PWD)/cmd/auth/.env

# Export all variable to sub-make
export

migration-create:
	@./bin/migrate create -ext sql -dir ./database/migrations $(name)

DB_URL=$(DB_ENGINE)://$(DB_USER):$(DB_PASSWORD)@$(DB_SERVER):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
migration-up:
	@echo $(DB_URL)
	@./bin/migrate -source file://database/migrations -database $(DB_URL) up $(count)

migration-down:
	@./bin/migrate -source file://database/migrations -database $(DB_URL) down $(count)


# SILENT MODE (avoid echoes)
.SILENT: all fmt test linter build

# PROCESS
all: fmt test linter build

fmt:
	@printf "$(ccyellow)Formatting files...$(ccend)\n"
	$(GOPATH)/bin/goimports -w -l -local gitlab.com/EDteam/shenlong .
	@printf "$(ccgreen)Formatting files done!$(ccend)\n"

fmt-one:
	@printf "$(ccyellow)Formatting files...$(ccend)\n"
	$(GOPATH)/bin/goimports -w -l -local gitlab.com/EDteam/shenlong $(file)
	@printf "$(ccgreen)Formatting files done!$(ccend)\n"

test:
	@printf "$(ccyellow)Testing files...$(ccend)\n"
	go test -race ./...
	@printf "$(ccgreen)Finished Testing files...$(ccend)\n"

vulnerability:
	@printf "$(ccyellow)Vulnerability check...$(ccend)\n"
	go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...
	@printf "$(ccgreen)Finished Vulnerability check...$(ccend)\n"

test-detail:
	@for d in $$(go list ./...); do \
		if go test -v -failfast $$d; then \
			printf "$(ccyellow)$$d test pass!!!$(ccend)\n"; \
		else \
			printf "$(ccred)$$d test failed :($(ccend)\n"; \
			exit 1; \
		fi; \
	done; \
	printf "$(ccgreen)All test pass!$(ccend)\n"

linter:
	@printf "$(ccyellow)Executing linter...$(ccend)\n"
	$(GOPATH)/bin/golangci-lint run --fast
	@printf "$(ccgreen)Linter finished!$(ccend)\n"

install-linter:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin latest

install-goimports:
	go install golang.org/x/tools/cmd/goimports@latest

install-migrate:
	sudo go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

copy-git-hooks:
	chmod +x scripts/git-hooks/*
	ln -s -f $(PWD)/scripts/git-hooks/pre-commit .git/hooks/pre-commit

copy-envs:
	cp .env.example cmd/api/.env

setup: install-linter install-migrate install-goimports copy-git-hooks copy-envs
	echo "Setup done!"

run:
	clear
	mkdir -p bin
	go build -o ./bin/api ./cmd/api
	ENV=local CONFIGURATION_FILEPATH=$(PWD)/cmd/api/.env ./bin/api
