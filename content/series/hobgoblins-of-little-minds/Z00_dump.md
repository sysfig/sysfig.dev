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
