run-publisher:
	(export $$(grep -v '^#' .env | xargs) && go run cmd/publisher/main.go)
run-subscriber:
	(export $$(grep -v '^#' .env | xargs) && go run cmd/subscriber/main.go $(ARGS))
serve-restapi:
	(export $$(grep -v '^#' .env | xargs) && go run cmd/restapi/main.go)