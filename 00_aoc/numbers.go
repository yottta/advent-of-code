package _0_aoc

func SumList(l []int) int {
	var sum int
	for _, i := range l {
		sum += i
	}
	return sum
}
