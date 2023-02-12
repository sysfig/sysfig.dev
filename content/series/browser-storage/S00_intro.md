---
title: Introduction
slug: intro
date: 2023-02-09T16:53:06-08:00
chapter: s
order: 00
draft: true
---

Modern web browsers offers many ways to persist state (i.e. store data) in the client:

- Query string
- Fragment
- Cache storage
- Cookies
- IndexedDB
- Web Storage API, of which there are two:
  - Local storage (`localStorage`)
  - Session storage (`sessionStorage`)
- File and Directory Entries API - allows your application to ask the browser to create a sandboxed file system just for your application
- File System Access API - allows your application to read, write, and manage files on a user's local device, or on a user-accessible network file system.

There are also dedicated JavaScript libraries that implements its own in-memory database, or is a wrapper that provides a layer of abstraction on the native API(s):

- [localforage](https://github.com/localForage/localForage) - a wrapper around asynchronous client-side storage APIs (including IndexedDB and WebSQL) but provides a simpler localStorage-like API. It tries to use IndexedDB by default, but falls back to localStorage for browsers that don't support IndexedDB
- [Apache CouchDB](https://couchdb.apache.org/)
- [PouchDB](https://pouchdb.com/) - a client-side implementation of CouchDB that uses IndexedDB
- [minimongo](https://github.com/mWater/minimongo) - an in-memory MongoDB-compatible storage, backed by localStorage. Can synchronize with a remote MongoDB instance.
- [TaffyDB](https://taffydb.com/)
- [RxDB](https://rxdb.info/) - a NoSQL client side database that can be used on top of IndexedDB. Supports indexes, compression and replication. Also adds cross tab functionality and observability to IndexedDB.
- [`idb`](https://github.com/jakearchibald/idb) (previously called 'IndexedDB Promised') - a thin wrapper around IndexedDB that provides usability improvements such as promises.
- [`idb-keyval`](https://github.com/jakearchibald/idb-keyval) - a key-value store backed by IndexedDB
- [Dexie.js](https://dexie.org) - a wrapper around IndexedDB
- [JsStore](https://jsstore.net/) - a wrapper around IndexedDB that you can query with SQL-like syntax

Some APIs have since been deprecated:

- `userData`
- `globalStorage`
- Web SQL Database API
- Flash Cookies, a.k.a. Locally Shared Objects (LSO)
- Silverlight's Isolated Storage

They differ in:

- how long they persist the state for - across reloads? What if the user close and reopens their browser?
- What kind of data it supports - simple strings? entire files?

Client-side storage allows applications to:

- store data offline
- persist data through refreshes
- persist data through closing and reopening the browser
- allows the application to continue to function with intermittent Internet connections
