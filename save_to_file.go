package main

import (
	"fmt"
	"os"
	"strings"
)

func SaveToFile(data interface{}, fileName string) error {
	err := os.WriteFile(fileName, []byte(toString(data)), 0666)
	if err != nil {
		return err
	}
	return nil
}

func ReadDeckFromFile(fileName string) error {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	read := string(b)
	fmt.Println("This is the content of file: ", read)
	return nil
}

func toString(data interface{}) string {
	var s []string

	if cards, ok := data.([]card); ok {
		for _, c := range cards {
			s = append(s, c.rank+c.suit)
		}
		return strings.Join(s, ", ")
	}

	if deck, ok := data.(Deck); ok {
		return toString([]card(deck))
	}

	ret := ""
	if hands, ok := data.([][]card); ok {
		for _, hand := range hands {
			ret = ret + toString(hand) + "\n"
		}
	}
	return ret
}
