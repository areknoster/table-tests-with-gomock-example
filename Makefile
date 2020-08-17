install-tools:
	cat tools.go | grep _ | grep \".*\" -o | xargs -tI % go install %


MOCKS_DESTINATION=mocks
.PHONY: mocks
# put the files with interfaces you'd like to mock in prerequisites
# wildcards are allowed
mocks: pkg/party/greeter.go pkg/party/visitor-lister.go
	@echo "Generating mocks..."
	@rm -rf $(MOCKS_DESTINATION)
	@for file in $^; do mockgen -source=$$file -destination=$(MOCKS_DESTINATION)/$$file; done