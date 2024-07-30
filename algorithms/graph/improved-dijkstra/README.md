# Dijkstra Improved

Why is it improved?

Because we're using a Priority Queue (under a Heap) to improve it's performance.

## Perforamnce

In the optimized implementation using a Min Heap (priority queue), the time complexity changes significantly:

1. Inserting into the Priority Queue: Each insertion into the priority queue (Min Heap) takes 𝑂(log 𝑉) time.
2. Extracting the Minimum Distance Vertex: Each extraction (pop) from the priority queue takes 𝑂(log 𝑉) time.
3. Updating the Distances of Adjacent Vertices: Each update operation involves adjusting the priority of a vertex in the priority queue, which also takes 𝑂(log 𝑉) time.

Comparison:

Naive implementation (without heap): O(V^2 + E), which can be approximated as O(V^2) in dense graphs.
Heap implementation: O((V + E) log V).

The heap implementation is more efficient, especially for large and sparse graphs (where E is much less than 𝑉^2).
For dense graphs, the benefit might be less pronounced, but it is still typically more efficient than the naive approach.
