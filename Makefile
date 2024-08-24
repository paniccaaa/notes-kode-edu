run:
	@go run cmd/notes-kode-edu/main.go

MIGRATE = migrate -path ./migrations -database "postgres://user:password@localhost:5432/notesdb?sslmode=disable"

migration_up:
	$(MIGRATE) up

migration_down:
	$(MIGRATE) down

migration_force:
	$(MIGRATE) force $(version)

migration_version:
	$(MIGRATE) version