.PHONY: lint
lint:
	golangci-lint run -c ./golangci.yml ./... --allow-parallel-runners --fix

.PHONY: docs
docs:
	@echo "Generating docs..."
	swag fmt
	swag init -g purebackend.go

.PHONY: run
run:
	go run purebackend.go

.PHONY: build
build:
	go build -o purebackend.exe purebackend.go

.PHONY: seed
seed:
	cd datastore/seeder && seeder

.PHONY: seed-test
seed-test:
	cd datastore/seeder && seeder test

.PHONY: ui
ui:
	cd ui && pnpm serve-static

.PHONY: ui-server
ui-server:
	cd ui && pnpm start

.PHONY: serve
serve:
	purebackend.exe

# TO BE USED AS `make -j deploy` to run parallel tasks
.PHONY: deploy
deploy: serve ui

.PHONY: deploy-server
deploy-server: serve ui-server

.PHONY: docker-build
docker-build:
	docker build --pull --rm -f "Dockerfile" -t purebackend:v0.0.2 "."

.PHONY: docker-run
docker-run:
	docker run -it --rm --env-file=.env -p 8080:8080/tcp purebackend:v0.0.2