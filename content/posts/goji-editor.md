---
title: The Story of the Goji Editor
slug: goji-editor
date: 2023-02-07T20:36:57.974Z
publishDate:
tags: []
draft: true
---

Over the last 6 months, I have been working full-time on a personal project that I have been itching to do for the last few years. It's called Goji and it's a self-learning platform that allows users to organize their notes as a series of dependent notes (i.e. B requires A), ultimately forming a directional [graph](https://en.wikipedia.org/wiki/Graph_(discrete_mathematics)). There's a lot more to it than that, but I'll let the product speak for itself when the time comes.

I've tried writing using WYSIWYG editors like the ones that power Notion and Medium, and it feels very smooth and polished. Compare that to Markdown editors like [Boost Note](https://boostnote.io/), which focuses more on function than on form (which is not a bad thing for its target audience).

I want to design a markup editor (for a custom markup language called GojiMark) with real-time syntax highlighting, but with the polished-look and experience of the Notion and Medium editors. Specifically, I wanted these features:

- When a user types some special character (Notion users commonly uses the [slash command](https://www.notion.so/Learn-the-shortcuts-66e28cec810548c3a4061513126766b0) <kbd>/</kbd>), a popup will appear that helps you fill in the markup. For example, GojiMark uses double braces (`[[` and `]]`) to mark an internal link. Between the brace should be the UUID of the note it's referring to, as well as an optional link title (i.e. `[[e36afb77-8762-4e17-8bba-c61bc678ff4a|API]]`, this is inspired by [free links](https://en.wikipedia.org/wiki/Help:Wikitext#Free_links) in [Wikitext](https://en.wikipedia.org/wiki/Help:Wikitext)). But having the user find the referred note, copy the UUID, and paste it into referring note is not a good user experience. With this feature, once the editor types the opening braces `[[`, a popup will appear that allows them to search/filter through existing notes; when they select the one they were looking for, then the editor will help them populate the stuff between the double braces.

  ![](/img/notion-editor-slash-command.png)

- Clean, no clutter feel

  ![](/img/medium-editor-blank.png)

- Allow users to rearrange their notes (when it's part of a series) using some kind of drag and drop (D&D) mechanism.

  ![](/img/notion-editor-dnd.png)

Having a look around, I couldn't find any open-source editors with a permissive license out there that supports all the features I wanted. The closest I came to was this self-professed [Notion Clone](https://github.com/konstantinmuenster/notion-clone) (see [demo](https://notion-clone.kmuenster.com)). It also uses [`react-contenteditable`](https://github.com/lovasoa/react-contenteditable), which I think unnecessarily introduce React in an area where it's not needed.

In the end, I _wanted_ to write my own editor because I _believe_ I can do a better job than what is available from the open-source space. I doubt it will become something as polished as the Notion editor (backed by a multi-billion dollar company with hundreds of employees), but better than any off-the-shelf solution.

# Syntax Highlighting

# contentEditable

Every browser implements contentEditable differently. Even if two contentEditables _look_ the same when rendered, the underlying HTML can be different.

Let's demonstrate this using the following web page:

```html
<!DOCTYPE html>
<html lang="en">
<body>
  <div id="target" contenteditable></div>
</body>
</html>
```

We will carry out the same set of operations inside the contentEditable on different browsers, and run `document.getElementById("target").outerHTML` to get the underlying HTML and see how they differ.

For the first test case, we will perform the following steps (Operation A):

1. Click on the contentEditable `div`
2. Type <kbd>f</kbd>
3. Type <kbd>o</kbd>
4. Type <kbd>o</kbd>

It turns out that Firefox, Safari, and Chrome all produced the same HTML markup.

```html
<div id="target" contenteditable="">foo</div>
```

So let's try a different set of operations (Operation B):

1. Click on the contentEditable `div`
2. Type <kbd>f</kbd>
3. Type <kbd>o</kbd>
4. Type <kbd>o</kbd>
5. Type <kbd>Return</kbd>

This time, some differences surface:

- Firefox

  ```html
  <div id="target" contenteditable="">
    <div>foo</div>
    <div><br></div>
  </div>
  ```

- Safari and Chrome

  ```html
  <div id="target" contenteditable="">
    foo
    <div><br></div>
  </div>
  ```

This little experiment demonstrates that browsers process user input into contentEditable differently. This means for our editor to work consistently across different browsers, the editor would have to ignore these differences (i.e. treat a bare text node the same way as one that is wrapped in a `div`).

Let's try another set of operations (Operation C):

1. Click on the contentEditable `div`
2. Type <kbd>f</kbd>
3. Type <kbd>o</kbd>
4. Type <kbd>o</kbd>
5. Hold <kbd>Shift</kbd> and type <kbd>Return</kbd>

This time, we get the following markup:

- Firefox and Chrome

  ```html
  <div id="target" contenteditable="">
    foo
    <br>
    <br>
  </div>
  ```

- Safari

  ```html
  <div id="target" contenteditable="">
    foo
    <div><br></div>
  </div>
  ```

Visually, what shows on the screen is the same regardless of whether you hold <kbd>Shift</kbd> or not, but the underlying HTML is different. This adds an additional constraint on our editor whereby it has to treat `<div>foo</div><div><br></div>` the same way as it treats a `foo<br><br>`.

Since what our editor stores is not the visual representation but just the markup text (the syntax highlighting is just to make the experience more intuitive), our editor needs to normalize the HTML before it passes it to the parser.

## ContentEditable To String

There's a Web API - the [`HTMLElement.innerText`](https://html.spec.whatwg.org/multipage/dom.html#the-innertext-idl-attribute) API - that's designed to give you a representation of text output as users would see it ( "as rendered"). It's roughly what you'd get if you highlighted the text and copied it to your clipboard. Because `HTMLElement.innerText` is based, roughly, on the visual representation of the text's layout, it is a good way to normalize the textual content of the contentEditable, and roll over the differences in the underlying HTML.

With `HTMLElement.innerText`, any type of new line (with or without holding the <kbd>Shift</kbd> key) is represented by a single newline character `"\n"` (this is good for normalizing newlines). A blank line is encoded as two consecutive newline characters (`"\n\n"`).

If we carry out Operations B and C and run `document.getElementById("target").innerText`, you'll see the output is identical:

- Firefox Operation B - `"foo\n\n"`
- Firefox Operation C - `"foo\n\n"`
- Chrome Operation B - `"foo\n\n"`
- Chrome Operation C - `"foo\n\n"`

### Whitespaces and Newlines

But life in browser-land is never as simple as it seems. As useful as `HTMLElement.innerText` is, it can't normalize every difference between how browsers work with contentEditables.

To demonstrate, carry out the following operation (Operation D):

1. Click on the contentEditable `div`
2. Type <kbd>f</kbd>
3. Type <kbd>o</kbd>
4. Type <kbd>o</kbd>
4. Type <kbd>Space</kbd>

You'll get the following HTML (the space character is replaced with `·` for better visualization) and the corresponding `innerText`:

- Firefox - `"foo\n"`

  ```html
  <div id="target" contenteditable="">
    foo·
    <br>
  </div>
  ```

- Chrome - `"foo "` (that whitespace character is a non-breaking space (NBSP, U+00A0))

```html
<div id="target" contenteditable="">
  foo&nbsp;
</div>
```

It seems like Firefox adds an additional `<br>` element to the end of the `<div>` if the text node ends with whitespace. Chrome converts the space (U+0020) to a non-breaking space (NBSP, U+00A0). Why do these browsers need to do that? Why not just keep it as `foo `?

Well, it has to do with how browsers treat whitespaces and newlines in HTML. If you load a HTML page with the following paragraph element in the body:

```html
<p> Hello   World!   </p>
```

You'll see the text `Hello World!` rendered on screen, _without_ the leading and trailing spaces, and the multiple spaces between the worlds `Hello` and `World` has also been collapsed to a single space. This quirky behavior is defined in the [CSS Text Module Level 3](https://www.w3.org/TR/css-text-3/) specification, specially [Section 3 - White Space and Wrapping: the white-space property](https://www.w3.org/TR/css-text-3/#white-space-property) and [Section 4 - White Space Processing & Control Characters](https://www.w3.org/TR/css-text-3/#white-space-processing).

Long story short, there's a CSS property called `white-space` that specifies how whitespaces (including spaces, tabs, line breaks, and other space separators such as non-breaking space) are to be handled. Valid values for `white-space` includes:

- `normal` (default) - multiple consecutive whitespaces are _collapsed_ into a single whitespace (as we saw with the whitespace between the word `Hello` and `World`). End-of-line spaces, tabs and newlines are removed, _but other space separators are preserved_. If a line is longer than its container, the text will wrap at an appropriate _soft-wrap opportunity_, except for trailing space separators.
- `nowrap` - same as `normal`, but if a line is longer than the container, it will not wrap and overflow.
- `pre` - all whitespaces are preserved and not collapsed or removed. Text do not wrap.
- `pre-wrap` - all whitespaces are preserved and text wrap. If the trailing spaces exceed the length of the container, the trailing spaces will _not_ wrap.
- `pre-line` - same as `normal` but newlines are preserved
- `break-spaces` - same as `pre-wrap` but trailing whitespaces are treated like normal characters and will wrap.

Since `whitespace: normal` is the default, when we typed `foo ` into our contentEditable, the browser would remove the end-of-line whitespaces and it won't be displayed. This behavior would be a pretty unintuitive for contentEditable. So the browsers try to get around this quirk differently. Chrome converts trailing spaces into non-breaking spaces, which are _not_ removed. With Firefox, I believe it adds a `<br>` element at the end of the `<div>` so the trailing spaces are no longer considered to be end-of-line.

> I could be wrong about why a `<br>` element is added after the text node. But my reasoning is because a `<br>` element at the end of a `<div>` (a block element) doesn't actually change the layout in any way, and can be ignored. So it is safe to add it before the closing `</div>` tag.

Because `HTMLElement.innerText` tries to mimic what the user sees on the screen, it will also apply the CSS `white-space` rules. In Chrome, the non-breaking space is treated like a normal character and so you get `"foo "` (the space is a non-breaking space). It's again a bit more quirky on Firefox - it takes the `<br>` element and convert it into a single `"\n"`. If the last character of the text node before the `<br>` element is a space, `innerText` will ignore it. This is why carrying out Operation D on Firefox gives you `"foo\n"`.

But the quirkiness on Firefox does not end here. If you then delete that trailing space, the `<br>` element is not removed. Try the following operation (Operation E):

1. Click on the contentEditable `div`
2. Type <kbd>f</kbd>
3. Type <kbd>o</kbd>
4. Type <kbd>o</kbd>
4. Type <kbd>Space</kbd>
4. Type <kbd>Backspace</kbd>

You'll get the following HTML:

```html
<div id="target" contenteditable="">foo<br></div>
```

And the `innerText` output is `"foo\n"`.

This means with the `innerText` implementation on Firefox, there can be cases where `foo` (with no trailing space) and `foo ` (with a trailing space) produces the same `innerText` output. This is a real blow in our effort for using `innerText` to process the contentEditable - our process requires different visual representations to provide different textual representations.

But we can actually forget about all these whitespace quirks if we change the `white-space` CSS property to `pre-wrap`, which is the same as `white-space: normal` except all whitespaces are preserved.

So let's update the contentEditable to the following:

```html
<div id="target" contenteditable style="white-space:pre-wrap"></div>
```

Now, if you carry out Operation D and run `document.getElementById("target").innerText`, you'll get `"foo "` (that whitespace is a normal space (U+0020)) on both Firefox and Chrome. Similarly, Operation E yields `"foo"` on both browsers.

Therefore, the takeaway lesson in this section is that when whitespace matters within a contentEditable, make sure that `white-space` is set to `pre-wrap` (or less commonly `break-space`) to preserve the whitespace.

### `<div>` vs `<p>` ContentEditables

The `contenteditable` attribute is a global attribute which means it can be set on any HTML element. But how a browser process changes within the contentEditable element depends on the type of the element. Block elements which normally holds inline elements (I'll call them leaf blocks from now on) behaves different those that acts as containers for other block elements (I'll call them container blocks).

Let's change our the type of our contentEditable element from `<div>` to `<p>` and repeat Operation B to see if there's a difference.

The `innerText` output is with the `<p>` elements is ``"foo\n\n"``, and the HTML looks like this:

```html
<p id="target" style="white-space:pre-wrap" contenteditable="">
  foo
  <br>
  <br>
</p>
```

With the `<div>` element, the the `innerText` output is the same (`"foo\n\n"`) and the HTML looks like this:

```html
<div id="target" style="white-space:pre-wrap" contenteditable="">
  <div>foo</div>
  <div><br></div>
</div>
```

But let's try a more complicated set of operations (Operation F):

1. Click on the contentEditable `div`
2. Type <kbd>f</kbd>
3. Type <kbd>o</kbd>
4. Type <kbd>o</kbd>
5. Type <kbd>Return</kbd>
6. Type <kbd>Return</kbd>
7. Type <kbd>b</kbd>
8. Type <kbd>a</kbd>
9. Type <kbd>r</kbd>

On Firefox, we get the `innerText` value of `"foo\n\n\nbar\n"` and the following HTML:

```html
<div id="target" style="white-space:pre-wrap" contenteditable="">
  <div>foo</div>
  <div><br></div>
  <div>
    bar
    <br>
  </div>
</div>
```

If we change the contentEditable element to a `<p>` element, we get the `innerText` value of `"foo\n\nbar\n"` and the following HTML:

```html
<p id="target" style="white-space:pre-wrap" contenteditable="">
  foo
  <br>
  <br>
  bar
  <br>
</p>
```

Note how with the `<div>` element, there are 3 newline characters between `foo` and `bar`, but with the `<p>` element, there are only 2. This is because with the `<div>` contentEditable, each block element (such as `<div>foo</div>`) is assumed to be styled to span the entire line, so visually, it would be as if there's a newline at the end of the text content. What complicates this somewhat is that a blank line is encoded as `<div><br></div>` within a `div`, because a `<div>` is a container for other block elements, whereas `<br>` is an inline element and so the browser wraps it in a `<div>`. So a blank line within a `<div>` contentEditable is `\n\n` - the first newline as substitute for the `<br>` element, and the second for the `<div>`.

Since our editor will only be dealing with text, it makes more sense to use a `<p>` contentEditable since it deals with a blank line in a more intuitive way. The reason why many contentEditables use a `<div>` is because it needs to support more than just text content, and it'd not be semantically correct to have, say, a `<div>` or `<table>` element within the contentEditable.

> This is a great relief, actually, because having to code a cross-browser-compatible editor that supports non-text elements is not trivial. The way the Medium editor deals with this cross-browser inconsistency is to [create an abstraction on top of the DOM](https://medium.engineering/why-contenteditable-is-terrible-122d8a40e480). They call this abstraction the Medium Editor Model. Essentially, when a user types something into the editor, that event is captured and the Medium Editor Model is updated, then the HTML representation of the model is rendered. This basically means that Medium uses their own implementation of contentEditable.
> But it would be crazy to write a program to capture and process _every_ possible keyboard event, given there are hundreds of thousands of possible characters in the many writing systems of the world. Instead, the Medium editor treats some user-initiated events specially (carriage returns, deletion, type-over (selecting text and typing over them), and paste), and leaving other events to default to the behavior of the native contentEditable.

## Pasted Content

Now that we've settled on using a `<p>` element as our contentEditable, we need to make sure that when someone paste something into the contentEditable that is not plain text (e.g. images, tables, etc.), that we convert it plain text first.

## Parsing

## Setting Caret

