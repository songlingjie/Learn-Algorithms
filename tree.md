##### 二叉树生成和遍历

```go
// 前序遍历
func preOrder(node *tree)  {
	if node == nil {
		return
	}

	fmt.Println(node.data, " ")
	preOrder(node.left)
	preOrder(node.right)
}

// 中序遍历
func inOrder(node *tree)  {
	if node == nil {
		return
	}

	preOrder(node.left)
	fmt.Println(node.data, " ")
	preOrder(node.right)
}

// 后序遍历
func postOrder(node *tree)  {
	if node == nil {
		return
	}

	preOrder(node.left)
	preOrder(node.right)
	fmt.Println(node.data, " ")
}
```

##### 不是用递归前序遍历二叉树
```php
$tree = create();

//辅助栈
$stack = [];
$stack[] = $tree;

while (!empty($stack)) {

    $tree = array_pop($stack);
    print_r($tree->data. "\n");

    if (!empty($tree->right)) {
        $stack[] = $tree->right;
    }

    if (!empty($tree->left)) {
        $stack[] = $tree->left;
    }
}
```
