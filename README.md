# sqlxwrapper

### Install

```
go get github.com/prongbang/sqlxwrapper
```

### Usage

- Count

```go
count := pqwrapper.Count(conn, query, args...)
```

- Select List

```go
typs := pqwrapper.SelectList[Type](conn, query, args...)
```

- Select One

```go
typ := pqwrapper.SelectOne[Type](conn, query, args...)
```

- Create

```go
tx, err := pqwrapper.Create(conn, query, []any{&data.ID}, data.Name)
```

- Update

```go
tx, err := pqwrapper.Update(conn, query, set, args...)
```

- Delete

```go
tx, err := pqwrapper.Delete(conn, query, args...)
```
