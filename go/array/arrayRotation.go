package array
import "fmt"

func rotate(nums []int, k int)  {
    n := len(nums)
    if k > n {
		k %= n
	}
    
    if k == 0 || k == n  {
        return 
    }
    
    var compute int = len(nums) - k
    
    A := reverse(nums[: compute])
    B := reverse(nums[compute : ])
    A = append(A, B...)
    
    
    fmt.Println(reverse(A))
   
}



func reverse(arr []int) []int {
    for i, j := 0, len(arr) - 1; i < j; i,j = i +1, j - 1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
    
    return arr
}