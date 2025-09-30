# Blog Aggregator ("gator")

## Overview

**gator** is a command-line blog/RSS aggregator written in Go. It allows users to register, add feeds, follow/unfollow feeds, and browse posts from followed feeds. The application uses PostgreSQL as its database backend.

## Prerequisites

Before running **gator**, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.24.4 or higher)
- [PostgreSQL](https://www.postgresql.org/download/)

## Setup

1. **Clone the repository:**
   ```sh
   git clone https://github.com/Samuel-Tarifa/blog-aggregator.git
   cd blog-aggregator
   ```

2. **Configure your database:**
   - Create a PostgreSQL database.
   - Apply the migrations in `sql/schema/` to set up the tables.

3. **Set up the config file:**
   - Create a file named `.gatorconfig.json` in your home directory:
     ```json
     {
       "db_url": "postgres://username:password@localhost:5432/dbname?sslmode=disable",
       "current_user_name": ""
     }
     ```
   - Replace `username`, `password`, and `dbname` with your PostgreSQL credentials.

## Building

Build the project using Go:

```sh
go build -o gator main.go
```

## Usage

Run the `gator` command with one of the supported subcommands:

```sh
./gator <command> [arguments...]
```

### Common Commands

- **register `<username>`**  
  Register a new user.
  ```sh
  ./gator register alice
  ```

- **login `<username>`**  
  Set the current user.
  ```sh
  ./gator login alice
  ```

- **addfeed `<name>` `<url>`**  
  Add a new feed and follow it.
  ```sh
  ./gator addfeed "My Blog" "https://myblog.com/rss"
  ```

- **feeds**  
  List all available feeds.
  ```sh
  ./gator feeds
  ```

- **follow `<feed_url>`**  
  Follow a feed by its URL.
  ```sh
  ./gator follow "https://myblog.com/rss"
  ```

- **following**  
  List feeds you are following.
  ```sh
  ./gator following
  ```

- **unfollow `<feed_url>`**  
  Unfollow a feed.
  ```sh
  ./gator unfollow "https://myblog.com/rss"
  ```

- **browse `[limit]`**  
  Browse posts from feeds you follow. Optionally specify the number of posts to show (default: 2).
  ```sh
  ./gator browse 5
  ```

- **agg `<interval>`**  
  Periodically fetch new posts from feeds.  
  Example: every 10 minutes
  ```sh
  ./gator agg 10m
  ```

- **reset**  
  Reset the database (delete all users).

## Notes

- The config file `.gatorconfig.json` stores your database connection string and the current user.
- Make sure your PostgreSQL server is running and accessible via the connection string in your config file.

## License

MIT