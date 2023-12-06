package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
	// "slices"
	"strconv"
)

func get_numbers(line string,part int) (nums []float64) {
    if part != 1 {
        line = strings.ReplaceAll(line," ","")
    }
    reg := regexp.MustCompile(`\d+`)
    num_strs := reg.FindAllString(line,-1)
    for _,str := range num_strs {
        num,_ := strconv.ParseFloat(str,64)
        nums = append(nums,num)
    }
    return nums
}


func main() {
    file,err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal("problem parsing first cmd line arg (or opening file)",err)
    }
    defer file.Close()

    part,err := strconv.Atoi(os.Args[2])
    if err != nil {
        log.Fatal("problem parsing second cmd line arg",err)
    }

    scanner := bufio.NewScanner(file)
    _ = scanner.Scan()
    times := get_numbers(scanner.Text(),part)
    _ = scanner.Scan()
    distances := get_numbers(scanner.Text(),part)
    
    ans := 1
    for i,time := range times { // solve second degree equation for speed zero-crossing
        if discriminant := time*time - 4.0*distances[i]; discriminant > 0 {
            discr_sqrt := math.Sqrt(discriminant)
            term1 := int(math.Ceil(0.5*(time - discr_sqrt)))
            term2 := int(math.Floor(0.5*(time + discr_sqrt)))
            // test for edge case when you zero-cross at an integer exactly
            // when sqrt is an integer AND time is even (so the sqrt and sum is even)
            if discr_sqrt - math.Round(discr_sqrt) < math.SmallestNonzeroFloat64 && 
                math.Mod(time,2.0) < math.SmallestNonzeroFloat64 {
                ans *= term2 - term1 - 1
            } else {
                ans *= term2 - term1 + 1
            }
        }
    }

    fmt.Println("The answer is: ",ans)


}
