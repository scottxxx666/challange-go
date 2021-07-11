package _39dailytemperatures

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	res := make([]int, len(temperatures))
	indexes := make(map[int]int)

	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			res[i] = 0
		} else {
			res[i] = indexes[stack[len(stack)-1]] - i
		}
		stack = append(stack, temperatures[i])
		indexes[temperatures[i]] = i
	}
	return res
}
