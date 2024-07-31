# Untitled TTRPG discord bot service thingie

For now setting this up as monorepo of services but might only have one.  

## .devcontainer

See <https://github.com/qdm12/godevcontainer/tree/master>

## DB

### Migration Scripts

<https://atlasgo.io/versioned/intro>

Inspect database

`atlas schema inspect -u "postgres://postgres:postgres@db:5432/chronicle?sslmode=disable"`

Inspect schema

`atlas schema inspect -u "postgres://postgres:postgres@db:5432/chronicle?search_path=public&sslmode=disable"`

```
atlas schema apply \
  -u "postgres://postgres:postgres@db:5432/chronicle?sslmode=disable" \
  --to file://schema.sql \
  --dev-url "postgres://postgres:postgres@db:5432/chronicle?sslmode=disable"
```

Initialize

```
atlas migrate diff initial \
  --to file://schema.sql \
  --dev-url "postgres://postgres:postgres@db:5432/chronicle?sslmode=disable" \
  --format '{{ sql . "  " }}'
```
