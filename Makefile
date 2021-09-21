
tidy:
	@go mod tidy

test:
	@go test ./... -v -count=1 -cover -covermode=atomic -coverprofile=cover.out

see:
	@go tool cover -html=cover.out