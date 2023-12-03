package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func parseline(line string) (numleft []int,numright []int,
                            nums []int,syms []int) {
    concat := false
    numstr := ""
    for i,c := range line {
        if unicode.IsDigit(c) {
           if !concat {//start of new number
                numleft = append(numleft,i)
                concat = true
            } 
            numstr += string(c)
        } else {
            if concat {//end of number
                concat = false 
                numright = append(numright,i-1)
                num,_ := strconv.Atoi(numstr)
                numstr = ""
                nums = append(nums,num)
            }
            if c != '.' { syms = append(syms,i) }// assume non '.' is a valid symbol
        }
    }
    if concat {// last number on line ends on the right edge
        numright = append(numright,len(line)-1)
        num,_ := strconv.Atoi(numstr)
        nums = append(nums,num)
    }
    return numleft,numright,nums,syms
}

func parseline2(line string) (numleft []int,numright []int,
                            nums []int,syms []int) {
    concat := false
    numstr := ""
    for i,c := range line {
        if unicode.IsDigit(c) {
           if !concat {
                numleft = append(numleft,i)
                concat = true
            } 
            numstr += string(c)
        } else {
            if concat {
                concat = false 
                numright = append(numright,i-1)
                num,_ := strconv.Atoi(numstr)
                numstr = ""
                nums = append(nums,num)
            }
            if c == '*' { syms = append(syms,i) }
        }
    }
    if concat { 
        numright = append(numright,len(line)-1)
        num,_ := strconv.Atoi(numstr)
        nums = append(nums,num)
    }
    return numleft,numright,nums,syms
}

func valid_number_logic(left int,right int,syms []int) bool {
    // valid number could have symbol(s) on lines above or below
    // at line positions start-1 to end+1
    // or on same line at start-1 and/or end+1
    for _,sym := range syms {
        if sym >= left-1 && sym <= right+1 {
            return true
        }
    }
    return false
}

func find_valid_number(numleft []int,numright[]int,nums []int,
                        syms_prev []int,syms_curr []int,syms_next []int, mode int,
                        ) (line_ans int) {
    // mode = 0, first line; mode = 1, normal; mode = 2, last line
    for i,num := range nums {
        adjacent := valid_number_logic(numleft[i],numright[i],syms_curr)
        if !adjacent && mode <= 1 {
            adjacent = valid_number_logic(numleft[i],numright[i],syms_next)
        }
        if !adjacent && mode > 0 {
            adjacent = valid_number_logic(numleft[i],numright[i],syms_prev)
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

func find_gear_ratio(numleft_prev []int,numleft []int,numleft_next []int, 
                    numright_prev []int,numright []int,numright_next [] int,
                    nums_prev []int,nums []int,nums_next []int,
                    syms []int, mode int) (line_ans int) {
    // mode = 0, first line; mode = 1, normal; mode = 2, last line
    for _,sym := range syms {
        gear_nums := gear_ratio_logic(numleft,numright,nums,sym,nil)
        if len(gear_nums) < 3 && mode <= 1 {
            gear_nums = gear_ratio_logic(numleft_next,numright_next,
                                            nums_next,sym,gear_nums) 
        }
        if len(gear_nums) < 3 && mode > 0 {
            gear_nums = gear_ratio_logic(numleft_prev,numright_prev,
                                            nums_prev,sym,gear_nums) 
        }
        if len(gear_nums) == 2 { line_ans += gear_nums[0]*gear_nums[1] }
    }
    return line_ans
}

func solution1(scanner *bufio.Scanner) (ans int) {
    // first line: read two lines and parse
    _ = scanner.Scan()
    current_line := scanner.Text()
    numleft,numright,nums,syms_curr := parseline(current_line)
    _ = scanner.Scan()
    next_line := scanner.Text()
    _,_,_,syms_next := parseline(next_line)
    ans += find_valid_number(numleft,numright,nums,[]int{0},
                                    syms_curr,syms_next,0)
    for ; ; {//for rest of lines; at start each iter: current=prev
        if !scanner.Scan() { // on last line
            _,_,_,syms_prev := parseline(current_line)
            numleft,numright,nums,syms_curr := parseline(next_line)
            ans += find_valid_number(numleft,numright,nums,
                            syms_prev,syms_curr,syms_next,2)
            break
        } else {
            _,_,_,syms_prev := parseline(current_line)
            current_line = next_line
            numleft,numright,nums,syms_curr := parseline(current_line)
            next_line = scanner.Text()
            _,_,_,syms_next := parseline(next_line)
            ans += find_valid_number(numleft,numright,nums,
                            syms_prev,syms_curr,syms_next,1)
        }
    }
    return ans
}

func solution2(scanner *bufio.Scanner) (ans int) {
    // first line: read two lines and parse
    _ = scanner.Scan()
    current_line := scanner.Text()
    numleft,numright,nums,syms_curr := parseline2(current_line)
    _ = scanner.Scan()
    next_line := scanner.Text()
    numleft_next,numright_next,nums_next,_ := parseline2(next_line)
    ans += find_gear_ratio([]int{0},numleft,numleft_next,[]int{0},
                                numright,numright_next,[]int{0},
                                nums,nums_next,syms_curr,0)

    for ; ; {//for rest of lines; at start each iter: current=prev
        if !scanner.Scan() { // on last line
            numleft_prev,numright_prev,nums_prev,_ := parseline2(current_line)
            numleft,numright,nums,syms_curr := parseline2(next_line)
            ans += find_gear_ratio(numleft_prev,numleft,numleft_next,
                                numright_prev,numright,numright_next,
                                nums_prev,nums,nums_next,syms_curr,2)
            break
        } else {
            numleft_prev,numright_prev,nums_prev,_ := parseline2(current_line)
            current_line = next_line
            numleft,numright,nums,syms_curr := parseline2(current_line)
            next_line = scanner.Text()
            numleft_next,numright_next,nums_next,_ := parseline2(next_line)
            ans += find_gear_ratio(numleft_prev,numleft,numleft_next,
                                numright_prev,numright,numright_next,
                                nums_prev,nums,nums_next,syms_curr,1)
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
