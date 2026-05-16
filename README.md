# dbcook

A CLI tool for generating and running SQL migrations, built in Go.

## How to use

**Initialize a project**

Creates a `dbcook.toml` config file in the current directory:

```bash
dbcook init
```

**Create a migration file**

```bash
dbcook generate posts title url
```

Creates the following migration file (filenames are always prefixed with a timestamp).

```sql
-- file: 1778764280802_create_posts.sql
CREATE TABLE IF NOT EXISTS posts (
  id INTEGER PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  url VARCHAR(255) NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
);
```

## Config file

The config file uses the **TOML** format and can be created by running the `dbcook init` command.

```toml
# file: dbcook.toml
output_path = "db/migrations"
```

The option **output_path** sets the output path for the generated migration files, relative to the location of the config file.

## Commands

This section describes the available commands of `dbcook`.

### Init

**Status:** *Not implemented yet*

### Generate

The `generate` (alias `g`) generates a migration file from its arguments.

**Usage:**

```bash
# Generates a migration file in the current directory or 
# relative to the path set in the dbcook.toml file, if present.
dbcook generate posts title url

# Use the --output flag to explicitly set the output path for
# the generated migration file.
dbcook generate posts title url --output ./tmp/migrations

# Use the --print flag to print to stdout instead of creating a file.
dbcook generate posts title url --print
```

The rules for determining the output location of the generated migration files are the following:

1. **`--output <dirpath>` flag** &mdash; explicitly sets the location to the provided path.
2. **`dbcook.toml` config file** &mdash; If present, it'll use a relative location defined by the `output_path`.
3. **Current working directory** &mdash; if no flag and no config are found, file is created in the current working directory.

### Migrate

**Status:** *Not implemented yet*

## Roadmap

- [x] Implement `init` command
- [x] Implement `generate` command
- [ ] Implement `migrate` command