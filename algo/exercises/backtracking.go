package main

func generateParenthesis(n int) []string {
	res := make([]string, 0, n)
	backtrack(n, n, &res, "")
	return res
}

func backtrack(left, right int, res *[]string, cur string) {
	if left == 0 && right == 0 {
		*res = append(*res, cur)
	}

	if right < left {
		return
	}

	if left > 0 {
		backtrack(left-1, right, res, cur+"(")
	}

	if right > 0 {
		backtrack(left, right-1, res, cur+")")
	}
}

func main() {
	generateParenthesis(2)
}
