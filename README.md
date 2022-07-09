# entgo-openapi-echo

Example app that using OpenAPI, ent and echo.

## Installation

Generate ent schema.

```
$ go generate ./ent
```

Generate OpenAPI schema.

```
$ cd ent
$ go generate
$ ls openapi.json
openapi.json

$ cd ..
```

Generate code from openapi.json.

```
$ go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
$ oapi-codegen -package main -generate server -old-config-style ent/openapi.json > oapi.go                   
```

Build

```
$ go build
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
