---
title: Dump
slug: dump
date: 2020-12-31T16:35:12-08:00
chapter: z
order: 99
tags:
    - programming
draft: true
---

## Common Sense

- Keep code DRY
- Avoid premature optimization. Keep code DRY but don't be a tyrant about it.
- Teaching is the best way to learn (i.e. the Feynman technique) - you don't really know you know something unless you can explain it to someone else, or for someone to question you.
- You can never know everything - this works two ways - it means you shouldn't kick yourself because you don't know something, but it also means don't think you know a lot, because there's always more to learn.


## Avoid Too Many Layers of Abstraction

Many people can only keep so many layers of abstractions and levels of indirection in their head.

When writing an web API, for example, I usually have an API routing layer, a handler layer that handles the request and response, and a data layer that interfaces with the database. If there are common logic shared by any of these layers, that logic gets abstracted into helper functions and placed in a `helpers/` directory.

But I've worked at a place where this is split into many layers:

1. API routing layer
2. Santization/validation layer
3. Faux handler layer that transform the request to a form acceptable by the engine layer, and transform the return value from the engineer to an appropriate response
4. Engine layer where the core logic lies
5. Models layer that mould data into the form acceptable by the database
6. Persistence layer that carries out the actual CRUD functions on the database

Whilst _some_ operations may require this complexity, it was definitely too many layers for what we were doing. It made it difficult for newcomers to get on board easily, and developers have to open 6 different files just to follow the logic of one request. When one operation may involve 5 or 6 requests, it's really hard to keep track of the flow.

In practice, developers didn't actually keep to the framework - some people wrote engine logic into the handler layer, and persistence logic to the models layer, etc. This was enough evidence to convince me that it's too many layers.

## Motivation Trumps Discipline

Many motivational videos tell you that discipline is more valuable than motivation, because motivation is in short supply, but discipline is like a muscle that you train. Whilst this works for some people, it doesn't work for me.

Work, for me, needs a purpose. If I am not motivated but disciplined, I can make sure I sit in front of the computer at 9am every morning, read articles, write code etc. But my work won't be any good because you need more than discipline to do _good work_. Programming is a creative problem-solving process, and it's supposed to be fun. If you feel a bit burnt out, or want to quit the field, what you need is not discipline, but a break - a break to rediscover your resolve, your motivation.

I think motivation comes from inspiration - you are inspired by something greater than you and your innate drive to be better motivates you to reach new heights. Motivation may come from reading about other people's work. Personally, I am fascinated by 3D animations, natural language processing, Linux, bug bounty hunting etc. and I am always in awe of the people in these fields and the work they do. When [self-driving cars](https://en.wikipedia.org/wiki/Self-driving_car), [deepfakes](https://en.wikipedia.org/wiki/Deepfake), and [AI-generated photos](https://generated.photos/) first came to light, I was amazed at what machine learning and AI technology can do today, and how little I understand it. It was the same fascination that brought me into programming, and it's this same fascination that's going to keep me going.

Apart from technical motivations, you also need to find a non-technical resolve - may that be family, friends, or a greater cause.

Discipline is important when shit needs to get done but we don't really feel like it. But just like muscles lifting weights, discipline is working _against_ something - it is tiring and has its limits. Motivation, on the other hand, when you can find it, pushes you _forward_ - it gives you energy even when you feel tired.

Don't depend solely on motivation or discipline - you need both.

## Charlatans and Sheep

Back in 2014, when I was just a few months in my web development career, when I was the bottom of the barrel in my first web development job, no one respected my time or my opinion. I wanted to be heard, to be respected. I wanted to be one of those people that write blog posts that people read, to be one of those speakers that gets invited all around the world to speak at prestigious conferences.

So I started a blog. In it, I write about what I have just learnt. I try to read many articles on the topic, condense the most important parts, discard the fluff and inaccuracies, research what's unclear, and produce a more polished article. I do it because Richard Feynman taught me that teaching others is the best way to learn. So if I am reading someone else's article and I ask "Why is that?", or "What does X mean?", then that means there's room for improvement on their article, that means it may be worth me writing a better one.

