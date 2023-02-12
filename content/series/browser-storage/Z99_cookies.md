---
title: Cookies
slug: cookies
date: 2023-02-09T17:04:09-08:00
chapter: z
order: 99
draft: true
---

Cookies were designed to storage small pieces of data given by the server from its response, so that the browser will *automatically* store it and send it back to the server in subsequent requests.

However, there's no stopping you from storing arbitrary data in a cookie even when that data is not sent from the server. Those cookies will be sent to the server on every request.

A cookie can store up to 4096 bytes of data.


- Layman
  - What are cookies
  - How are cookies used
- Technical
  - What are cookies
    - [Each attribute]
  - How are cookies used
    - [Each use]

Difference - Domain, host, origin, site

- domain (`example.com`)
- host (`localhost`, `example.com`)
- origin - scheme, host, and port (`https://example.com:8080`)

Cross-Origin Resource Sharing (CORS) deals with origins
`SameSite` cookies, XSS, CSRF deals with sites.

Site will soon include the scheme (http and https will be considered different sites)

https://web.dev/same-site-same-origin/

> **What are cookies?** Cookies are small pieces of text that a web server gives to your browser, with the intention that the browser will store it and present it back to the server on subsequent requests.

Typically, a website is visited by many visitors, many of whom will be anonymous (not logged in). Now, each of these anonymous visitor may visit multiple pages from the same website during the same visit. Because the core of the HTTP protocol is stateless, from the perspective of the web server, each HTTP request from the visitor's browser is independent, meaning the web server is not able to deduce that these HTTP requests are made from the same browser.

If a server wants to keep track of which requests belong to the same visitor, it can group requests based on the same IP address. But multiple visitors may share the same IP address (think cafes, libraries, universities). Furthermore, this scheme requires the server to keep state on the server which may require a lot of compute, memory, and storage resources.

Cookies (a.k.a. HTTP cookies, web cookies, browser cookies, Internet cookies, computer cookies) are one way for servers to identify visitors, even anonymous ones, across multiple HTTP requests, by preserving some of that state on the visitor's browser. In this scheme, the collective state is distributed on the client's browser, not centrally on the server.

## Uses

### Authentication

This is known as cookies-based authentication.

> There are other kinds of identification; some, such as session-based identification, can work in parallel with cookies (you can have session-based cookies-based authentication).

### User Preferences / Personalization

In this scheme, the server wants to store some small identifier on the user's browser (e.g. session ID) so that the browser don't need to send a large amount of user information (e.g. user preferences) every time. The server does this by setting a cookie.

Preferences:

- Color scheme (dark theme, `theme=dark`)
- Whether the user has agreed to some terms (`terms-agreed=1`)
- Whether the user has dismissed some notice/pop-up (`notice-acknowledge=1`)

Cookies are string data that are given to browsers by web servers. When the browser send subsequent HTTP(S) request to the same server (identified by domain), it will include all cookies that the server has given it previously.

There are several common usages of cookies, but they all revolve around maintaining some kind of identity. Whether it's to keep a user authenticated after they've logged in, or keeping track of an anonymous browsing session on an eCommerce site.

When a user first access a website, the browser would not have a cookie from the website's server, and so the first request the browser sends to the server do not contain any cookies. From this, the server would assume that this user is a 'new' user. In the response to the first request, the server can send the browser a cookie containing something that can identify the user.

In the case of a user logging in, it may be some kind of token. The browser needs to send this token along with every request to protected endpoints (those that requires the user to be logged in, such as the profile page) in order to access them. The token is used for authentication in place of the username/password combination. This is how you're able to log in to a site once, and continue browsing without entering your credentials on every page.

In the case of an anonymous user, it may be an arbitrary session ID. The server may implement some logic to keep track of actions the user has performed on its website, using the session ID to distinguish between different users. For example, an e-commerce site may use a cookie to see the products the user have viewed (the anonymous session data), in order to provide more relevant recommendations on subsequent pages.

## Setting Cookies

A web server and a browser exchange cookies by using two HTTP headers:

