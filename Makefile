DB_DRIVER=postgres
DB_HOST=127.0.0.1
DB_PORT=5432
DB_NAME=db_yuyuid
DB_USER=yuyuid
DB_PASS=
DB_TZ=Asia/Jakarta
DB_SSL_MODE=disable

MIGRATE_CMD=migrate

DB_DSN=$(DB_DRIVER)://$(DB_USER)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)

# migration path
MIGRATION_DIR=$(shell pwd)/platform/migration

.PHONY: all migrate-up migrate-down migrate-force migrate-drop migrate-new run build

run:
	go run cmd/app/main.go

build:
	go build cmd/app/main.go

migrate-up:
	$(MIGRATE_CMD) -path $(MIGRATION_DIR) -database "$(DB_DSN)" up
migrate-down:
	$(MIGRATE_CMD) -path $(MIGRATION_DIR) -database "$(DB_DSN)" down
migrate-force:
	$(MIGRATE_CMD) -path $(MIGRATION_DIR) -database "$(DB_DSN)" force $(version)
migrate-drop:
	$(MIGRATE_CMD) -path $(MIGRATION_DIR) -database "$(DB_DSN)" drop -f

migrate-new:
	$(MIGRATE_CMD) create -ext sql -dir $(MIGRATION_DIR) -seq $(name)
