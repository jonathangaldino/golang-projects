# Sorting Algorithms

- QuickSort
- Selection Sort


# Quicksort

QS algorithm uses the divide to conquer strategy, usually with Recursion.
It's faster then Selection Sort and way more used by programmers.

In the worst occasion, QS has time complexity of O(n^2), making it slower as Selection Sort.
However, QS has medium time complexity of O(n log n).

# Selection Sort

Selectio Sort is not that fast. We will create another list with the ordered items.
That means we will need to check every element in the list once - thus having O(n).


##  Merge Sort vs Quick Sort

We usually don't take constants too serious when speaking about time complexity (Big O), however, when we compare QS and MS, QS has a smaller constant then MS.
They both have O(n log n) but QS is usually faster in practice because QS usually performs in a medium time complexity case.

The worst case of Quicksort depends on the choice of the Pivot.
If the array is already ordered, QS will try to order it again.

### Occasions

Worst scenario: Already ordered array and the pivot is the first item. QS will run and check all items. Thus performing O(n).

Best scenario:  Choosing an element as the Pivot in a random way, QS will complete in a medium execution of O(n log n).


In the worst case ever, there are O(n) levels (call stack).
