package main

import (
	"fmt"
	"strings"
)

func WeirdAlgo(n int) {
	fmt.Print(n, " ")
	if n == 1 {
		return
	} else if n%2 == 0 {
		WeirdAlgo(n / 2)
	} else {
		WeirdAlgo(n*3 + 1)
	}
}

func MissingNumber(n int, nums []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	missing := 0

	for i := 1; i <= n; i++ {
		if !m[i] {
			missing = i
		}
	}
	return missing
}

func Repetition(s string) int {
	i := 0
	n := len(s)
	m := make(map[byte]int)
	mx := 0
	for j := 0; j < n; j++ {
		m[s[j]]++
		if len(m) == 1 {
			mx = max(mx, m[s[j]])
		}
		if len(m) > 1 {
			for len(m) > 1 {
				m[s[i]]--
				if m[s[i]] == 0 {
					delete(m, s[i])
				}
				i++
			}
		}
	}
	return mx
}

func IncreasingArray(nums []int) int {
	count := 0
	for j := 0; j < len(nums)-1; j++ {
		if nums[j] > nums[j+1] {
			count += nums[j] - nums[j+1]
		}
	}
	return count
}

func Permutations(n int) {
	var even, odd []string
	if n == 2 || n == 3 {
		fmt.Println("No Solution")
		return
	}
	if n == 1 {
		fmt.Println(n)
	}
	for i := n; i > 0; i-- {
		if i%2 == 0 {
			even = append(even, fmt.Sprintf("%d", i))
		}
	}
	for i := n; i > 0; i-- {
		if i%2 != 0 {
			odd = append(odd, fmt.Sprintf("%d", i))
		}
	}
	fmt.Println(strings.Join(append(even, odd...), " "))

}

func main() {
	WeirdAlgo(3)

	miss := []int{3, 2, 5, 1, 7}
	fmt.Println()
	fmt.Println("Missing number = ", MissingNumber(5, miss))

	s := "ATTCGGGA"
	fmt.Println("Most number of times a word Repeated : ", Repetition(s))

	arr := []int{3, 2, 5, 1, 7}
	fmt.Println("Increasing array : ", IncreasingArray(arr))

	Permutations(5)
}
