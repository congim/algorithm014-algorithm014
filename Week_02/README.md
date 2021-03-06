
### 总结：

-   哈希表，集合
-   树，二叉树，二叉搜索树
-   堆，二叉堆
这几种特殊的数据结构都是通过最基本的数据结构，例如数组，通过函数以进行实现，例如哈希表是通过哈希函数实现。为什么要抽象出不同的数据类型呢？最主要的目的就是为了提高运行时间，以空间换时间。

而树的结构其实是由数学，或者生活中其他的常见的现象得出来的，覃超老师讲的一个很有意思的点其实是：每一个树的节点可以理解为时间轴上面不同选择而变成了某一个状态。每一个节点是一种**状态**。

#### 哈希表 - Set

-   在 python 里面，哈希表的实现就是通过 dict, key-value pair, 一个字典里面只有的 key 是不会重复的。
-   Set 具体实现的方式和哈希表类似，每个 set 里面的值是唯一的，同比 dict 里面的 key

#### 树 - 二叉树 - 二叉搜索树：

-   之前的数据结构，基本都是一维结构，一条直线。
-   树结构已经升维成了二维的数据结构，Node, Child Node, Sibling, Root.
-   二叉树，就是一个 Node，有左右子树。无序的数其实要遍历整棵树。
-   遍历方式组成可以分成为前序，中序，后续遍历。
-   二叉搜索树：左子树所有的节点的值都小于根节点的值。右子树的所有节点的值都大于根节点的值。-- 但是中序遍历的时候，这是一个升序的遍历。
-   查询和插入操作都是 logn

#### 堆 heap 和 二叉堆 binary heap:

-   特性，能够快速找到一堆数中最大值或者最小值
-   入门级的堆是二叉堆，delete, insert = logn find-min/max = O(1)
-   二叉堆 通过完全二叉树实现的，除了叶子节点都是满的
-   假设第一个元素在数组中索引为 0，左孩子的索引是 2*i +1, 右孩子是 2 * i + 2, 索引为 i 的父节点索引是 floor((i-1)/2)
-   模拟插入，删除。heapify up and heapify down. 用优先队列进行实现。
