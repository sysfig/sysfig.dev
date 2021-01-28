PostgreSQL is an open-source, object-relational database management system (ORDBMS) based on [POSTGRES](https://dsf.berkeley.edu/postgres.html). The POSTGRES project was developed by a team at the University of California at Berkeley Computer Science Department, led by Professor Michael Stonebraker. The project lasted from 1986 to 1994, and its last release was v4.2.

In 1994, Andrew Yu and Jolly Chen rewrote POSTGRES in ANSI C and replaced the original query interpreter (PostQUEL) with a SQL interpreter. This became Postgres95 and later PostgreSQL (or simply 'Postgres')

Postgres runs as a _master_ backend (server-side) process (usually the program is named `postgres`) that listens at a specific port (defaults to `5432`) for TCP/IP network connections from clients. After accepting and establishing a connection (a.k.a. _session_), the client can send database operations (i.e. SQL queries) to the master process, who will parse the query, check for permissions and, if allowed, generates an _execution plan_, execute the plan on the client's behalf, and return the results.

Whenever the server accepts a new TCP connection, it spawns/forks a new process. This means there's a process for the master process and separate processes for each connection. Therefore, if your database is being used by 50 different clients, then at least 51 processes would be running on the web server.

A single Postgres database server can manage many databases. A collection of databases managed by a single PostgreSQL server instance can be referred to as a database _cluster_. Typically, you'd use a separate database for each project. Within a database, you may create different _tables_ (or _relations_ in formal SQL speak) that hold different kinds of data (e.g. you may have `users` and `products` table). A table is made up of columns of a specific data type (e.g. numeric type) and rows (where each row is a distinct entry in the table)

## Users

When you connect to a Postgres server, you choose what Postgres user to authenticate as.

Note that PostgreSQL user accounts are completely independent from the (operating) system's users accounts. However, many tools default to using your system's username when authenticating with the Postgres server, so you may want to (for convenience) set your Postgres username to be the same as your system's username.

The user who created the database is a _superuser_, which is not subject to access controls.

## Databases

VERIFY?? There's always a database with the same name as the name of the _system_'s user that started the master process. Typically, this is a user named `postgres`, and so when you create a new database, you'd often find the `postgres` database already on the server.??VERIFY

Data are stored as files

## SQL

- SQL commands end with a semi-colon (`;`).
- Whitespaces like the number of spaces and indentation carries no semantic meaning
- `--` can be used for comments (everything from the right of `--` to the next newline character are not evaluated)
- Keywords and identifiers are case-insensitive, except when they are double-quoted

```sql
CREATE TABLE <table_name> (
  <column_1_name>  <column_1_type>, -- comments
  <column_2_name>  <column_2_type>
);
```

```sql
DROP TABLE <table_name>;
```

The order of the columns matter because SQL supports the syntax of inserting rows where the values are implicitly mapped to the column in order.

```sql
INSERT INTO weather VALUES ('San Francisco', 46, 50, 0.25, '1994-11-27');
```

You can also explicitly map the columns to the values.

```sql
INSERT INTO weather (city, temp_lo, temp_hi, prcp, date)
    VALUES ('San Francisco', 43, 57, 0.0, '1994-11-29');
```

Values that are not simple numeric types are quoted using single quotes (`'`)


- [`COPY`](https://www.postgresql.org/docs/current/sql-copy.html) - copy data into a table from text files
- `SELECT` - retrieve information from the database

### `SELECT`

```sql
SELECT <columns> FROM <tables> <qualifications>
```

`<columns>` are the columns to return. `<tables>` are the tables where these columns belongs to. `<qualifications>` are an optional set of rules which can restrict the entries returned, or alter the way the results are presented.

Another way to look at it is that `<columns>` determines which columns are returned, and `<qualifications>` determines which and how rows are returned.

`SELECT DISTINCT` - remove duplicates from the results

#### Qualifications

- `WHERE`
- `ORDER BY`
