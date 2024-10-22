deep:
	dep-tree entropy cmd/main.go
.PHONY: deep

run:
	docker compose up --build
.PHONY: run

cover-html: ### run test with coverage and open html report
	go test -coverprofile=coverage.out -covermode=branch ./...
	go tool cover -html=coverage.out
	rm coverage.out
.PHONY: coverage-html

cover: ### run test with coverage
	go test -coverprofile=coverage.out -covermode=branch ./...
	go tool cover -func=coverage.out
	rm coverage.out
.PHONY: coverage