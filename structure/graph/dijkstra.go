package graph

import (
	"math"
)

const INF = math.MaxInt32

//dijkstra算法
//参考文章：http://www.cnblogs.com/skywang12345/p/3711516.html

/*
 * Dijkstra最短路径。
 * 即，统计图中"顶点vs"到其它各个顶点的最短路径。
 *
 * 参数说明：
 *       vs -- 起始顶点(start vertex)。即计算"顶点vs"到其它顶点的最短路径。
 *     prev -- 前驱顶点数组。即，prev[i]的值是"顶点vs"到"顶点i"的最短路径所经历的全部顶点中，位于"顶点i"之前的那个顶点。
 *     dist -- 长度数组。即，dist[i]是"顶点vs"到"顶点i"的最短路径的长度。
 */
func Dijkstra(matrix [][]int, vs int) []int {

	vexnum := len(matrix)
	dist := make([]int, vexnum)
	// flag[i]=true表示"顶点vs"到"顶点i"的最短路径已成功获取。用此变量可以区分S和U集合。
	flag := make([]bool, vexnum)

	// 初始化
	for i := 0; i < vexnum; i++ {
		// 顶点i的最短路径还没获取到。
		flag[i] = false

		// 顶点i的最短路径为"顶点vs"到"顶点i"的权。
		dist[i] = matrix[vs][i];
	}

	// 对"顶点vs"自身进行初始化
	flag[vs] = true
	dist[vs] = 0

	// 遍历G.vexnum-1次；每次找出一个顶点的最短路径。
	for i := 1; i < vexnum; i++ {
		// 寻找当前最小的路径；
		// 即，在未获取最短路径的顶点中，找到离vs最近的顶点 k。
		min := INF
		var k int
		for j := 0; j < vexnum; j++ {
			if flag[j] == false {
				if dist[j] < min {
					min = dist[j]
					k = j
				}
			}
		}
		// 标记"顶点k"为已经获取到最短路径
		flag[k] = true
		// 修正当前最短路径和前驱顶点
		// 即，当已经"顶点k的最短路径"之后，更新"未获取最短路径的顶点的最短路径和前驱顶点"
		for j := 0; j < vexnum; j++ {
			var tmp = INF
			if matrix[k][j] != INF {
				tmp = min + matrix[k][j]
			}
			if flag[j] == false {
				if tmp < dist[j] {
					dist[j] = tmp
				}
			}
		}
	}

	return dist
}
