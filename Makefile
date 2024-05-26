CURRENT_DIR=$(shell pwd)

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	./scripts/gen-proto.sh	${CURRENT_DIR}


DB_URL := "postgres://postgres:1234@localhost:5432/users?sslmode=disable"

migrate-up:
	migrate -path migrations -database $(DB_URL) up

migrate-down:
	migrate -path migrations -database $(DB_URL) down

migrate-file:
	migrate create -ext sql -dir migrations/ -seq insert_table_workers

migrate-dirty:
	migrate -path ./migrations/ -database $(DB_URL) force 0

run:
	go run cmd/main.go
