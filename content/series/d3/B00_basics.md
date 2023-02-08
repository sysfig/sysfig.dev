---
title: Basics
slug: basics
date: 2023-02-07T20:49:51.230Z
publishDate:
chapter: a
order: 1
tags:
    - d3
draft: true
---

# Scaffolding

We are going to learn d3 from scratch, so let's start with a blank HTML5 scaffold. Create a new file that ends with `.html` (e.g. `d3.html`) and paste in the following.

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Understanding d3.js</title>
</head>
<body>
  <!-- Our code will go here-->
</body>
</html>
```

Open the file using a browser, you should see a blank screen.

d3 is a JavaScript library that exports an object called `d3`, we can make it available to our page by adding a `<script>` tag inside the `<head>` element.

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Understanding d3.js</title>
  <script src="https://d3js.org/d3.v6.min.js"></script>
</head>
<body>
  <!-- Our code will go here-->
</body>
</html>
```

To test that d3 is correctly added, we can try to run `console.log()` on `d3`.

```html
<body>
  <script>console.log(d3)</script>
</body>
```

After you've made the change, open the console of your browser's developer tools and refresh the page. You should see some output like:

```console
Object { format: l(t), formatPrefix: formatPrefix(t, n), timeFormat: format(t), timeParse: parse(t), utcFormat: utcFormat(t), utcParse: utcParse(t), Adder: class g, Delaunay: class nu, FormatSpecifier: uc(t), InternMap: class y, … }
```

This means d3 is successfully added to the `window` global object of the page, and we can start using it.

# Selecting Elements

In the next section, we are going to use d3 to add a `<p>` element to the DOM. But before we can do that, we must first learn how to _select_ elements from the page. This is because the DOM has a tree-link structure, which means every element apart from the root element (`<html>`) must exist _within_ another element. So we must first learn how to select the parent element that we will append new elements to.

The `d3` object has a `select` method, which allows you to select a single DOM element from the document. It accepts a single argument, which can either be:

