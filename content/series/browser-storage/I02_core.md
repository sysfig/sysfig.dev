---
title: Core Flow
slug: core
date: 2023-02-11T17:03:33-08:00
chapter: i
order: 2
draft: true
---

0. Check IndexedDB support
1. Open a connection to a database
2. Create an object store in the database
3. Create a transaction
4. Set up event handlers for that listens for DOM events targetted at the transaction
5. Obtain an object store within the context of a transaction
6. Manipulate data to an object store by making requests within a transaction
7. Set up event handlers on each request
8. Read the results of the requests from the request's event handlers

### Checking for IndexedDB support

According to `caniuse.com`, IndexedDB is fully or partially supported on [98.43%](https://caniuse.com/indexeddb) of all browsers.

Firefox supports IndexedDB, but notably **not** in its Private Browsing mode. If you enter `window.indexedDB` in Firefox's console whilst in Private Browsing mode, you'll get a `null` (this behavior may have changed since the time of this writing - check for yourself!).

So even though most browsers support IndexedDB, if your application depends on IndexedDB for key functionality, it would still be wise to still check for browsers support.

```ts
if (!window.indexedDB) {
  // Perhaps ask the user to use a supported browser
  console.log('IndexedDB is not support on this browser')
}
```

### Opening a connection

Once we've confirmed that IndexedDB is support, the next step is to open a connection to a database by calling `window.indexedDB.open(name[, version])`, passing in the name (required) and version (optional) of the database.

```ts
const openRequest = window.indexedDB.open("name", 1);
```

Each origin can use multiple databases, and the name is a string that uniquely identifies the database. If the database with that name does not exist within the origin, it will be created.

#### Database versions

To understand what the version is, let's imagine you've launched an application that uses IndexedDB to store some tasks, where a typical task object looks like this:

```ts
{
  id: 452,
  body: "Vacuum room"
}
```

People are using your application and they've stored some data into the database that is persisting in their browser. But in the new version of the application, you've:

- changed the name of the `body` property to `name`
- added a `type` property
- created an index on the `body` property so that tasks can be found by specifying the `type`

Without versioning, the user will get the new version of the application when they open the app in a new tab or window, or if they refresh the page. However, existing applications will continue to read and write data to the database using the old schema. This results in some data with one structure and other data with another structure.

With versioning, you can provide a positive, non-zero integer (`unsigned long long`) as the version, and increment the version whenever you need to make an incompatible schema change. Then, when the updated application opens the connection to the database, IndexedDB will see a mismatch between the old database version (used for the existing data), and the new (higher) version (used by the updated application).

If there are open connections from the old version of the application. IndexedDB will emit a `versionchange` DOM event to any `IDBDatabase.onversionchange()` event handler(s) of the **old** application(s), and emit a `block` DOM event to the `IDBDatabase.onblock()` event handler(s) of the **new** application(s).

The `versionchange` change event follows `IDBVersionChangeEvent` interface, an extension of `Event` that contains the `oldVersion` and `newVersion` properties.

The event handlers should contain logic to handle the version mismatch. Nothing can happen with regards to IndexedDB in the new application until it is unblocked by having all open old connections closed. How the application handles a mismatch is application-dependent. Typically, the `onversionchange()` logic may choose to clear existing data, or save any unsaved data, close the connection, and ask the user to refresh or close the page; the `onblock()` may ask the user to close other pages with the application running.

If there are no open connections using the old version, IndexedDB will emit a `upgradeneeded` DOM event to the `IDBDatabase.onupgradeneeded()` event handler of the **new** application. In `onupgradeneeded()`, the application can migrate any data from old versions and make schema changes.

In short, versioning enables you to 'migrate' from an older schema to a newer one.

If you do not specify a version when calling `window.indexedDB.open()`, it will keep the existing version of the database if it exists, or, if the database doesn't exists, create the new database and set its version to `1`.

#### `IDBOpenDBRequest`

As with all operations in IndexedDB, opening a database connection is asynchronous, and the `open()` call doesn't return with the connection. Instead it returns with a `IDBOpenDBRequest` object immediately, and opens the connection asynchronously in the background.

> `IDBOpenDBRequest` is a sub-class of `IDBRequest` object. Most asynchronous IndexedDB operations return with some form of `IDBRequest` object.

Once the connection request has complete, the relevant DOM event would be fired with the `IDBOpenDBRequest` object as its `target`:

- `error` - occurs when there's an error opening the database connection. It could be:
  - a user's browser setting disables the use of IndexedDB
  - the browser runs out of memory
  - the browser does not support IndexedDB
- `blocked` - occurs when an application calls `open()` on a database with a greater version than the current version **and there are other open connections to the same database elsewhere** (e.g. on another tab), then the new connection will be blocked until all the existing open connections have closed.
- `versionchange` - occurs when another application on another tab/window is trying to open a connection to the database with a higher version than the current application
- `upgradeneeded` - occurs when an application is trying to open a new database connection with a higher version than the existing version **and there are no existing connections using the old version**.
- `success` - the connection opened successfully

These events can be listened by defining event handlers (`(event) => {}`) on the `IDBOpenDBRequest` object:

- `onerror` - the error code will be made available via the `errorCode` property on the `IDBOpenDBRequest` object.
- `onsuccess` - a `IDBDatabase` object representing the connection will be made available via the `result` property on the `IDBOpenDBRequest` object.
- `onblocked`
- `onversionchange`
- `onupgradeneeded` - this is the only place where you can alter the structure of the database, by adding and deleting object stores and indexes

```ts
let db;
const openRequest = window.indexedDB.open("name", 1);
openRequest.onerror = (event) => {
  console.log(event);
};
openRequest.onblocked = (event) => {
  console.log('Please close existing instances of this application')
};
openRequest.onversionchange = (event) => {
  console.log('Please close this application')
};
openRequest.onupgradeneeded = (event) => {
  // If new database is created -> create the database schema - object stores and indexes
  const store = db.createObjectStore(...);

  // If database version is incremented -> migrate old data to comply with new schema. This may involve creating new object stores and deleting old ones
};
openRequest.onsuccess = (event) => {
  db = openRequest.result; // or event.target.result
};
```

Ultimately, the goal of opening a database connection is to obtain an `IDBDatabase` object, representing the database.

### Creating a Store

To store any data, we must create an object store. We can do this by calling `IDBDatabase.createObjectStore(name, options)` within the `onupgradeneeded` handler (this is the only place we can create object stores and indexes). The `name` identifies the object store within the database, and the `options` object specifies how the key should be determined. The `options` object can have any of the following properties:

- `keyPath` (_`string`_) - tells IndexedDB to uses one of the properties of the stored value as the key. These keys are called _in-line keys_ because the key is within the stored value. `keyPath` and in-line keys only works if the stored value is a JavaScript object.
- `autoincrement` (_`boolean`_) - tells IndexedDB to use an auto-incrementing number as the key. The keys for the object store would starts at `1` and increments by 1 for each new record. The job of generating and incrementing they key falls to a _key generator_, but that is a browser implementation details which we don't need to deal with.

If you don't specify an option, then you must supply the key each time you create a record.

Interestingly, you can also specify **both** `autoincrement` and `keyPath`; in this case, the key is generated using the key generator and then stored in-line into the stored value (which must be a JavaScript object) at the specified key path. If a value already exists at the key path, that value is used as the key instead of the auto-generated value.

```ts
const store = db.createObjectStore("name"});
// ...or...
const store = db.createObjectStore("name", { keyPath: "foo.bar" });
// ...or...
const store = db.createObjectStore("name", { autoincrement: true });
// ...or...
const store = db.createObjectStore("name", { autoincrement: true, keyPath: "foo.bar" });
```

> In some implementations it is possible for the implementation to run into problems after queuing a task to create the object store after the createObjectStore() method has returned. For example in implementations where metadata about the newly created object store is inserted into the database asynchronously, or where the implementation might need to ask the user for permission for quota reasons. Such implementations must still create and return an IDBObjectStore object, and once the implementation determines that creating the object store has failed, it must abort the transaction using the steps to abort a transaction using the appropriate error. ~ [`createObjectStore(name, options)` in the Indexed Database API 3.0 W3C specification](https://w3c.github.io/IndexedDB/#dom-idbdatabase-createobjectstore)

### Creating a Transaction

`IDBDatabase.transaction`

`db.transaction(scope[, mode[, options]])`

`scope` is an array of the names of object stores that you want to access

Modes:

- `readonly` (default) - allows you read from an existing object store
- `readwrite` - allows you make changes to an existing object store

Options

- `durability` -  https://w3c.github.io/IndexedDB/#transaction-durability-hint

`readwrite` transactions places a lock on the object store. Thus, only one `readwrite` transaction can run on a particular object store at any one time. For transactions in `readwrite` mode, it's best practice to minimize the scope as much as possible to ensure `readwrite` transactions with non-overlapping scopes can run concurrently.

`versionchange` - allows you to change the schema - adding and deleting object stores and indexes (can't be specified in the `IDBDatabase.transaction` call)

### Setting up DOM event handlers

Results of requests of a transaction can have one of 3 outcomes, each triggering a DOM event and can be handled with a handler:

- `error` / `onerror` - the default behavior is to abort the transaction (thus triggering the `abort` event). You can stop by calling `stopPropagation()`
- `abort` / `onabort` - causes any changes in the transaction to be rolled back. Can be triggered by manually calling `abort()`, or if an `error` event is not handled.
- `complete` / `oncomplete` - all pending requests have completed

```ts
transaction.oncomplete = (event) => {};
transaction.onerror = (event) => {};
transaction.onabort = (event) => {};
```

### Manipulating Data

You must first obtain the object store from the transaction

```ts
store.add({ username:'none', password:'none'}, 1);
```


```ts
let db;
const request = window.indexedDB.open("name", 1);
request.onerror = (event) => { ... };
request.onsuccess = (event) => { ... };
request.onupgradeneeded = (event) => {
  const productStore = db.createObjectStore("products");
  const userStore = db.createObjectStore("users");
  store.transaction.oncomplete = (event) => {
    const transaction = db.transaction(["products", "users"], "readwrite")

    // Note that whilst both `txUsesrStore` and `userStore` are `IDBObjectStore` objects
    // `txUsesrStore` has an additional read-only `transaction` property set to a `IDBTransaction`
    // indicating that this `IDBObjectStore` object can only 'work' within the scope of the transaction
    const txUserStore = transaction.objectStore("users");

    // Manipulate data in the store(s)
    txUserStore.add( ... )
  };
};
```

Like every IndexedDB operation, adding data is asynchronous. Thus, to get the results of a request, we must set up event handlers.

- `error` / `onerror`
- `success` / `onsuccess`

Different types of requests yields different results:

- `add()` - adds a value, and requiring that no other objects in the object store has the same key. Returns the key of the stored value
- `put()` - adds or updates (i.e. upsert) a value, overriding an existing value with the same key if it exists.
- `delete(key)` - deletes a value by key
- `get(key)` - get a specifiy value by key
- `getAll()` - get all values in the object store. The result is an array of all values (not keys).
- `getAllKeys()` - get all keys in the object store. The result is an array of all keys (no values)
- `openCursor()` - get a cursor to iterate through multiple values in the object store. This can be all the values, or a range of values as constrained by a _key range_. The cursor object contains the `key` and `value` properties representing the key and value of the record pointed to by the cursor. The cursor also has a `continue()` method, which you can call to request the next record; this will trigger another `success` event on the event and the next result can similarly be handled with the `onsuccess` handler.

  If no values matches the key range, or if the cursor has reached the end of the range, then `cursor` will be undefined.

```ts
const request = objectStore.add(customer);
request.onerror = (event) => {
  // Handle errors!
};
request.onsuccess = (event) => {
  // Use `request.result`, or `event.target.result` if request is not available (if you didn't save `request` into a variable and chained it)
  // event.target.result === customer.ssn;
};
```

#### Adding Data


#### Getting Data

##### Getting a single value

##### Iterating through multiple values

https://developer.mozilla.org/en-US/docs/Web/API/IndexedDB_API/Using_IndexedDB#specifying_the_range_and_direction_of_cursors

#### Updating Data

```ts
const objectStore = db.transaction(["customers"], "readwrite").objectStore("customers");
const request = objectStore.get("444-44-4444");
request.onerror = (event) => {
  // Handle errors!
};
request.onsuccess = (event) => {
  // Get the old value that we want to update
  const data = event.target.result;

  // update the value(s) in the object that you want to change
  data.age = 42;

  // Put this updated object back into the database.
  const requestUpdate = objectStore.put(data);
  requestUpdate.onerror = (event) => {
    // Do something with the error
  };
  requestUpdate.onsuccess = (event) => {
    // Success - the data is updated!
  };
};
```

#### Deleting Data


## Non-Core Flows

### Adding an Index

To specify the property to use for the index, you specify a key path.

Within `openRequest.onupgradeneeded`

```ts
const index = objectStore.createIndex("name", "name");
```

Then, anywhere you want to use the index

```ts
const index = objectStore.index("name")
```

Then, you can use all the methods you could have with an object store like `get()` and `delete()`, but the `name` value will be used instead of the key.

```ts
// Using a normal cursor to grab whole customer record objects
index.openCursor().onsuccess = (event) => {
  const cursor = event.target.result;
  if (cursor) {
    // cursor.key is a name, like "Bill", and cursor.value is the whole object.
    console.log(`Name: ${cursor.key}, SSN: ${cursor.value.ssn}, email: ${cursor.value.email}`);
    cursor.continue();
  }
};
```

Indexes have an additional `openKeyCursor()` method, which returns a _key cursor_ with two properties:

- `primaryKey` - the key of the object
- `key` - the property used for the index

The value is not available with the key cursor.

```ts
// Using a key cursor to grab customer record object keys
index.openKeyCursor().onsuccess = (event) => {
  const cursor = event.target.result;
  if (cursor) {
    // cursor.key is a name, like "Bill", and cursor.value is the SSN.
    // No way to directly get the rest of the stored object.
    console.log(`Name: ${cursor.key}, SSN: ${cursor.primaryKey}`);
    cursor.continue();
  }
};
```

### Updating Schema

Call the `IDBFactory.open` method with a newer version specified. This will trigger `request.onupgradeneeded`, and in there, you can make changes to the schema.


#### Handling multiple connections

`db` object
- `onblocked`
- `onversionchange` - another application have requested a version change on the database. Typical behavior is to close the database connection and prompt the user to reload or close the tab.

### Closing a database connection

`db.close()` or when th tab or browser if closed

Triggers `db.onclose` when 

### Deleting a Database

`IDBFactory.deleteDatabase`

## Libraries

### `idb`

## TODOs

- https://web.dev/indexeddb/
- https://w3c.github.io/IndexedDB/
- Synchronizing IndexedDB based on UI changes - https://web.dev/indexeddb-uidatabinding/
- Synchronizing between local and remote database
- https://web.dev/indexeddb-best-practices/
- https://javascript.info/microtask-queue
- https://rxdb.info/slow-indexeddb.html
- https://www.codemag.com/article/1411041/Introduction-to-IndexedDB-The-In-Browser-Database
