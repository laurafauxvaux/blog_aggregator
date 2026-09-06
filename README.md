# Blog Aggregator

A command-line RSS feed aggregator written in Go as part of the Boot.dev backend curriculum.

## Requirements

To run the application, you need to have:
- PostGreSQL 
- Go

## Installation
```bash
go install github.com/laurafauxvaux/blog_aggregator@latest
```

## Configuration

Create a `.gatorconfig.json` file in your home directory:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

Update `db_url` with the connection information for your PostgreSQL database.

## Run

```bash
gator <command> [arguments]
```

## Commands

- `register <username>` — register a user
- `reset` — remove all users from the database
- `addfeed <name> <url>` — add a feed to the database
- `feeds` — shows the feeds already existing in the database
- `follow <url>/unfollow <url>` — follow or unfollow a feed
- `agg <duration>` — continuously collects posts from feeds
- `browse [limit]` — displays the most recent saved posts
