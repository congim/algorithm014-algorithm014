学习笔记

### 递归版本
```golang
type TreeNode struct {
    left  *TreeNode
    right *TreeNode
    value int
}

var visited map[node.value]node
func dfs(node) {
    if _, ok := visited[node.value]; ok {
        return
    }
    visited[node.value] = node
    dfs[node.left]
    dfs[node.right]
}
```
### 改进版本

```python
visited = set

def dfs(node, visited)
 if node in visited: 
    # alreadyi visited
    return
 visited.add(node)
 # process current node here
 ...
 for next_node int node.children()
    if not next_node in visited:
     dfs(next_node, visited)
```

### 非递归写法
```python

def DFS(self, tree):
    if tree.root is None:
        return []
    visited, stack = [], [tree.root]

    while stack:
        node = stack.pop()
        visited.add(node)

        process(node)
        nodes = generate_related_nodes(node)
        stack.push(nodes)
    # other processing work    
```

### 广度优先遍历
一层一层向下下扩散

```python

# Python
def BFS(graph, start, end):
    visited = set()
	queue = [] 
	queue.append([start]) 
	while queue: 
		node = queue.pop() 
		visited.add(node)
		process(node) 
		nodes = generate_related_nodes(node) 
		queue.push(nodes)
	# other processing work 
```

### 贪心算法

贪心算法是一种在每一步选择中逗采取在当前状态下最好或者最优的选择，从而希望导致结果是全局最好或者最优的算法。
贪心算法与动态规划的不同在于它对每个自问题的解决方案逗做出选择，不能回退。动态规划则会保存以前的运算结果，并根据以前的结果对当前进行选择，有回退功能。