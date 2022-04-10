

d3 manipulates the DOM based on data. For example, if we have an array of 3 data points, we can pass this array to d3, who may add 3 `<li>` elements to an `<ol>` element.


We will first look at the DOM manipulation aspect of d3.js, then moving on to how to bind data to DOM elements, and finally at some of d3's more advanced


Because everything d3 creates is part of the DOM, we can integrate d3 with other DOM manipulation libraries, we can style the elements using CSS, add event listeners on them with regular JavaScript.

---

Data sources

https://www.kaggle.com/
https://dataverse.harvard.edu/


---

```html
<body>
  <script>
    const getPosts = () => {
      return Array.from({length: 12}, () => ({
        id: Math.random().toString(36).substring(2,3),
        content: Math.floor(Math.random() * 200000),
      }));
    }
    setInterval(() => {
      const selection = d3.select("body")
      .selectAll("p")
      .data(getPosts(), p => p.id)
      console.log(selection.enter().size()); // 12
      console.log(selection.exit().size()); // 12
      
      selection
        .join(
          enter => enter.append("p").style("color", "green"),
          update => update.style("color", "black"),
        )
        .text(p => p.content);
    }, 5000)
  </script>
</body>
```

---

Courses

https://embermap.com/topics/d3/intro-101
https://scrimba.com/learn/d3js
