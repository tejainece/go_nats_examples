package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nats-io/go-nats"
)

func main() {
	fmt.Println("Utility to publish to NATS:")

	conn, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	stdinLines := bufio.NewScanner(os.Stdin)

	for {
		if !stdinLines.Scan() {
			os.Exit(0)
		}
		subj := stdinLines.Text()
		if !stdinLines.Scan() {
			os.Exit(0)
		}
		msg := stdinLines.Text()

		err = conn.Publish(subj, []byte(msg))
		if err != nil {
			panic(err)
		}
		conn.Flush()
	}
}
