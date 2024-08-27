.PHONY: docs
docs:
	swag init -g ./cmd/app/main.go

.PHONY: api
api: docs
	go run ./cmd/app/main.go

.PHONY: ent
ent:
	go generate ./ent