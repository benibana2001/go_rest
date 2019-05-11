##Runtests
test:
	go test
db_init:
	go run database/migrations/DropTable.go
	go run database/migrations/CreateTable.go
