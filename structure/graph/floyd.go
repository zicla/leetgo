package graph

//Floyd算法
//参考：http://blog.51cto.com/ahalei/1383613
func Floyd(matrix [][]int, source int, dest int) int {

	N := len(matrix)

	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if matrix[i][j] > matrix[i][k]+matrix[k][j] {
					matrix[i][j] = matrix[i][k] + matrix[k][j];
				}
			}
		}
	}

	return matrix[source][dest]

}
