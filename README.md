# Untitled TTRPG discord bot service thingie

For now setting this up as monorepo of services but might only have one.  

## .devcontainer

See <https://github.com/qdm12/godevcontainer/tree/master>

## DB

### Go Definitions

`sqlc generate`

### Migration Scripts

<https://atlasgo.io/versioned/intro>

Inspect database

`atlas schema inspect -u "postgres://postgres:postgres@db:5432/chronicle?sslmode=disable"`

Inspect schema

`atlas schema inspect -u "postgres://postgres:postgres@db:5432/chronicle?search_path=public&sslmode=disable"`

```shell
atlas schema apply \
  -u "postgres://postgres:postgres@db:5432/chronicle?sslmode=disable" \
  --to file://schema.sql \
  --dev-url "postgres://postgres:postgres@db:5432/chronicle?sslmode=disable"
```

Initialize

```shell
atlas migrate diff initial \
  --to file://schema.sql \
  --dev-url "postgres://postgres:postgres@db:5432/chronicle?sslmode=disable" \
  --format '{{ sql . "  " }}'
```

Update schema

```shell
atlas migrate diff add_commits \
  --to file://schema.sql \
  --dev-url "postgres://postgres:postgres@db:5432/chronicle?sslmode=disable" \
  --format '{{ sql . "  " }}'
```

Run all migrations on a new database

```shell
atlas migrate apply \
  --url "postgres://postgres:postgres@db:5432/chronicle?sslmode=disable"  \
  --dir "file://migrations" 
```

Migrate from a baseline, the baseline will be ignored

```shell
atlas migrate apply \
  --url "postgres://postgres:postgres@db:5432/chronicle?sslmode=disable"  \
  --dir "file://migrations" \
  --baseline "20240731025606"
```

Using the free version of Atlas, we need to manually add extensions.

Right now we are adding them to the root migration (probably should make this a separate first migration).
