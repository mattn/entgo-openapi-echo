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
$ entgen -driver sqlite3 -dsn ./entry.sqlite -rplural
2022/07/09 23:40:48 Generating ent/schema/entry.go  

```

Generate ent client.

```
$ go generate ./ent
```

Write entc.go

```
$ cat > ent/entc.go
//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entoas.NewExtension()
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ex))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
^D
```

Then modify ent/generate.go like below.

```
//go:generate go run -mod=mod entgo.io/ent/cmd/ent@latest generate ./schema
//go:generate go run -mod=mod entc.go
```

Generate OpenAPI schema.

```
$ go generate ./ent
$ ls ent/openapi.json
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
