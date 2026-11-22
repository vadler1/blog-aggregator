# ensure server is running
start
```sh
sudo service postgresql start
```
status
```sh
sudo service postgresql status
```

# connect to db
Use psql in cli to connect to db:
```sh
psql 'postgres://postgres:<PASSWORD>@localhost:5432/gator'
```

# migration using goose
run migration
```sh
goose postgres -dir ./sql/schema/ '<connection_string>' up
```

get list of tables in psql
```sh
\dt
```
# generate typesafe query functions
```sh
sqlc generate
```
