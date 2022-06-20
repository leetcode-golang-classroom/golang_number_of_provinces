# golang_number_of_provinces

There are `n` cities. Some of them are connected, while some are not. If city `a` is connected directly with city `b`, and city `b` is connected directly with city `c`, then city `a` is connected indirectly with city `c`.

A **province** is a group of directly or indirectly connected cities and no other cities outside of the group.

You are given an `n x n` matrix `isConnected` where `isConnected[i][j] = 1` if the `ith` city and the `jth` city are directly connected, and `isConnected[i][j] = 0` otherwise.

Return *the total number of **provinces***.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2020/12/24/graph1.jpg](https://assets.leetcode.com/uploads/2020/12/24/graph1.jpg)

```
Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
Output: 2

```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/12/24/graph2.jpg](https://assets.leetcode.com/uploads/2020/12/24/graph2.jpg)

```
Input: isConnected = [[1,0,0],[0,1,0],[0,0,1]]
Output: 3

```

**Constraints:**

- `1 <= n <= 200`
- `n == isConnected.length`
- `n == isConnected[i].length`
- `isConnected[i][j]` is `1` or `0`.
- `isConnected[i][i] == 1`
- `isConnected[i][j] == isConnected[j][i]`

## 解析

題目給了一個 n by n 矩陣 isConnected 

每個 entry isConnected[i][j] 有兩個值

isConnected[i][j]  = 1 代表 $vertex_i$  與 $vertex_j$ 相連

isConnected[i][j]  = 0 代表 $vertex_i$  與 $vertex_j$ 不相連

要求寫出一個演算法來計算有多少相連的 component

類似 [323. ****Number of Connected Components in an Undirected Graph****](https://www.notion.so/323-Number-of-Connected-Components-in-an-Undirected-Graph-3fd878585c85470db45d21418402280e) 透過 edge 來使用 Union/Find Algorithm

![](https://i.imgur.com/KqoNIYv.png)

## 程式碼
```go
package sol

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	parent := make([]int, n)
	rank := make([]int, n)
	for node := 0; node < n; node++ {
		parent[node] = node
		rank[node] = 1
	}
	var find = func(node int) int {
		p := parent[node]
		for p != parent[p] {
			parent[p] = parent[parent[p]]
			p = parent[p]
		}
		return p
	}
	var union = func(node1, node2 int) int {
		p1 := find(node1)
		p2 := find(node2)
		if p1 == p2 {
			return 0
		}
		if rank[p1] > rank[p2] {
			parent[p2] = p1
			rank[p1] += rank[p2]
		} else {
			parent[p1] = p2
			rank[p2] += rank[p1]
		}
		return 1
	}
	result := n
	for n1 := 0; n1 < n; n1++ {
		for n2 := 0; n2 < n; n2++ {
			if isConnected[n1][n2] == 1 {
				result -= union(n1, n2)
			}
		}
	}
	return result
}

```
## 困難點

1. 要透過 UnionFind Algorithm 的概念來優化演算法

## Solve Point

- [x]  透過 isConnected 找出 相連的 edge
- [x]  透過 edge 來使用 Union/Find Algorithm