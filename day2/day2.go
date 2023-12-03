package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func get_int_from_str(str string) (num int) {
    _,err := fmt.Sscanf(strings.TrimSpace(str),"Game %d",&num)
    if err != nil {
        log.Fatal("failed to read num,color pair ",err)
    }
    return num
}

func get_pair_from_str(str string) (color string,num int) {
    _,err := fmt.Sscanf(strings.TrimSpace(str),"%d %s",&num,&color)
    if err != nil {
        log.Fatal("failed to read game num ",err)
    }
    color,_ = strings.CutSuffix(color,",")
    return color,num
}

func solution1(line string, totalColors map[string]int) int { 
    base_split := strings.Split(line,":")
    game_info,cube_info := base_split[0],base_split[1]  
    shownColors := make(map[string]int) 
    for key := range totalColors { shownColors[key] = 0 }

    for _,sub := range strings.Split(cube_info,";"){
        for _,sub_sub := range strings.Split(sub,",") {
            color,num := get_pair_from_str(sub_sub)
            shownColors[color] = num
        }
        for key,val := range shownColors {
            if val > totalColors[key] { return 0 }
            shownColors[key] = 0
        }
    }
    return get_int_from_str(game_info)
}

func solution2(line string, totalColors map[string]int) int { 
    base_split := strings.Split(line,":")
    cube_info := base_split[1]  
    shownColors := make(map[string]int) 
    for key := range totalColors { shownColors[key] = 0 }

    for _,sub := range strings.Split(cube_info,";"){
        for _,sub_sub := range strings.Split(sub,",") {
            color,num := get_pair_from_str(sub_sub)
            shownColors[color] = max(num, shownColors[color])
        }
    }
    ans := 1
    for _,val := range shownColors { ans *= val }
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
    
    ans := 0
    totalColors := map[string]int{"red": 12, "green": 13, "blue": 14}

    for scanner := bufio.NewScanner(file); scanner.Scan(); {//for each line 
        if part == 1 {
            ans += solution1(scanner.Text(),totalColors)
        } else {
            ans += solution2(scanner.Text(),totalColors)
        }
    }
    fmt.Println("The answer is: ", ans)
}
