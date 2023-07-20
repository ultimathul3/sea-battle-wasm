run:
	DEVELOPMENT=1 go run cmd/app/main.go

lint:
	golangci-lint run

wasm:
	@GOOS=js GOARCH=wasm go build -o sea-battle.wasm cmd/app/main.go
