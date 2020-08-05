install-tools:
	cat tools.go | grep _ | grep \".*\" -o | xargs -tI % go install %

generate-mocks: service/service.go
	@echo "Generating mocks..."
	@for file in $^; do mockgen -source=$$file -destination=mocks/mock_$$file; done%