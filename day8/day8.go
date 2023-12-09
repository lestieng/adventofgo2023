package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
        "adventofgo2023/lib"
)

type Node struct {
    children map[byte]string
    name string
}

func gcd (a int, b int) int {
    if a >= b {
        if b == 0 {return a}
        return gcd(a%b,b)
    }
    if b > a {
        if a == 0 {return b}
        return gcd(a,b%a)
    }
    return 1
}

func lcm(a int,b int) int {
    return a*(b/gcd(a,b))
}

func lcm_multi(nums []int) int {
    if len(nums)==2 {return lcm(nums[0],nums[1])}
    return lcm(nums[0],lcm_multi(nums[1:]))
}

func parse_input(scanner *bufio.Scanner) (Nodes []Node,
                nodeMappings map[string]int,starting_points []int) {
    reg := regexp.MustCompile(`[1-2|A-Z]{3}`) //test case uses digits
    Nodes = make([]Node,0)
    nodeMappings = make(map[string]int)
    starting_points = make([]int,0)
    for i := 0; scanner.Scan(); {
        nodes := reg.FindAllString(scanner.Text(),-1)
        if nodes[0][2] == 'A' { starting_points = append(starting_points,i) }
        nodeMappings[nodes[0]] = i
        Nodes = append(Nodes,Node{children: map[byte]string{
            'L': nodes[1], 'R': nodes[2]},
            name: nodes[0]})
        i++
    }
    return Nodes,nodeMappings,starting_points
}

func main() {
    file,part := lib.File_part_from_cmdline(os.Args)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    instruction := scanner.Text()
    scanner.Scan()
    num_instructions := len(instruction)
    steps := 0
    steps_to_complete := make([]int,0)
    ans := 0
    Nodes,nodeMappings,starts := parse_input(scanner)
    if part == 1 {
        current := nodeMappings["AAA"]
        for ; current != nodeMappings["ZZZ"]; steps++ {
            current_instruction := instruction[steps%num_instructions]
            current = nodeMappings[Nodes[current].children[current_instruction]]
        }
        ans = steps
    } else {
        count := 0
        // Find the steps needed for each path
        for _,current := range starts {
            for ; Nodes[current].name[2] != 'Z'; steps++ {
                current_instruction := instruction[steps%num_instructions]
                current = nodeMappings[Nodes[current].children[current_instruction]]
                count ++
            }
            if len(steps_to_complete) == 0 {
                steps_to_complete = append(steps_to_complete,count)
                count = 0
            } else {
                steps_to_complete = append(steps_to_complete,count)
                count = 0
            }
        }
        // Steps to complete all at the same time is lowest common multiple of
        // the individual steps needed
        ans = lcm_multi(steps_to_complete)
    }
    fmt.Println("The answer is ",ans)
}
