setup:
	@echo " --- Setup and generate configuration --- "
	@cp internal/config/example/env.yml.example internal/config/env/env.yml
	@echo " --- Done Setup and generate configuration --- "

build: setup
	@echo "--- Building binary file ---"
	@go build -o ./main main.go

run:
	@echo "--- Run Program ---"
	@go run main.go