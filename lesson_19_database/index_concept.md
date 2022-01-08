What is Database index ? How is it implemented under the hood?

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

- B+Tree are self-balance, BST are not self-balance => BST can lead be a very tall tree
- B+Trees are much shorter than other balanced binary tree structures like AVL => 
- If we want to retrieve a range of entries in order by key, B+Trees is much efficient than AVL, BST or Hash Table, because each node contained tons of keys with sorted

What is B+Tree indexes’s disadvantage?

- More complex than Hash indexes
- For equality query, Hash Indexes much faster O(n) on average, B+Tree O(logN)

Given B+Tree index  (C1, C2, C3), can B+Tree support query on C2 ?

No, because B+Trees just support leftmost prefix

Compare B+Tree Index and Hash Index ?

- Match part of first column => ex: column **Name** => find name is “Ngh” => B+Tree support all name has Name start with “Ngh”
- ` `Support match range of value (>=, >, <=, <)

What happens to index when database machine crashes?

DBMS uses an Data Structure to write on disk called write-ahead log WAL

- Append-only file which every B+Tree modification must be written before it can be apply
- Append-only => fast

When database come back after crash, this logs is readed to restone B+Trees

What diff between clustered index and non\_clustered index ?

- Clustered index: the index that define the order that data is stored on disk
  - When creating table, if you don’t have specific primary key, DBMS often create its own hidden row primary key
  - Just has only one clustered index




- Non-Cluster Index: the index does not sort the physical cata’s order
  - Non-clustered index only has pointer to the data
  - Can have multiple non-clustered indexes.

Should we index boolean column ?

No, because selectivity is too small

Selectivity = number of distinct indexed value / number of rows in table




