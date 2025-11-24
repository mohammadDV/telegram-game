test:
	go test ./... -v

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-service:
	go test ./internal/service/... -v

test-entity:
	go test ./internal/entity/... -v

test-integration:
	TEST_INTEGRATION=true go test ./... -v