---
title: Data Representation - Numbers
slug: data-representation-numbers
date: 2021-01-01T15:50:36-08:00
chapter: a
order: 1
tags:
    - computer-science
draft: true
---

We mentioned previously that everything is stored as binary numbers. So how are numbers represented and interpreted? By that I mean given a sequence of binary data, say `0001111010110110`, how do we know if it's 16 1-digit binary numbers, or 8 2-digit numbers, or 4 4-digit numbers, or a mix?

In the context of data, each digit in a sequence of binary data is called a _bit_ (short for binary digits). We can store numbers as n-bit binary sequences.

So if we use 4 bits for a number, the binary sequence `0001111010110110` represents 4 values:

- binary 0001/decimal 1
- binary 1110/decimal 14
- binary 1011/decimal 11
- binary 0110/decimal 6

But if we choose to represent numbers using 16 bits, then the same sequence of `0001111010110110` would be 2 values:

- binary 00011110/decimal 30
- binary 10110110/decimal 182

?? How do we know if the number is 8-bit, 16-bit, 32-bit or 64-bit number ??

?? How do we know if a number is signed or unsigned ?? 2's compliment format ??





Binary digits can be grouped together into bytes. There are two popular methods for converting binary to denary.


Integers and Floating Point Numbers

## Text

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

And this is exactly what happened in the 1960s. A working group within the American Standards Association's (ASA) developed the _American Standard Code for Information Interchange_ (_ASCII_) (pronounced /ˈæskiː/, or ASS-kee). ASCII encodes 128 characters into 7-bit integers.

![ASCII table showing the mapping of characters to decimal numbers](/img/ascii-decimal-chart.png)
(screenshot from asciichart.com)

As you can see, the ASCII standard defines character encoding for alphanumeric characters, punctuation, space, and a bunch of what are called 'control characters'. So using ASCII encoding, everyone can agree that the number sequence `72 101 108 108 111` represents the text `Hello`.

8 bits is called a _byte_.

But ASCII only have a limited number of characters it can represent (128), and doesn't support accented characters (e.g. `á`, `ø̄`), or many non-latin alphabets like Chinese (e.g. `地`, `雨` ), Thai (e.g. `ฆ`, `ฐ`), Arabic (), etc.


?? UTF-8, UTF-16 ??

Unicode supports emojis. What's being stored is not the image itself, but a numeric code point in Unicode. When you send an emoji to someone, you're sending a number which the receiver's program would interpret as an emoji and display an image. This is also why emojis look different on an Android machine compared to an Apple iPhone - because it's not the image that is transmitted, it's a number.

So how are new emojis added? Well, each year, a ??committee?? sits down and decide what new emojis to add to the standard. In 2020, 117 new emojis were added, including Bubble Tea (or Boba) and the transgender flag (
🏳️‍⚧️)

## Images

## Videos

## Audio

## Computer Programs

Computer programs are compiled or interpreted into assembly code, which the assembler translates into machine code - binary code that can execute on the CPU.


