package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"net/http"
	"os"
)

var inputQueue KeyQueue
var lastKey rune = 'w'

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/start", Start)
	http.HandleFunc("/move", Move)
	http.HandleFunc("/end", End)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	// Add filename into logging messages
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("Running server on port %s...\n", port)
	log.Printf("Keep this terminal focused to recieve input from the keyboard.\nUse the arrow keys to move the snake.")

	go http.ListenAndServe(":"+port, nil)

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Println("Press ESC to quit")
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		} else if key == keyboard.KeyEsc {
			break
		}
		// fmt.Printf("You pressed: %q\r\n", char)
		inputQueue.Enqueue(char)
	}
}
