package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type lineData struct {
    left []int 
    right []int 
    nums []int 
    syms []int
    active bool
}

func parseline(linestr string) (line lineData) {
    concat := false
    numstr := ""
    for i,c := range linestr {
        if unicode.IsDigit(c) {
           if !concat {//start of new number
                line.left = append(line.left,i)
                concat = true
            } 
            numstr += string(c)
        } else {
            if concat {//end of number
                concat = false 
                line.right = append(line.right,i-1)
                num,_ := strconv.Atoi(numstr)
                numstr = ""
                line.nums = append(line.nums,num)
            }
            if c != '.' { line.syms = append(line.syms,i) }
            // assume non '.' is a valid symbol
        }
    }
    if concat {// last number on line ends on the right edge
        line.right = append(line.right,len(linestr)-1)
        num,_ := strconv.Atoi(numstr)
        line.nums = append(line.nums,num)
    }
    line.active = true
    return line
}

func parseline2(linestr string) (line lineData) {
    concat := false
    numstr := ""
    for i,c := range linestr {
        if unicode.IsDigit(c) {
           if !concat {
                line.left = append(line.left,i)
                concat = true
            } 
            numstr += string(c)
        } else {
            if concat {
                concat = false 
                line.right = append(line.right,i-1)
                num,_ := strconv.Atoi(numstr)
                numstr = ""
                line.nums = append(line.nums,num)
            }
            if c == '*' { line.syms = append(line.syms,i) }
        }
    }
    if concat { 
        line.right = append(line.right,len(linestr)-1)
        num,_ := strconv.Atoi(numstr)
        line.nums = append(line.nums,num)
    }
    line.active = true
    return line 
}

func valid_number_logic(left int,right int,syms []int) bool {
    // valid number (between left and right) could have symbol(s) 
    // on lines above or below at line positions left-1 to right+1
    // or on same line at left-1 and/or right+1
    for _,sym := range syms {
        if sym >= left-1 && sym <= right+1 {
            return true
        }
    }
    return false
}

func find_valid_number(prev lineData, curr lineData, next lineData) (
                        line_ans int) {
    for i,num := range curr.nums {
        adjacent := valid_number_logic(curr.left[i],curr.right[i],curr.syms)
        if !adjacent && next.active {
            adjacent = valid_number_logic(curr.left[i],curr.right[i],next.syms)
        }
        if !adjacent && prev.active {
            adjacent = valid_number_logic(curr.left[i],curr.right[i],prev.syms)
        }
        if adjacent { line_ans +=  num }
    }
    return line_ans
}

func gear_ratio_logic(numleft []int,numright []int,nums []int,sym int,
                            gear_nums []int) ( gear_nums_out []int) {
    gear_nums_out = gear_nums
    for j,left := range numleft {
        if sym >= left-1 && sym <= numright[j]+1 {
            gear_nums_out = append(gear_nums_out,nums[j])
        } 
        if len(gear_nums_out) > 2 { break }
    }
    return gear_nums_out
}

func find_gear_ratio(prev lineData, curr lineData, next lineData) (
                        line_ans int) {
    for _,sym := range curr.syms {
        gear_nums := gear_ratio_logic(curr.left,curr.right,curr.nums,sym,nil)
        if len(gear_nums) < 3 && next.active {
            gear_nums = gear_ratio_logic(next.left,next.right,
                                        next.nums,sym,gear_nums) 
        }
        if len(gear_nums) < 3 && prev.active {
            gear_nums = gear_ratio_logic(prev.left,prev.right,
                                        prev.nums,sym,gear_nums) 
        }
        if len(gear_nums) == 2 { line_ans += gear_nums[0]*gear_nums[1] }
    }
    return line_ans
}

func solution1(scanner *bufio.Scanner) (ans int) {
    // first line: read two lines and parse
    _ = scanner.Scan()
    current_line := scanner.Text()
    current := parseline(current_line) 
    _ = scanner.Scan()
    next_line := scanner.Text()
    var prev lineData; prev.active = false
    next := parseline(next_line)
    ans += find_valid_number(prev,current,next)
    for ; ; {//for rest of lines; at start each iter: current=prev
        if !scanner.Scan() { // on last line
            prev := parseline(current_line)
            current := parseline(next_line)
            next.active = false
            ans += find_valid_number(prev,current,next)
            break
        } else {
            prev := parseline(current_line)
            current_line = next_line
            current := parseline(current_line)
            next_line = scanner.Text()
            next := parseline(next_line)
            ans += find_valid_number(prev,current,next)
        }
    }
    return ans
}

func solution2(scanner *bufio.Scanner) (ans int) {
    // first line: read two lines and parse
    _ = scanner.Scan()
    current_line := scanner.Text()
    current := parseline2(current_line)
    _ = scanner.Scan()
    next_line := scanner.Text()
    next := parseline2(next_line)
    var prev lineData; prev.active = false
    ans += find_gear_ratio(prev,current,next)

    for ; ; {//for rest of lines; at start each iter: current=prev
        if !scanner.Scan() { // on last line
            prev := parseline2(current_line)
            current := parseline2(next_line)
            next.active = false
            ans += find_gear_ratio(prev,current,next)
            break
        } else {
            prev := parseline2(current_line)
            current_line = next_line
            current := parseline2(current_line)
            next_line = scanner.Text()
            next := parseline2(next_line)
            ans += find_gear_ratio(prev,current,next)
        }
    }
    return ans
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

    if part == 1 {
        fmt.Println("The answer is: ", solution1(scanner))
    } else {
        fmt.Println("The answer is: ", solution2(scanner))
    }
}
