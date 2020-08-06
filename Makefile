install-tools:
	cat tools.go | grep _ | grep \".*\" -o | xargs -tI % go install %

generate-mocks: pkg/party/helloer.go pkg/app/service.go
	@echo "Generating mocks..."
	@for file in $^; do mockgen -source=$$file -destination=mocks/$$file; done