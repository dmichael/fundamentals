# Fundamentals

What exactly are the fundamentals when working with computers? There are is an incredible diversity of topics and techniques - from data structures to distributed systems

This repository aims to collect some computer science fundamentals implemented in Golang.

## Data Structures (Abstract Data Types)

* Linked Lists
* Stacks (LIFO), Queues (LIFO)
  * Dijkstra's Two Stack Algorithm for Expression Evaluation
* Trees and Graphs
  * Binary Search Tree
  * Traversal
    * In-Order
    * Pre-Order
    * Post-Order
  * Heap (Priority Queue)
    * Max-Heap
    * Min-Heap  
  * Tries (Prefix Trees)
* Vectors, ArrayLists
* Hash Tables

## Algorithms

* Searching
  * Breadth-first Search
  * Depth-first Search
  * Binary Search
* Sorting
  * Merge Sort
  * Quick Sort

## Concepts

* Bit Manipulation
* Memory (Stack vs Heap)
* Recursion
* Dynamic Programming
* Big O

# Subjects

* Base 2

## Linked Lists

#### Definition

 > A linked list is a linear collection of data elements, called nodes, each pointing to the next (and previous) node by means of a pointer.

The principal benefit of a linked list over a conventional array is that the list elements can easily be inserted or removed without reallocation or reorganization of the entire structure.

#### Performance

* Constant time insertions and deletions from head and tail `O(1)`
* Value access is `O(N)` - every node must be visited

#### When To Use

* Use a `linked list` to optimize insertions and deletions from the front or back of a collection
* Can be used to implement a `stack` and a `queue`

```go
type Element struct {
  Next, Prev *Element
  Value interface{}
}
```

### Traversal

In Go, the [`container/list`](https://golang.org/pkg/container/list/) package implements the doubly-linked list

```go
for e := l.Front(); e != nil; e = e.Next {
	// do something with e.Value
}
```

#### Linked List References

* https://en.wikipedia.org/wiki/Linked_list

## Heap

#### Definition

> **A heap is a specialized tree-based data structure that satisfies the heap property**: If A is a parent node of B, then the key (the value) of node A is ordered with respect to the key of node B with the same ordering applying across the heap

Interestingly, this data structure can be implemented using an `Array`!

#### When To Use

Use a heap when you need quick access to the largest (or smallest) item, because that item will always be the first element in the array or at the root of the tree.

#### Performance

* Access at constant time `O(1)`
* Re-balance at `O(logN)`

#### References

* https://en.wikipedia.org/wiki/Heap_(data_structure)
* https://www.youtube.com/watch?v=t0Cq6tVNRBA

## String Manipulation
