# jumpcloud-hasher
A simple Golang password hasher

# Development notes
- Not implementing a database, as it violates the "use only standard library" instruction, though by allowing a DB driver, this project could implement a simple SQL database (mySQL, Postgres, SQLite, etc)
- Instead of a database using an in-memory hashmap, this will have fast performance, but will not have persistance
- Implemented Delay using standard sleep, but processing using channels and go routines
- Logging middleware is illustrative, but would otherwise adhere to logging standards that already exist
- There was no indication of a return from `/shutdown` endpoint, so I added some basic information
- To adhere to the outline, I returned only the stats requested, though, in production I might add additional stats: pending/processing count (instead of just total, for instance), size of datastore, etc
- HTTP rest testing using only the standard library was becoming a significant time sink, However, unit tests have been included to illustrate a commitment to testing. In a production environment I would likely use something like gorilla to help with testing.

# Installation
```
git clone https://github.com/radean0909/jumpcloud-hasher.git
cd jumpcloud-hasher
docker-compose up
```

# Testing
```
docker-compose run app go test -v ./utils/database
```

# Endpoints
### Request: POST /hash
`curl -i -d "password=angryMonkey" http://localhost:8080/hash`

**Response:** 
```
HTTP/1.1 202 Accepted
Date: Thu, 07 Mar 2019 20:41:14 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8

2
```

### Request: GET /hash/1
`curl -i http://localhost:8080/hash/1`

**Response:**
```
HTTP/1.1 200 OK
Date: Thu, 07 Mar 2019 20:41:55 GMT
Content-Length: 89
Content-Type: text/plain; charset=utf-8

ZEHhWB65gUlzdVwtDQArEyx-KVLzp_aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A-gf7Q==
```

### Request: GET /stats
`curl -i http://localhost:8080/stats`

**Response:**
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 07 Mar 2019 20:42:42 GMT
Content-Length: 25

{"total":2,"average":58}
```

### Request: GET /shutdown
`curl -i http://localhost:8080/shutdown`

**Response:**
```
HTTP/1.1 202 Accepted
Content-Type: application/json
Date: Thu, 07 Mar 2019 20:43:37 GMT
Content-Length: 32

{"pending":0,"estimatedTime":0}
```
