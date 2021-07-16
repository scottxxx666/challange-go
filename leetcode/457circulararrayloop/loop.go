package _57circulararrayloop

func circularArrayLoop(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		isVisit := make([]bool, len(nums))
		for j := 0; j < len(nums); j++ {
			isVisit[j] = false
		}
		direction := 1
		if nums[i] < 0 {
			direction = -1
		}

		prev := -1
		temp := i
		res := true
		for !isVisit[temp] {
			isVisit[temp] = true
			if prev == temp {
				res = false
				break
			}
			if nums[temp]*direction < 0 {
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
