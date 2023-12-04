package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseline(line string) (winning []int, mine []int) {
    base_split := strings.Split(line,"|")
    first,second := strings.Split(base_split[0],":")[1],base_split[1]
    for _,num := range strings.Split(strings.TrimSpace(first)," ") {
        parsednum,_ := strconv.Atoi(strings.TrimSpace(num))
        winning = append(winning,parsednum)    
    }
    for _,num := range strings.Split(strings.TrimSpace(second)," ") {
        if num == "" {continue}
        parsednum,_ := strconv.Atoi(strings.TrimSpace(num))
        mine = append(mine,parsednum)    
    }
    return winning,mine
}

func get_score(winning []int,mine []int,part int) (score int) {
    for _,card := range winning {
        for _,my := range mine {
            if my == card {
                score += 1
            }
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
        if cards_won[i] > 0 && memo[i] == 0{
            memo[i] = cards_seen(cards_won,memo,i+1,i+cards_won[i]) 
            seen += memo[i]
        } else if cards_won[i] > 0 {
            seen += memo[i]
        }
    }
    return seen
}

func solution(file *os.File,part int) (ans int) {
        cards_won := make([]int,0)
        for scanner := bufio.NewScanner(file); scanner.Scan(); {
            winning_cards,my_cards := parseline(scanner.Text())
            ans += get_score(winning_cards,my_cards,1)
            cards_won = append(cards_won,get_score(winning_cards,my_cards,2))
        }
        if part == 1 {
            return ans
        } else {
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
