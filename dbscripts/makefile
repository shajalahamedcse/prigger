DB_HOST:=localhost
DB_USER:=myuser
DB_NAME:=mydb
DB_PASSWORD:=mypassword
CREAT_DB_SQL_FILE:=create_db.sql

create_db:
	@ PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -f create_db.sql

delete_db:
	@ PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -f delete_db.sql

create_notify:
	@ PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -f notify_fn.sql

delete_notify_fn:
	@ PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -f delete_notify_fn.sql

delete: delete_db delete_notify_fn
	@ echo "Deleting All"

create: create_db create_notify
	@ echo "Creating All"
