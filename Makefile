install-tools:
	cat tools.go | grep _ | grep \".*\" -o | xargs -tI % go install %
# put the files with interfaces you'd like to mock in prerequisites
mocks: pkg/party/greeter.go pkg/app/service.go
	@echo "Generating mocks..."
	@for file in $^; do mockgen -source=$$file -destination=mocks/$$file; done