package main

import (
	"io"
	"log"
	"os"
)

func FunctionWithDefer() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 2048)

	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
	//defer can be anywhere
	//but it will get executed
	//once the entire function is done
	defer f.Close()
}
