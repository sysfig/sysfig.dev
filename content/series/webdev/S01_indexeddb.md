IndexedDB is an in-browser database that can be used to persist data for offline uses. IndexedDB is a type of database called an object store; in particular, IndexedDB stores JavaScript objects against a unique key.

## IndexedDB vs. X

Unlike the Web Storage APIs, IndexedDB is asynchronous, which means it does not block the main thread.

## Basic CRUD

Not relational, and you do not query it with SQL. There are tools? that builds on top of IndexedDB that allows you to do that.

IndexedDB Shim
IndexedDB Promised
[localforage](https://github.com/localForage/localForage) - a wrapper around asynchronous client-side storage APIs (including IndexedDB and WebSQL) but provides a simpler localStorage-like API
Dexie.js
Taffydb

## Transactions

## Object Stores

An IndexedDB database can store data directly, but you can also create _object stores_ and add data with the same core format inside those object stores. Object stores allows you to:

- group related data together, similar to tables in a relational database
- create indexes that allows you to find objects quickly based on a key that is not the primary key

