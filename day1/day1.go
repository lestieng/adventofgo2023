package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseline(first string, second string) int {
    if first == "" {
	log.Fatal("first is empty wut...")
    } 
    if second == "" { //only one digit in line
	second = first
    }
    num, err := strconv.Atoi(first + second)
    if err != nil {
	log.Fatal("problem parsing concat string: ",err)
    }
    return num
}

func part1_logic (line string) int { 
    first,second := "",""
    for _,c := range line { // for each character
	if _,err := strconv.Atoi(string(c)); err != nil {
	    continue 	
	}
	if first == "" {
	    first = string(c)
	} else {
	    second = string(c) 
	} 
    }
    res := parseline(first,second)
    return res
}


func part2_logic(line string, digits []string, numdigs int) int {
    num1,num2 := -1,-1
    first,second := "",""
    for i,digit := range digits {
	ind := strings.Index(line,digit)
	if ind < 0 { continue }
	if num1 == -1 || ind < num1 {
	    num1 = ind 
	    first = strconv.Itoa(i%numdigs + 1)
	}
	if ind2 := strings.LastIndex(line,digit); ind2 > num2 {
	    num2 = ind2
	    second = strconv.Itoa(i%numdigs + 1)
	}
    }
    res := parseline(first,second)
    return res
}

func main(){
    file,err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal("problem parsing first cmd line arg (or opening file)",err)
    }
    defer file.Close()

    part,err := strconv.Atoi(os.Args[2])
    if err != nil {
        log.Fatal("problem parsing second cmd line arg",err)
    }

    ans := 0
    digits := [...]string{"one", "two", "three", "four", "five", "six",
			"seven", "eight", "nine", "1", "2", "3", "4", "5",
			"6", "7", "8", "9"}
    for scanner := bufio.NewScanner(file); scanner.Scan(); { //for each line
	if part == 1 {
	    ans += part2_logic(scanner.Text(),digits[9:18],9) //part1_logic(scanner.Text())
	} else {
	    ans += part2_logic(scanner.Text(),digits[:],9)
	}
    }
    fmt.Println("The total sum is: ", ans)
}