- a web server can give cookies to browsers by adding a `Set-Cookie` HTTP header in the response. To set multiple cookies, you either:
  - add multiple `Set-Cookie` headers, or
  - use commas to separate the cookies within the same `Set-Cookie` header
- the browser can attach relevant cookies with subsequent requests by adding the `Cookie` HTTP header in the request. Multiple cookies are separated with a semi-colon. (A cookie's attributes are not send back to the server, because the server doesn't care. The attributes are only there for the browser to know _when_ and to _whom_ to send back the cookie)

```
Set-Cookie: <name>=<value>
```

It is up to the browser to accept those cookies and stores them. On subsequent requests, the cookies set by the same origin will be included in a single `Cookie` header (multiple cookies are joined, separated by a semi-colon).

### Anatomy of a Cookie

Specification for Cookies.
https://tools.ietf.org/html/rfc2109
https://tools.ietf.org/html/draft-ietf-httpbis-rfc6265bis-07
https://tools.ietf.org/html/draft-west-cookie-incrementalism-01
https://tools.ietf.org/html/draft-west-first-party-cookies-07

From the perspective of the server, the name and value of a cookie is all it cares about. But it can set additional _attributes_ that instructs the browser of _when_ it should send a cookie back.

- `Domain=<domain>` - Specifies which domains this cookie should be sent to. A site can only set the `Domain` attribute to a domain under the same site as the current domain. This means a cookie sent by `www.section.example.com` can be sent back to:
  - the current domain (`www.section.example.com`) - this is the default behavior
  - any of its parents (`section.example.com`, `example.com`) if `Domain=section.example.com` or `Domain=example.com`
  - a subdomain under its parent (e.g. `cdn.example.com`, `another.section.example.com`) if `Damain=cdn.example.com` or `Damain=another.section.example.com`

  Here, `app.section.example.com`, `example.com`, and `cdn.example.com` are considered by the browser to be part of the same _site_.

  In the beginning, browsers determines whether two domains are part of the same site by assuming that any subdomains that share the same domain under the top-level domain (TLD, such as `.com`) belongs to the same organization, and it's also assumed that the organization would not attack itself. So `example.com` and `app.example.com` are assumed to belong to the same organization, and thus considered the same site.

  However, this assumption becomes more complicated as many registrars offers second-level domain (SLD or 2LD, e.g. `.co.uk`) that anyone can register. So browsers should not assume that `foo.co.uk` is controlled by the same organization as `bar.co.uk`.
  
  The solution, then, is for browsers to keep a list of all 'official' TLDs and SLDs.

  However, this separation based on TLDs and SLDs can no longer be assumed. Take the case of Heroku, where users can deploy applications under the subdomain `*.herokuapp.com`. Therefore, `appa.herokuapp.com` and `appb.herokuapp.com` are no longer be considered to be under the control of the same organization, and so this separation of sites at the TLD level no longer works.

  Therefore, instead of separating sites based on , the Mozilla Foundation created a new Public Suffix List, which lists the domains under which people can register their own (sub)domains. Nowadays, all major browsers such as Firefox, Chrome and Opera use the list to determine if two domains are part of the same site.

  So a more concise way to determine the valid values for the `Domain` attribute is that a web server can set a cookie for its own domain as well as any parents' domain, as long as that parent domain is not a _[public suffix](https://publicsuffix.org/)_. This means that `news.example.com.` can set the cookie for `example.com.`, but not `com.`.

  The Public Suffix List allows browsers to know that `sysfig.github.io` and `example.github.io` are different sites, so are `appa.herokuapp.com` and `appb.herokuapp.com`, but `www.reddit.com` and `old.reddit.com` are the same site. This is because `github.io` and `herokuapp.com` are in the Public Suffix list, whereas `reddit.com` is not.

  > Note the distinction between **origin** and **site**. `sysfig.github.io` and `example.github.io` are different origins _and_ different sites. `www.reddit.com` and `old.reddit.com` are different origins but the same site.

  A site should not, however, set `Domain` to a different site (`foo.com` cannot set the `Domain` attribute of its cookies to `bar.com`), or to a subdomain of the current hostname (`app.example.com` should not set the `Domain` attribute of its cookies to `foo.app.example.com`). If the `Domain` attribute is invalid, the browser must reject the cookie.
