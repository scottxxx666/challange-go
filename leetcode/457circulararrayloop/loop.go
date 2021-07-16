package _57circulararrayloop

func circularArrayLoop(nums []int) bool {
	isVisit := make([]int, len(nums))
	for j := 0; j < len(nums); j++ {
		isVisit[j] = -1
	}
	for i := 0; i < len(nums); i++ {
		direction := 1
		if nums[i] < 0 {
			direction = -1
		}

		prev := -1
		temp := i
		res := true
		for isVisit[temp] != i {
			if isVisit[temp] != -1 {
				res = false
				break
			}
			if nums[temp]*direction < 0 {
				res = false
				break
			}
			isVisit[temp] = i
			if prev == temp {
				res = false
				break
			}

			prev = temp
			temp = (temp + nums[temp]%len(nums) + len(nums)) % len(nums)
		}

		if prev != temp && res {
			return true
		}
	}
	return false
}
