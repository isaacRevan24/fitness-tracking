generate-mock:
	echo "Generating mocks"
	go generate -v ./...

run-tests:
	echo "Running all"
	go test ./test/... -v