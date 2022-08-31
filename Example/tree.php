<?php

// 使用非递归遍历二叉树

class tree
{
    public $data;

    public $right;

    public $left;

    public function __construct($data = 0, $right = null, $left = null)
    {
        $this->data  = $data;
        $this->right = $right;
        $this->left = $left;
    }
}

function create(): tree
{
    $arr = [7,4,13,1,15,6,3];

    $tree = new tree(array_shift($arr));
    foreach ($arr as $item) {
        $tree = insert($tree, $item);
    }

    return $tree;
}

/**
 * @param  $tree
 * @param int  $data
 * @return tree
 */
function insert($tree, int $data): tree
{
    if (empty($tree)) {
        $tree = new tree($data);
    }

    if ($tree->data == $data) {
        return $tree;
    }

    if ($tree->data > $data) {
        $tree->left = insert($tree->left, $data);
        return $tree;
    }
    $tree->right = insert($tree->right, $data);
    return  $tree;
}

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