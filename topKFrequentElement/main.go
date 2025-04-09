package topKFrequentElement

import "fmt"

func main() {
	nums := []int{1, 1, 2, 3, 4, 4, 5}
	fmt.Printf("%v\n", topKFrequent(nums, 2))
}

type ratedNum struct {
	num   int
	count int
}

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]*int)
	top := make([]int, 0, k)

	for _, num := range nums {
		if counts[num] == nil {
			counts[num] = new(int)
			*counts[num] += 1
		}
		*counts[num]++
	}

	for i := 0; i < k; i++ {
		maxRatedNum := getMax(counts)
		top = append(top, maxRatedNum.num)
		counts[maxRatedNum.num] = nil
	}

	return top
}

func getMax(counts map[int]*int) (numMax ratedNum) {
	for num, count := range counts {
		if count == nil {
			continue
		}

		if *count > numMax.count {
			numMax.num = num
			numMax.count = *count
		}
	}
	return numMax
}
