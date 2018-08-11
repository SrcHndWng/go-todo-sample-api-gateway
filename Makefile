build:
	GOOS=linux go build -ldflags="-s -w" -o bin/create handlers/create/main.go
	GOOS=linux go build -ldflags="-s -w" -o bin/list handlers/list/main.go
	GOOS=linux go build -ldflags="-s -w" -o bin/get handlers/get/main.go
	GOOS=linux go build -ldflags="-s -w" -o bin/update handlers/update/main.go
	GOOS=linux go build -ldflags="-s -w" -o bin/delete handlers/delete/main.go