## Minimal Spanning Tree by using Heap i.e Prim's Algorithm
Prim's algorithm finds the Minimum Spanning Tree (MST) using a vertex-growing approach. It starts from an arbitrary vertex and gradually expands the tree by always adding the cheapest edge that connects a vertex in the tree to a vertex outside the tree.


### Key Properties
✅ For: Undirected, weighted, connected graphs

✅ Greedy algorithm: Makes locally optimal choices

✅ Uses: Priority Queue (Min-Heap) data structure

✅ Time: O(E log V) with binary heap, O(E + V log V) with Fibonacci heap

✅ Best for: Dense graphs (E ≈ O(V²))

✅ Works with: Positive, negative, and zero weights

✅ Produces same MST as Kruskal's (just different approach)
