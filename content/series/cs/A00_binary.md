---
title: Binary
slug: binary
date: 2021-01-01T13:45:13-08:00
chapter: a
order: 0
tags:
    - computer-science
draft: true
---

A computer is a machine that takes some numeric input data, perform calculations on that data, and produces some numeric output data. Although we can watch videos, browse through photos, listen to music, and read text files on the computer, those videos, photos, audio, and texts are just sequences of numbers.

> In computing, all data is represented as numbers.

When we work with numbers in our everyday lives, we typically use the _decimal_ system. The word 'decimal' comes from the latin word 'decimus', which means 'tenth'. As such, the decimal system uses 10 different symbols to represent numbers - 0, 1, 2, 3, 4, 5, 6, 7, 8, and 9; because of this, the decimal system is also called _base 10_. In other words, each digit can represent 10 different values.

If we want to represent a number higher than 9, then we use more digits. But there are different ways to represent multi-digit numbers. For example, we can interpret the number 284 as 2 + 8 + 4 = 14, but that's not a very good counting system because:

- There are multiple ways to represent the same value (284 would be the same as 482 or 293)
- It would take a lot of digits to represent large numbers

Instead, we use a system where we read the number from right to left, and multiply each digit by base^n, where n is the number of digits to the right of that digit. So the decimal number 284, when read right-to-left, represents (4 × 10^0) + (8 × 10^1) + (2  × 10^2), which simplifies to (4 × 1) + (8 × 10) + (2  × 100), or 4 + 80 + 200. This system allows each value to have a unique representation, and allows the largest amount of values to be represented with the fewest digits.

But computers don't count using the decimal system; instead, at least in classical computing (as opposed to quantum computing), numbers are stored as _binary_, or _base 2_. This means the system uses two different symbols - 0 and 1 - to represent numbers. But the logic of reading the value of the number is the same. So for binary number 101010, when read right-to-left, represents (0 × 2^0) + (1 × 2^1) + (0 × 2^2) + (1 × 2^3) + (0 × 2^4) + (1 × 2^5), which simplifies to (0 × 1) + (1 × 2) + (0 × 4) + (1 × 8) + (0 × 16) + (1 × 32), or 0 + 2 + 0 + 8 + 0 + 32, which is the decimal number 42.

So how come we use base 10 but we designed computers to use base 2? Well, we probably ended up with the decimal system because we have 10 fingers. Computers don't have fingers. Computers are made up of electronic parts which runs on electricity. The electronic components that hold data inside processors are called transistors, and they have two possible states - on or off - which fits well into a binary system.

> Computers represent data as binary numbers.

So we can now be more concrete and say that a computer is a machine that performs calculations on binary data.

So if all data are represented as binary numbers, how does that work? How do we translate some number, text, images, and audio to binary and back? Well, the method differs by different data types, in the following posts, we'll look at each data type individually.