I recently read an article on HTTP cookies called [How HTTP Cookies Work](https://thoughtbot.com/blog/lucky-cookies). In the first paragraph, the author writes:

> When starting this work, I knew very little about how HTTP cookies actually worked. We’ll explore what I learned about cookies and how they are implemented.

Writing up about what you've learnt is, in my opinion, the best way to solidify what you've learnt and also to produce the clearest content. The author is upfront about his experience, and he doesn't claim to be an expert - just sharing what he's learnt. But what I have seen more and more are people who regurgitates other people's content, without asking questions, without doing more research, without explaining things in a different (better) way, and claiming it as their own thoughts. What ends up happening is an echo chamber that started off with a few articles, then suddenly hundreds of articles on the same topic, with mostly the same content.

What's worse is that because these content copiers do not properly process the information they are regurgitating, they don't really understand what they are teaching others. What ends up happening is a game of Telephone, where inaccuracies creeps into the content. So what started off as a few decent articles, spawn into hundreds of bad or harmful articles.

During my research on cookies, I came across another article on `SameSite` cookies, in it, the author wrote:

> Both in anticipation and in reaction to SameSite’s activation in browsers, posts started sprouting on blogs all over the Web to spread the word about the mechanics of the “new” cookie attribute.
> Some of those posts were of admirable precision, such as Rowan Merewood’s web.dev piece entitled “SameSite cookies explained”. Unfortunately, relatively few of the posts about SameSite went to the effort of clarifying the concept of site, from which the concepts of same-site request and cross-site request are of course derived.
> Moreover, many posts, including those produced by influential members of the infosec community, appeared to use the terms “origin” and “site” interchangeably or, at least, somewhat loosely.
> Back in February 2019, Kristian Bremberg wrote the following on the venerable Detectify blog:

    The SameSite attribute is rather new and provides excellent protection against CSRF attacks. If a cookie uses the SameSite attribute, the web browser will make sure that the request made with the cookie came from the origin that sat [sic] the cookie.
> Infosec superstar Troy Hunt himself, in a seminal post entitled “Promiscuous Cookies and Their Impending Death via the SameSite Policy” and published in early January 2020, described the effects of the different SameSite attribute values as follows:

    None: what Chrome defaults to today without a SameSite value set
    Lax: some limits on sending cookies on a cross-origin request
    Strict: tight limits on sending cookies on a cross-origin request



"But you said earlier that teaching others is the best way to learn, so a you're now saying beginners should write articles?" No, that's not what I am saying at all. Even if you're a beginner, there's always something you can teach others. In fact, beginners have the benefit of seeing things from a beginner's perspective, so when they explain a concept to others, they are less likely to skip a step or make assumption about what others already know.

But authors, regardless of experience, should only write about topics that they've thoroughly research; better yet - have verified themselves. When I wrote an article about setting up NGINX, Express and PM2, I didn't just gather code snippets from various articles and stuck them together, I tested those code snippets by actually setting it up myself and tested that it works. When I wrote an article on how Docker containers work, I didn't just reiterate analogies made by other people ("a container is just like a lightweight virtual machine"), I researched into cgroups, namespaces, processes, etc. and only when I am confident in knowing how containers isolate themselves from other processes on the host did I begin to write my article.

There are a lot of good, bad, and truly terrible content out there. Before you write a new article, ask yourself honestly "Can I truly do a better job than what's been done already?" If the answer is yes, give it your best shot; if not, help spread the good content because you know the creator put a lot of effort into making it. But please don't add to the bad and truly terrible. You don't need to be an expert to write articles, you don't need to know everything. Write about what you _do_ know, even if it seems like a speck of sand in an endless ocean - a diamond, no matter how small, is still a diamond.

But here comes the even more subject part of this post (what?!). I believe that if your motivations for writing is truly to learn, then whatever you write is going to be good; this is because, in your learning, you would have asked a lot of questions, researched to find the answers, and your writing would be a documentation of your learning. When you're learning, you can't skip a step, because if you do, you wouldn't have reached the end (i.e. you wouldn't have understood).

But if your motivation is to _look_ like you understand, to _look_ like an expert, then whatever you produce will only ever (at best) _look_ like it's good content, and, at worse, expose you as a fraud, a _charlatan_.

A lot of people in the software development industry suffers from Imposter Syndrome, where they feel like they are not good enough for their job. I have certainly felt this way before - there's just so much I do not know, how can anyone feel competent (or even adequate) know how little they know? Well, in my opinion, know just how much you don't know is part of wisdom. Fools always think they know everything because they lack the wisdom (and humility) to recognize what they don't know. Or worse, they're trying to lie to themselves and others because their sense of worth is based on how much they know. So if you're suffering from Imposter Syndrome and feel not good enough for your job, just know that many of your peers and experts don't know a lot of things as well, and we all had to start off somewhere. Even chess grandmasters didn't know how the pieces moves at one time in their lives. You're doing fine as long as you keep trying, keep learning.

But there are _actual_ imposters in our industry. Those who pretend to know a lot, who pretends to be experts, but are actually not worth their weight in salt. You'll know this is you if you're constantly anxious about being found out. I know this because I was an imposter a few times in my career.

One time (quite a few times actually), I volunteered to give a presentation but didn't prepare for it sufficiently, leaving myself less than 8 hours to prepare. Obviously, when it was time to present, I hadn't validated everything I want to talk about, and so was constantly nervous about saying the wrong thing, or making a mistake.

As a side note, if you feel nervous about presenting, a lot of it is because you don't feel like you know the subject well enough. You are nervous because you feel like you don't bring enough value to warrant your audience's time. I find that if I fully prepare for presentations, if I truly know what I am talking about, when I know I can bring value to my audience, most of the nerve disappears.

So I am calling imposters who teaches others things they don't actually know _charlatans_.

And charlatans can profit because there are enough sheep to follow them. They profit through their fame by selling courses, writing blog posts or producing videos with adverts, sponsored/affiliate links, etc. As long as enough people read their content, this model works for them. Their content don't have to be accurate, since most people who clicks on "How cookies work?" don't actually know how cookies work. Most people falls prey to the fallacy of appealing to authority - if the author _looks_ authoritative enough (e.g. a self-professed _senior_ full-stack developer) then what they say must be right. If it _sounds_ like it makes sense, then there's no need to validate it.

Of course, I am not talking about you, my dear reader. I am talking everyone else.

All jokes aside, there are enough people who are happy to accept what they read without validating themselves, or consulting with a more authoritative source. These people I call _sheep_.

Being admired without putting in the work is tempting, and that's why there are so many charlatans out there. But resist the urge - the 'fame', the 'attention', it's not worth it. In the long run, you may even find it to be a burden. A person who no one knows can be anybody, whereas a person who everyone knows have an obligation to continue being that person.

Don't be a charlatan, don't be a sheep.
