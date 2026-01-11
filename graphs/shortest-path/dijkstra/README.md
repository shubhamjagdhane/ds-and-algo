## Shortest Path using Dijkstra's Algorithm

Dijkstra's algorithm determines the shortest paths from a starting node to all other nodes in a weighted graph with non-negative edge weights. It's a greedy algorithm that explores the graph from the source using a Min-Heap, ensuring the shortest distance is always processed first. Nodes already visited via the shortest path are marked, eliminating longer paths.

### Key Properties

✅ Finds: Shortest path from a single source to all vertices

✅ Works for: Weighted graphs with non-negative edge weights

❌ Does NOT work: Graphs with negative edge weights

### Graph-1
![graph-1](../../../images/undirected_weighted_graph_1.png)

### Graph-2
![graph-2](../../../images/undirected_weighted_graph_2.png)
