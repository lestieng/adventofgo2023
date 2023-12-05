package main

import (
	"bufio"
	"strings"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type rangeMap struct {
    input_ranges [][]int
    output_ranges [][]int
    from string
    to string
}

func read_seeds1(seeds []string) (seednums []int) {
    seednums = make([]int,0)
    for _,seed := range seeds { 
        temp,_ := strconv.Atoi(seed)
        seednums = append(seednums,temp)
    }
    return seednums
}

func build_rangemap(scanner *bufio.Scanner) (newMap rangeMap) {
    _ = scanner.Scan()
    namereg := regexp.MustCompile(`\w+`)
    line := strings.ReplaceAll(scanner.Text(),"-"," ")
    names := namereg.FindAllString(line,-1)
    newMap.from = names[0]
    newMap.to = names[2]
    reg := regexp.MustCompile(`\d+`)
    for ; scanner.Scan(); {
        line = scanner.Text()
        if line == "" { break }
        nums := reg.FindAllString(line,-1)
        start_out, _ := strconv.Atoi(nums[0])
        start_in, _ := strconv.Atoi(nums[1])
        length, _ := strconv.Atoi(nums[2])
        newMap.input_ranges = append(newMap.input_ranges,[]int{start_in,start_in+length-1})
        newMap.output_ranges = append(newMap.output_ranges,[]int{start_out,start_out+length-1})
    }
    return newMap
}

func build_loop(scanner *bufio.Scanner) (myRange []rangeMap, myMap map[string]int){
    myRange = make([]rangeMap,0)
    myMap = make(map[string]int,0)
    testRange := build_rangemap(scanner)
    myRange = append(myRange,testRange)
    myMap[testRange.from] = 0
    // fmt.Printf("%+v\n",myRange[0])
    for i:=0; ;i++ {
        testRange := build_rangemap(scanner)
        myRange = append(myRange,testRange)
        myMap[testRange.from] = i+1
        // fmt.Printf("%+v\n",myRange[i+1])
        if testRange.to == "location" { break }
    }
    return myRange,myMap
}

func translate(input int, to rangeMap) (output int) {
    for i,startstop := range to.input_ranges {
        if input >= startstop[0] && input <= startstop[1] {
            return to.output_ranges[i][0]+(input-startstop[0])
        }
    }
    return input
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
    first_line := scanner.Text()
    reg := regexp.MustCompile(`\d+`)
    seeds := reg.FindAllString(first_line,-1)
    _ = scanner.Scan()
    myRange,_ := build_loop(scanner)
    var location int
    if part == 1 {
        seednums := read_seeds1(seeds)
        for i,seed := range seednums {
            val := seed
            for _,rmaps := range myRange {
                val = translate(val,rmaps)
            }
            if i==0 {
                location = val
            } else {
                location = min(location,val)
            }
        }
    } else {
        first := true
        for i:=0; i<len(seeds); {
            seedstart,_ := strconv.Atoi(seeds[i])
            i++
            seedlength,_ := strconv.Atoi(seeds[i])
            i++
            for seed := seedstart; seed<seedstart+seedlength-1; seed++ {
                val := seed
                for _,rmaps := range myRange {
                    val = translate(val,rmaps)
                }
                if first {
                    location = val
                    first = false
                } else {
                    location = min(location,val)
                }

            }
        } 
    }
    // fmt.Println(translate(seednums[0],myRange[myMap["seed"]]))
   
    fmt.Println("The answer is: ", location)
}

