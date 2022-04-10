In traditional computing (as opposed to quantum computing), all data, be it videos, images, audio, or texts, are stored as binary numbers (`0`s and `1`s).

To store a piece of string (a sequence of characters used to represent text), you store each character sequentially. But you can't just store the character `a` in the hard drive; computers and digital storage devices only understand binary numbers. Therefore, in order to store or transmit a character, you must first convert it into a number. Converting information (e.g. character) into code - a data format that can be stored and transmitted (e.g. numbers) - is known as encoding. Since we are encoding characters, this process is more specifically known as _character encoding_.

Generally, character encoding uses a mapping that assigns each character to an integer. For example, if we want to encode each letter of the English alphabet (including capital letters), one mapping may assign the character `a` to the integer `0`, `b` to the integer `1` etc. whist another mapping may assign the character `A` to `0`, pushing the character `a` to `27`. These different mappings shares the same _character set_ (the set of characters that can be use) but maps each character differently.

Traditionally, the integer (in its binary form) would then be stored and/or transmitted directly, and so these mappings are also known as _character encoding scheme_ (or, if the context is clear, _encoding scheme_ for short). For a program to read text written by another program properly, they must use the same character encoding scheme. If the decoding program uses an incompatible encoding, it will lead to garbled text (often called _[mojibake](https://en.wikipedia.org/wiki/Mojibake)_, romanized Japanese for 文字化け)

> Each character encoding scheme is associated with a character set, but a character set can have multiple character encoding schemes.

But it would be chaotic if every software developer uses its own custom encoding scheme. This is why vendors came together to agree on a standard character encoding scheme. One of the most popular standardized character encoding scheme in the early days of computing is _ASCII_, an abbreviation of _American Standard Code for Information Interchange_.

## ASCII

The American Standards Association's (ASA) X3.2 subcommittee began development of ASCII in 1960, and the first version was published in 1963. There has been several revisions since then, with the latest change in 1986.

ASCII encodes 128 characters into 7-bit integers ($2^7 = 128$). In character encoding terminology, we can say that the _codespace_ - the range of integers available for encoding characters - of ASCII is from 0 to 128 (or 0x00 to 0x7F in hexadecimals). Each integer within the codespace is called a _code point_.

Of the 128 code points:

- The first 32 (0 - 31 decimal, 0x00 - 0x1f hexadecimal) and the last (127 dec, 0x7f hex) code points encoded for non-printing _control characters_.

  The design of ASCII were based of older teleprinter encoding systems, and were also partly designed to replace those systems. As such, ASCII includes characters that are not meant to be printed, but provide metadata and instructions for devices, such as printers, that uses ASCII.

  For example, ASCII code point 10 (dec) represents the _line feed_ (␊) instruction, which instructs the printer or teleprinter to advance/roll its paper by one line. Similarly, ASCII code point 13 (dec) represents the _carriage return_ (␍) instruction, which instructs (tele)printers to return to the beginning of the line.

  ␊ and ␍ are still in use today, but others have become obsolete. For example, the bell code (␇) instructed the device to audibly alert the operator.

- 10 (48 - 57 decimal, 0x30 - 0x39 hexadecimal) are used for numbers 0 - 9
- 26 (65 - 90, 0x41 - 0x5a hexadecimal) for capital letters in the English alphabet
- 26 (97 - 122, 0x61 - 0x7a hexadecimal) for lowercase letters in the English alphabet
- The rest (25) fills up the space with punctuation characters

As its name implies, ASCII was intended to be used only in the USA, and lacks non-English characters and non-dollar currency symbols.

![](https://upload.wikimedia.org/wikipedia/commons/4/4f/ASCII_Code_Chart.svg)

### Design Considerations

#### Upper- and Lowercase Letters

The more keen-eyed of you may have noticed that there's a gap of 6 punctuation characters between the character `Z` and the character `a`. Why didn't the lowercase letters follow immediately after the capital letters? To understand this, we must look at the binary form of the code points.

The code point of the character `A` is `1000001` in 7-bit binary; for `a`, it's `1100001`. As you can see, they differ only in the second-most-significant-bit. This makes it easier to convert between upper- and lowercase letters since you just need to flip one bit. Similarly, it also makes it easier to compare letters in a case-insensitive manner, as you only have to test for equality of the 5 least significant bits.

#### Compatibility with Typewriters

The set of punctuation preceding the numbers correspond to their shifted positions on mechanical typewriters. For example, if you pressed <kbd>Shift</kbd> + <kbd>3</kbd> on a typewriter like Remington No. 2, you will get the `#` character. In ASCII, the characters `3` and `#` differ from each other by 16 places (1 stick).

#### Sorting

The alphabet characters were in numeric order in ASCII, which made sorting text alphabetically easy, as the computer only needs to sort the data numerically. Punctuation which acted as separators (such as `,` and `.`) appears before the digits and alphabet for similar reasons.

#### Maximizing Hamming Distance

In the early days of telecommunication, data transmission were error-prone. To minimize the effect of corrupted transmission, the code points that are deemed essential for data transmission (start of message (SOM), end of address (EOA), end of message (EOM), end of transmission (EOT), "who are you?" (WRU), "are you?" (RU), a reserved device control (DC0), synchronous idle (SYNC), and acknowledge (ACK)) are spaced out to maximize their Hamming distance.

Essentially, the greater the Hamming distance, the more errors must be introduced to convert on code point to the other. Therefore, by maximizing the Hamming distance of these code points, you minimize the risk of an error in transmission that leads to an EOM character being mistaken for an EOT character, for example.

## Deviations from ASCII

ASCII originally used 7 bits to store its data, but as 8-, 16-, and 32-bit computers became the norm, many variants of ASCII arose which made use of the extra bit to encode 128 more characters. These variants were collectively known as ASCII extensions, or Extended ASCII.

> Extended ASCII is not a single encoding, but rather a collective term.

Extended ASCII encodings differed from each other greatly, and each country seemed to have its own variant to include its own alphabet and currency symbols. This means we ended up with the same problem as before - where different software developers and vendors use their own variant of Extended ASCII and programs become less interoperable.

The next big attempt at standardization was the ANSI standard, but that simply documented all these variants into _code pages_. Nowadays, most of the popular character encoding schemes uses the _Unicode_ character set.

## Unicode

![](https://home.unicode.org/wp-content/uploads/2019/12/Unicode-Logo-Final-Blue-95x112.jpg)

[Unicode](https://home.unicode.org/) aims to be a universal character set that includes every character in all the world's writing systems. As such, the codespace for Unicode is massive and ranges from 0 to 1114111 (0x000000 to 0x10ffff). It has the usual `a` and `Z`, but also characters like `꫞`.

> You can view the core specification for Unicode v13 at [unicode.org/versions/Unicode13.0.0](https://www.unicode.org/versions/Unicode13.0.0/UnicodeStandard-13.0.pdf)

Because its codespace is so large, there are bound to be characters which look very similar to each other. For example, the characters `` ` `` and `⸌` look very similar. It would be impractical for Unicode users to refer to these characters by informal names because people will use different names for the same character. For example, many refer to the `` ` `` character as a backtick. Therefore, each character is given a unique _character name_ string as well as an integer code point.

For example, the Chinese full stop character `。` is given the unique character name 'Ideographic Full Stop', and the similar-looking `𑙃` character is given the character name Modi Abbreviation Sign.

Unicode is slightly different from the character encoding schemes of the past. Traditionally, characters are mapped to integer code points and the binary form of those code points are stored/transmitted directly, making the mapping itself the encoding scheme. Unicode also maps characters to code points, but these code points must themselves be encoded into a binary form called _code units_ to be stored/transmitted.

For example, the character `a` (Latin Small Letter A) maps to the code point 97, often written as `U+0061` (0x61 is hexadecimal for 97). This code point can be encoded using UTF-8 to yield the binary `01100001` (`61` in hexadecimals, and `97` in decimals). You can also encode the same code point using UTF-16BE, which gives the binary `00000000 01100001` (or `00 61` in hexadecimals).

Popular encoding schemes for Unicode characters include UTF-8 (**U**nicode **T**ransformation **F**ormat **8**-bit), UTF-16BE, UTF-16LE, UTF-32BE, and UTF-32LE.

### Encoding Unicode

#### UTF-8, UTF-16BE, UTF-16LE, UTF-32BE, and UTF-32LE

[UTF-8](https://www.utf8.com/) is a _variable-width encoding_ that encodes code points in 1, 2, 3, or 4 bytes. It is a variable-width encoding because it will use 1 bytes if it can store it in 1 byte; as opposed to _fixed-width encoding_ (e.g. UCS-4) that will use the same number of bytes (e.g. 4 bytes) even though it can store it in 1.

UTF-16BE and UTF-16LE are also variable-width encoding, but encodes code points in 2 or 4 bytes. UTF-16BE/UTF-16LE doesn't encode to 1- or 3-bytes for historical reasons. Between 1991 and 1995, Unicode only used a 16-bit, fixed-width encoding scheme called UCS-2, which accommodates 65,536 different characters. This was fine for Unicode v1.1, which had 40,635 designated and 24,901 reserved code points (for a total of 65,536 code points). However, Unicode v2.0 dramatically increased the supported code points to [1,114,112](https://www.unicode.org/versions/stats/charcountv2_0.html), which means new characters with high code points can no longer be encoded using UCS-2. As such, UTF-16, a variable-width encoding, was introduced as part of Unicode 2.0 to accommodate the new characters, but to as ensure backwards-compatibility with programs that used UCS-2, who assumed every code point maps to 2 bytes, UTF-16 had to use a _code unit_ of 16-bits, which means it encodes every character into either 2 bytes or 4 bytes.

Using UTF-16BE or UTF-16LE can potentially use more space than UTF-8, as characters that can be encoded in 1 byte must use 2 bytes with UTF-16BE/UTF-16LE; but it's also true that some characters encode to 3 bytes with UTF-8 but only 2 bytes with UTF-16BE/UTF-16LE. For example, the traditional Chinese character for 'machine' is 機, which encodes to `E6 A9 9F` with UTF-8 and `6A 5F` with UTF-16BE.

UTF-32LE and UTF-32BE are fixed-width encodings (4 bytes). ??UTF-32 is fast for internal memory representation?? (See https://stackoverflow.com/questions/496321/utf-8-utf-16-and-utf-32)

https://www.unicode.org/faq//utf_bom.html

#### Other Encoding Schemes

There are other encoding schemes out there:

- UCS-2 - encodes every Unicode code point to 16 bits (i.e. 2 octets/bytes) - allows you to encode , which may not cover all of the 143,859 characters in Unicode v13.0.
- UCS-4 - encodes every Unicode code point to 32 bits (i.e. 4 octets/bytes) - allows you to encode 4,294,967,295 different characters, way more than the 143,859 characters in Unicode v13.0. But uses 4 bytes for each character, which can waste a lot of storage, especially for English texts, where the characters (e.g. `a`) can be encoded into 1 byte with other encoding schemes.
- ISO 8859-1 (a.k.a. Latin-1) - encodes 191 characters from the Latin script of Western European languages into a single byte. It was the default encoding used in many web technologies (e.g. `latin1` is the default character set used in MySQL 5.7<sup>[ref](https://dev.mysql.com/doc/refman/5.7/en/charset.html)</sup>, MySQL 8.0 moved to using UFT-8). First published in 1987, its character set is a predecessor to, and forms the first 2 blocks of, Unicode.

  Some may still opt to use ISO 8859-1 (a.k.a. Latin-1) over UTF-8 as it can be more space-efficient. Characters from code points 0-127 are encoded identically with both encoding schemes (into 1 byte), but UTF-8 uses 2 bytes for the 128 remaining characters above code point 127, whilst ISO 8859-1 uses just 1. For example, the Latin small letter æ with code point `U+00E6` (`230` in decimal) encodes to `C3 A6` with UTF-8, but to only `E6` with ISO 8859-1.

- UTF-7 - similar to UTF-8 but guarantees that the _most significant bit_ (MSB, or _high bit_) is always `0`. This is needed for backwards compatibility with old systems that ??(ASCII was 7 bit)??

New versions of Unicode are released regularly and each new version includes newly-added characters. There are currently 143,859 characters in Unicode v13.0, increasing to 144,697 with the upcoming Unicode v14.0. See [Unicode® Statistics](https://www.unicode.org/versions/stats/) for statistics on other versions.

### Organizing Unicode

Although each character has a unique code point, it's often useful to categorize related code points together so that you can refer to them as a group.

One way of organizing Unicode characters is to split them into 17 equal groups of 65536 ($2^16$) characters each. Each of these groups is called a _plane_, and each is given a number between 0 and 16:

- Plane 0 - _[Basic Multilingual Plane](https://unicode.org/roadmaps/bmp/)_ (BMP)
- Plane 1 - _[Supplementary Multilingual Plane](https://www.unicode.org/roadmaps/smp/)_ (SMP)
- Plane 2 - _[Supplementary Ideographic Plane](https://www.unicode.org/roadmaps/sip/)_ (SIP)
- Plane 3 - _[Tertiary Ideographic Plane](https://www.unicode.org/roadmaps/tip/)_ (TIP)
- Plane 4 - 15 - Although the Unicode codespace is large (1,114,112 code points), not every code point is assigned to a character. As of Unicode v13, only 281,392 code points (~25%) are assigned. Therefore, many of the planes are left empty.
- Plane 14 - _[Supplementary Special-purpose Plane](https://www.unicode.org/roadmaps/ssp/)_
- Plane 15 - 16 - Supplement­ary Private Use Area (PUA) Planes

Every plane apart from the BMP are also called _astral planes_, or _supplementary planes_.
https://en.wikipedia.org/wiki/Plane_(Unicode)

A finer grouping of code points can be achieved with _blocks_, where each blocks is a continuous, non-overlapping range of code points, containing a multiple of 16 code points, and starting at a location that is a multiple of 16.

Another way to group code points is by slotting into one of 7 _basic types_, which are:

- Graphic - maps to characters that can be written and read by humans. In other words, a character which can be associated with one or more glyphs. Examples include numbers, letters of an alphabet, symbols, spaces, etc.
- Format - invisible characters that affects neighboring characters. Examples include line breaks
- Control - used by external protocols not defined in the Unicode standards
- Private-use
- Surrogate
- Noncharacter
- Reserved

Yet another way of grouping code points is with _General Categories_, which categorizes characters based on its usage. Each General Category has a 2-letter alias, where the first letter specifies its _major class_, whilst the second letter denotes its subclass.

- `Lu` - Letter, uppercase
- `Ll` - Letter, lowercase
- `Lt` - Letter, titlecase
- `Lm` - Letter, modifier
- `Lo` - Letter, other
- `Mn` - Mark, nonspacing
- `Mc` - Mark, spacing combining
- `Me` - Mark, enclosing
- `Nd` - Number, decimal digit
- `Nl` - Number, letter
- `No` - Number, other
- `Pc` - Punctuation, connector
- `Pd` - Punctuation, dash
- `Ps` - Punctuation, open
- `Pe` - Punctuation, close
- `Pi` - Punctuation, initial quote (may behave like Ps or Pe depending on usage)
- `Pf` - Punctuation, final quote (may behave like Ps or Pe depending on usage)
- `Po` - Punctuation, other
- `Sm` - Symbol, math
- `Sc` - Symbol, currency
- `Sk` - Symbol, modifier
- `So` - Symbol, other
- `Zs` - Separator, space
- `Zl` - Separator, line
- `Zp` - Separator, paragraph
- `Cc` - Other, control
- `Cf` - Other, format
- `Cs` - Other, surrogate
- `Co` - Other, private use
- `Cn` - Other, not assigned (including noncharacters)

In cases where a character can have multiple uses (Greek letters can be used as numbers as well as be symbols in mathematical expressions), it is categorized based on its primary use.


https://www.compart.com/en/unicode/category

Combining characters
https://dmitripavlutin.com/what-every-javascript-developer-should-know-about-unicode/#25-combining-marks
https://stackoverflow.com/questions/27331819/whats-the-difference-between-a-character-a-code-point-a-glyph-and-a-grapheme

#### Basic Multilingual Plane (BMP)

The first plane - the Basic Multilingual Plane - contains characters from most 

## Further Reading

Every effort has been made to ensure the accuracy of the terminology used, but errors may remain. If in doubt, refer to the [Glossary of Unicode Terms](https://www.unicode.org/glossary/) for the official definition.

If you see a character or symbol in print and would like to identify its Unicode code point, try drawing it on [shapecatcher](https://shapecatcher.com/), which is a Unicode recognition tool and supports 11817 printable characters.

https://deliciousbrains.com/how-unicode-works/

---

## Unicode and JavaScript

[ECMAScript 2020](https://262.ecma-international.org/11.0/) stores strings as a sequence of unsigned 16-bit integers.

uses UTF-16 to encode all its strings

> You can encode the actual JavaScript file (e.g. `main.js`) in any encoding you want (e.g. UTF-8), but when the JavaScript actually gets run by the engine (e.g. V8), the strings are stored as a sequence of unsigned 16-bit integers.

Since UTF-16 is a variable-width encoding, a character may be stored as 2 or 4 bytes. When a character is encoded and stored in 4 bytes (as 2 16-bit code units), the pair of code units are called a _surrogate pair_. The first unit is called the _lead surrogate_ (or _high surrogate_), the second the _trail surrogate_ (or _low surrogate_).

Traditionally, JavaScript treats surrogate pairs as two separate characters. For example, the personal computer emoji 💻 (Unicode code point `U+1F4BB`) encodes to `D8 3D DC BB` with UTF-16BE, but because JavaScript's `String`'s `length` property counts the number of UTF-16 code units instead of the number Unicode code points, `'💻'.length` evaluates to `2` instead of `1`.

```js
> '💻'.length
2
```

We can take this to the extreme with _[Zalgo text](https://en.wikipedia.org/wiki/Zalgo_text)_, which are text that has been modified using Unicode's _[combining characters](https://en.wikipedia.org/wiki/Combining_character)_ that is used to stack diacritics above and below letters.

```js
> 'Ḩ̵̨̡̢̡̢̨̨̢̢̡̨̢̢̛̛̛̗̠̤̖̰̼̯͖̩̞̲̝͉̘̪̣̘̼̬̺̮͉̠͔͔̘͉̳͓̼͎̮̩̭̰̼̱͈̳̩͕͕͍͕̬̰̭͉̬̻̻̥̺̞͔̜̩̗̠͈̖͖͓͔̳͎̺͉͈̝͖͈̻͕̭͖̘͕̪̬̬̹̖̖̤̠̝̘̳̣͕͇̭̘͕͉͖̯͐̿̆͒̉̌͒̏͊̈́͂͒̐̔̃̓̓̊͛̔͌͆̾͐̒̊̂̊̿̌̍̐͆̿̌̈́͒̒̑̊̈́̄̐͗̅̏̏̏͛̊̓̄̄̾͂̋̊̈̎̒̅̽̾͌̉͊̆͂̓̿͆͛̑̑͒̿͒̋̆̉͊̏͗̍͒̕̕̚̚͜͜͝͝͝͠͝͝͝ͅé̶̢̧̢̡̨̢̧̧̡̨̧̧̨̢̢̨̨̛̛̺͇͎̗͔͖̬͕͇̱͎͉̰̖͎͈̤̳̯̩̥͇̤̼̮̺̘̰̪̳̹͉͍͙̤͓̮̼̗̼͎͚͇̤͙̺̰̯̻͉͈̥̰̱̲̮͚̼̦͎͍̣̠̘͍̱͉̤̲̪̮̹̪̬̭̞͉͙̼͍̞̥͇̫̮̼̥̺̠̥͚͎͔̫͓̲̫͍̪͍̻̻̪͎̮̜̻͍̣̜̗̺̻̟̟̝͓̬̟̦̱̱̼̣̖̦͕̮̞͈̰̩̼̼͈̦̠̟̗̤͉̞͈̤͇̞̲̱͎̺̻̺̟̲̲̼̪̦͉̣̠̼̯̱̤̰̬̩̼̗̩͓̫͚͈̝͕̪̜͕͗̿͆́͛̀̍͂̿͊̔͌̓̓́̾͊̒̆̽͆̀̈͊̅͑̏̒̾̾̈́̒̔̒̇́̅͛̈́͊͐̉͗̓͛̑̍͂̅͂́̈̒͌̿͊͒̂̈̍̊̉̎̊̿̽̅̅͊̂̌̐̒̇̂̃̌̒̋͒̔̃̇͆̍͐͛̽̑͒̍̐͒͂̏̊̽̌̑̆͒̆̒̅͗̿̂̽̍͑͋̿͒̓̿̌̃̉̄̾̄̊͌͑̑̌͗̑̒̒̍͒͊̾̃͗͐̈̕͘͘̚̚͘̕͘͜͜͜͜͜͜͝͠͝͠͝͝͝͝͝͠͠͠ͅͅͅͅͅͅl̵̨̛̛̛̛̠̞̬̲̝͔͖̹̞͖̥̠͖͇̩͉̠͔̹̳͈̟͚͕̺̯̦̤͔̲̞̟̲͉̜̳̙̙̖̳̞͙̼͓̣͓̥̮̪̈́͊͆̉̋̀̿͆̑̏͒̅̈̑̌͂̂̿̽̾̔̀͊͂̀͆̈́̇̈́̊̾̿́̒͊͑͊̈̇̓̄̃͑́̓͑̓̎͐̇̿̓̌̉̔̇̈͂̐̔́̆̅͒͗̏̆͑̎̏̿̋̌͂̅͂̆̓͋͐̒̃̽̿͐̆̉̋̔̈̍̔̃͂͒̊̊͐̃͐̀͛̂̽͂̆͛̃̒͗̄̇͐̆̄͛̎͋̾̌̎̃̾̆̑̆͌̐̊̅̏̎̂̍̔͆̄̇̍̽̽̄͂̈̀̄̀̉̆̀͂͊̏͋͋͑̌̆̌́͒̋̓̓͒͐̾͌̂͆͒̏̐̾̅̃̈̋̒̀̽͂̉̐̄͒͘̚̕͘͘͘̚̚͘̚͘̕̚̚͝͝͝͝͝͝͝͝ͅl̶̢̧̨̡̡̢̢̧̧̨̛̛̛̛̛̛̛̛̹͕̙̤̰̟̭͖͙̗̻̖͉͇͕̤̲̥̖̬̪̤͓̗̪̲̲̞̺̹̠̟̳̫̣̳̺͖̼̘̝̫̙͎̻̬̜͓̱̭̟̗̮͇̺͔̺̯̺̻̞̙̤̖͎͖̼̣̹̙̩͇̩̰̜̘͓̤̯͎̮͖̗̺̬̘̼͍̠͇̲̯̺̻̞͓̬̗̞̥̮͕̯̥̤̘͚̩̯͍̻̻̣͔͎̦͉͇̬̩̪̦̻̠͉̻̮̖̱͚̺̙̱̼̩̪͍̩̻̬͇̬̻̬̣̟̟̣̬̯͓͇̜̥͈̪͓̳̩̤̹͙͓̝̖̳͙͚̼̩̬͇͒̏̆̄̀̆̄̏͛̔͐̋̉͊̃̀̀͐̍̇̇̿͆̍̓͗͗̐̔͊̊͋́̂͑̀͋̿̈́͊̋̍̿̒͐͂̒̂͐̆̓̇̈̍̍̿̄̽͌̑͑̈̇̈̑̾̍̋̉̑̆̓̊̑͂̓̍͂́̉̑̎̌́͆͂̂͒̑͌̇̓͛̎̌͆̄͛̋͒̓͒̉͒͛̿̾͛̑̆̍̽̎̆͒͋̅͐̌̔̾̑͊̐̉̾̍͒̎̔͑̅͊͂͗͗̒̔̊͛̈̓̃̒̌̀̿̽̍̆̎̉̆̚͘̕̕̚̚̕͘͜͜͜͜͜͠͠͝͝͝͠͝͝͝͝͠͠͝ͅͅͅͅͅͅͅͅǫ̷̨̨̡̨̨̨̡̢̨̡̨̨̨̨̨̢̢̧̢̛̛̛̛̛̛̛̼͈͈̘̥̠̟̼̹̪̻͕͔̻͙̣̦̖̰̣̪͈̦̼̺̳̮̱̥͇̰̮͖͙̯̱̩̼͇̘̝͈̙̻͙̟͙̫̤̪͖̮̱̲̰̥̘͖̖̠̱̹̥̥̗͕̗̦̼̖͚͍̭͖̙̤͎̼̰̠͎̰͍͓͍̰͖͍̞̗̼̤̺͈͔͔̭͎̭̠̜̼͍͙̠̺͓̮̮̳̲͇̮͕̤̮̳̹͕̘͚̫̰̯͔̩͍̩̠̯̩͓̣̹̻͈͖̠͉͔̼̣͈̘̜͕͔͖̝̯̮̙̮͔̣̬̫̱̘̱̝̥̘̠͓͍͇̦̰̭̮̝͈̤͔̦͓̳̉̎͐̾͛͛̅̀͑̋̉͋̓̔̔̈̑̀̅̆̈́̔̃̆̈́̉̉́̌͆̉͐͌͐̉̊͂̍̏̂̔̋̌͊͂͂̂̓̄̑̅͑́̃̆̀̎͊̂̉̾͋̿͆̇̃̔͒͗̿͋̋̌̿̋̑̑͗̎̓́͂͛́̑̔́̃̾̅́̉͒̀̄͑̏̉̑͋̈͐̇̒̾͒̓̍̂̄̂̾̐͌͋̒͑͊͒͛͋͆̆̎͐̑̋͒͐̐̈͑͆͂͐̐̔̐̅͗̃̐̏̐̏̓̎͆́̆͐̀̉̌̏̑̊͊̓͌̄̋͋̾̾̾͂͋͒̋̔͑̃̑̒̌̃̿̽͌̎̊̔̀͐͛̋̋̐͌̄̆̆̌̃̊͛̈͗̀͊̽̏̽̏̂̔͛̒͗̇̄̐̾̽͘̚̚̕̚̕̚̚̚͘̕̕̚̕̕̕̕̕͘̕͜͜͝͠͝͝͝͝͠͝͠͝͝͝͝͝͠͝ͅͅͅͅͅ'.length
1772
```

To correctly get the right number of characters, you can use the string's `String.prototype[@@iterator]()` method.

```js
> [...'💻'].length
1
```

Difference between [`charAt()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/charAt), [`codePointAt()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/codePointAt), and [`charCodeAt()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/charCodeAt)

Resources by good authors:

https://mathiasbynens.be/notes/javascript-unicode

Resources of unsure quality:

- https://flaviocopes.com/javascript-unicode/
- https://kevin.burke.dev/kevin/node-js-string-encoding/
- https://dmitripavlutin.com/what-every-javascript-developer-should-know-about-unicode/

---

`content-type: text/html; charset=UTF-8`

---

Collation is the standardized ordering of written information. Collation is related to code points because many collation uses the numeric order of the code points to order text.

---

UTF-8

> UTF-8 was designed on a placemat in a New Jersey diner one night in September 1992. (https://www.cl.cam.ac.uk/~mgk25/ucs/utf-8-history.txt)

https://utf8everywhere.org/
https://www.youtube.com/watch?v=MijmeoH9LT4

---


https://2ality.com/2013/09/javascript-unicode.html

---

Unicode security

https://www.unicode.org/reports/tr36/
https://www.unicode.org/reports/tr39/
https://websec.github.io/unicode-security-guide/
