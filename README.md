# Swag perf problems

## Clone and fix swag

`git clone https://github.com/swaggo/swag.git && cd swag`

`git checkout v1.7.3`

`git apply annotation_error.patch`

`go build -o swag_fix ./cmd/swag`

## Clone this repository

`cd ..`

`git clone https://github.com/matrixik/swag-perf.git`

`cd swag-perf`

Generate Swagger spec file:

With https://github.com/swaggo/swag

`time ../swag/swag_fix init -o swagger/docs --md swagger/markdown --parseDependency true`

With https://github.com/go-swagger/go-swagger

`time swagger generate spec -o ./swagger-ui/swagger.json --scan-models`

Start app:

`go run main.go`

Open swagger:

For swag output: http://localhost:8080/swagger/index.html

For go-swagger output: http://localhost:8080/swagger-ui/
