lint:
	golangci-lint run -c ./golangci.yml ./...

.PHONY: docs
docs:
	@echo "Generating docs..."
	swag fmt
	swag init -g purebackend.go

run:
	go run purebackend.go

build:
	go build -o purebackend.exe purebackend.go

seed:
	cd datastore/seeder && seeder

seed-test:
	cd datastore/seeder && seeder test

.PHONY: ui
ui:
	cd ui && pnpm serve-static

ui-server:
	cd ui && pnpm start

serve:
	purebackend.exe

# TO BE USED AS `make -j deploy` to run parallel tasks
deploy: serve ui

deploy-server: serve ui-server