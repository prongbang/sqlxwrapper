# sqlxwrapper

### Install

```
go get github.com/prongbang/sqlxwrapper
```

### Usage

- Count

```go
pqwrapper.Count(conn, query, args...)
```

- Select List

```go
pqwrapper.SelectList[Type](conn, query, args...)
```

- Select One

```go
pqwrapper.SelectOne[Type](conn, query, args...)
```

- Update

```go
pqwrapper.Update(conn, query, set, args...)
```

- Delete

```go
pqwrapper.Delete(conn, query, args...)
```
