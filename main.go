package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("aplicação online")
	fmt.Println("realizando operação 1")
	fmt.Println("realizando operação 2")

	name, err := os.Hostname()

	getppid := os.Getppid()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("PID", getppid)
	fmt.Println("Hostname", name)

	time.Sleep(time.Minute * 5)
	log.Fatalln("saindo")

}
