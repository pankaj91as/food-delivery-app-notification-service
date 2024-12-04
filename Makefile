run-publisher:
	(export $$(grep -v '^#' .env | xargs) && go run cmd/publisher/main.go)
run-subscriber:
	(export $$(grep -v '^#' .env | xargs) && go run cmd/subscriber/main.go $(ARGS))
serve-restapi:
	(export $$(grep -v '^#' .env | xargs) && go run cmd/restapi/main.go)
build-restapi:
	(GOOS=linux CGO_ENABLED=0 go build -tags netgo -a -v -o scripts/build/restapi cmd/restapi/main.go)
restapi:
	(export $$(grep -v '^#' ./scripts/build/.env | xargs) && ./scripts/build/restapi)
