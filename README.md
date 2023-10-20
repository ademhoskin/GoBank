# GoBank!
Bank System using Go and Postgres

# To start DB
## Must have make, migrate (go) AND Docker Installed
- "make postgres" starts postgres16 alpine container
- "make createdb" creates bank database
- "make dropdb" drops bank database if needed
-  'migrate -path db/migration -database "postgres://root:secret@localhost:5432/bank_db?sslmode=disable" -verbose up' will migrate our schema to our bank database
  
