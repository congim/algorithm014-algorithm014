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