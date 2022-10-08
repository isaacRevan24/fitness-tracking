generate-mock:
	echo "Generating mocks"
	go generate -v ./...

run-test-from-scratch:
	echo "Generating mocks"
	go generate -v ./...
	echo "Running all"
	go test ./test/... -v

run-tests:
	echo "Running all"
	go test ./test/... -v