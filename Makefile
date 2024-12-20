mock/generate:
	sh ./tools/mockgen.sh

run/management:
	go run cmd/http/main.go

run/bff/api:
	go run cmd/bff/http/main.go
