# dbcook

A CLI tool for generating and running SQL migrations, built in Go.

## How to use

**Initialize a project**

Creates a `dbcook.toml` config file in the current directory:

```bash
dbcook init
```

**Generate a migration**

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

**OPTIONS**

**output_path** &mdash; sets the output path relative to the location of the config file.

## Generated files

The rules for determining the output location of the generated migration files are the following:

1. **`--output <dirpath>` flag** &mdash; explicitly sets the location to the provided path.
2. **`dbcook.toml` config file** &mdash; If present, it'll use a relative location defined by the `output_path`.
3. **Current working directory** &mdash; if no flag and no config are found, file is created in the current working directory.


## Usage Examples

```bash
# No config, migration file is created in the current working directory
dbcook generate posts title url

# With a dbcook.toml file, the file is placed at a relative location defined in by the output_path config
dbcook generate posts title url

# Output location set by the --output flag
dbcook generate posts title url --output ./tmp/migrations
```
