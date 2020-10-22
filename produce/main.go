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

	err = conn.Send("food", "", []byte("Satender Gautam"), nil)
	if err != nil {
		fmt.Println(err)
	}
	conn.Disconnect()

}
