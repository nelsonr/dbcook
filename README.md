# dbcook

A CLI tool for generating and running SQL schemas, built in Go.

## How to use

**Initialize a project**

Creates a `dbcook.toml` config file in the current directory:

```bash
dbcook init
```

**Generate a schema**

```bash
dbcook generate posts title url
```

Creates the following schema file below. Filenames are always prefixed with a timestamp.

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

The config file uses the **toml** format and can be automatically created by running the `dbcook init` command.

```toml
# dbcook.toml
output_path = "db/schema"
```

**OPTIONS**

**output_path** &mdash; sets the output path relative to location of the config file.

## Generated files

The rules for determining the location of the generated schema files are the following:

1. **`--out <dir>` flag** &mdash; explicitly sets the location to the provided path.
2. **`dbcook.toml` config file** &mdash; If present, it'll use relative location defined by the `output_path`.
3. **Current working directory** &mdash; if no flag and no config are found, files are created at the location where the command was run.


## Usage Examples

```bash
# No config, files are placed at the location where the command was run
dbcook generate posts title url

# With a dbcook.toml file, the files are placed at a relative location defined in by the output_path config
dbcook generate posts title url

# Override for a one-off
dbcook generate posts title url --out ./tmp/migrations
```
