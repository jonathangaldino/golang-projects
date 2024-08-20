# K Nearest Neighboors

The k-Nearest Neighbors (KNN) algorithm is simple and effective for many classification and regression tasks, but it has several drawbacks:

**High Computational Cost:** KNN requires calculating the distance from a query point to all points in the training set, making it computationally expensive, especially for large datasets.

**Storage and Memory Intensive:** Since KNN is a lazy learning algorithm, it stores the entire training dataset, which can consume significant memory.

**Slow Predictions:** Predictions can be slow because KNN requires scanning the entire dataset at prediction time to find the nearest neighbors.

**Sensitive to the Choice of k:** The performance of KNN heavily depends on the choice of the hyperparameter k (number of neighbors). Too small a k can lead to overfitting, while too large a k may underfit.

**Curse of Dimensionality:** As the number of dimensions (features) increases, the distance between points becomes less meaningful due to the curse of dimensionality, leading to poor performance.

**Imbalanced Data:** KNN struggles with imbalanced datasets, as it may be biased toward the majority class because of the voting mechanism.

**Sensitive to Noise:** Noisy data can significantly affect the performance, as nearby noisy data points can alter the classification.

## Definition of characteristics

Religion can be: catolic, cristian, etc
Hobby can be: sing, reading, dance, etc
On vacation, prefer to: travel, go to the beach, rest, etc

## Characteristics of the data

Transforming the data into numerical values. Example:

- Religion: 0 (catolic), 1 (cristian)
- Hobby: 0 (sing), 1 (reading), 2 (dance)
- On vacation: 0 (travel), 1 (go to the beach), 2 (rest)

## Non cardinal characteristics

- Singing: [1, 0, 0]
- Reading: [0, 1, 0]
- Dancing: [0, 0, 1]
- ... etc

With more characteristics, the dataset will be more complex.

The same applies to `On vacation, prefer to`. For example:

- Travel: [1, 0, 0]
- Go to the beach: [0, 1, 0]
- Rest: [0, 0, 1]
- ... etc

## Combining characteristics

Combining characteristics in a single vector for each person.
For example, for a person with the following characteristics:

- Religion: catolic (1)
- Hobby: reading (0, 1, 0)
- On vacation: go to the beach (0, 1, 0)

### Applying weights for each characteristic

- Religion has a weight of 3 => 1*3=3
- Hobby has a weight of 1 => (0, 1, 0) * 1 = (0, 1, 0)
- On vacation has a weight of 2 => (0, 2, 0)

Combined vector: [3, 0, 1, 0, 0, 2, 0]


## Calculating the distance between two vectors

We can calculate the distance between two vectors using **Euclidean Distance**, or even Manhattan Distance.

**Euclidean Distance** is a measure of the distance between two points in a multi-dimensional space. It is calculated as the square root of the sum of the squared differences between the corresponding components of the two points.

### Example of Euclidean Distance

Let's say we have two vectors:

Vector 1: [1, 2, 3]
Vector 2: [4, 5, 6]

The distance between these two vectors is calculated as:

sqrt((1 - 4)^2 + (2 - 5)^2 + (3 - 6)^2)

= sqrt(1 - 4 + 4 + 9 + 16)

= sqrt(22)

= 4.47214

### Example of Manhattan Distance

Let's say we have two vectors:

Vector 1: [1, 2, 3]
Vector 2: [4, 5, 6]

The distance between these two vectors is calculated as:

abs(1 - 4) + abs(2 - 5) + abs(3 - 6)

= 1 - 4 + 2 - 5 + 3 - 6

= 1 + 2 + 3

= 6
