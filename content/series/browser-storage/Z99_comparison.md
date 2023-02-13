https://web.dev/storage-for-the-web/

Cookies pros:

- Convenience - values are automatically sent to the server on every request

Cookies cons:

- Values are automatically sent to the server on every request

Web API storage Pros:

- Can store a relatively large amount of data (5MB per domain in Chrome, 3MB in Opera)

Web API storage Cons:

- Synchronous
- Can only store strings; other data types must be serialized and deserialized to another format. For example, you may be able to store objects as JSON using `JSON.stringify()` to serialize and `JSON.parse()` to deserialize.

`sessionStorage` pros:

- Allows you to isolate data belonging to the same site but in different sessions. For example, it allows you open one tab to search for flights on a certain date, and open another tab and search for flights on a different date, and have and filters or sort parameters persisted across refreshes.

IndexedDB pros:

- Provides built-in schema versioning, giving you the opportunity to migrate existing data from the old version to a new one.
- Asynchronous API - will not block the main thread

IndexedDB cons:

- low-level, more complicated/verbose than, say, Web Storage API. You'd have to set up databases, object stores, and indexes; with `localStorage` and `sessionStorage`, these objects are available directly under `window`, and there's no set up needed to start using it.
- You can't use IndexedDB in Firefox Private Browsing Mode
  - https://bugzilla.mozilla.org/show_bug.cgi?id=781982
  - https://bugzilla.mozilla.org/show_bug.cgi?id=1639542
---



Common uses

Offline
Caching

- Web Storage
  - Persisting user input - so if their browser crash or they refresh or close the tab by mistake, their inputs are not lost.
