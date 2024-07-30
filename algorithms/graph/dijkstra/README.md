# Dijkstra

This algorithm is used to find the shortest path from a node to another.

## Performance

In terms of complexity, it performs two operations:

1. Finding the Minimum Distance Vertex: This operation takes O(V) time, where V is the number of vertices.
2. Updating the Distances of Adjacent Vertices: For each vertex, we might update the distances of all its adjacent vertices, which takes O(E) time in total, where E is the number of edges.

Since we repeat the process of finding the minimum distance vertex for each of the V vertices, the overall time complexity is:
𝑂(𝑉2+𝐸)

In dense graphs where E is close to 𝑉ˆ2, this can be approximated as: 𝑂(𝑉^2)

This is possible to improve, thus we have implementations using Priority Queue.
