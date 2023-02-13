---
title: Web Storage API
slug: web-storage-api
date: 2023-02-11T17:20:43-08:00
chapter: w
order: 0
---

The _[Web Storage API](https://html.spec.whatwg.org/multipage/webstorage.html)_ is a web API, maintained by [https://whatwg.org/](WHATWG), for storing string-only key-value pairs locally in the browser.

The Web Storage API provides two ways to store data - `localStorage` and `sessionStorage`. Both using the same API, but differs in how long the data persists for.

Data stored in `sessionStorage` persists only for that session, which is tied to a particular tab or window. If you open the same page on a new tab or window, that new tab/window is a new and completely different session, and have its own `sessionStorage` storage. Users can refresh the page without losing data, but when the tab/window is closed, the data is cleared.

On the other hand, data stored in `localStorage` persists until the user manually clears it via their browser settings, or until your application clears it.

In short, `sessionStorage` is meant for storage that spans a single window/tab, and `localStorage` is meant for storage that spans multiple windows/tabs

## Check for support

According to [caniuse.com](https://caniuse.com/namevalue-storage), the Web Storage API is supported in 98.86% of browsers in use. However, some browsers (notably Opera Mini) does not support it, and many browsers disables it in incognito or private browsing modes.

Therefore, it's advisable to check for support by checking for the existence of the `localStorage` or `sessionStorage` object(s) under `window`.

```ts
function localStorageIsAvailable() {
  if (!'localStorage' in window) {
    return false;
  }
  try {
    const testString = ')*_()_a_random_string_(*&%GUU)'
    storage.setItem(testString, testString);
    storage.removeItem(testString);
    return true
  } catch (e) {
    return false
  }
}

if (!localStorageIsAvailable()) {
  // Unsupported browser, ask user to use a supported browser
}
```

It's not enough to check for the existence of `window.localStorage` since some browsers implements `localStorage` but allocates it no storage quota. Therefore, the surest way to ensure `localStorage` is available is to **use** it, by setting and deleting values.

Once you've confirmed that `localStorage` or `sessionStorage` is available, we are ready to start adding, reading, updating, and deleting data. Unlike other in-browser storage solutions like IndexedDB, there's no set-up required - you can just start using `localStorage`/`sessionStorage`.

## CRUD

The `localStorage`/`sessionStorage` objects follows the [`Storage`](https://html.spec.whatwg.org/multipage/webstorage.html#the-storage-interface) interface, which defines the following methods:

- `setItem(key, val)` - sets a key-value pair, replacing the existing record if it exists. Can be used to create a new record or to update an existing one.
- `getItem(key)` - used to get the value of a single record. If there are no records with that key, `null` is returned.
- `removeItem(key)` - removes a key-value pair by key
- `clear()` - clears all records
- `key(n)` - returns the key of the `n`th record, or `null` if `n` is equal or higher than the number of records. The iteration is not defined, and can change on most mutations. You should not rely on a record keeping the same index across mutations.

Apart from using these methods, you can also treat `localStorage` as a dictionary (a string-only object) and manipulate its properties by accessing it directly. However, the return values may be different.

```ts
const key = 'foo'
localStorage[key] = 'bar'   // => 'bar'
localStorage[key]           // => 'bar'
delete localStorage[key]    // => `true`
localStorage[key]           // => `undefined` (not `null`)
```

The `Storage` interface specifies an additional read-only `length` property, which tells you the number of key-value pairs in `localStorage`/`sessionStorage`.

## Iterating through all records

There are no native way to iterate over all items in `localStorage`/`sessionStorage`, but you can do it with a `for` or `while` loop and using the `key(n)` method.

For example, we can create a 'snapshot' of `localStorage` using the following `for` loop:

```ts
const snapshot = {}
for (let i = 0; i < localStorage.length; i++){
  const key = localStorage.key(i);
  const value = localStorage.getItem(key);
  snapshot[key] = value
}
```

And here's the same functionality, but using a `while` loop:

```ts
const snapshot = {}
let key;
let i = 0;
while (key = localStorage.key(i)) {
  const value = localStorage.getItem(key);
  snapshot[key] = value
  i++
}
```

## Events

Whenever a change is made to `localStorage`, the browser engine will fire a `storage` DOM event at `window` of **other** pages opened at the same origin, which other pages can listen to by adding an event handler with `window.addEventListener()`.

The `storage` event abides by the [`StorageEvent`](https://html.spec.whatwg.org/multipage/webstorage.html#storageevent) interface, making the following readonly properties available:

- `key` - the key of the record that has changed, or `null` if the event is caused by the `clear()` method
- `oldValue` - the value of the record before the change, which is `null` if the value is being created
- `newValue` - the value of the record after the change, which is `null` if the value is being deleted
- `url` - the URL of the document whose the storage item changed
- `storageArea` - the `localStorage` or `sessionStorage` object

> The `StorageEvent` interface is an extension to the [`Event`](https://dom.spec.whatwg.org/#interface-event) interface, so all the properties and methods available in `Event` is also available in `StorageEvent`.

A common use case for listening to this event is to synchronize data between multiple tabs/windows running the same application.

> Events are **not** fired for `sessionStorage`, since different tabs/windows are considered different sessions.

## Errors

Most browsers set a limit to the amount of data that each origin can store in `localStorage`/`sessionStorage`. If your application attempts to run `setItem(key, val)` that would cause this limit to be exceeded, the `QuotaExceededError` `DOMException` will be thrown (this error is called `NS_ERROR_DOM_QUOTA_REACHED` on Firefox). This error will also be thrown if the user has disabled storage for the origin.

A `SecurityError` `DOMException` can also be thrown if the browser prohibits the use of Web Storage API.
