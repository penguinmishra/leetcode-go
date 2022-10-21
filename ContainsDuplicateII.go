package main
import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
    m := make(map[int]int)
    for i, v := range nums {
		val, present := m[v]
		if present && i - val <= k {
			return true
		}
		m[v] = i
	}
	return false
}

func main(){
	nums := []int{1,2,3,1,2,3}
	k := 2
	fmt.Println(containsNearbyDuplicate(nums, k))
}