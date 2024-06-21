docker-compose up --build

psql -h localhost -U myuser -d mydb
DB_HOST=localhost DB_USER=myuser DB_PASSWORD=mypassword DB_NAME=mydb \
nodemon --exec "go run main.go" --signal SIGTERM
