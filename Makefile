webdev:
	npm install && node esbuild.js --dev

srvdev:
	cd back && DIR=../ihm go run .

build:
	docker compose build

start:
	docker compose up -d

stop:
	docker compose down