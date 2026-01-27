.PHONY: web
web:
	cd ihm && npm install && node esbuild.js --dev

.PHONY: srv
srv:
	cd back && DIR=../ihm go run .

.PHONY: start
start:
	docker compose up -d

.PHONY: stop
stop:
	docker compose down