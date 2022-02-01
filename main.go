package main

import (
	"log"

	tgfun "github.com/Sagleft/tgfun/lib"
)

func main() {
	data := getFunnelData()
	script := getFunnelScript()

	funnel, err := tgfun.NewFunnel(data, script)
	if err != nil {
		log.Fatalln(err)
	}
	funnel.Run()

	forever := make(chan struct{})
	//run in background
	<-forever
}
