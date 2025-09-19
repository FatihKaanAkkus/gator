# gator

`gator` is a CLI tool for managing and consuming RSS feeds. Uses PostgreSQL and is written in Go.

## Prerequisites

Before you begin, make sure you have the following installed:

- [Go](https://go.dev/dl/) (version 1.20+ recommended)
- [PostgreSQL](https://www.postgresql.org/download/)

You should also have a running Postgres instance and a database created for `gator`.

## Installation

Clone this repository and install the CLI:

```bash
git clone https://github.com/FatihKaanAkkus/gator.git
cd gator
go build
go install
```

This will place the `gator` binary in your `$GOPATH/bin` (make sure it’s on your `$PATH`).

## Configuration

`gator` needs a config file to know how to connect to your database and which user is logged in.

Create a file named `config.json` in `$HOME/.gator/` folder:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
}
```

* `db_url` --> your Postgres connection string

## Usage

Once installed and configured, you can run commands like:

```bash
gator register <username>
gator login <username>
gator addfeed <feed-name> <feed-url>
gator following
gator browse
```

## Example

1. Register a user

    ```bash
    gator register fatih
    ```

1. Login as that user

    ```bash
    gator login fatih
    ```

1. Add a feed

    ```bash
    gator addfeed "fthkn.com blog" http://fthkn.com/blog/feed/
    ```

1. List feeds you’re following

    ```bash
    gator following
    ```

1. Browse posts

    ```bash
    gator browse
    ```