- `Path=<path>` - The path prefix under which the cookie should be sent. For example, setting `Path` to `/` will cause the browser to send the cookie to all paths (provided it satisfies the other attributes). Setting the `Path` to `/v2/` will only send that cookie for paths under `/v2/`, such as `/v2/profile`, but not `/profile` or `/v3/profile`.
  `Path` allows the server to distinguish between applications (or different parts of the same application) that shares the same hostname.
- `Expires=<date>` - sets an expiration date after which the browser should no longer send the cookie and may delete the cookie. The date must be formatted as a HTTP Date (`Date: <day-name>, <day> <month> <year> <hour>:<minute>:<second> GMT`).
  Since cookie attributes are used only by browsers to determine whether it should send a cookie, the date specified in the `Expires` attribute is evaluated based on the client's operating system's clock. So if the client's system clock is not accurate (perhaps due to _[clock skew](https://en.wikipedia.org/wiki/Clock_skew)_), then it's possible for the browser to send an already-expired cookie.
- `Max-Age=<duration>` - An alternative to `Expires` that limits the validity of the cookie. Instead of setting an absolute date (as `Expires` does), `Max-Age` sets a relative time, in seconds, after the response is received for which the cookie should expire.
  A zero or negative value will expire the cookie immediately.
  If the `Expires` and `Max-Age` attributes are both set, then `Max-Age` takes precedence.
- `Size`
- `HttpOnly` - Typically, JavaScript code can read and write cookies from the same origin by accessing the `document.cookie` property. Setting the `HttpOnly` attribute restricts access to the data from JavaScript from other origins, and prevents cross-site scripting (XSS) attacks.
- `Secure` - Instructs the browser to only send this cookie back over HTTPS (some browsers make exceptions for `localhost` and local, private addresses)
- `SameSite=<value>` - determines whether a browser should send a cookie in a cross-site request

  To reiterate what we said for the `Domain` attribute - a cookie can only be sent back to the same domain that set it (the origin), a parent of the origin, or a subdomain under the parent, depending on the value of `Domain`. A cookie will _never_ be sent to a different site.

  The `SameSite` attribute, then, is not concerned about _where_ the cookie is sent to, but rather _who_ can _initiate_ the request.

Attributes are added to the end of the `Set-Cookie` header, separated by semi-colons.

```
Set-Cookie: <name>=<value>; <attribute-name>=<attribute-value>;
```

limit its availability using the `Domain`, `Path`, `Secure`, and `HttpOnly` flags.

The other attributes exists for security - to ensure that the browser only sends back cookies to servers who have set them. This is important because cookies are often used to store tokens used for authentication; if the browser was tricked to send cookies to a malicious party (rather than the legitimate server that set the cookie in the first place), then the malicious party now have the token and can masquerade as that user.

In the following section, we will explore how browsers implement cookies security.