- a [W3C selector](https://www.w3.org/TR/selectors-api/) string, like the ones used in CSS (e.g. `div.someClass`), or
- an [`Element`](https://developer.mozilla.org/en-US/docs/Web/API/Element) object (since `Element` is the parent interface of `HTMLElement` and `SVGElement`, you can pass any HTML or SVG element on the page)

For instance, we can select the `<body>` HTML element by passing in the CSS selector string `body`:

```js
d3.select("body")
```

Or by passing in the element object itself.

```js
d3.select(document.getElementsByTagName("body")[0])
```

`d3.select` returns a `Selection` object that represents the selected elements, as well as provide d3-specific methods such as:

- `append()` - to add a new element to each element in the selection
- `attr()` - to add or modify attributes of each element in the selection
- `on()` - to add event listeners on each element of the selection
- `select()` - to select child elements within the selection
- `size()` - to get the number of matched elements in the selection
- `text()` - to get or set the text within each matched element of the selection

So you can view `d3.select()` as a wrapper that adds d3-specific methods to the matching HTML element. To try out one of these methods, replace the `<body>` element in your `d3.html` file with the following:

```html
<body>
  <script>
    d3.select("body")
      .style("background-color", "teal");
  </script>
</body>
```

Save the file and refresh the page; you'll see that the background has now changed to a teal color.

`d3.select` selects only a single element. If multiple elements matches the selector string, `d3.select` will scan the entire page and return the first matching element it finds. If no elements matches, it will still return with a `Selection` object, but it will be _empty_ (the selection's `empty()` method returns `true`).

```js
d3.select("body").size() // 1
d3.select("script").size() // 1, even though we have 2 <script> tags, because `d3.select()` only selects the first matching element.
d3.select("notexists").size() // 0
d3.select("body").empty() // false
d3.select("notexists").empty() // true
```

There's another top-level selection method called `d3.selectAll`. Whereas `d3.select` returns the first matching element, `d3.selectAll` returns _all_ matching elements.

```js
d3.select("script").size() // 1
d3.selectAll("script").size() // 2
```

Similar to `d3.select`, `d3.selectAll` accepts either a selector string, or an array/psuedo-array/iterable of `Element`s.

> If you've used jQuery, `d3.selectAll` is similar to the `jQuery` function (often written as `$` and used like `$( "div.foo" )`).

Selections in d3 are a bit more complicated than this. For instance, a selection can include DOM elements which are not yet on the page (which will be added later). We'd like to avoid introducing unnecessary complexity too early, so for now, simply know that you can use `d3.select()` and `d3.selectAll()` to select elements from the DOM and provide them with d3-specific methods.

# Adding DOM Elements

Now we know how to select elements, let's select the `<body>` element and use the `Selection`'s `append()` method to append an `<p>` inside it. Replace the `<body>` element in your `d3.html` with the following:

```js
<body>
  <script>
    d3.select("body")
      .append("p");
  </script>
</body>
```

If we refresh the page and use our browser's inspector to check our document, we'll see that a `<p>` element has indeed been added to the document.

```html
<body>
  <script>
    d3.select("body")
      .append("p");
  </script>
  <p></p>

</body>
```

However, the `<p>` element is empty, which is not very useful. But how do we add some text to it?

It's useful to know that `append()` returns a new `Selection` object that represents the elements it's just appended to the DOM. So we can use the `Selection`'s `text()` method to add some text within the `<p>` element.

```js
d3.select("body") instanceof d3.selection // true
d3.select("body").append("p") instanceof d3.selection // true
```

Update the `<body>` of your `d3.html` document to the following:

```html
<body>
  <script>
    d3.select("body")
      .append("p")
      .text("Hello, World!");
  </script>
</body>
```

Save the file and refresh your browser. You should now see the text `Hello, World!` printed.

We should make it clear that the `append()` method returns a new `Selection` of _all_ the objects it's just appended. To understand what this means, let's use an example. Update the `<body>` once again with the following:

```html
<body>
  <div></div>
  <div></div>
  <script>
    d3.selectAll("div")
      .append("p")
      .text("Hello, World!");
  </script>
</body>
```

We have added two new `<div>` elements within the `<body>` and are using `d3.selectAll` to select both of them. `d3.selectAll` returns with a `Selection` object representing the two `<div>` elements. We are then calling the `append()` method on the `Selection` instance, and `append()` will act on each element of the `Selection` and append a `<p>` element within it. If we run the script up to this point, we would get the following result:

```html
<body>
  <div><p></p>
  </div>
  <div><p></p>
  </div>
  <script>
    d3.selectAll("div")
      .append("p")
      .text("Hello, World!");
  </script>
</body>
```

The `append()` method then returns with a _new_ `Selection` object representing the _two_ `<p>` elements we've just added. So when we call the `text()` method on this new `Selection`, the text is added within the `<p>` elements, and not the `<div>` elements.

Now that we know how to add elements and text nodes with `append()` and `text()`, let's see how we can add nodes based on data.

# Data

What makes d3 more than just a DOM manipulation library is its feature to manipulate the DOM _based on data_. For example, we can use d3 to draw a bar chart that charts the sales data for a particular month. If we then add a date-picker that allows users to pick which month's data they want to visualize, d3 can plot a new graph based on the new data from that month.

But let's start off with something simpler. Let's say we have an array of integers, representing the number of items sold on an online store.

```html
<body>
  <script>
    const data = [44847, 63823, 70593, 82203, 109742, 129422, 175832, 182371, 149292, 79420, 59382, 29318]
  </script>
</body>
```

To keep things as simple as possible, we want d3 to append a `<p>` element for each of the integers in the array. Inside each `<p>` element, we want a text node displaying the corresponding data value. In essence, we want to achieve the following result:

```html
<body>
  <script>...</script>
  <p>44847</p>
  <p>63823</p>
  <p>70593</p>
  <p>82203</p>
  <p>109742</p>
  <p>129422</p>
  <p>175832</p>
  <p>182371</p>
  <p>149292</p>
  <p>79420</p>
  <p>59382</p>
  <p>29318</p>
</body>
```

We will now show you how this can be achieved, and then we will go through each command step-by-step. It turns out, the following chain of d3 commands gets what we want.

```html
<body>
  <script>
    const data = [44847, 63823, 70593, 82203, 109742, 129422, 175832, 182371, 149292, 79420, 59382, 29318];
    d3.select("body")
      .selectAll("p")
      .data(data)
      .join("p")
      .text(d => d);
  </script>
</body>
```

First, we use `d3.select("body")` to get a d3-representation of the `<body>` element, just like before.

The next `selectAll()` method is actually subtly different to the `d3.selectAll()` we have seen before. Whereas the `selectAll()` method of the `d3` object (i.e. `d3.selectAll()`) scans the entire document looking for matching elements, the `selectAll()` method of a selection (i.e. `d3.selection.prototype.selectAll()`) looks for matching elements only within each element of the selection. In this case, `d3.select("body").selectAll("p")` means "return a selection of `<p>` elements within the `<body>` element.

But we currently have no `<p>` elements within the `<body>` element, so this `selectAll` call will return with an empty selection.

```js
d3.select("body")
  .selectAll("p")
  .empty(); // true
```

So what good is an empty selection? Here's a good time to introduce an important caveat about selections. Unlike JavaScript's native `querySelectorAll()` or jQuery's `$`, selections in d3 not only selects elements that currently exists in the document, but also matching elements that we may add in the future.

With this in mind, we then call the `data()` method of this empty selection. The `data()` method binds an array of data with the selected elements; in our case, we are binding our `data` array of integers with each `<p>` element in our selection. But since our selection is empty, there's nothing for the data to bind to, so all our attempts at binding are unsuccessful.

In general, when binding an array of data to an array of selected elements, one of three things can happen to each element/datum:

- a selected element exists for a given datum, in which case binding is successful
- a selected element exists, but no data is available to bind to it, in which case binding is unsuccessful. This will happen when the size of the selection is larger than the size of the data array.
- a given datum exists, but no selected element is available to bind to it, in which case binding is unsuccessful. This will happen when the size of the data array is larger than the size of the selection.

Remember back when we covered the `append()` method, we said that it returns with a selection. Well, when you call `data()`, d3 actually creates _three_ different selections:

- _Update_ selection - containing all existing elements that has successfully bound to data
- _Exit_ selection - containing all existing elements which has no data to bind to
- _Enter_ selection - containing all the elements which _does not yet exist_, but for which data is available

The names of these selections makes more sense when you think about how d3 will update the DOM if we later change the data we pass in.

If the element (e.g. a `<rect>` element representing a bar in a bar chart) exists but the data has changed (e.g. if we change the arbitrary value of `5` to `10` ), then d3 doesn't need to remove and re-create the element, it only needs to update the element's `width` attribute. Hence, all existing elements for which existing data exists but needs updating, they are added to the update selection.

On the other hand, an existing element bound to an existing datum should be removed if that datum no longer exists in the new data. Think about a timelines chart of all animals in an animal rescue center. If an animal gets adopted, its datum would be removed from the data. Thus, the corresponding element no longer has a datum bound to it and should be removed. It is part of the exit selection since the element needs to 'exit the stage'.

By the same logic, if a new datum is added to the data array, a new element should be added to the DOM to represent it. When this happens, d3 adds a placeholder element into the enter selection, because a new element is about to 'enter the stage'.

When we call `data()` on a selection and bind some data to it, `data()` returns with the _update selection_. We can illustrate this with the following example:

```html
<body>
  <p></p>
  <p></p>
  <script>
    const data = [44847, 63823, 70593, 82203, 109742, 129422, 175832, 182371, 149292, 79420, 59382, 29318];
    d3.select("body")
      .selectAll("p")
      .data(data)
      .size(); // 2
  </script>
</body>
```

Because we have two `<p>` elements in our document, `d3.select("body").selectAll("p")` has a size of 2. When we try to bind our `data` array (with 12 elements), only 2 bindings are successful because there are only 2 elements. Therefore, the first two data goes into the update selection that is returned, whilst the 10 unbound data goes into the enter selection. We then call the `size()` method of the returned update selection to yield an output of `2`.

But what if we want to access the enter or exit selections from `data()`? Well, the selection returned from `data()` has an `enter()` and `exit()` method that allows you to retrieve the corresponding selection.

We can once again illustrate this using the following example:

```html
<body>
  <p></p>
  <p></p>
  <script>
    const data = [44847, 63823, 70593, 82203, 109742, 129422, 175832, 182371, 149292, 79420, 59382, 29318];
    const selection = d3.select("body")
      .selectAll("p")
      .data(data)
    selection.size(); // 2
    selection.enter().size(); // 10
    selection.exit().size(); // 0
  </script>
</body>
```

`data()` splits the data and elements into 3 selections so that it knows which elements it needs to update, add, and remove.

Now you know how `data()` works, let's return to our original example:

```html
<body>
  <script>
    const data = [44847, 63823, 70593, 82203, 109742, 129422, 175832, 182371, 149292, 79420, 59382, 29318];
    d3.select("body")
      .selectAll("p")
      .data(data)
      .join("p")
      .text(d => d);
  </script>
</body>
```

After `d3.select("body").selectAll("p").data(data)`, nothing is added to the DOM - all the selections still exists solely in JavaScript-land. The `join()` method reconciles any discrepancies between the current DOM elements with the data. It does this by updating elements in the update selection, adding new elements from the enter selection, and removing elements in the exit selection. The result is a set of elements that matches the data.

`join()` is a convenient way of updating, adding, and removing elements with a single method call. The alternative is to make calls to `exit()`, `remove()`, `enter()`, `append()`, `merge()`, and `order()` yourself:

```html
<body>
  <script>
    const data = [44847, 63823, 70593, 82203, 109742, 129422, 175832, 182371, 149292, 79420, 59382, 29318];
    let selection = d3.select("body")
      .selectAll("p")
      .data(data);
    
    selection.exit().remove();
    selection.enter().append("p").merge(selection).order().text(d => d);
  </script>
</body>
```

In the version above, we first get the exit selection with `exit()` and then remove those elements from the DOM with `remove()`. Then, we get the enter selection, append a `<p>` element for each element in the selection, and then merge the enter selection with the update selection (elements from the enter selection are added to the end). Then, we must also call the `order()` method that orders the merged selection according to the order in the data array (if we don't, the enter selection elements stays at the end). Finally, we call the `text()` method on the merged selection to add text to each element.

Finally, the `text()` method can take a string and will append the same string to every element of the selection. Here, however, we are passing in a function, which will be called for each selection element, with the element as its first parameter. The function should return with the string (or something that can be stringified) that's appended as a [`Text`](https://developer.mozilla.org/en-US/docs/Web/API/Text) node to the element.

# Updating Data

We have used d3 to draw elements (`<p>`) on the DOM based on some constant data. However, one of the benefits of making diagrams with d3 over drawing diagrams manually with tools like Adobe Illustrator, Inkscape, or Microsoft Publisher, is that they can be _interactive_.

d3 allows the consumers of your diagrams to change parameters or select different data sets, and d3 will update the diagram with the new parameters and data.

To update a d3 diagram based on new data, we can simply re-run the same `select()`, `data()`, `join()` chain.

```html
<body>
  <script>
    const getData = () => {
      return Array.from({length: 12}, () => Math.floor(Math.random() * 200000));
    }
    const drawDiagram = () => {
      d3.select("body")
        .selectAll("p")
        .data(getData())
        .join("p")
        .text(d => d);
    }
    setInterval(drawDiagram, 1000)
    drawDiagram()
  </script>
</body>
```

Now when you open the webpage, you'll see the numbers within the `<p>` updating every second.

But let's solidify what we covered in the last section and see what's happening behind the scenes with regards to the update, enter, and exit selections.

```html
<body>
  <script>
    const getData = () => {
      return Array.from({length: 12}, () => Math.floor(Math.random() * 200000));
    }
    const drawDiagram = () => {
      const selection = d3
        .select("body")
        .selectAll("p")
        .data(getData());
      
      console.log(selection.enter().size()); // 12 the first time, 0 after
      console.log(selection.exit().size()); // 0

      selection
        .join("p")
        .text(d => d);
    }
    setInterval(drawDiagram, 1000)
    drawDiagram()
  </script>
</body>
```

When `drawDiagram()` first runs, because there are no `<p>` elements in the document, all 12 integers in our array have no elements to bind to, and so they enter the enter selection. When `join("p")` was called, each element in the enter selection are added to the document as a `<p>` element.

The next time `drawDiagram()` is run, the `<p>` elements already exists, and d3 will treat the first selected `<p>` element as corresponding to the first datum in the data array and update it. So all elements and data are added to the update selection, hence the reason why the enter and exit selections are empty.

There are common scenarios where you don't want existing elements to update, but instead for the old ones to be removed and new ones added. We will cover how we can do that later at a more opportune time.

# Charts

But most of you are learning d3 not to programmatically add 'boring' elements like `<p>` to the DOM. You are here to use d3 to draw charts like scatterplots, treemaps, Sankey diagrams, timeline charts, network graphs, etc.

![Example of a treemap](https://raw.githubusercontent.com/d3/d3-hierarchy/master/img/treemap.png)

Well, we have the ability to draw charts and have them update with changing data with what we have covered so far. But it's not going to be simple.

Nonetheless, let's start with the simplest scenario there is - given another set of monthly product sales data, plot a bar graph that visualizes the sale of each product. We will use the following dataset:

```js
const data = [
  {
    product: "cap",
    sales: 39,
  },
  {
    product: "shorts",
    sales: 14,
  },
  {
    product: "socks",
    sales: 56,
  },
  {
    product: "jacket",
    sales: 6,
  },
]
```

To draw our chart, we can use SVG.

First, we will define a new `<svg>` element with a `viewBox` attribute of `-10 0 110 110` and `width` and `height` of `400`.

```js
const svg = d3
  .select("body")
  .append("svg")
  .attr("viewBox", "-10 0 110 110")
  .attr("height", "400")
  .attr("width", "400");
```

This gives us a 100 units by 100 units grid for us to draw the bars, plus 10 units for the axis labels.

Next, we draw some `<rect>` elements of uniform width, but where the height is proportional to the sales figures.

```js
const svg = ...
const rect = svg
  .append("g")
  .selectAll("rect")
  .data(data)
  .join("rect")
  .attr("fill", "black")
  .attr("x", (datum, index) => {
    return index * 25
  })
  .attr("width", 25)
  .attr("y", (datum, index) => {
    return 100 - datum.sales
  })
  .attr("height", (datum, index) => {
    return datum.sales
  })
```

![](/img/series-d3-bar-chart-rects.png)

We can then add some more `<text>` elements for the axis labels, and `<line>` elements for a simple grid.

```html
<body>
  <script>
    const data = [
      {
        product: "cap",
        sales: 39,
      },
      {
        product: "shorts",
        sales: 14,
      },
      {
        product: "socks",
        sales: 56,
      },
      {
        product: "jacket",
        sales: 6,
      },
    ]

    const svg = d3
      .select("body")
      .append("svg")
      .attr("viewBox", "-10 0 110 110")
      .attr("height", "400")
      .attr("width", "400");

    const yGrid = svg
      .append("g")
      .selectAll("line")
      .data([...Array(10).keys()].map(n => n*10))
      .join("line")
      .attr("stroke", "grey")
      .attr("stroke-width", "0.5")
      .attr("x1", "0")
      .attr("y1", d => d || -0.25)
      .attr("x2", "100")
      .attr("y2", d => d || -0.25)

    const rect = svg
      .append("g")
      .selectAll("rect")
      .data(data)
      .join("rect")
      .attr("fill", "black")
      .attr("x", (datum, index) => {
        return index * 25
      })
      .attr("width", 25)
      .attr("y", (datum, index) => {
        return 100 - datum.sales
      })
      .attr("height", (datum, index) => {
        return datum.sales
      })
    
    const xLabels = svg
      .append("g")
      .selectAll("text")
      .data(data)
      .join("text")
      .attr("text-anchor", "middle")
      .attr("font-size", "5")
      .attr("x", (datum, index) => {
        return (index * 25) + 12.5
      })
      .attr("y", (datum, index) => {
        return 105
      })
      .text((datum, index) => {
        return datum.product
      })
    
    const yLabels = svg
      .append("g")
      .selectAll("text")
      .data([...Array(10).keys()].map(n => n*10))
      .join("text")
      .attr("dominant-baseline", "middle")
      .attr("text-anchor", "end")
      .attr("font-size", "5")
      .attr("x", "-4")
      .attr("y", (datum, index) => {
        return 100 - datum
      })
      .text(d => d)
  </script>
</body>
```

![](/img/series-d3-bar-chart-manual.png)

Drawing the grid and axis manually may be fine if we want to draw a bar chart specific for this set of data, but what if we want to re-use the bar chart for a different set of data? Or what if the sales of socks boomed during a promotion and is now selling 300 pairs a month?

```js
    const data = [
      { ... },
      { ... },
      {
        product: "socks",
        sales: 300,
      },
      { ... },
    ]
```

![](/img/series-d3-bar-chart-manual-overshot.png)

The bar for `socks` overshot the grid area for the graph. And so to make the new data fit within the graph area, we'd have to redefine the grid lines and our labels.

```html
<body>
  <script>
    const data = [
      {
        product: "cap",
        sales: 39,
      },
      {
        product: "shorts",
        sales: 14,
      },
      {
        product: "socks",
        sales: 300,
      },
      {
        product: "jacket",
        sales: 6,
      },
    ]

    const svg = d3
      .select("body")
      .append("svg")
      .attr("viewBox", "-10 0 110 110")
      .attr("height", "400")
      .attr("width", "400");

    const yGrid = svg
      .append("g")
      .selectAll("line")
      .data([...Array(30).keys()].map(n => n*10))
      .join("line")
      .attr("stroke", "grey")
      .attr("stroke-width", "0.25")
      .attr("x1", "0")
      .attr("y1", d => (d / 3))
      .attr("x2", "100")
      .attr("y2", d => (d / 3))

    const rect = svg
      .append("g")
      .selectAll("rect")
      .data(data)
      .join("rect")
      .attr("fill", "black")
      .attr("x", (datum, index) => {
        return index * 25
      })
      .attr("width", 25)
      .attr("y", (datum, index) => {
        return 100 - (datum.sales / 3)
      })
      .attr("height", (datum, index) => {
        return (datum.sales / 3)
      })
    
    const xLabels = svg
      .append("g")
      .selectAll("text")
      .data(data)
      .join("text")
      .attr("text-anchor", "middle")
      .attr("font-size", "5")
      .attr("x", (datum, index) => {
        return (index * 25) + 12.5
      })
      .attr("y", (datum, index) => {
        return 105
      })
      .text((datum, index) => {
        return datum.product
      })
    
    const yLabels = svg
      .append("g")
      .selectAll("text")
      .data([...Array(30).keys()].map(n => n*10))
      .join("text")
      .attr("dominant-baseline", "middle")
      .attr("text-anchor", "end")
      .attr("font-size", "3")
      .attr("x", "-4")
      .attr("y", (datum, index) => {
        return 100 - (datum / 3)
      })
      .text(d => d)
  </script>
</body>
```

![](/img/series-d3-bar-chart-manual-rescaled.png)

But wouldn't it be great if the grid, labels, and bar heights and width automatically adjusted to the data? Well, with d3, you can!

To understand how, let's first understand how the d3 library is structured. revisit the `<script>` tag that we used to import the d3 library:

```html
<script src="https://d3js.org/d3.v6.min.js"></script>
```

This is actually a minified bundle of numerous d3 _modules_. Since v4, the functionality provided by d3 is split into many smaller modules. In fact, every method we have used so far are all encapsulated into the [`d3-selection`](https://github.com/d3/d3-selection) module. You can confirm this by changing the `<script>` element to download only the `d3-selection` module.

```html
<script src="https://cdn.jsdelivr.net/npm/d3-selection@2"></script>
```

When you refresh the page, everything should still work.

The benefit of using individual modules as opposed to the bundle is the reduction in the file size transferred. As of v6.7.0, that bundle includes the modules [`d3-array`](https://github.com/d3/d3-array), [`d3-axis`](https://github.com/d3/d3-axis), [`d3-brush`](https://github.com/d3/d3-brush), [`d3-chord`](https://github.com/d3/d3-chord), [`d3-color`](https://github.com/d3/d3-color), [`d3-contour`](https://github.com/d3/d3-contour), [`d3-delaunay`](https://github.com/d3/d3-delaunay), [`d3-dispatch`](https://github.com/d3/d3-dispatch), [`d3-drag`](https://github.com/d3/d3-drag), [`d3-dsv`](https://github.com/d3/d3-dsv), [`d3-ease`](https://github.com/d3/d3-ease), [`d3-fetch`](https://github.com/d3/d3-fetch), [`d3-force`](https://github.com/d3/d3-force), [`d3-format`](https://github.com/d3/d3-format), [`d3-geo`](https://github.com/d3/d3-geo), [`d3-hierarchy`](https://github.com/d3/d3-hierarchy), [`d3-interpolate`](https://github.com/d3/d3-interpolate), [`d3-path`](https://github.com/d3/d3-path), [`d3-polygon`](https://github.com/d3/d3-polygon), [`d3-quadtree`](https://github.com/d3/d3-quadtree), [`d3-random`](https://github.com/d3/d3-random), [`d3-scale`](https://github.com/d3/d3-scale), [`d3-scale-chromatic`](https://github.com/d3/d3-scale), [`d3-selection`](https://github.com/d3/d3-selection), [`d3-shape`](https://github.com/d3/d3-shape), [`d3-time`](https://github.com/d3/d3-time), [`d3-time-format`](https://github.com/d3/d3-time), [`d3-timer`](https://github.com/d3/d3-timer), [`d3-transition`](https://github.com/d3/d3-transition), and [`d3-zoom`](https://github.com/d3/d3-zoom) (you can find the list by examining the `d3` repository's [`package.json`](https://github.com/d3/d3/blob/master/package.json)) file. Together, this add up to a bundle size of 264.34 kB (82.15 kB when compressed). Although it's not huge, it does add to the load speed of the site. In contrast, the `d3-selection` script is only 12.93 kB in size (4.77 kB when compressed).

So, going back to making it easier to draw our bar chart, d3 provides several useful modules:

- [`d3-array`](https://github.com/d3/d3-array)
- [`d3-axis`](https://github.com/d3/d3-axis)
- [`d3-scale`](https://github.com/d3/d3-scale)

```html
<script src="https://cdn.jsdelivr.net/npm/d3-selection@2"></script>
<script src="https://d3js.org/d3-array.v2.min.js"></script>
<script src="https://d3js.org/d3-color.v2.min.js"></script>
<script src="https://d3js.org/d3-format.v2.min.js"></script>
<script src="https://d3js.org/d3-interpolate.v2.min.js"></script>
<script src="https://d3js.org/d3-time.v2.min.js"></script>
<script src="https://d3js.org/d3-time-format.v3.min.js"></script>
<script src="https://d3js.org/d3-scale.v3.min.js"></script>
```

Next, we will use d3 to dynamically position and scale the bars. After we've achieved this, we will similarly use d3 to dynamically draw the axises and gridlines.

```js
const y = d3
  .scaleLinear()
  .domain([0, d3.max(data, d => d.sales)])
  .range([0, 100])
```

`scaleLinear()` returns a function that maps values from the input _domain_ to its corresponding value in the output _range_. In our example, the input domain is the sales figures of our products, and the output range is the user units we want the y-axis to span. Since the maximum sales figure is `300` and our chart is `100` user units high, the returned function (`y`) will divide every value by 3.

```js
y(3) // 1
y(12) // 4
```

But if we change the sales figures so they span a different set of numbers.

```js
const data = [
  { ... },
  { ... },
  {
    product: "socks",
    sales: 66,
  },
  { ... },
]
```

Then, the `y` function will automatically adjust to map this new set of input domains to the same output range so that the highest number from the input always reaches the top of the bar chart.

```js
y(3) // 4.615384615384616
y(12) // 18.461538461538463
```

This `y` scaling function effectively replaces the `/ 3` we had manually added to position and scale the height of each bar. So we can replace:

```js
const rect = svg
  ...
  .attr("y", (datum, index) => {
    return 100 - (datum.sales / 3)
  })
  .attr("height", (datum, index) => {
    return (datum.sales / 3)
  })
```

with this:

```js
const rect = svg
  ...
  .attr("y", d => (100 - y(d.sales)))
  .attr("height", d => y(d.sales))
```

Now, try changing the sales figure for different products, you'll find that the height of each bar is automatically proportional to each bar's underlying value, and that the highest bar have a height of 100.

It's great that we got rid of some magic numbers, but looking at the `rect` constant, there are still some magic numbers we can get rid of.

```js
const rect = svg
  .append("g")
  .selectAll("rect")
  .data(data)
  .join("rect")
  .attr("fill", "black")
  .attr("x", (d, i) => i * 25)
  .attr("width", 25)
  .attr("y", d => (100 - y(d.sales)))
  .attr("height", d => y(d.sales))
```

```js
const chartHeight = 100;
const chartWidth = 100;

const rect = svg
  .append("g")
  .selectAll("rect")
  .data(data)
  .join("rect")
  .attr("fill", "black")
  .attr("x", (d, i) => i * chartWidth / data.length)
  .attr("width", chartWidth / data.length)
  .attr("y", d => (chartHeight - y(d.sales)))
  .attr("height", d => y(d.sales))
```

```js
const x = d3.scaleBand()
  .domain(data.map(d => d.product))
  .range([0, chartWidth])

const rect = svg
  .append("g")
  .selectAll("rect")
  .data(data)
  .join("rect")
  .attr("fill", "black")
  .attr("x", d => x(d.product))
  .attr("width", x.bandwidth())
  .attr("y", d => (chartHeight - y(d.sales)))
  .attr("height", d => y(d.sales))
```

```js
const data = [
  { ... },
  { ... },
  { ... },
  { ... },
  {
    product: "keychain",
    sales: 25,
  }
]
```

![](/img/series-d3-bar-chart-5th-product.png)

Now that we got the bars to have the right size and be in the right place, let's shift our attention to making the axis and labels dynamic as well.

```js
const axisLabelMargin = 5;
const xLabels = svg
  .append("g")
  .selectAll("text")
  .data(data)
  .join("text")
  .attr("text-anchor", "middle")
  .attr("font-size", Math.ceil(x.bandwidth() / 5))
  .attr("x", (d, i) => x(i) + x.bandwidth() / 2)
  .attr("y", chartHeight + axisLabelMargin)
  .text(d => d.product)
```

```js
const xLabels = svg
  .append("g")
  .selectAll("text")
  .data(data)
  .join("text")
  .attr("text-anchor", "middle")
  .attr("font-size", Math.ceil(x.bandwidth() / 5))
  .attr("x", (d, i) => x(i) + x.bandwidth() / 2)
  .attr("y", chartHeight + axisLabelMargin)
  .text(d => d.product)
```

```js
const xAxis = svg.append("g")
  .attr("transform", `translate(0,${chartHeight})`)
  .call(d3.axisBottom(x).tickSize(1))
  .attr("stroke-width", "0.25")
  .attr("font-size", "3")
  .attr("color", "black")
```

```js
const yGrid = svg
  .append("g")
  .selectAll("line")
  .data([...Array(30).keys()].map(n => n*10))
  .join("line")
  .attr("stroke", "grey")
  .attr("stroke-width", "0.25")
  .attr("x1", "0")
  .attr("y1", d => (d / 3))
  .attr("x2", "100")
  .attr("y2", d => (d / 3))

const yLabels = svg
  .append("g")
  .selectAll("text")
  .data([...Array(30).keys()].map(n => n*10))
  .join("text")
  .attr("dominant-baseline", "middle")
  .attr("text-anchor", "end")
  .attr("font-size", "3")
  .attr("x", "-4")
  .attr("y", (datum, index) => {
    return 100 - (datum / 3)
  })
  .text(d => d)
```

```js
const y = d3
  .scaleLinear()
  .domain([0, d3.max(data, d => d.sales)])
  .range([chartHeight, 0])

const yAxis = svg.append("g")
  .call(d3.axisLeft(y).tickSize(1))
  .attr("stroke-width", "0.25")
  .attr("font-size", "3")
  .attr("color", "black")

const rect = svg
  ...
  .attr("y", d => y(d.sales))
  .attr("height", d => (chartHeight - y(d.sales)))
```

https://observablehq.com/@d3/lets-make-a-bar-chart/4

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Understanding d3.js</title>
  <script src="https://cdn.jsdelivr.net/npm/d3-selection@2"></script>
  <script src="https://d3js.org/d3-array.v2.min.js"></script>
  <script src="https://d3js.org/d3-color.v2.min.js"></script>
  <script src="https://d3js.org/d3-format.v2.min.js"></script>
  <script src="https://d3js.org/d3-interpolate.v2.min.js"></script>
  <script src="https://d3js.org/d3-time.v2.min.js"></script>
  <script src="https://d3js.org/d3-time-format.v3.min.js"></script>
  <script src="https://d3js.org/d3-axis.v2.min.js"></script>
  <script src="https://d3js.org/d3-scale.v3.min.js"></script>
</head>
<body>
  <script>
    const data = [
      {
        product: "cap",
        sales: 39,
      },
      {
        product: "shorts",
        sales: 14,
      },
      {
        product: "socks",
        sales: 12,
      },
      {
        product: "jacket",
        sales: 6,
      },
      {
        product: "keychain",
        sales: 25,
      }
    ]

    const chartHeight = 100;
    const chartWidth = 100;
    const axisLabelMargin = 5;

    const x = d3.scaleBand()
      .domain(data.map(d => d.product))
      .range([0, chartWidth])

    const y = d3
      .scaleLinear()
      .domain([0, d3.max(data, d => d.sales)])
      .range([chartHeight, 0])

    const svg = d3
      .select("body")
      .append("svg")
      .attr("viewBox", "-10 0 110 110")
      .attr("height", "400")
      .attr("width", "400");
    
    const rect = svg
      .append("g")
      .selectAll("rect")
      .data(data)
      .join("rect")
      .attr("fill", "black")
      .attr("x", d=> x(d.product))
      .attr("width", x.bandwidth())
      .attr("y", d => y(d.sales))
      .attr("height", d => (chartHeight - y(d.sales)))

    const xAxis = svg.append("g")
      .attr("transform", `translate(0,${chartHeight})`)
      .call(d3.axisBottom(x).tickSize(1))
      .attr("stroke-width", "0.25")
      .attr("font-size", "3")
      .attr("color", "black")
    

    const yAxis = svg.append("g")
      .call(d3.axisLeft(y).tickSize(1))
      .attr("stroke-width", "0.25")
      .attr("font-size", "3")
      .attr("color", "black")
    
  </script>
</body>
</html>

```

```
<body>
  <script>
    fetch("http://localhost:8080/files/happiness-gdppc-2019.csv")
      .then(response => response.text())
      .then(d3.csvParse)
      .then(data => {})
  </script>
</body>
```

```js
[
  {
    Country: "Afghanistan",
    "Log GDP per capita": "7.697",
    Score: "2.375"
  }, ...
]
```

For example, d3 provides you with methods to draw treemaps in its [`d3-hierarchy`](https://github.com/d3/d3-hierarchy/) module.
