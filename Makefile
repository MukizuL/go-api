run: build
	@docker run --rm -p 8000:8000 --name backend-project MukizuL/backend-project

build:
	@docker build -t MukizuL/backend-project .

test:
	@go test -v ./...

test_with_server: build
	@echo "Starting Go server..."
	@docker run -d --rm -p 8000:8000 --name backend-project MukizuL/backend-project
	@timeout /t 3
	@echo "Running tests..."
	@trap 'docker stop backend-project' EXIT; \
    	pytest Tests/*