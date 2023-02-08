---
title: Learning Abstractions
slug: learning-abstractions
date: 2020-12-31T16:35:12-08:00
chapter: z
order: 99
tags: []
draft: true
---

A lot of people ask whether they should learn JavaScript before learning jQuery. More recently, people ask if they should use React or vanilla JavaScript, CSS, and HTML.

My answer is always "Learn it properly from first principle - don't use a tool until you know why you need it."

I still agree with that answer, but now I have a different (but complementary) reason.

jQuery and React are abstractions from the vanilla JS, CSS, and HTML. But more importantly - they are imperfect, incomplete abstractions.

For example, with JSX, you can't use `class` as the name of a component's attribute, even though JSX is supposed to be an abstraction on HTML. The reason for this restriction is that `class` is a reserved keyword in JavaScript.

So, although JSX is supposed to be an abstraction on HTML, it is a leaky abstraction. For the user of this abstraction, they'll have to know about all the areas in which the abstraction leaks. This means to use the leaky abstraction, they'll have to learn about both the abstraction's interface, as well as some of the underlying complexity that the abstraction is trying to hide. Sometimes, this is a bigger learning curve than learning and using the underlying technology.

Thus, it's not always faster learning an abstraction over learning the underlying technology. But once you've learnt the underlying technology, the abstraction can make it easier to perform the same tasks faster in the future.

For teaching and learning, it means that when you teach a leaky abstraction, you should try to teach _how_ the abstraction works first, before teaching how to _use_ the abstraction later.
