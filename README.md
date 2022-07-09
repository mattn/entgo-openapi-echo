# entgo-openapi-echo

Example app that using OpenAPI, ent and echo.

## Installation

Create database.

```
$ sqlite3 entry.sqlite < entry.sql
```

Generate ent's schema.

```
$ go install github.com/mattn/entgen@latest
$ entgen -driver sqlite3 -dsn ./entry.sqlite
2022/07/09 23:40:48 Generating ent/schema/entries.go  

```

Generate ent client..

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
