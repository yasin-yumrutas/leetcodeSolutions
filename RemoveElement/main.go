package main

import "fmt"

func removeElement(nums []int, val int) int {
    slow := 0

    for fast := 0; fast < len(nums); fast++ {
        if nums[fast] != val {
            nums[slow] = nums[fast]
            slow++
        }
    }

    return slow
}

func main() {
    nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
    val := 2
    k := removeElement(nums, val)

    fmt.Println("Remaining count:", k)
    fmt.Println("Modified array:", nums[:k]) // İlk k elemanını gösteriyoruz
}
