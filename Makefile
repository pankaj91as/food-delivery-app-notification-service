run-publisher:
	(export $$(grep -v '^#' .env | xargs) && go run cmd/publisher/main.go)
run-subscriber:
	(export $$(grep -v '^#' .env | xargs) && go run cmd/subscriber/main.go)
serve-restapi:
	(export $$(grep -v '^#' .env | xargs) && go run cmd/restapi/main.go)