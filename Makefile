test:
	go test -v -count=100 ./...

cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

#gen:
#    mockgen -source=test_with_mock/repository/repository.go \
#    -destination=test_with_mock/mocks/mock_repository.go
