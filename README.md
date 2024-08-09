# golang-triangle-weight

You are given an array of arrays of integers, where each array corresponds to a row in a triangle of numbers. For example, [[1], [2, 3], [1, 5, 1]] represents the triangle:

```yaml
    1
  2   3
1   5   1
```
We define a path in the triangle to start at the top and go down one row at a time to an adjacent value, eventually ending with an entry on the bottom row. For example, 1 -> 3 -> 5. The weight of the path is the sum of the entries.

Write a program that returns the weight of the maximum weight path.


## Think

for each level we could know the adjacent list if current level is not leave level (i != len(triangles))

for each node we could use DEPTH First Search to calculate each route. And only when all possible route then result could be found

and here one greedy solution

sum(level, pos) = triangle[level][pos] + Math.max(sum(level+1, pos), sum(level+1, pos+1))

time complexity is O(n)

