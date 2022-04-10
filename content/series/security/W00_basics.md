# Ideas

- Defence-in-depth - https://en.wikipedia.org/wiki/Defense_in_depth_(computing)

# Cross-Site Scripting (XSS)

Cross-Site Scripting (XSS) is a vulnerability of a website that allows unsanitized user input to be directly injected onto a web page's HTML markup. This is most commonly done by badly-implemented comment sections. Instead of writing plain text comments, an attacker can inject a script, that when incorporated into the HTML that is served to other clients, will run the script on unsuspecting users.

This allows an attacker to execute an arbitrary script as their victims. The possibilities of what harm can be done is endless:

- [XSS Worms](https://en.wikipedia.org/wiki/XSS_worm)
- The injected script has access to the user's cookies for that site, and if such cookies are used for authentication, the script can perform actions on the site on behalf of the victim
- The injected script can read the user's cookies for that site (except `HttpOnly` cookies) and can send those cookies to the malicious party. (This is why sensitive cookies such as those used for authentication should be `HttpOnly`)
- The injected script can rewrite the page's HTML, which, itself, can have many uses:
  - rewrite the destination for which a form should be sent (by changing the `action` attribute of the `<form>` element) to send potentially personal information (credit card details) to malicious party.
  - To propagate false news on a legitimate site, fooling its readers
  - To vandalize a site

# Cross-Site Request Forgery (CSRF)

Cross-Site Request Forgery (CSRF) is where a malicious site, or a vulnerable site susceptible to XSS, makes a request, on behalf of the victim, to a legitimate site (e.g. a banking site) to perform something nefarious.

Prevention: https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html#token-based-mitigation
