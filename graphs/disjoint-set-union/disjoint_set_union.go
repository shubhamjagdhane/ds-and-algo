package disjointsetunion

func find(x int, parent []int) int {
	if x == parent[x] {
		return x
	}

	parent[x] = find(parent[x], parent)
	return parent[x]
}

func union(x, y int, parent, rank []int) {
	xParent, yParent := find(x, parent), find(y, parent)

	if rank[xParent] > rank[yParent] {
		parent[yParent] = xParent
	} else if rank[xParent] < rank[yParent] {
		parent[xParent] = yParent
	} else {
		parent[xParent] = yParent
		rank[yParent]++
	}
}
