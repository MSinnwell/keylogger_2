package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
)

func main() {
	f, err1 := os.Create("output.txt")
	if err1 != nil {
		fmt.Println("Error on create: ", err1)
	}
	err2 := keyboard.Open()
	if err2 != nil {
		fmt.Println("Error on open: ", err2)
	}
	for {
		char, _, err3 := keyboard.GetKey()
		if err2 != nil {
			fmt.Println("Error on GetKey: ", err3)
		}
		_, err4 := f.WriteString(string(char))
		_, err5 := f.WriteString("\n")

		if err4 != nil || err5 != nil {
			fmt.Println("Error on write: ", err4, err5)
		}

	}
	func() {
		_ = keyboard.Close()
	}()
}
