## Multi Source BFS (Breadth First Search)
Multi-source BFS is a variant of the standard Breadth-First Search algorithm where instead of starting from a single source node, you initialize the queue with multiple source nodes simultaneously.

### Core Concept
Think of it like dropping multiple stones into a pond at the same time - the ripples from all stones expand outward simultaneously. The first ripple to reach any point determines the distance to that point.

### Advantages
- More efficient than running separate BFS from each source
- Guarantees shortest paths in unweighted graphs
- Natural parallelization potential
