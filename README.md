# SDT

[![CircleCI](https://circleci.com/gh/sensiblehq/sdt.svg?style=svg)](https://circleci.com/gh/sensiblehq/sdt)

SDT is a schema translation utility so you can:

- Copy table definitions between database providers
- Keep your data warehouse in sync with external sources
- Export table definitions in a format that can be analyzed

## Usage

```
sdt --target redshift \
    --table users \
    --schema public \
    --connection-string postgres://username:password@localhost:5432/database

CREATE TABLE users (
  id INTEGER,
  first_name VARCHAR(100),
  last_name VARCHAR(100)
)
```
