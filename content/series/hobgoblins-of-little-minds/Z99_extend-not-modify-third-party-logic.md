## Extend, Not Modify

Having recently joined a new company, I was in the process of acquainting myself with the new codebase. Whilst studying the Express-based API, I noticed that all of our endpoints are using [`res.jsonp()`](https://expressjs.com/en/api.html#res.jsonp) to send back JSON data instead of [`res.json()`](https://expressjs.com/en/api.html#res.json).

_JSONP_ (short for _JSON with padding_) is a legacy method of circumventing the _same-origin policy_ (SOP) imposed by brwosers. JSONP was popular from around 2005 to 2014, when the _cross-origin resource sharing_ (_CORS_) specification became a W3C recommendation, and became the more secure way to circumvent SOP.

It was 2022, so there's absolutely no reason why our API should be using JSONP. What's more confusing, the front-end application that interfaces with the API is not using JSONP (there are no callback functions defined).

Not understanding how the API is working as it is, my immediate thought was - "Let's just change it to `app.json()`". When I tried that, however, the response formatted the body as a string instead of a JSON object, and the content type was `text/plain` instead of `application/json`.

Somehow, `app.jsonp()` was acting like `app.json()`, and `app.json()` was acting like `app.text()`. I was bamboozled. I have been using Express.js for a long time, and I have never see such strange behavior. I searched Google and Stack Overflow for answers, and it seems I am the only person who has ever encountered this problem.

I thought it was because we were using an outdated version of Express.js. As with many other issues, updating the software often fixes odd bugs. But not this time. No dice. Maybe there's an estoric configuration parameter that I am not aware of? So I searched through the documentation. Still, no luck.

After many hours of digging, I realized it's not a configuration parameter or an outdated version of Express.js, it was because, in a random file somewhere, we are overriding the `app.jsonp()` and `app.send()` methods. Since `app.json()` calls `app.send()` in the background, it is affected too.

How `app.jsonp()` and `app.send()` are modified is not important. What is important is that instead of *extending* the functionality of Express.js by creating *new* methods (e.g. `app.customJSON()` or `app.customSend()`), the developers of the API *modified* existing, well-established methods/behavior.

With `app.customJSON()`, readeres of the code would have expected it to be a custom implementation. With `app.send()`, readers of the code would have expected it to work the same way as is documented in the official documentation.

Had the developers extended the Express API with `app.customJSON()`, the expectation they set would have aligned with reality. By modifying existing methods, however, the expection is misaligned with reality, and lead to confusion.

In conclusion, if you are using some third-party libraries/SDKs, do not modify its behavior. If you must, extend it.
