package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func FileSaver(data interface{}, fileName string) error {

// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		fmt.Println("Couldn't create a new file: ", err.Error())
// 		return err
// 	}

// 	defer file.Close()

// 	_, err = file.WriteString(stringConverter(data))
// 	if err != nil {
// 		fmt.Println("Couldn't write into file: ", err.Error())
// 		return err
// 	}
// 	return nil

// }

// func stringConverter(data interface{}) string {
// 	res := ""
// 	if cards, ok := data.([]Card); ok {
// 		for _, card := range cards {
// 			res = res + card.Type + card.Suit + ", "
// 		}
// 		return res
// 	}

// 	if deck, ok := data.(Deck); ok {
// 		for i := 0; i < len(deck); i++ {
// 			res = res + deck[i].Type + deck[i].Suit + ", "
// 		}
// 		return res
// 	}

// 	hands := data.([][]Card)
// 	for i, cards := range hands {
// 		res = res + "player " + strconv.Itoa(i+1) + " cards: " + stringConverter(cards) + "\n"
// 	}
// 	return res
// }
