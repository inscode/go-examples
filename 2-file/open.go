package main

import (
	"os"
	"fmt"
)

func main() {
	file, err := os.Open("C:\\gitNote\\Go-tour-book\\src\\2-file\\access.log")
	if err != nil {
		fmt.Println("err is: ", err)
		return
	}
	defer file.Close()

	var tmp = make([]byte, 64)
	_, err = file.Read(tmp)
	if err != nil {
		fmt.Println("err is:", err)
	}

	fmt.Println(string(tmp[:]))

}
