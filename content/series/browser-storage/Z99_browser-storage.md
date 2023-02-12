---
title: Browser Storage
slug: browser-storage
date: 2020-12-15T16:17:54-08:00
chapter: z
order: 99
draft: true
---


When building any non-trivial browser-based web application, there's usually a need to persist data across page refreshes. You may want to persist data to:

- Provide a better user experience (you don't want the user to log in after every page now, do you?)
- Provide better performance (if you cache a copy of some data you got before, you can serve these immediately, whilst you send a new network request to get the new data, if it's different; this strategy is known as stale-while-revalidate)
- Provide offline support (for example, a writing app that saves your work locally, allow you to work on a plane, and syncs it with the server once connection is re-established.)

- Cache Storage API - file-based contents such as HTML, CSS, JavaScript files. ??Not data??
- Web Storage API
  - synchronous (blocks the main thread)
  - Limit of 5MB
  - Can only contain strings
  - Web workers and service workers cannot access it
  - Two types:
    - LocalStorage:
    - SessionStorage - Same as LocalStorage, the only difference being that the life of the data is tied to the life of the tab
- Cookies
  - synchronous (blocks the main thread)
  - Can only contain strings
  - Sent with every HTTP request
  - Web workers cannot access it
- Indexed DB
  - asynchronous (does not block the main thread)
  - Web workers and service workers can access it
- File and Directory Entries API
  - asynchronous (does not block the main thread)
  - Allows you to read and write to files in a sandboxed filesystem
  - Only available on Chromium-based browsers
- File System Access API
  - Allows you to read and write to files in the user's 'normal' filesystem (as opposed to a sandboxed filesystem)
  - Requires permissions from the user
- WebSQL
  - Should use IndexedDB instead
- Application Cache
  - Should use service workers and the Cache API instead

Use IndexedDB with a library https://www.npmjs.com/package/idb

You can use the StorageManager API to determine how much storage quota your browser can allow. Note that this API is not supported on Safari or Internet Explorer.


https://web.dev/persistent-storage/ - you don't want to lose your drafts

"""Starting in iOS and iPadOS 13.4 and Safari 13.1 on macOS, there is a seven-day cap on all script writable storage, including IndexedDB, service worker registration, and the Cache API. This means Safari will evict all content from the cache after seven days of Safari use if the user does not interact with the site. This eviction policy does not apply to installed PWAs that have been added to the home screen. See Full Third-Party Cookie Blocking and More on the WebKit blog for complete details."""
