package disjointsetunion

func Find(x int, parent []int) int {
	if x == parent[x] {
		return x
	}

	parent[x] = Find(parent[x], parent)
	return parent[x]
}

func Union(x, y int, parent, rank []int) {
	xParent, yParent := Find(x, parent), Find(y, parent)

	if rank[xParent] > rank[yParent] {
		parent[yParent] = xParent
	} else if rank[xParent] < rank[yParent] {
		parent[xParent] = yParent
	} else {
		parent[xParent] = yParent
		rank[yParent]++
	}
}
