# Graph Analyzer - API
The API is responsible for calculating the graph properties. 
It fetches the data from the Neo4j database and makes the calculated 
properties available via a REST interface.


## Run Locally

### Checkout the project
```zsh
git clone https://github.com/Graph-Analyzer/api.git
```

### Configuration
Create a `.env` file locally, based on the provided `.env.example`.

### Run

```zsh
go build && ./api
```

### Access

- [Swagger](http://localhost:8080/api/doc/index.html)
- [API](http://localhost:8080/api)


### Mocks

Tests use mocks generated by [mockery](https://github.com/vektra/mockery).
Please follow the installation instructions to set it up.

Regenerate the mocks by running this command.

```zsh
mockery
```

### Code Formatting
Code formatting is done using [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports).
Installation instructions can be found in the official documentation.

Run the code formatting with the following command:

```zsh
goimports -w .
```

## OpenAPI v2 documentation

### Installation
Follow the guide here: [https://github.com/swaggo/swag](https://github.com/swaggo/swag)

### Access
The swagger docs are accessible under `/api/doc/index.html`.

### Generation
You can regenerate the docs with the following command.

```zsh
swag init
```

### Formatting
You can format the annotations in the code with the following command.

```zsh
swag fmt
```

## gRPC client
Copy the generated files over from the data-collector

## Authors

- [@lribi](https://github.com/lribi)
- [@pesc](https://github.com/pesc)

## License

This project is licensed under the [MIT](https://github.com/Graph-Analyzer/api/blob/main/LICENSE) License.

### Third Party Licenses

Third party licenses can be found in `THIRD-PARTY-LICENSES.txt`.
Regenerate them with this command.

```zsh
# Make sure you have a compiled version of the project in the project root
go install github.com/uw-labs/lichen@latest

./scripts/go-licenses.sh
```
