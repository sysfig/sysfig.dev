---
title: Core Concepts
slug: concepts
date: 2023-02-11T17:03:33-08:00
chapter: i
order: 1
---

## Origins, Databases, Object Stores, Values, Keys, and Indexes

Each _origin_ (e.g. `https://example.com`) can have its own set of IndexedDB _databases_. Each database is identified with a name that is unique for that origin. Within each database is one of more _object stores_, which is where you store the actual data. Each object store must have a name that unique within the database.

> You can think of an object store akin to a 'table' in a relational database, or 'collections' in MongoDB.
> You can design your application so that different types of data are stored in separate databases, each with a single store. Or you can store different types of data in different object stores within a single database. Or something in between. Typically, you'd create a database for each application that falls within your origin. For example, `https://example.com/app` would have a different database to `https://example.com/blog`. Then, within each application, you'd create separate object stores each data type (e.g. `users`, `posts`, `categories`).

Data is stored in object stores as key-value pairs. Despite the name 'object store', the stored values can be primitive values (e.g. strings, numbers, booleans, `undefined`, and `null`), objects (e.g. `Date`, `RegExp`, `File`, and `ArrayBuffer`), as well as arrays where the elements of one of the supported types.

When you store a value, in addition to specifying the object store, you must also assign the object a _key_. The key is a data value that is unique within the object store, and by which stored values are sorted and referenced.

IndexedDB will automatically sort and index an object using its key. This index is what allows your application to search for a value in the object store by key.

Additional indexes can be manually created on any object stores that hold JavaScript objects as its stored values. Similar to the index for keys, the keys of these indexes will also be automatically sorted. These additional indexes allows you to search stored values based on any specified property of the stored object other than the key. For example, if you have an object store of users with properties `id`, `firstname`, `lastname`, and `id` is the key, you can create additional indexes on `firstname` and `lastname` to gain the ability to search the object store using the user's first name or last name.

You can also use indexes to enforce constraints. A common constraint is the unique constraints, which ensure a specific property in this object store is unique amongst all stored values. For example, if you have an object store of users with properties `id`, `email`, `firstname`, and `lastname`, you can create an index with a unique constraint to ensure all users in the object store have unique email addresses.

Collectively, the sum of all object stores and their indexes are referred to as a database's _schema_.

## Connections, Requests, DOM Events, and Handlers

The first step to working with IndexedDB is to open a _connection_ to a database. Multiple connections can be opened to the same database at the same time; this allows multiple instances of your application, on multiple tabs and/or windows, to operate on the same database concurrently.

Once a connection to a database is opened, operations, such as creating/deleting object stores and reading/writing data, are sent to the database as asynchronous _requests_. Requests are asynchronous, meaning you don't get the results of the operation straight away. Instead, you are sending the database a request and it will notify you, via _DOM `Events`_, once that request is complete. Asynchronous requests ensures IndexedDB does not block the main thread whilst data is being read or written. To make use of the result of a request, you can attach _event handlers_ to the request.

## Transactions

All requests execute within the context of a _transaction_. Transactions ensures:

- atomicity - that either all requests are completed successful, or none of the requests are complete (completed requests are rolebacked), and
- multiple instances of an application (from two different tabs or windows) do not interfere with each other's modifications

Transactions 'belong' to the database connection and can operate on multiple object stores. When creating a transaction, you must specify its _scope_, which is the set of object stores that this transaction can interact with. You must also specify a _mode_, which specifies what actions the transaction can do on those object stores - whether the transaction can write to the object store(s) or only read from them.

A transaction has a limited lifetime. You can keep a transaction active by sending requests. If a transaction is not in use by the time the execution flow returns to the main event loop, it becomes inactive. To keep a transaction active, make sure it has at least one pending (incomplete) request. Trying to use an inactive transaction will throw an exception with the `TRANSACTION_INACTIVE_ERR` error code.

When all requests belonging to a transaction complete, the transaction is auto-committed. This commitment is done automatically; you cannot manually commit a transaction.

## Results and Cursors

IndexedDB provides the `add()`, `get()`, `put()`, `delete()` methods on `IDBIndex.objectStore` that enables you to create, read, update, and delete individual records within a store. But the IndexedDB API also provides you with a way to read multiple records using a _cursor_.

A cursor points to one particular record at a time, allowing you to read its value. But it also has a `continue()` method, that allows you to 'continue' to the next record and iterate through the results.
