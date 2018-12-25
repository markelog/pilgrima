install:
	@echo "[+] installing dependencies"
	@dep ensure
.PHONY: install

start: install
	@echo "[+] start"
	@docker-compose up
.PHONY: start

db: 
	@echo "[+] set up db"
	@docker-compose up --renew-anon-volumes -d db
.PHONY: db

dev: db
	@echo "[+] start in development mode (docker)"
	@docker-compose up app
.PHONY: dev

test:
	@echo "[+] test"
	@docker-compose up -d db
	@go test -race -test.parallel 1 ./...
.PHONY: test

watch-test:
	@echo "[+] watch tests"
	@docker-compose up --renew-anon-volumes -d db
	@watchexec --restart --exts "go" --watch . "go test ./..."
.PHONY: watch-test

watch:
	@echo "[+] start in development mode"
	@watchexec --restart --exts "go" --watch . "go run main.go"
.PHONY: watch
