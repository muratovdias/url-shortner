run:
	docker build -t url-shortner .
	docker run --name url-shortner-container -p 8080:8080 url-shortner -d