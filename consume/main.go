package main

import (
	"fmt"

	"github.com/go-stomp/stomp"
)

func main() {
	conn, err := stomp.Dial("tcp", "localhost:61613")
	if err != nil {
		fmt.Println(err)
	}

	sub, err := conn.Subscribe("food", stomp.AckAuto)
	if err != nil {
		fmt.Println(err)
	}
	for {
		msg := <-sub.C
		fmt.Println(string(msg.Body))
	}

}
