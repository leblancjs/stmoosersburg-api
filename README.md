[![Build Status](https://travis-ci.org/leblancjs/stmoosersburg-api.svg?branch=master)](https://travis-ci.org/leblancjs/stmoosersburg-api)
[![codecov](https://codecov.io/gh/leblancjs/stmoosersburg-api/branch/master/graph/badge.svg)](https://codecov.io/gh/leblancjs/stmoosersburg-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-purple.svg)](https://opensource.org/licenses/MIT)

# St-Moosersburg API
The API behind the board game with wealthy and influential moose.

## Configuration
To run the service, some environment variables need to be set.

It is recommended to keep them in a *dotenv* file named `.env`, which is listed in the `.gitignore` file to keep it out of version control, since database credentials and API keys aren't meant to be shared with the world.

### Database
There are currently two kinds of databases supported: *in memory* and *Postgres*.

By default, an in memory database is used, which is not suitable when the service is deployed on a server, since all data is lost when it is shutdown.

To use a real database, such as Postgres, the `DB_TYPE` variable must be set.

```
# Defaults to "inmemory"
DB_TYPE=inmemory|postgres
```

#### Postgres
A few additional environment variables need to be set to connect to a Postgres database: `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, and `DB_SSL_MODE`.

Their default values configure a connection to a Postgres database hosted on localhost.

```
# Defaults to "localhost"
DB_HOST=www.domain.com

# Defaults to "5432"
DB_PORT=1234

# Defaults to "postgres"
DB_USER=username

# Defaults to an empty string ""
# Omit, or leave blank if the database user doesn't have a password
DB_PASSWORD=password

# Defaults to "stmoosersburg"
DB_NAME=database

# Defaults to "disable" for databases hosted on localhost
DB_SSL_MODE=required|disable
```

##### Schema
An SQL script is provided to create the tables required to setup the database schema. It can be found [here](db/postgres/schema.sql).

## Build and Run
### Using Go Run
To facilitate running the service from a terminal or a command prompt, a shell script and a batch file are provided. They take care of setting the environment variables defined in the `.env` file (see the [Configuration](#Configuration) section for more details).

#### macOS and Linux
Open a terminal and enter the following command:

```
./run.sh
```

#### Windows
> **NOTE:** The batch file *has not* been tested.

Open a command prompt and enter the following command:

```
run.bat
```

### Using Visual Studio Code
A Visual Studio Code (VS Code) launch configuration is provided. It automatically sets the environment variables defined in the `.env` file.

This method is recommended for debugging.

## Test
> **NOTE:** It may be necessary to set the `GO111MODULE` environment variable to *on*, depending on where the repository was cloned.

To run the tests, open a terminal or a command prompt and enter the following command:

```
go test ./... -coverprofile=coverage.out -covermode=atomic
```

The test coverage report can be viewed in a web browser by using the following command:

```
go tool cover -html=coverage.out
```