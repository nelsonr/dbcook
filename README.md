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

**Run database migrations**

```bash
dbcook migrate db/blog.db
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

The `init` command initializes a `dbcook.toml` config file in the current working directory.

**Usage:**

```bash
dbcook init
```

### Generate

The `generate` (alias `g`) command generates a migration file from its arguments.

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

The rules for determining the output location of the migration files are the following:

1. **`--output <dirpath>` flag** &mdash; explicitly sets the location to the provided path.
2. **`dbcook.toml` config file** &mdash; If present, it'll use a relative location defined by the `output_path`.
3. **Current working directory** &mdash; if no flag and no config are found, file is created in the current working directory.

### Migrate

The `migrate` command runs the migrations files on the given SQLite database.

It will try to obtain the path to the migrations folder from the config file. 
If not found, it'll look for migration files in the current working directory.

The migrations are run by alphabetical order of the `.sql` files.

**Usage:**
```bash
# The second argument is the path to the database file.
dbcook migrate db/blog.db

# Connected to database: db/blog.db
# 
# Found 3 migration file(s)
# 
# Running migration: 1778958053498_create_posts.sql
# Running migration: 1778959968418_create_photos.sql
# Running migration: 1778962504881_create_comments.sql
# 
# Completed 3 migrations successfully!
```

## Roadmap

- [x] Implement `init` command
  - [ ] Add option for the database file path
- [x] Implement `generate` command
  - [ ] Add support for additional data types
- [x] Implement `migrate` command
  - [ ] Allow overriding the path to the migrations directory