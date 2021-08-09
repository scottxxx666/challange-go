package _03kthlargestinstream

type KthLargest struct {
	list []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{list: nums, k: k}
}

func (kth *KthLargest) Add(val int) int {
	kth.list = append(kth.list, val)
	return kth.find()
}

func (kth *KthLargest) find() int {
	i := 0
	j := len(kth.list) - 1
	for i <= j {
		index := kth.index(i, j)
		if index == kth.k-1 {
			return kth.list[index]
		} else if index <= kth.k-1 {
			i = index + 1
		} else {
			j = index - 1
		}
	}
	return -1
}

func (kth *KthLargest) index(start int, end int) int {
	pivot := kth.list[start]
	k := start
	for k <= end {
		if kth.list[k] >= pivot {
			k++
		} else {
			kth.list[k], kth.list[end] = kth.list[end], kth.list[k]
			end--
		}
	}
	kth.list[k-1], kth.list[start] = kth.list[start], kth.list[k-1]
	return k - 1
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
