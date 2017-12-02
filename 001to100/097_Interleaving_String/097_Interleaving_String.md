097. Interleaving String

Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.

For example,
Given:
s1 = "aabcc",
s2 = "dbbca",

When s3 = "aadbbcbcac", return true.
When s3 = "aadbbbaccc", return false.


```

暴力递归方法

func isInterleave(s1 string, s2 string, s3 string) bool {

	N1 := len(s1)
	N2 := len(s2)
	N3 := len(s3)
	if N1+N2 != N3 {
		return false
	}

	result := []bool{false}

	isInterleaveRecursive(s1, N1, 0, s2, N2, 0, s3, N3, 0, result)

	return result[0];
}

func isInterleaveRecursive(s1 string, N1 int, n1 int, s2 string, N2 int, n2 int, s3 string, N3 int, n3 int, result []bool) {

	if result[0] {
		return
	}

	if n3 == N3 {
		result[0] = true
	}

	if n1 < N1 && s3[n3] == s1[n1] {

		isInterleaveRecursive(s1, N1, n1+1, s2, N2, n2, s3, N3, n3+1, result)
	}

	if n2 < N2 && s3[n3] == s2[n2] {
		isInterleaveRecursive(s1, N1, n1, s2, N2, n2+1, s3, N3, n3+1, result)
	}

}



```