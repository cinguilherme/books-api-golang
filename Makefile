build-image:
	docker build . -t go_api_book:latest

run-docker:
	docker run -p 3000:3000 go_api_book:latest

compose:
	docker-compose up