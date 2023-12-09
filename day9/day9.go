package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
        "adventofgo2023/lib"
)

func get_nums(line string,reg *regexp.Regexp) (nums []int) {
    for _,num := range reg.FindAllString(line,-1) {
        temp,_ := strconv.Atoi(num)
        nums = append(nums,temp)
    }
    return nums
}

func predict_next_in_sequence(nums []int) int {
    if len(slices.Compact(slices.Clone(nums)))==1 { //checks number of unique
        return nums[len(nums)-1]
    }
    newnums := make([]int,0)
    for i:=1; i<len(nums); i++ {
        newnums = append(newnums,nums[i]-nums[i-1])
    }
    return nums[len(nums)-1] + predict_next_in_sequence(newnums)
}

func predict_previous_in_sequence(nums []int) int {
    if len(slices.Compact(slices.Clone(nums)))==1 {
        return nums[0]
    }
    newnums := make([]int,0)
    for i:=1; i<len(nums); i++ {
        newnums = append(newnums,nums[i]-nums[i-1])
    }
    return nums[0] - predict_previous_in_sequence(newnums)
}

func main() {   
    file,part := lib.File_part_from_cmdline(os.Args)
    defer file.Close()

    reg := regexp.MustCompile(`-?\d+`)
    ans := 0 
    for scanner := bufio.NewScanner(file); scanner.Scan(); {
        if part == 1{
            ans += predict_next_in_sequence(get_nums(scanner.Text(),reg))
        } else {
            ans += predict_previous_in_sequence(get_nums(scanner.Text(),reg))
        }
    }
    fmt.Println("The answer is ",ans)
} 
