package main

import "fmt"

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    slow := 0

    for fast := 1; fast < len(nums); fast++ {
        if nums[fast] != nums[slow] {
            slow++
            nums[slow] = nums[fast]
        }
    }
    
    return slow + 1
}

func main() {
    nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
    k := removeDuplicates(nums)
    fmt.Println("Unique count:", k)
    fmt.Println("Modified array:", nums[:k]) // İlk k elemanı gösteriyoruz
}
