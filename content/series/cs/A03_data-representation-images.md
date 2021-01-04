---
title: Data Representation - Images
slug: data-representation-images
date: 2021-01-01T17:27:50-08:00
chapter: a
order: 3
tags:
    - computer-science
draft: true
---

There are two types of images - bitmaps (a.k.a. raster) and vector.

Bitmap images are made up of a grid of pixels, where each pixel is a color.

Vector images are a bit more complicated conceptually, as they are instructions that tells the computer how to draw the image on the screen - what paths to draw and what colors or pattern shapes should be filled in with. But to display a vector image, the computer must draw it on the computer screen, which consists of a grid of pixels. So ultimately, any pictures or graphics you see on the screen are bitmap images.

So how are bitmap images stored in the computer?

Well, just as text is composed of a sequence of characters, a bitmap image is simply a grid of pixels, each representing a color. And just as we use character encoding to map a character to a number, we can use another set of convention to map a color to a number.

There are numerous conventions out there:

- RGB
- RGBA


![](/img/6633FF-1.png)

A single-pixel Portable Network Graphics (PNG) image filled with a purple color represented by the Hex notation of #6633FF.

```console
$ xxd -b 6633FF-1.png
00000000: 10001001 01010000 01001110 01000111 00001101 00001010  .PNG..
00000006: 00011010 00001010 00000000 00000000 00000000 00001101  ......
0000000c: 01001001 01001000 01000100 01010010 00000000 00000000  IHDR..
00000012: 00000000 00000001 00000000 00000000 00000000 00000001  ......
00000018: 00000001 00000011 00000000 00000000 00000000 00100101  .....%
0000001e: 11011011 01010110 11001010 00000000 00000000 00000000  .V....
00000024: 00000011 01010000 01001100 01010100 01000101 01100110  .PLTEf
0000002a: 00110011 11111111 00110010 01001001 00000000 11110101  3.2I..
00000030: 00000000 00000000 00000000 00001010 01001001 01000100  ....ID
00000036: 01000001 01010100 01111000 10011100 01100011 01100010  ATx.cb
0000003c: 00000000 00000000 00000000 00000110 00000000 00000011  ......
00000042: 00110110 00110111 01111100 10101000 00000000 00000000  67|...
00000048: 00000000 00000000 01001001 01000101 01001110 01000100  ..IEND
0000004e: 10101110 01000010 01100000 10000010                    .B`.
```

You can see the 3 bytes starting with the last byte of line `00000024` is `01100110 00110011 11111111`, which is `102 51 255` in decimal (or `6633FF` in hexadecimal), representing the proportion of red, green, and blue that makes up the color of our 1-pixel image.

The structure of a simple PNG file consists of:

- PNG signature - a unique 8-byte (64 bit) sequence which 'marks' this file as a PNG file.
- Image header
- Image data
- Image end

https://en.wikipedia.org/wiki/Portable_Network_Graphics

The rest of the file consists of 


Videos are just a lot of images being displayed rapidly at a certain frame rate (e.g. 24 frames per second)
