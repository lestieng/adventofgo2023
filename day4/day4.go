package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse_get_score(line string, part int) (score int) {
    base_split := strings.Split(line,"|")
    first,second := strings.Split(base_split[0],":")[1],base_split[1]
    winning := make(map[int]int,0)
    // get winning cards and put in (hash)map
    for _,num := range strings.Split(strings.TrimSpace(first)," ") {
        parsednum,_ := strconv.Atoi(strings.TrimSpace(num))
        winning[parsednum] = 0
    }
    // check if my cards are among winning cards
    for _,num := range strings.Split(strings.TrimSpace(second)," ") {
        if num == "" {continue}
        parsednum,_ := strconv.Atoi(strings.TrimSpace(num))
        if _,found := winning[parsednum]; found {
            score += 1
        }
    }
    if part == 1 && score > 0 { 
        return 1<<(score-1) 
    } else {
        return score
    }
}

func cards_seen(cards_won []int,memo []int,start int, stop int) (seen int) {
    for i:=start; i <= stop; i++ {
        seen += 1
        if cards_won[i] > 0 && memo[i] == 0{//haven't been down this subtree
            memo[i] = cards_seen(cards_won,memo,i+1,i+cards_won[i]) 
            seen += memo[i]
        } else if cards_won[i] > 0 {//have been down this subtree, use memo
            seen += memo[i]
        }
    }
    return seen
}

func solution(file *os.File,part int) (ans int) {
        if part == 1 {
            for scanner := bufio.NewScanner(file); scanner.Scan(); {
                ans += parse_get_score(scanner.Text(),1)
            }
            return ans
        } else {
            cards_won := make([]int,0)
            for scanner := bufio.NewScanner(file); scanner.Scan(); {
                cards_won = append(cards_won,parse_get_score(scanner.Text(),2))
            }
            stop := len(cards_won)-1
            memo := make([]int,stop)
            return cards_seen(cards_won,memo,0,stop)
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
   
    fmt.Println("The answer is: ", solution(file,part))
}
