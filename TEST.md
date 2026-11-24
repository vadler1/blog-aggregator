# How to start
Written in Go - so make sure to have it installed

## Installation
Can be installed with
```sh
go install
```
This should retrieve the required dependencies

# Database
Requires a postgresql instance.

## ensure server is running (example: linux)
start
```sh
sudo service postgresql start
```
status
```sh
sudo service postgresql status
```

## connect to db
Use psql in cli to connect to db:
```sh
psql 'postgres://postgres:<PASSWORD>@localhost:5432/gator'
```

## migration using goose
run migration
```sh
goose postgres -dir ./sql/schema/ '<connection_string>' up
```

get list of tables in psql
```sh
\dt
```
## generate typesafe query functions
```sh
sqlc generate
```

# usage
The gator cli has the following commands

- `gator reset`: deletes users, and with that all attached data (including feeds and posts)
- `gator login <user>`: sets the provided user name as active
- `gator register <user>`: creates a new user with provided name
- `gator users`: lists all registered users
- `gator addfeed <name> <url>`: add a new feed (and automatically follow it)
- `gator feeds`: list all feeds
- `gator follow <url>`: follow provided url
- `gator unfollow <url>`: unfollow provided url
- `gator following`: list followed feeds
- `gator agg <interval>`: runs continuously with provided duration (e.g. 5m will iterate every 5 minutes), getting the oldest not yet fetched feed, and stores the retrieved posts
- `gator browse <?limit=2>` retrieves up to limit posts from followed feeds, sorted by publishing date
