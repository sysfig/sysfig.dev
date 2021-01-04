---
title: Data Representation - Text
slug: data-representation-text
date: 2021-01-01T17:26:24-08:00
chapter: a
order: 2
tags:
    - computer-science
draft: true
---

A string of text (e.g. `Hello`) is simply an ordered sequence of characters (i.e. `H` followed by `e`, followed by `l`, and so on). So to store text, we actually store just store an ordered sequence of characters.

But since everything is stored as numbers. So to store those characters, we must first convert into a numeric representation and store the number instead. Then, to read those characters back, we must have a way to do the reverse - convert those numbers back into characters.

The simplest conversion method is to simply have a table that maps each character with an integer. For example, I can invent my own convention whereby I'll use the binary number 0 to represent the capital letter A, 1 for B, 2 for C, and so on. For lowercase letters I'll carry on from where the capital letters ended, with lowercase 'a' being mapped to 26, 'b' to 27, and so on.

```
0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25
A  B  C  D  E  F  G  H  I  J  K  L  M  N  O  P  Q  R  S  T  U  V  W  X  Y  Z

26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51
a  b  c  d  e  f  g  h  i  j  k  l  m  n  o  p  q  r  s  t  u  v  w  x  y  z
```

This method of mapping a character to a number is called _character encoding_. Using this convention, the text `Hello` can be _encoded_ into the numbers `7`, `30`, `37`, `37`, `40`.

But what if someone uses a different encoding to me? If they start their table with lowercase letters first, their `Hello` text may be encoded as the numbers `33`, `4`, `11`, `11`, `14`.

So for people and programs to be able to read text stored by another program, the writer and the reader must agree on the same character encoding standard.

## ASCII

And this is exactly what happened in the 1960s. A working group within the American Standards Association's (ASA) developed the _American Standard Code for Information Interchange_ (_ASCII_) (pronounced /ˈæskiː/, or ASS-kee). ASCII encodes 128 characters into 7-bit integers.

![ASCII table showing the mapping of characters to decimal numbers](/img/ascii-decimal-chart.png)
(screenshot from asciichart.com)

As you can see, the ASCII standard defines character encoding for alphanumeric characters, punctuation, space, and a bunch of what are called 'control characters'. So using ASCII encoding, everyone can agree that the number sequence `72 101 108 108 111` represents the text `Hello`.

So if we create a txt file called `greet` and write `Hello` inside of it (without any newline characters at the end).

```txt
Hello
```

Then we can use command-line tools like [`xxd`](https://linux.die.net/man/1/xxd) to dump the contents the file in binary format.

```console
$ xxd -b test
00000000: 01001000 01100101 01101100 01101100 01101111           Hello
```

The first part of the output (`00000000:`) is the line number of the output, it is there simply to make it easier to reference locations of the bits in the sequence. The next part of the sequence (`01001000 01100101 01101100 01101100 01101111`) is the actual binary data stored in the file. `01001000` is binary for the decimal 72, `01100101` is binary for decimal 101, etc. And `xxd` is trying to be even more helpful by translating every 8 bits (called a _byte_) into its ASCII equivalent.

## Unicode

But ASCII only have a limited number of characters it can represent (128), and there's no more space to add support for accented characters (e.g. `á`, `ø̄`), or many non-latin scripts like Chinese (e.g. `地`, `雨` ), Thai (e.g. `ฆ`, `ฐ`), Arabic (e.g. `ش`, `غ`), etc.


?? UTF-8, UTF-16 ??

Unicode supports emojis. What's being stored is not the image itself, but a numeric code point in Unicode. When you send an emoji to someone, you're sending a number which the receiver's program would interpret as an emoji and display an image. This is also why emojis look different on an Android machine compared to an Apple iPhone - because it's not the image that is transmitted, it's a number.

So how are new emojis added? Well, each year, a ??committee?? within the Unicode Consortium sits down and decide what new emojis to add to the standard. In 2020, 117 new emojis were added to [Emoji 13.0](https://www.unicode.org/Public/emoji/13.0/), including Bubble Tea (or Boba), the transgender flag, and many gender-neutral emojis.

The current specification has 3245 emojis
