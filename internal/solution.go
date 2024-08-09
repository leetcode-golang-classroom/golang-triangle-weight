package internal

func FindLargestWeight(triangle [][]int) int {
	// do dfs search
	height := len(triangle)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dfs func(level int, pos int) int
	dfs = func(level int, pos int) int {
		if level == height || pos < 0 && pos >= len(triangle[level]) {
			return 0
		}
		val := triangle[level][pos]
		return val + max(dfs(level+1, pos), dfs(level+1, pos+1))
	}
	return dfs(0, 0)
}
