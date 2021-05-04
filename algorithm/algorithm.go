package algorithm

func BubbleSort(data []int) []int {
	n := len(data)
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if data[j] < data[j-1] {
				dummy := data[j]
				data[j] = data[j-1]
				data[j-1] = dummy
			}
		}
	}
	return data
}

func InsertionSort(data []int) []int {
	n := len(data)
	for i := 1; i < n; i++ {
		for k := i; k > 0; k-- {
			if data[k] < data[k-1] {
				dummy := data[k]
				data[k] = data[k-1]
				data[k-1] = dummy
			}
		}
	}
	return data
}

func SelectionSort(data []int) []int {
	n := len(data)
	for i := 0; i < n; i++ {
		k := i
		for j := i + 1; j < n; j++ {
			if data[j] < data[k] {
				k = j
			}
		}
		dummy := data[i]
		data[i] = data[k]
		data[k] = dummy
	}
	return data
}

func ShellSort(data []int) []int {
	var gap [8]int
	n := len(data)
	k := 0
	gap[0] = (int)(n / 2)
	for gap[k] > 1 {
		k++
		gap[k] = (int)(gap[k-1] / 2)
	}
	for i := 0; i <= k; i++ {
		step := gap[i]
		for j := step; j < n; j++ {
			temp := data[j]
			p := j - step
			for p >= 0 && temp < data[p] {
				data[p+step] = data[p]
				p = p - step

			}
			data[p+step] = temp
		}
	}
	return data
}
