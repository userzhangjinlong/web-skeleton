package main

import (
	"linkr-frame/app/console"
	"linkr-frame/bootstrap"
	"log"
)

func main() {
	bootstrap.InitSystemConfig()
	bootstrap.InitMysql()
	err := console.Execute()

	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
