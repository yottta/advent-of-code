package main

import (
	"fmt"
	aoc "github.com/yottta/advent-of-code/00_aoc"
	"slices"
	"strconv"
	"strings"
)

func main() {
	aoc.BasicRun(part1, part2)
}

type handType int

const (
	highCard handType = iota + 1
	onePair
	twoPairs
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type cardType int

const (
	joker cardType = iota + 1
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
	ace
)

var part1CardTypes = map[string]cardType{
	"2": two,
	"3": three,
	"4": four,
	"5": five,
	"6": six,
	"7": seven,
	"8": eight,
	"9": nine,
	"T": ten,
	"J": jack,
	"Q": queen,
	"K": king,
	"A": ace,
}
var part2CardTypes = map[string]cardType{
	"2": two,
	"3": three,
	"4": four,
	"5": five,
	"6": six,
	"7": seven,
	"8": eight,
	"9": nine,
	"T": ten,
	"J": joker,
	"Q": queen,
	"K": king,
	"A": ace,
}

func part1(content []string) {
	hands := parseHands(content, part1CardTypes)
	for i, h := range hands {
		hands[i].hType = getHandType(h.cardCounts)
	}
	sortHands(hands)
	var res int
	for i, h := range hands {
		res += h.bid * (i + 1)
	}
	fmt.Println(res)
}

func part2(content []string) {
	hands := parseHands(content, part2CardTypes)
	for i := 0; i < len(hands); i++ {
		h := hands[i]
		h.hType = getHandType(h.cardCounts)
		h.hType = applyJokers(h)
		hands[i] = h
	}
	sortHands(hands)
	var res int
	for i, h := range hands {
		res += h.bid * (i + 1)
	}
	fmt.Println(res)
}

type card struct {
	value    string
	cardType cardType
}

type hand struct {
	cards      []card
	cardCounts map[cardType]int
	bid        int
	hType      handType
}

func parseHands(content []string, cTypes map[string]cardType) []hand {
	var hands []hand
	for _, line := range content {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		rawCards := strings.Split(strings.TrimSpace(parts[0]), "")
		var cards []card
		cardCounts := make(map[cardType]int)
		for _, rawCard := range rawCards {
			cType := cTypes[rawCard]
			cards = append(cards, card{value: rawCard, cardType: cType})
			cardCounts[cType]++
		}

		bid, err := strconv.Atoi(parts[1])
		aoc.Must(err)
		hands = append(hands, hand{cards: cards, bid: bid, cardCounts: cardCounts})
	}
	return hands
}

func getHandType(cardCounts map[cardType]int) handType {
	// check for five of a kind
	var pairCount int
	for _, count := range cardCounts {
		if count == 5 {
			return fiveOfAKind
		}
		if count == 4 {
			return fourOfAKind
		}
		if count == 3 {
			for _, count2 := range cardCounts {
				if count2 == 2 {
					return fullHouse
				}
			}
			return threeOfAKind
		}
		if count == 2 {
			pairCount++
		}
	}
	if pairCount == 2 {
		return twoPairs
	}
	if pairCount == 1 {
		return onePair
	}
	return highCard
}

func applyJokers(h hand) handType {
	noOfJokers := h.cardCounts[joker]
	if noOfJokers == 0 {
		return h.hType
	}
	switch h.hType {
	case fourOfAKind, fullHouse:
		return fiveOfAKind
	case threeOfAKind:
		switch noOfJokers {
		case 1:
			return fourOfAKind
		case 3:
			return fourOfAKind
		}
	case twoPairs:
		if noOfJokers == 2 {
			return fourOfAKind
		}
		return fullHouse
	case onePair:
		switch noOfJokers {
		case 1:
			return threeOfAKind
		case 2:
			return threeOfAKind
		}
	case highCard:
		return onePair
	}
	return h.hType
}

func sortHands(hands []hand) {
	slices.SortFunc(hands, func(a, b hand) int {
		if a.hType != b.hType {
			return int(b.hType - a.hType)
		}
		for i := 0; i < len(a.cards); i++ {
			if a.cards[i].cardType == b.cards[i].cardType {
				continue
			}
			return int(b.cards[i].cardType - a.cards[i].cardType)
		}
		return 0
	})
	slices.Reverse(hands)
}
