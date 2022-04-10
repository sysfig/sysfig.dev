SVG, short for Scalable Vector Graphics, is a image format for 2D vector images where the data is stored as an Extensible Markup Language (XML) file.

Here's a simple SVG file that produces the image below.

```xml
<svg style="border:1px solid yellow" width="100" height="100">
    <circle r=10 fill='white' cx=30 cy=35 />
    <circle r=10 fill='white' cx=70 cy=35 />
    <path d="M 30 60 C 30 80 70 80 70 60" fill=none stroke="white" stroke-width="6" />
  </g>
</svg>
```

<svg style="border:1px solid yellow" width="100" height="100">
    <circle r=10 fill='white' cx=30 cy=35 />
    <circle r=10 fill='white' cx=70 cy=35 />
    <path d="M 30 60 C 30 80 70 80 70 60" fill=none stroke="white" stroke-width="6" />
  </g>
</svg>

# Short Introduction to XML in the Context of SVG

XML is a markup language where each element is represented as a pair of _opening_ and _closing tags_.

```xml
<element-name></element-name>
```

For example, an element named `g` can be represented in XML as:

```xml
<g></g>
```

Each element can have zero or more _attributes_.

```xml
<element-name attribute-name="attribute-value"></element-name>
```

For example, we can set the `r` attribute of the `circle` element to `10` using the code below:

```xml
<circle r="10"></circle>
```

In SVG, the types of attributes you can assign are different for different elements. For example, the `circle` element has the `r` attribute that represents the radius of the circle, but the `rect` element (representing a rectangle) does not have the `r` attribute.

Some attributes, called _global attributes_, can be applied to all SVG elements:

https://developer.mozilla.org/en-US/docs/Web/SVG/Element/g#global_attributes

- `id` - a unique identifier for a specific element
- `class` - a group identifier that can be shared by multiple elements
- `style` - inline CSS styles
- `transform` - https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/transform
  - `matrix`
  - `translate`
  - `scale`
  - `rotate`
  - `skewX`
  - `skewY`
- _presentation attributes_:
  - `stroke`
  - `stroke-width`
  - `fill`

These global attributes work in the same way as HTML.

In XML, just like HTML, elements are organized in a tree-like structure where each element can contain one or more children. Children are placed between the opening and closing tags of the parent element.

```xml
<parent>
  <child></child>
  <child></child>
</parent>
```

For example, in the code below, the `g` parent element has two children - `circle` and `rect`:

```xml
<g>
  <circle r="10"></circle>
  <rect width="20" height="20"></rect>
</g>
```

Any transformations and styles applied to the parent (using `transform` and `style` attributes) are inherited by its children.

For sibling elements, the first element that appears are drawn first. New elements are drawn _on top of_ existing elements.

> The SVG 2.0 specification had once toyed with supporting a `z-index` attribute that allows you to override the order in which elements are drawn, but this plan has since been shelved.

# The Elements

Now we know how elements are represented (in XML), let's take a look at the kinds of elements available in SVG:

- `<svg>` - the root element, similar to `<html>` in HTML.
- `<g>` - the group/container element. It groups multiple elements together, which allows you to manipulate its children (e.g. translate them) in the same way. This is similar to `<div>` in HTML.
- `<rect>`
- `<circle>`
- `<ellipse>`
- `<line>` - draws a straight line between 2 points
- `<path>` - draws any arbitrary path between 2 points
- `<text>` - 
  - `<tspan>`
- `<clipPath>` - a container for child elements where the area the children occupies can be used as a . An element can then use the `clipPath` element as a clip by specifying `clipPath` element via the `clip-path` attribute.
- `<defs>` - defines a custom element which can be reused in other parts of the document

 <a>, <foreignObject>, <image>, <polygon>, <polyline>,  <switch>, and <use>

https://developer.mozilla.org/en-US/docs/Web/SVG/Element

## The Root `<svg>` Element

The root `svg` element defines a new _SVG viewport_ (note an SVG viewport is different from the browser 'viewport' used in CSS).

This rectangular SVG viewport has a width and a height (set by the `width` and `height` attributes), just like any other HTML element. This width and height are used to externally position the `<svg>` element in relation to the rest of the HTML document.

However, each viewport also has an _internal_ Cartesian (a.k.a. two-dimensional) _user coordinate system_. In this coordinate system, the positive x direction is to the right; the positive y direction is down. This is unlike mathematical graphs, where y's positive direction is up.

By default, the origin (i.e. position `0,0`) is at the top-left corner of the SVG image. Also by default, the size of the coordinate system matches the size of the image. This means an SVG image of width and height of `100` has a coordinate system that starts at `0,0` at the top-left, and ends at `100,100` at the bottom-right.

But these default settings for the coordinate system can be overridden by using the `viewBox` attribute on the `<svg>` element.

The `viewBox` attribute is a space-or-comma-separated string of 4 numbers:

