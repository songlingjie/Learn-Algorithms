package main

import "fmt"

type Queue struct {
	list []int
}
func NewQueue() *Queue{
	list := []int{}
	return &Queue{list:list}
}
func (q *Queue)Push(data int) {
	q.list = append(q.list, data)
}

func (q *Queue)Pop() int{
	if len(q.list)==0 {
		return 0
	}
	res := q.list[0]
	q.list = q.list[1:]
	return res
}


func main()  {

	// 使用邻接矩阵构建无向图
	graph := newGraph()

	fmt.Printf("%c", graph)

	//bfs(graph, 0,7)
}

// 广度优先 从s 到 t 的最短路劲
func bfs(graph [][]int, s int, t int)  {
	var visited []bool
	var queue = NewQueue()
	var prev []int

	visited[s] = true
	queue.Push(s)

	for len(queue.list) > 0{
		w := queue.Pop()
		for i := 0; i < len(graph[w]); i++ {
			q := graph[w][i]
			if !visited[q]{
				prev[w] = q
				if t == q {
					fmt.Printf("%c", prev)
					return
				}
			}
			queue.Push(q)
		}
	}
}

func newGraph() [][]int {
	/**
	0	1	2
	3	4	5
		6	7
	*/
	var graph [][]int
	addNode(0,1,graph)
	addNode(0,3,graph)
	addNode(1,2,graph)
	addNode(1,4,graph)
	addNode(3,4,graph)
	addNode(4,5,graph)
	addNode(4,6,graph)
	addNode(6,7,graph)
	addNode(5,7,graph)
	return graph
}

func addNode(i int, j int, graph [][]int)  {
	graph[i][j] = 1
	graph[j][i] = 1
}