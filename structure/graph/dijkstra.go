package graph

//dijkstra算法
//参考文章：http://www.cnblogs.com/skywang12345/p/3711516.html
func Dijkstra(matrix [][]int, v0 int) []int {
	//int i,j,k;
	//int min;
	//int tmp;
	//int flag[MAX];      // flag[i]=1表示"顶点vs"到"顶点i"的最短路径已成功获取。
	//
	//// 初始化
	//for (i = 0; i < G.vexnum; i++)
	//{
	//flag[i] = 0;              // 顶点i的最短路径还没获取到。
	//prev[i] = 0;              // 顶点i的前驱顶点为0。
	//dist[i] = G.matrix[vs][i];// 顶点i的最短路径为"顶点vs"到"顶点i"的权。
	//}
	//
	//// 对"顶点vs"自身进行初始化
	//flag[vs] = 1;
	//dist[vs] = 0;
	//
	//// 遍历G.vexnum-1次；每次找出一个顶点的最短路径。
	//for (i = 1; i < G.vexnum; i++)
	//{
	//// 寻找当前最小的路径；
	//// 即，在未获取最短路径的顶点中，找到离vs最近的顶点(k)。
	//min = INF;
	//for (j = 0; j < G.vexnum; j++)
	//{
	//if (flag[j]==0 && dist[j]<min)
	//{
	//min = dist[j];
	//k = j;
	//}
	//}
	//// 标记"顶点k"为已经获取到最短路径
	//flag[k] = 1;
	//
	//// 修正当前最短路径和前驱顶点
	//// 即，当已经"顶点k的最短路径"之后，更新"未获取最短路径的顶点的最短路径和前驱顶点"。
	//for (j = 0; j < G.vexnum; j++)
	//{
	//tmp = (G.matrix[k][j]==INF ? INF : (min + G.matrix[k][j])); // 防止溢出
	//if (flag[j] == 0 && (tmp  < dist[j]) )
	//{
	//dist[j] = tmp;
	//prev[j] = k;
	//}
	//}
	//}
	//
	//// 打印dijkstra最短路径的结果
	//printf("dijkstra(%c): \n", G.vexs[vs]);
	//for (i = 0; i < G.vexnum; i++)
	//printf("  shortest(%c, %c)=%d\n", G.vexs[vs], G.vexs[i], dist[i]);

	return nil;
}
