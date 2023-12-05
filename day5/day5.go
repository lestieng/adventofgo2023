package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type rangeMap struct {
    input_starts []int
    input_stops []int
    increments []int

}

type multiRange struct {
    starts []int
    stops []int
}

// lazy bubble sort implementation, no built-in argsort 
// so need to do it manually in order to re-order the other slices
func rangemap_sort(myMap *rangeMap) {
    swaps := 1
    for ; swaps > 0; { // while swaps > 0
        swaps = 0
        for i:=0; i<len(myMap.input_starts)-1; i++ {
            if myMap.input_starts[i] > myMap.input_starts[i+1] {
                myMap.input_starts[i], myMap.input_starts[i+1] = 
                    myMap.input_starts[i+1], myMap.input_starts[i]
                myMap.input_stops[i], myMap.input_stops[i+1] = 
                    myMap.input_stops[i+1], myMap.input_stops[i]
                myMap.increments[i], myMap.increments[i+1] = 
                    myMap.increments[i+1], myMap.increments[i]
                swaps++
            }
        }
    }
}

func read_seeds(seeds []string) (seednums []int) {
    seednums = make([]int,0)
    for _,seed := range seeds { 
        temp,_ := strconv.Atoi(seed)
        seednums = append(seednums,temp)
    }
    return seednums
}

func build_rangemap(scanner *bufio.Scanner) (newMap rangeMap, to string, maxval int) {
    _ = scanner.Scan()
    namereg := regexp.MustCompile(`\w+`)
    line := strings.ReplaceAll(scanner.Text(),"-"," ")
    names := namereg.FindAllString(line,-1)
    to = names[2]
    reg := regexp.MustCompile(`\d+`)
    for ; scanner.Scan(); { // this loops breaks on EOF or a blank line
        line = scanner.Text()
        if line == "" { break }
        nums := reg.FindAllString(line,-1)
        start_out, _ := strconv.Atoi(nums[0])
        start_in, _ := strconv.Atoi(nums[1])
        length, _ := strconv.Atoi(nums[2])
        maxval = max(maxval,start_out+length-1)
        newMap.input_starts = append(newMap.input_starts,start_in)
        newMap.input_stops = append(newMap.input_stops,start_in+length-1)
        newMap.increments = append(newMap.increments,start_out-start_in)
    }
    rangemap_sort(&newMap) 
    return newMap,to,maxval
}

func build_loop(scanner *bufio.Scanner) (myRange []rangeMap, maxval int){
    myRange = make([]rangeMap,0)
    testRange,to,_ := build_rangemap(scanner)
    myRange = append(myRange,testRange)
    for i:=0; ;i++ {
        testRange,to,maxval = build_rangemap(scanner)
        myRange = append(myRange,testRange)
        if to == "location" { break }
    }
    // maxval is the largest value in the final output mapping
    return myRange,maxval
}

func translate(input int, to rangeMap) (output int) {
    for i,start := range to.input_starts {
        if input >= start && input <= to.input_stops[i] {
            return input + to.increments[i]
        }
    }
    return input
}

func partition(from *multiRange,ind int,part int) {
    if ind == len(from.starts)-1 {
        from.starts = append(from.starts,part+1)
        from.stops = append(from.stops,from.starts[ind])
        from.stops[ind] = part
    } else {
        from.starts = slices.Insert(from.starts,ind+1,part+1)
        from.stops = slices.Insert(from.stops,ind,part)
    }
}

func translate_range(from *multiRange, to rangeMap) {
    // check for overlap of range in from with range from to 
    // if partial overlap, partition the range in from
    for k := 0; k<len(from.starts); k++ { // loop extends if partitioned
        for i,start := range to.input_starts {
            tostops := to.input_stops[i]
            if from.starts[k] >= start && 
                from.starts[k] <= tostops {
                if from.stops[k] <= tostops {
                    from.starts[k] += to.increments[i]
                    from.stops[k] += to.increments[i]
                } else {
                    partition(from,k,tostops)
                    from.starts[k] += to.increments[i]
                    from.stops[k] += to.increments[i]
                }
                break
            }
        }
    }
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
    var location int
    if part == 1 {
        seednums := read_seeds(seeds)
        myRange,maxval := build_loop(scanner)
        location = maxval
        for _,seed := range seednums {
            val := seed
            for _,rmaps := range myRange {
                val = translate(val,rmaps)
            }
            location = min(location,val)
        }
    } else {
        myRange,maxval := build_loop(scanner)
        location = maxval
        for i:=0; i<len(seeds); i+=2 {
            seedstart,_ := strconv.Atoi(seeds[i])
            seedlength,_ := strconv.Atoi(seeds[i+1])
            seedRange := multiRange {
                starts: append([]int{},seedstart),
                stops:  append([]int{},seedstart+seedlength-1),
            }
            for _,rmaps := range myRange {
                translate_range(&seedRange,rmaps)
            }
            location = min(location,slices.Min(seedRange.starts))
        } 
    }
    fmt.Println("The answer is: ", location)
}

