PORT=8080
DSN="host=localhost port=5432 user=root password=secret dbname=ecommerce_db sslmode=disable timezone=UTC connect_timeout=5"

DB_DOCKER_CONTAINER=ecommerce_db_container
BINARY_NAME=ecommerceplatform
SQLX_DOCKER_IMAGE=rust

ifeq ($(OS),Windows_NT)
    CURRENT_DIR = $(shell cd)
else
    CURRENT_DIR = $(shell pwd)
endif

# install make and run the commands one by one for the changes to take place.

# make create-network
create-network:
	docker network create ecommerce 

# creating the container with postgres software
# make postgres
postgres:
	@echo "${DB_DOCKER_CONTAINER}"
	docker run --network ecommerce --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

# creating ecommerce_db database inside the postgres container
# make createdb 
createdb:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=root --owner=root ecommerce_db 

# make stop_containers
stop_containers:
	@echo "Stopping other docker containers..." \
	if [ $$(docker ps -q) ]; then \
		echo "found and stopped containers..."; \
		docker stop $$(docker ps -q); \
	else \
		echo "no active containers found..."; \
	fi \

# make start-docker
start-docker:
	docker start ${DB_DOCKER_CONTAINER}

# make create_migrations
create_migrations:
	@echo "Current directory: ${CURRENT_DIR}"
	docker run --network ecommerce --rm -v "${CURRENT_DIR}:/app" -w /app ${SQLX_DOCKER_IMAGE} bash -c "cargo install sqlx-cli && sqlx migrate add -r init"

# make migrate-up
migrate-up:
	@echo "Running migrations..."
	docker run --network ecommerce --rm -v "${CURRENT_DIR}:/app" -w /app ${SQLX_DOCKER_IMAGE} bash -c "cargo install sqlx-cli && sqlx migrate run --database-url 'postgres://root:secret@ecommerce_db_container:5432/ecommerce_db?sslmode=disable'"

# make migrate-down
migrate-down:
	docker run --network ecommerce --rm -v "${CURRENT_DIR}:/app" -w /app ${SQLX_DOCKER_IMAGE} bash -c "cargo install sqlx-cli && sqlx migrate revert --database-url 'postgres://root:secret@ecommerce_db_container:5432/ecommerce_db?sslmode=disable'"

# make build
build:
	@echo "Building backend api binary"
	go build -o ${BINARY_NAME} cmd/server/*.go 
	@echo "Binary built ..."

# make start    [first runs build then run stop_containers then start-docker then execute this statements]
start: build stop_containers start-docker
	@set PORT=${PORT} DSN=${DSN} .\\${BINARY_NAME}.exe & \
	@echo "api started!" \
 	go run cmd/server/main.go 

# make stop
stop: 
	@echo "Stopping backend"
	@-pgkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped backend"

restart: stop start  