According to the [RFC6265](https://datatracker.ietf.org/doc/rfc6265/) spec, browsers should allow each cookie to be at least 4096 bytes in size (any more than that is browser-dependent). Each domain can set at least 50 cookies, and browsers should support at least 3000 cookies in total.

>    Practical user agent implementations have limits on the number and
   size of cookies that they can store.  General-use user agents SHOULD
   provide each of the following minimum capabilities:

   o  At least 4096 bytes per cookie (as measured by the sum of the
      length of the cookie's name, value, and attributes).

   o  At least 50 cookies per domain.

   o  At least 3000 cookies total.


## Types of Cookies

We can categorize cookies based on several dimensions:

- Duration - how long the cookie is stored on the browser
- Provenance - who set the cookie
- Purpose - what is the cookie used for

### Duration

Session Cookies vs Persistent


Session cookies are stored in memory and not written to disk. They are automatically deleted after the session is finished (what a session means is dependent on the browser implementation; typically, a session ends when the browser tab is closed).

Persistent cookies have an expiration date and can persist after the user has quit and reopened the browser. If neither the `Expires` nor `Max-Age` attributes are set, the cookie defaults to a session cookie.

### Provenance

First-Party vs Third-Party Cookies

A first-party cookie is given by the web server serving the web page the user is directly viewing.

A third-party cookie are given by servers that are serving some content on the web page being viewed (such as adverts), but not the web page itself. In the case of adverts, if an advertiser have a large reach (i.e. they have their adverts on many websites, such as Google AdSense), then when a user visits two of those sites (e.g. gardening blog and eCommerce site), then the advertiser is able to identify that user as the same person, and may deduce that the user is a gardening enthusiast with spending power, and so the next time that user's browser sends a request to the advertiser's server (to retrieve the ads), the advertiser's server can return with more relevant ads (this is known as _targeting_).

Because of this ability for third-party cookies to collect user information from multiple sites, it's often seen as intrusive and privacy-infringing.



## Cookies Security

https://humanwhocodes.com/blog/2009/05/12/cookies-and-security/

### Scoping Cookies

In the most basic scenario, browsers should only send back cookies to servers that set them in the first place (i.e. the origin server).

This means if `legitimate.com` set a cookie containing a session ID I can use to authenticate, when I visit `malicious.com`, my browser must not send `malicious.com` my `legitimate.com` cookie. Doing so will allow whoever controls `malicious.com` to steal my session ID and use it access `legitimate.com` masquerading as me.

Cookie scoping is a security feature under the _Same-origin policy_ umbrella that are built-in to all major browsers.

However, the `Set-Cookie` has a `Domain` attribute that allows websites to set cookies for any of its parent domains and sub-domains. This means `malicious.example.com` is able to set cookies for `legitimate.example.com` and `example.com`. It can do so by setting the following `Set-Cookie` response header:

```
Set-Cookie: <name>=<value>; Domain=example.com
Set-Cookie: <name>=<value>; Domain=legitimate.example.com
```

### Session Fixation / session hijacking

https://www.abortz.net/papers/session-integrity.pdf

Cookie stuffing

### Cross-Site Requests

A cross-site request is where a request to origin A is initiated on a page served by origin B. For example, I've included a picture of the flag of Amsterdam taken from Wikimedia Commons.

![](https://upload.wikimedia.org/wikipedia/commons/thumb/6/6d/Flag_of_Amsterdam.svg/320px-Flag_of_Amsterdam.svg.png)

The image is not hosted on my server (`sysfig.dev`); rather, I am linking directly to the Wikimedia Commons server where I've found it (`https://upload.wikimedia.org/wikipedia/commons/thumb/6/6d/Flag_of_Amsterdam.svg/320px-Flag_of_Amsterdam.svg.png`)

When your browser parsed the link and sent a request to `upload.wikimedia.org` to load the image, this is an example of a cross-site request. A request for the image is made to `upload.wikimedia.org` (origin A) but the request is initiated on a web page served by `sysfig.dev` (origin B).

#### Cross-Site Requests and Cookies

Now, if you've previously went on `upload.wikimedia.org` and they've set a cookie, when your browser sends a cross-site request to fetch the image, the browser must decide whether to include the cookie set previously. Its decision is based on the cookie's `SameSite` attribute:

- `Strict` - the cookie must not be sent in cross-site requests
- `Lax` - the cookie must not be sent in cross-site requests, except in the case of navigation. For example, if you click the link [wikimedia.org](https://www.wikimedia.org/) now, your browser will send over any cookies set with the attribute `SameSite=Lax` to the `wikimedia.org` server, even though the request initiated from `sysfig.dev`, because clicking the link counts as navigation.
- `None` - the browser sends cookies for both same-site and cross-site requests.

Nowadays, if the `SameSite` attribute is not set, browsers default to using the `Lax` value; but previously, the default was `None`.

#### Cross-Site Request Forgery (CSRF) and `SameSite`

When the `SameSite` attribute is set to `None` (or defaults to `None` in older browsers), cookies are attached to all requests to the origin, regardless of what site initiated the request. This means if you visit a malicious site that wants to use your `SameSite=None` cookie to authorize some action on your behalf, all they need to know is what request to send.

Consider the following example:

- You log in to `bank.com`
- `bank.com` issues you a cookie (`session=a1b2c3`) that you can use to authenticate with `bank.com`'s server. This saves you from typing in you ID and password every time. However, `bank.com` has set the `SameSite` attribute to `None`.
- You visit `malicious.com`. That site contains an image link to `bank.com/transfer?to=malicious&amount=999`
- Your browser looks up cookies for `bank.com` and finds the `session=a1b2c3` cookie. Because `SameSite` is set to `None` for that cookie, your browser will attach that cookie with the request.
- The `bank.com` server sees a request to transfer money, and it checks the cookie to see that it's a valid session associate to you. Therefore, the `bank.com` server accepts the request as valid and process the transfer, meaning you just lost $999 to the malicious party.

We've outlined a simplified scenario for a bank, but it can be used to forge post on social media (impersonate), or delete files on file-hosting sites (destruction). 

Now, back to our banking example - if the `bank.com` server has set the `SameSite` attribute to `Lax`, it would be slightly better, but not by much. The difference now is instead of the request being sent automatically, the malicious website must now coerce or trick you into navigating to the URL.

This is not difficult, as they can disguise the link as something innocuous, such as "Click [here](bank.com/transfer?to=malicious&amount=999) to learn more". This is known as _click-jacking_.

The only way to prevent this type of attacks is to set the `SameSite` attribute to `Strict`, meaning requests to the `bank.com` server must always be initiated from the `bank.com` website itself.

However, not every application can use `SameSite=Strict` for every requests. For example, a video hosting site may want to allow third-party sites to embed an `iframe` of their video on their sites, but still allow logged-in uses to 'Like' the video. The video-hosting site can set an authentication cookie when the user first logs in, and whenever the user clicks the 'Like' button in the `iframe`, sends over the same cookie to authenticate the user. If `SameSite=Strict`, this won't be possible; they must use `SameSite=None`.

Embedded content is just one example, other examples include """widgets, affiliate programs, advertising, or sign-in across multiple sites"""(https://web.dev/samesite-cookies-explained/).

Note that CSRF attacks only applies to cookies that are used for authenticating or authorizing a user. You can have a cookie that sets the theme of the page (e.g. `theme=dark`) and that cookie can have `SameSite=None` since it can't be used to perform any actions on behalf of a user.

## Privacy

Building profiles so even if you never enter any Personal Identifiable Information (PII), they can know your race, gender, social-economic class, nationality, interests, etc.

### Tracking pixels

1 pixel by 1 pixel transparent image. When your browser requests this pixel from the server, the server sets a cookie with a unique identifier for you, and also records in its database which website you were visiting.

## Storing Cookies

Typically, you don't need to care about where cookies are stored as it is managed by your browser. How browsers store cookies are also implementation-dependent. Old browsers used to store cookies as text files - one file for each domain and one line for each cookie. Some stores cookies in a lightweight database such as SQLite

## Legal Aspects

GDPR - https://gdpr.eu/cookies/


Strictly necessary cookies are the only type of cookies which do not require consent from the user

## History

Before cookies, if you logged in to a website and closed your browser, you'd have to log in again the next time you go on the website. This is because there's no way to keep state on the browser. This is a usability issue.

So in June 1994, Lou Montulli, the ninth employee of Netscape Communications, the company behind the Netscape browser, came up with the idea of persisting some stateful data on the client side. Initially, this was aptly termed 'persistent client state object', but Montulli soon opted for the name 'cookies', after an even older computing term called 'magic cookies', which is a piece of information that computer programs pass between each other without changing its content.

In his proposal, the cookie was used to store shopping cart information locally in the browser.

```
Cookie: CUSTOMER=WILE_E _COYOTE; PART_NUMBER=ROCKET_LAUNCHER_0001
```

The state of the cart (items in the cart) is now on the browser, not the server. This is great because it saves the server from having to store that information indefinitely, not knowing when the client has actually abandoned the cart.

In 1997, the IETF recommended that all browsers should block third-party cookies unless the user explicitly allows it. But by then, the Internet marketing industry has solidified its hold on the Internet, and the two major browsers at the time - Netscape and Internet Explorer - did not act on the IETF recommendation, and cookies continue to be accepted and sent back to.

In 1999, DoubleClick announced the purchase of a company called Abacus Direct, which holds the information on 88 million catalog shoppers. DoubleClick's intention is to merge the online profiles it has gathered with the offline information of real people. This merge would give DoubleClick the complete profile or an enormous pool of people. However, after public outcry and a Federal Trade Commission inquiry, DoubleClick dropped the plan.

https://www.innovatorsunder35.com/the-list/lou-montulli/

## Notes

Locally Shared Objects (LSO) (Flash Cookies / Super Cookies) - similar to browser cookies, Flash cookies are text files created by Flash programs. They are stored permanently on your computer. It's not very secure as any Flash program can read LSO set by any other Flash programs from other website ??VERIFY??.

Since LSOs are not browser-dependent, if LSOs are used for tracking, it can track you across different browsers on the same machine.

Since LSOs are not deleted when you clear your browser cookies, they can be used as a redundant backup store of your browser cookies. So after you've deleted your browser cookies, the Flash program can use the data in the LSO to reinstate the previously-deleted browser cookies. These are known as _[Zombie Cookies](https://en.wikipedia.org/wiki/Zombie_cookie)_.

But Flash cookies is a thing of the past, since Flash is disabled on all major browsers. https://www.blog.google/products/chrome/saying-goodbye-flash-chrome/


### Uses

- Personal preferences (e.g. dark theme)
- Store data for non-registered users (e.g. shopping cart) - although there are better mechanisms for that nowadays like Web Storage API and IndexedDB
- Authentication tokens - keeping an authenticated session - by issueing a random, unique ID, the server is assuming that a hacker won't be able to guess a legitimate ID. So that ID is good enough to authenticate with. But that's not 100% secure. To make it even more secure, you can encrypt the user's data, and then base64-encode it so it's ASCII.

  You can also use base64-encoded JWTs
- Tracking data - early days, companies like DoubleClick and Engage place adverts on many different sites, tracking users across these sites using a shared cookie and building an anonymous profile of each user's interests. This means the advertisers can serve more targeted ads, or to ensure the same user do not see the same advert too many times.
  If the tracked user then registers on a site that then sells the user data to advertisers, the advertiser can now put a name to the anonymous profile. This information may be valuable to certain parties for nefarious purposes. For example, if Person A frequents websites about alternative medicine that uses advertiser B, and registers at a social media site that then sells user data to advertiser B, then advertiser B can sell information about Person A to Company C, who may exploit Person A's lack of scientific knowledge and sell him/her bogus medicinal treatments that is not helpful (or even harmful) because they know Person A is susceptible to scams.

  Tracking works because when you make a request from a web page, your browser may send a `Referer` header to the server telling them where the request is initiated from. For resource request, the `Referer` is the URL of the page that includes the resource; for a linked page, it's the URL of the page that contains the link.

  A page can, however, limit either the `Referer` header is sent by setting the Referrer Policy, which can be set using the `Referrer-Policy` response header or in a HTML `<meta name="referrer">` tag.
- Track Unique visitors
- Track how long between each page (cookie tracks time)

### SPA

Cookies are mostly used for server-side rendered pages (PHP)

https://web.dev/samesite-cookie-recipes/
https://web.dev/schemeful-samesite/
https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies

---

Without cookies, every time you end the session (e.g. by closing the tab or browser), you'd have to re-authenticate or rebuild your shopping cart.

---

Clients can manually remove cookies, or cookies can expire. The server, if it wants to remove the cookie, can override the original cookie by replacing it with an expired cookie; it may also be prudent to set the value of the cookie to an empty string.

---
https://publicsuffix.org/learn/

---

eTLD+1
