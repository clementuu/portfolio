.PHONY: web
web:
	cd ihm && npm install && node esbuild.js --dev

.PHONY: srv
srv:
	cd back && DIR=../ihm go run .

.PHONY: build
build:
	docker compose build

.PHONY: start
start:
	docker compose up -d

.PHONY: build-and-start
build-and-start: build start

.PHONY: stop
stop:
	docker compose down

.PHONY: test-back
test-back:
	cd back && go test ./... -coverprofile=coverage.out

.PHONY: test-ihm
test-ihm:
	cd ihm && npm run test:coverage

.PHONY: back-coverage
back-coverage:
	cd back && go tool cover -html=coverage.out

.PHONY: deploy
deploy:
	cd deploy && ansible-playbook -i hosts.ini deploy.yml

.PHONY: ssl
ssl:
	cd deploy && ansible-playbook -i hosts.ini ssl.yml