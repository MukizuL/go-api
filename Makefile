run: build
	@./bin/api

build:
	@go build -o bin/api
	@go build -o bin/api.exe

test:
	@go test -v ./...

testpy:
	pytest Tests/tests.py

test_with_server: build
	@echo "Starting Go server..."
	@./bin/api &
	@echo "Running tests..."
	@pytest Tests/tests.py