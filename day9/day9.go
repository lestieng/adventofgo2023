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

func differences(nums []int) (diffs []int) {
    for i:=1; i<len(nums); i++ {
        diffs = append(diffs,nums[i]-nums[i-1])
    }
    return diffs
}
//Need to Clone because sometimes Compact changes underlying slice ?????
func predict_next_in_sequence(nums []int) int {
    if len(slices.Compact(slices.Clone(nums))) == 1 { return nums[len(nums)-1] }
    return nums[len(nums)-1] + predict_next_in_sequence(differences(nums))
}

func predict_previous_in_sequence(nums []int) int {
    if len(slices.Compact(slices.Clone(nums))) == 1 { return nums[0] }
    return nums[0] - predict_previous_in_sequence(differences(nums))
}

func main() {   
    file,part := lib.File_part_from_cmdline(os.Args)
    defer file.Close()

    reg := regexp.MustCompile(`-?\d+`)
    ans := 0 
    if part == 1{
        for scanner := bufio.NewScanner(file); scanner.Scan(); {
            ans += predict_next_in_sequence(get_nums(scanner.Text(),reg))
        }
    } else {
        for scanner := bufio.NewScanner(file); scanner.Scan(); {
            ans += predict_previous_in_sequence(get_nums(scanner.Text(),reg))
        }
    }
    fmt.Println("The answer is ",ans)
} 
