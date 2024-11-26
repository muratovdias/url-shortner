run:
	docker build -t url-shortner .
	docker run --name url-shortner-container -p 8080:8080 url-shortner -d
docs:
	swag init --parseDependency --parseInternal --parseDepth 5 -g cmd/main.go -o swagger -d .
test:
	go test -coverpkg=./... ./tests/component_test.go -cover