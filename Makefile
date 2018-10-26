install:
	@echo "[+] installing dependencies"
	@dep ensure --vendor-only
.PHONY: install

start: install
	@echo "[+] start"
	@docker-compose up
.PHONY: start

dev:
	@echo "[+] start in development mode"
	@docker-compose up -d
	@watchexec --restart --exts "go" --watch . "docker-compose restart app"
.PHONY: dev