- `min-x` - defines the x-coordinate on the top-left of the SVG image, defaults to `0`.
- `min-y` - defines the y-coordinate on the top-left of the SVG image, defaults to `0`.
- `width` - defines the width of the coordinate system in _user units_, all elements within the `svg` will use these coordinates when positioning and drawing shapes.
- `height` - defines the height of the coordinate system in _user units_, all elements within the `svg` will use these coordinates when positioning and drawing shapes.

Let's go back to the smiley face SVG file we showed you at the beginning.

```xml
<svg style="border:1px solid yellow" width="100" height="100">
    <circle r=10 fill='white' cx=30 cy=35 />
    <circle r=10 fill='white' cx=70 cy=35 />
    <path d="M 30 60 C 30 80 70 80 70 60" fill=none stroke="white" stroke-width="6" />
  </g>
</svg>
```

Since the SVG viewport has a width and height of `100`, the `viewBox` value defaults to `0 0 100 100`. We can add this attribute in and still obtain an identical image.

```xml
<svg style="border:1px solid yellow" viewBox="0 0 100 100" width="100" height="100">
    <circle r=10 fill='white' cx=30 cy=35 />
    <circle r=10 fill='white' cx=70 cy=35 />
    <path d="M 30 60 C 30 80 70 80 70 60" fill=none stroke="white" stroke-width="6" />
  </g>
</svg>
```

<svg style="border:1px solid yellow" viewBox="0 0 100 100" width="100" height="100">
    <circle r=10 fill='white' cx=30 cy=35 />
    <circle r=10 fill='white' cx=70 cy=35 />
    <path d="M 30 60 C 30 80 70 80 70 60" fill=none stroke="white" stroke-width="6" />
  </g>
</svg>

The benefit of having an internal user coordinate system is that you can draw elements within the `<svg>` relative to the defined user coordinate system, without having to care about the size of the `<svg>` element.

We can demonstrate this by removing the `viewBox` attribute and changing the `width` and `height` of the `<svg>` element (the viewport) to `200`.

```xml
<svg style="border:1px solid yellow" width="200" height="200">
    <circle r=10 fill='white' cx=30 cy=35 />
    <circle r=10 fill='white' cx=70 cy=35 />
    <path d="M 30 60 C 30 80 70 80 70 60" fill=none stroke="white" stroke-width="6" />
  </g>
</svg>
```

<svg style="border:1px solid yellow" width="200" height="200">
    <circle r=10 fill='white' cx=30 cy=35 />
    <circle r=10 fill='white' cx=70 cy=35 />
    <path d="M 30 60 C 30 80 70 80 70 60" fill=none stroke="white" stroke-width="6" />
  </g>
</svg>

