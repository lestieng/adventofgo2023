package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const ONE_OF_KIND = 0
const ONE_PAIR = 1
const TWO_PAIR = 2
const THREE_OF_KIND = 3
const FULL_HOUSE = 4
const FOUR_OF_KIND = 5
const FIVE_OF_KIND = 6

var ofAKind = map[int]int {
    2: ONE_PAIR, 3: THREE_OF_KIND, 4: FOUR_OF_KIND, 5: FIVE_OF_KIND,
}

type cardHand struct {
    card_scores []int
    bid int 
    type_score int 
    hand string 
}

func set_type_score1(type_tracker []int) (type_score int) {
    switch num := len(type_tracker); num {
    case 0:
        type_score = ONE_OF_KIND
    case 1:
        type_score = ofAKind[type_tracker[0]]
    case 2:
        if type_tracker[0] != type_tracker[1] {
            type_score = FULL_HOUSE
        } else {
            type_score = TWO_PAIR
        }
    }
    return type_score
} 

func set_type_score2(init_score int,type_tracker []int,jcount int) (type_score int) {
    switch ; init_score {
    case ONE_OF_KIND:
        type_score = ONE_PAIR
    case FULL_HOUSE:
        type_score = FIVE_OF_KIND
    case TWO_PAIR:
       if jcount == 1 {
            type_score = FULL_HOUSE
        } else {
            type_score = FOUR_OF_KIND
        }
    case ofAKind[type_tracker[0]]:
        type_score = ofAKind[type_tracker[0]+1]
    default: 
        type_score = init_score
    }
    return type_score
} 

func set_scores(hand string, cardValues map[rune]int) (
                type_score int, card_scores []int) {
    type_tracker := make([]int,0)
    remaining := hand
    for _,c := range hand {
        card_scores = append(card_scores,cardValues[c])
        if len(remaining) > 0 {
            count := strings.Count(remaining,string(c))
            if count > 1 {
                type_tracker = append(type_tracker,count)
            }
            remaining = strings.ReplaceAll(remaining,string(c),"")
        }
    }
    return set_type_score1(type_tracker),card_scores
}

func set_scores2(hand string, cardValues map[rune]int) (
                type_score int, card_scores []int) {
    type_tracker := make([]int,0)
    jcount := 0
    remaining := hand
    for _,c := range hand {
        card_scores = append(card_scores,cardValues[c])
        if c == 'J' { jcount++ }
        if len(remaining) > 0 {
            count := strings.Count(remaining,string(c))
            if count > 1 {
                type_tracker = append(type_tracker,count)
            }
            remaining = strings.ReplaceAll(remaining,string(c),"")
        }
    }
    type_score = set_type_score1(type_tracker)
    if jcount > 0 && jcount < 5 {
        type_score = set_type_score2(type_score,type_tracker,jcount)
    }
    return type_score,card_scores
}

func set_hand(hand *cardHand, line string, cardValues map[rune]int) {
    base_split := strings.Split(line," ")
    hand.hand = base_split[0]
    hand.bid,_ = strconv.Atoi(base_split[1])
    hand.type_score,hand.card_scores = set_scores(hand.hand,cardValues)
}

func set_hand2(hand *cardHand, line string, cardValues map[rune]int) {
    base_split := strings.Split(line," ")
    hand.hand = base_split[0]
    hand.bid,_ = strconv.Atoi(base_split[1])
    hand.type_score,hand.card_scores = set_scores2(hand.hand,cardValues)
}

func should_insert(newHand cardHand, oldHand cardHand) bool {
    if newHand.type_score < oldHand.type_score { return true }
    if newHand.type_score > oldHand.type_score { return false }
    for i,cardscore := range newHand.card_scores {
        if cardscore < oldHand.card_scores[i] { return true }
        if cardscore > oldHand.card_scores[i] { return false }
    }
    return false
}

func insert_sorted(hands []cardHand,newHand cardHand) (newHands []cardHand) {
    if len(hands) == 0 { return append(hands,newHand) }
    for i,hand := range hands {
        if should_insert(newHand,hand) {
            newHands = slices.Insert(hands,i,newHand)
            return newHands
        }
    }
    newHands = append(hands,newHand)
    return newHands
}

func main(){

    part,err := strconv.Atoi(os.Args[1])
    if err != nil {
        log.Fatal("problem parsing cmd line arg",err)
    }
    var file *os.File
    if part == 0 {
        file,err = os.Open("test_input.txt")
    } else {
        file,err = os.Open("day7_input.txt")
    }
    if err != nil {
        log.Fatal("problem or opening file",err)
    }
    defer file.Close()

    var cardValues map[rune]int
    Hands := make([]cardHand,0)

    if part == 1 {
        cardValues = map[rune]int {
            'A': 12, 'K': 11, 'Q': 10, 'J': 9, 
            'T': 8, '9': 7, '8': 6, '7': 5, 
            '6': 4, '5': 3, '4': 2, '3': 1, 
            '2': 0,
        }
        for scanner := bufio.NewScanner(file); scanner.Scan();  {
            newHand := cardHand{}
            set_hand(&newHand,scanner.Text(),cardValues)
            Hands = insert_sorted(Hands,newHand)
        }
    } else {
        cardValues = map[rune]int {
            'A': 12, 'K': 11, 'Q': 10, 'T': 9,
            '9': 8, '8': 7, '7': 6, '6': 5, 
            '5': 4, '4': 3, '3': 2, '2': 1, 
            'J': 0, 
        }
        for scanner := bufio.NewScanner(file); scanner.Scan();  {
            newHand := cardHand{}
            set_hand2(&newHand,scanner.Text(),cardValues)
            Hands = insert_sorted(Hands,newHand)
        }
    }

    ans := 0
    for i,hand := range Hands {
        ans += hand.bid*(i+1)
    }

    fmt.Println("The answer is: ",ans)

}
