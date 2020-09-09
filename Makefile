NAME=orchestrator-tester
VERSION=0.0.1

# OUTSIDE DOCKER
################
start:
	docker-compose run --service-ports $(NAME)

restart:
	docker container prune -f
	docker-compose down --volumes --rmi all
	docker-compose run --service-ports $(NAME)

stop:
	docker container prune -f
	docker-compose down --volumes --rmi all


# INSIDE DOCKER
################
.PHONY: init
init:
	@go mod init $(NAME)

.PHONY: build
build:
	@go build -o build/$(NAME)

.PHONY: run
run: build
	@./build/$(NAME)

.PHONY: clean
clean:
	@rm -r build
