# Chat Server

Chat using websockets and concurrency by channels.

### Setup

1. Deploy a Postgres container. 
```
docker run --name some-postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres
```
2. Create a database e.g. chat_experiment
3. Add .env file in root project
```
SIGNING_KEY=custom_key
DB_HOST=localhost
DB_PORT=5432
DB_NAME=chat_experiment
DB_USER=postgres
DB_PASSWORD=password
```
4. Run the project.