Now the coordinate system has doubled in size (it's largest coordinate is now `200,200`); but instead of the `<circle>` and `<path>` elements resizing with the `<svg>` element, they've stayed in place.

However, if we re-introduce the `viewBox` attribute, then the coordinate system will always have the same size, regardless of how big the `<svg>` element is from the perspective of the HTML document, meaning the elements within the viewport resizes with the viewport.

```xml
<svg style="border:1px solid yellow" viewBox="0 0 100 100" width="200" height="200">
    <circle r=10 fill='white' cx=30 cy=35 />
    <circle r=10 fill='white' cx=70 cy=35 />
    <path d="M 30 60 C 30 80 70 80 70 60" fill=none stroke="white" stroke-width="6" />
  </g>
</svg>
```

<svg style="border:1px solid yellow"  viewBox="0 0 100 100" width="200" height="200">
    <circle r=10 fill='white' cx=30 cy=35 />
    <circle r=10 fill='white' cx=70 cy=35 />
    <path d="M 30 60 C 30 80 70 80 70 60" fill=none stroke="white" stroke-width="6" />
  </g>
</svg>


The `viewBox` scales to fit _within_ the viewport. We can demonstrate this by stretching the viewport into a rectangle but keeping the `viewBox` a square.

```xml
<svg style="border:1px solid yellow" viewBox="0 0 100 100" width="400" height="200">
    <rect width="100%" height="100%" fill="red" />
    <circle r=10 fill='white' cx=30 cy=35 />
    <circle r=10 fill='white' cx=70 cy=35 />
    <path d="M 30 60 C 30 80 70 80 70 60" fill=none stroke="white" stroke-width="6" />
  </g>
</svg>
```

<svg style="border:1px solid yellow" viewBox="0 0 100 100" width="400" height="200">
    <rect width="100%" height="100%" fill="red" />
    <circle r=10 fill='white' cx=30 cy=35 />
    <circle r=10 fill='white' cx=70 cy=35 />
    <path d="M 30 60 C 30 80 70 80 70 60" fill=none stroke="white" stroke-width="6" />
  </g>
</svg>

```xml
<svg style="border:1px solid yellow" viewBox="0 0 100 100" width="200" height="400">
    <rect width="100%" height="100%" fill="red" />
    <circle r=10 fill='white' cx=30 cy=35 />
    <circle r=10 fill='white' cx=70 cy=35 />
    <path d="M 30 60 C 30 80 70 80 70 60" fill=none stroke="white" stroke-width="6" />
  </g>
</svg>
```

<svg style="border:1px solid yellow" viewBox="0 0 100 100" width="200" height="400">
    <rect width="100%" height="100%" fill="red" />
    <circle r=10 fill='white' cx=30 cy=35 />
    <circle r=10 fill='white' cx=70 cy=35 />
    <path d="M 30 60 C 30 80 70 80 70 60" fill=none stroke="white" stroke-width="6" />
  </g>
</svg>

The red rectangle indicates the size of the `viewBox` whilst the yellow border outlines the shape of the viewport.

By default, when the `viewBox` and the viewport have a different aspect ratio, the `viewBox` will scale up as much as possible whilst still fitting within the viewport, and then center-aligned.

However, this behavior can be changed using the `preserveAspectRatio` attribute.

https://css-tricks.com/scale-svg/
https://www.sarasoueidan.com/blog/svg-coordinate-systems/

## Rectangles

- `x` - x-coordinate of the top-left
- `y` - y-coordinate of the top-left
- `width`
- `height`
- _presentation attributes_

## Lines

- `x1`
- `y1`
- `x2`
- `y2`
- _presentation attributes_

## Ellipses

- `cx` - x-coordinate of the center
- `cy` - y-coordinate of the center
- `rx` - horizontal radius
- `ry` - vertical radius
- _presentation attributes_

## Paths

Drawing with only straight lines, rectangles and ellipses is quite limiting. So for everything else there's the `<path>` element. With the `<path>` element, you can draw everything we've just described, and more.

The `<path>` element uses the `d` attribute to specify how the path should be drawn. Specifically, the `d` attribute contains a list of _commands_ for each path segment.

Each command begins with a letter, which can be uppercase or lowercase. An uppercase letter indicates that the coordinates specified within the command are _absolute_ (they are relative to the origin (i.e. `0,0`, the top-left corner of the `<svg>`). A lowercase letter indicates that the coordinates specified within the command are relative to ??end location?? of the previous command.

- `M` (move) - `x y` (e.g. `M 0 0`)
- `L` (line) - draw a straight line from the current position to the specified x- and y- coordinates (e.g. `L25,50`)
- `A` (arc) - rx ry rotation arc-flag sweep-flag x y
- `Q` (quadratic bezier) - x1 y1 x2 y2 - it's call the quadratic bezier because it uses the previous position and 2 points (see https://pomax.github.io/bezierinfo/)
- `C` (cubic bezier) - x1 y1 x2 y2 x3 y3 - it's call the quadratic bezier because it uses the previous position and 3 points (see https://pomax.github.io/bezierinfo/)
- `Z` - close the path to form a shape (this will join the start and end of the path with a straight line)

> Here, we used spaces to separate different parameters of the same command. Other people and software uses commas. Both are equivalent (e.g. `M 0 0` is the same as `M0,0` and `M 0,0`)
> Note also that no separators are needed between the command letter and its parameter, nor _between_ commands(e.g. `M 0 0 L 10 10` can be shorted to `M0 0L10 10`). Removing superfluous characters can reduce the size of the file but can also makes it harder for humans to read.

## Text

- `x` - 
- `y`
- `dominant-baseline` - https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/dominant-baseline
- `text-anchor` - https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/text-anchor

![](https://svg-tutorial.github.io/images/text-position.svg)

### `<tspan>`

The `tspan` element allows you to apply different styles to text within a `text` element.

- `dx`
- `dy`

### Fonts

If the font does not exist on the viewer's computer, then a fallback font would be used (the fallback font is configurable by the viewer).

To avoid the font not being available, you should:
- provide a list of fonts instead of just one
- use a web-safe font like Arial, Times New Roman, Courier, etc.

## Clip Path

"""
```xml
<clipPath id="my-clip-path">
  <path d="M 50 50 L 250 50 L 250 250 L 50 250 Z"/>
</clipPath>
<image clip-path="url(#my-clip-path)" href="..."/>
```
"""(https://svg-tutorial.github.io)

# Embedding SVGs

There are two ways to embed an SVG image onto a webpage:

- Via the `src` attribute of an `img` element:
  - cannot select elements within the SVG image, this means:
    - stylesheets will not affect the styling of the SVG, even if it matches the `id` and `class` of the elements within the image.
    - JavaScript cannot be used to manipulate the SVG elements, meaning interactivity is not possible
  - if the SVG uses external fonts (e.g. Google Fonts), most browsers will not allow access to the external resource and thus a fallback font will be used ??VERIFY??
- Directly embed the SVG into the HTML document

---

All modern browsers (see https://caniuse.com/svg) can parse SVG elements as DOM elements, more specifically - SVG Document Object Model or SVGDOM. This means you are able to work with SVG elements in the same way as HTML elements - adding event handlers, programmatically changing attributes, style with CSS, animate with CSS and/or JavaScript etc.




gradients, rotations, filter effects, animations

---

[SVG Tiny 1.2](https://www.w3.org/TR/SVGTiny12/) - 
