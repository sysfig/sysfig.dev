---
title: Views
slug: views
date: 2022-04-30T12:11:00+01:00
chapter: z
order: 99
tags:
  - postgres
draft: true
---


There are two places you can specify which queries are to be run on the database. The first is in the client code using a client library such as C's [`libpq`](https://www.postgresql.org/docs/current/libpq.html). Another is in the database itself in the form of _views_.

Views are simply named SQL `SELECT` queries. You can refer to the view (the output of the `SELECT` query) as a normal table, which means you can even have views that builds on top of other views.

Although the concept of views are quite simple, it has many uses:

- convenience - views stop you from typing out the same query (often complex) every time
- encapsulation - allows you to change the underlying structure of your table (add columns, change data type) without changing the view (in the case of changing data type, you can use [`CAST`](https://www.postgresql.org/docs/current/sql-createcast.html) to convert the table's column data type to a compatible type). This allows database administrators to optimize the data structure but also give time to developers to update their code.
- DRY - if there are multiple applications (perhaps written in different languages that can't share common code) that reads the same set of data from the same database table, instead of defining the query in the application code multiple times, it may be DRYer to define a view.
- security - instead of granting users permissions for a table (which may include certain rows and/or columns you don't want certain users to see), you can create a more restrictive view and grant permissions only for that view
- performance - instead of allowing application developers to send arbitrary (non-performant) queries to your database, you provide a view which you can optimize.

```sql
CREATE VIEW <view_name> AS <sql_select_query>;
```

## updatable views

## Why Not Views

View joins?
derived table?
table-valued function == parametized view
**stored procedures

What about updates/deletes

---

Adding indexes to views
