package main

import (
	"log"

	mcrcon "github.com/Kelwing/mc-rcon"
)

func main() {
	conn := new(mcrcon.MCConn)
	err := conn.Open("IP:PORT", "rconPassword")
	if err != nil {
		log.Println("Connection failed!", err)
	}
	defer conn.Close()

	err = conn.Authenticate()
	if err != nil {
		log.Println("Authorization failed!", err)
	}

	result, err := conn.SendCommand("list")
	if err != nil {
		log.Println("Command not found or broken!", err)
	}

	log.Println(result)
}
