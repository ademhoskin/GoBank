# GoBank!
Bank System using Go and Postgres

# To start DB
## Must have make, migrate (go) AND Docker Installed
- "make postgres" starts postgres16 alpine container
- "make createdb" creates bank database
- "make dropdb" drops bank database if needed
-  'make migrateup' to migrate bank db up, 'migratedown' to migrate down

#Tests
- "make tests" will run all tests for project
  
