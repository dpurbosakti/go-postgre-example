build:
	GOOS=linux GOARCH=amd64 go build -o belajar-echo ./app

run:
	go run app/main.go

seeder:
	go run database/seeders/seeder.go