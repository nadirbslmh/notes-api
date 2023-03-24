## notes-api

REST API application to manage notes. Implemented with clean architecture

## How to use

1. Clone this repository.

2. Copy the `.env` file.

```sh
cp .env.example .env
```

3. Fill the values inside the `.env` file for the database configurations.

4. Create a new database called `notes_api`.

```sql
CREATE DATABASE notes_api;
```

5. Run the application with this command:

```sh
go run main.go
```

## Additional Notes

There are two branches in this repository:

- `with-db`: Includes clean architecture implementation with Database. Authentication feature is not provided.
- `main`: Will includes clean architecture complete implementation.

## Resources

- [Graceful Shutdown Implementation](https://medium.com/tokopedia-engineering/gracefully-shutdown-your-go-application-9e7d5c73b5ac).
