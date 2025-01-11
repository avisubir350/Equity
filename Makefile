FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp
AUTH_BINARY=authApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	cd project && docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_auth
	@echo "Stopping docker images (if running...)"
	cd project && docker-compose down
	@echo "Building (when required) and starting docker images..."
	cd project && docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	cd project && docker-compose down
	@echo "Done!"

# build_front builds the front end binary
build_front:
	@echo "Building front end binary..."
	cd ./front-service && env GOOS=darwin CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"
	@echo "Moving binary to cmd/web directory..."
	cd ./front-service && mv ./${FRONT_END_BINARY} ./cmd/web/${FRONT_END_BINARY}
	@echo "Done!"

# build_broker: builds the broker binary
build_broker:
	@echo "Building broker binary..."
	cd ./broker-service && env GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

# build_auth: builds the authentication binary
build_auth:
	@echo "Building auth binary..."
	cd ./authentication-service && env GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

# start_app builds the front end lib and starts the app
start_app: build_front
	@echo "Starting the app Equity Insights"
	cd ./front-service/cmd/web && ./${FRONT_END_BINARY} &

# stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./front-service/cmd/web/${FRONT_END_BINARY}"
	@echo "Stopped front end!"


