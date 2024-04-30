package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var input []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите число")
		scanner.Scan()
		input = strings.Split(scanner.Text(), "")
		start()
	}

}

func start() {
	var indx int
	var out string
	for i, v := range input {
		indx = i
		out += v
	}

	if indx > 0 && input[indx-1] == "1" && input[indx] != "1" || indx >= 1 && input[indx-1] == "1" && input[indx] == "1" || input[indx] == "0" {
		fmt.Println(out, "компьютеров")
	} else if input[indx] == "1" {
		fmt.Println(out, "компьютер")
	} else if input[indx] == "2" || input[indx] == "3" || input[indx] == "4" {
		fmt.Println(out, "компьютера")
	} else if input[indx] == "5" || input[indx] == "6" || input[indx] == "7" || input[indx] == "8" || input[indx] == "9" {
		fmt.Println(out, "компьютеров")
	}
}
