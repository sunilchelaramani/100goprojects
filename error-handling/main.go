//custom error handling in project

package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	//customerr := errors.New("this is a custom error")
	//fmt.Println("err happened", customerr)
	//fmt.Println("err happened", err)
	if err != nil {
		//fmt.Println("err happened", err)
		log.Println("err happened", err)
		//log.Fatalln(err)
		//panic(err)
	}
}
