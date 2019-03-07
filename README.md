# jumpcloud-hasher
A simple Golang password hasher

# development notes
- Not implementing a database, as it violates the "use only standard library" instruction, though by allowing a DB driver, this project could implement a simple SQL database (mySQL, Postgres, SQLite, etc)
- Instead of a database using an in-memory hashmap, this will have fast performance, but will not have persistance
- Implemented Delay using standard sleep, but processing using channels and go routines
- Logging middleware is illustrative, but would otherwise adhere to logging standards that already exist
- There was no indication of a return from `/shutdown` endpoint, so I added some basic information
- To adhere to the outline, I returned only the stats requested, though, in production I might add additional stats: pending/processing count (instead of just total, for instance), size of datastore, etc

# installation
```
git clone https://github.com/radean0909/jumpcloud-hasher.git
cd jumpcloud-hasher
docker-compose up
```