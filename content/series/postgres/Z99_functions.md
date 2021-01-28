# Functions

## Window Functions

A _window function_ perform calculations across a _window_ of rows. A window is a set of rows that are, in some way, related to the current row. For example, in a table of students, a window may be the set of students that belong to the same class as the current student.

You can, for example, get the average numeric value for a particular column over a set of rows. This is similar to aggregation functions but instead of the returning a single value in the results, a window function is meant to be used to populate a separate column in the results.



## User-Defined Functions


user-defined SQL functions (?? is it the same as stored procedures??)
user-defined functions and types

aggregate functions computes a single result from multiple rows. For example, you can use an aggregate function to:

- get the total value of all inventory in your warehouse (`sum`)
- get the average rating of all the chess players in a tournament (`avg`)
- get the total number of entries in a raffle (`count`)
- get the ID of the participant who completed the most challenges (`max`)
- get the lowest recorded temperature from historic weather data (`min`)

Aggregate functions cannot be used in the `WHERE` clause since the `WHERE` clause is used to determine which rows are included in the aggregation.

- `HAVING` - similar to `WHERE` but works on aggregate data rather than raw table rows. You can also look at it this way - the `WHERE` clause filters pre-aggregation, `HAVING` filters post-aggregation
