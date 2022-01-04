**What is Database index ? How is it implemented under the hood?**

A database index is a data structure  that improve the speed of data retrieval operations on a database table.

Why don’t we index every columns to support fast read?

Increasing at the cost of additional writes and storage space to maintain the index data structure

How does Hash Index work ? What is its disadvantages?

- Hash indexes work based on a Hash Table with:
  - Key: is hash code of index columns
  - Value: pointer to corresponding rows

Disadvantages are 

- Hash indexes don’t support particle key matching
- Hash indexes support only equality comparisons
- Hash indexes don’t stable (because hash table data structure doesn’t stable)

Compare B+Tree and B-Tree ?

`	`B-Tree similar with B+Tree, except:

- Inner node can contain pointer to record
- No duplicated in B-Tree
- Left-node doesn’t connect each other by **Doubly Linked List**

Why use B+Tree to index, not BST/AVL/Red-Black Tree?

- D

What is B+Tree indexes’s disadvantage?

- D

Given B+Tree index  (C1, C2, C3), can B+Tree support query on C2 ?

- D

Compare B+Tree Index and Hash Index ?

- D

**What happens to index when database machine crashes?**

- D

What diff between clustered index and non\_clustered index ?

- D

**Should we index boolean column ?**

- D




10/2018
