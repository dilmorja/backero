TESTFLAGS = -v -cover

test18:
	@go1.18 test $(TESTFLAGS) ./...

test:
	@go test $(TESTFLAGS) ./...