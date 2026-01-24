## Minimal Spanning Tree by sorting weights i.e Kruskal's Algorithm
Kruskal's algorithm finds the Minimum Spanning Tree (MST) of an undirected, weighted graph using a greedy approach. It builds the MST by sorting edges by weight and adding them in ascending order, avoiding cycles.

### Key Properties
✅ For: Undirected, weighted, connected graphs

✅ Greedy algorithm: Makes locally optimal choices

✅ Uses: Union-Find (Disjoint Set Union) data structure

✅ Time: O(E log E) or O(E log V) with sorting + Union-Find

✅ Works with: Positive, negative, and zero weights

✅ May produce multiple MSTs if edges have equal weights
