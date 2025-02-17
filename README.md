# 🐊 Blog Aggregator 🐊

Blog Aggregator is an RSS feed aggregator written in Go. It allows users to register, add RSS feeds, follow/unfollow feeds, and browse posts from the feeds they follow.

## Features

- User registration and login
- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/NachoGz/blog-aggregator.git
   cd blog-aggregator
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Set up the database:

   ```sh
   # Ensure you have a PostgreSQL database running (Postgres v15 or later.)
   # Update the database URL in the configuration file `~/.gatorconfig.json`.
   ```

4. Run database migrations:

   ```sh
   cd sql/schema
   goose up
   ```

5. Build the binary
   ```sh
   go build ./gator
   ```

## Usage

Note: you can replace `./gator` with `go run .`

### Commands

- **login**: Set the current user.

  ```sh
  ./gator login <username>
  ```

- **register**: Add a new user to the database.

  ```sh
  ./gator register <username>
  ```

- **users**: List all users in the database.

  ```sh
  ./gator users
  ```

- **reset**: Reset the state of the database (delete all users and associated records).

  ```sh
  ./gator reset
  ```

- **agg**: Fetch RSS feeds, parse them, and store them as posts in the database.

  ```sh
  ./gator agg <time-between-reqs>
  ```

- **addfeed**: Add a feed to the database.

  ```sh
  ./gator addfeed <feed-name> <feed-url>
  ```

- **feeds**: List all feeds in the database.

  ```sh
  ./gator feeds
  ```

- **follow**: Follow a feed.

  ```sh
  ./gator follow <feed-url>
  ```

- **following**: List all feeds that the current user is following.

  ```sh
  ./gator following
  ```

- **unfollow**: Unfollow a feed.

  ```sh
  ./gator unfollow <feed-url>
  ```

- **browse**: View all posts from the feeds the user follows.
  ```sh
  ./gator browse <limit>
  ```

### Configuration

Create a configuration file `~/.gatorconfig.json` in your home directory with the following content:

The database URL should be of this format `protocol://username:password@host:port/database?sslmode=disable`

```json
{
  "db_url": "your_database_url",
  "current_user_name": ""
}
```
