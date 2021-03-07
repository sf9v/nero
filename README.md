[![GoDoc Reference](https://pkg.go.dev/badge/github.com/sf9v/nero)](https://pkg.go.dev/github.com/sf9v/nero)
![Github Actions](https://github.com/sf9v/nero/workflows/test/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/sf9v/nero/badge.svg?branch=main)](https://coveralls.io/github/sf9v/nero?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/sf9v/nero)](https://goreportcard.com/report/github.com/sf9v/nero)

# Nero

A library for generating the repository pattern.

## Motivation

We heavily use the _[repository pattern](https://threedots.tech/post/repository-pattern-in-go/)_ in our codebases and we often [write our queries manually](https://golang.org/pkg/database/sql/#example_DB_QueryContext). It becomes tedious and repetitive as we have more tables/models to maintain. So, we decided to experiment on creating this library to generate our repositories automatically.

## Goals

- Decouple implementation from the `Repository` interface
- Easy integration with existing codebase
- Minimal API

## Installation

```console
$ go get github.com/sf9v/nero
```

## Example

See the [integration test](./test/integration/playerrepo) for a more complete example.

```go
import (
    "database/sql"

    _ "github.com/lib/pq"

    // import the generated package
    repo "github.com/sf9v/nero-example/repository"
)

func main() {
    dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
    db, err := sql.Open("postgres", dsn)
    ...

    ctx := context.Background()


    // initialize a product repository and enable debug
    productRepo := repo.NewPostgresRepository(db).Debug()

    // create
    creator := repo.NewCreator().Name("Product 1")
    productID, err := productRepo.Create(ctx, creator)
    ...

    // query
    queryer := repo.NewQueryer().Where(repo.IDEq(product1ID))
    product, err := productRepo.QueryOne(ctx, queryr)
    ...

    // update
    now := time.Now()
    updater := repo.NewUpdater().Name("Updated Product 1").
        UpdatedAt(&now).Where(repo.IDEq(product1ID))
    _, err = productRepo.Update(ctx, updater)
    ...

    // delete
    deleter := repo.NewDeleter().Where(repo.IDEq(product1ID))
    _, err = productRepo.Delete(ctx, deleter)
    ...
}
```

## Supported back-ends

Below is the list of supported back-ends.

| Back-end             | Library                                                       |
| -------------------- | ------------------------------------------------------------- |
| PostgreSQL           | [lib/pq](http://github.com/lib/pq)                            |
| SQLite (soon)        | [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)       |
| MySQL/MariaDB (soon) | [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) |

If your your back-end is not yet supported, you can implement your own [custom back-end](#custom-back-ends).

## Custom back-ends

Implementing a custom back-end is very easy. In fact, you don't have to use the official back-ends. You can implement custom back-ends (Oracle, MSSQL, BoltDB, MongoDB, etc.) by implementing the [_Templater_](./template.go) interface.

See official [postgres template](./pg_template.go) for reference.

## Standing on the shoulders of giants

This project wouldn't be possible with the amazing open-source projects it was built on.

- [Masterminds/squirrel](https://github.com/Masterminds/squirrel) - Fluent SQL generation in golang
- [lib/pq](https://github.com/lib/pq) - Pure Go Postgres driver for database/sql
- [pkg/errors](https://github.com/pkg/errors) - Simple error handling primitives
- [hashicorp/multierror](https://github.com/hashicorp/go-multierror) - A Go (golang) package for representing a list of errors as a single error.
- [jinzhu/inflection](https://github.com/jinzhu/inflection) - Pluralizes and singularizes English nouns
- [stretchr/testify](https://github.com/stretchr/testify) - A toolkit with common assertions and mocks that plays nicely with the standard library

Also, the following have a huge influence on this project and deserves most of the credits:

- [ent](https://github.com/facebook/ent) - An entity framework for Go. Simple, yet powerful ORM for modeling and querying data.
- [SQLBoiler](https://github.com/volatiletech/sqlboiler) - Generate a Go ORM tailored to your database schema.

## Contributing

Any suggestions and ideas are very much welcome, feel free to [open an issue](https://github.com/sf9v/nero/issues) or [make a pull request](https://github.com/sf9v/nero/pulls)!

## License

[Apache-2.0](LICENSE)
