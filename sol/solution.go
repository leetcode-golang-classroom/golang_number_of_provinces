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
