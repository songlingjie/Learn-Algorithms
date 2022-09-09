<?php

/**
 * 使用邻接矩阵 存储如下图 （ps 邻接表是用链表存储）
 * 0    1    2
 * 3    4    5
 *      6    7
 */
$graph       = [];
$graph[0][1] = 1;
$graph[0][3] = 1;
$graph[1][2] = 1;
$graph[1][4] = 1;
$graph[2][1] = 1;
$graph[2][5] = 1;
$graph[3][4] = 1;
$graph[4][1] = 1;
$graph[4][3] = 1;
$graph[4][5] = 1;
$graph[4][6] = 1;
$graph[6][4] = 1;
$graph[6][7] = 1;
$graph[5][2] = 1;
$graph[5][4] = 1;
$graph[5][7] = 1;

// 广度优先寻找最短路径
function bfs($graph, $j, $s): array
{

    $visited = []; //判断节点是否 被访问过
    $queue   = [];
    $path    = []; // 访问路径

    $queue[]     = $j;
    $visited[$j] = true;
    while (!is_null($q = array_shift($queue))) {
        foreach ($graph[$q] as $key => $item) {
            if ($visited[$key]) {
                continue;
            }
            $path[] = $key;
            if ($key == $s) {
                return $path;
            }
            $queue[]       = $key;
            $visited[$key] = true;
        }
    }
    return $path;
}


function dfs($graph, $j, $s): array
{
    $visited = [];
    $path    = [];
    $found = false;

    _dfs($j, $s, $graph, $visited, $path, $found);
    return $path;
}
function _dfs($j, $s, $graph, &$visited, &$path, &$found)
{
    if ($j == $s) {
        $found = true;
    }
    if ($found) {
        return true;
    }
    $visited[$j] = true;

    foreach ($graph[$j] as $key => $item) {
        if (!$visited[$key]) {
            $path[] = $key;
             _dfs($key, $s, $graph, $visited, $path, $found);
        }
    }
}

print_r(dfs($graph, 0, 7));

