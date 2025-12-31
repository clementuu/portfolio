.PHONY: webdev
webdev:
	npm install && node esbuild.js --dev

.PHONY: srvdev
srvdev:
	cd back && DIR=../ihm go run .

start:
	docker compose up -d

stop:
	docker compose down