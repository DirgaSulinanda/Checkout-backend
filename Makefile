PROJECT_PATH="${GOPATH}/src/github.com/DirgaSulinanda/Checkout-Backend"

start:
	@sudo docker-compose up -d

stop:
	@sudo docker-compose down

rebuild:
	@sudo docker-compose build
