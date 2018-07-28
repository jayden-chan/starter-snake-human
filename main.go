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
var snakeColor string = "#84D5E2"
var snakeTaunt string = "NOT A HUMAN..."

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/start", Start)
	http.HandleFunc("/move", Move)
	http.HandleFunc("/end", End)

	port := "9000"

	// Parse --color command line argument
	for i, _ := range os.Args {
		if i == 0 {
			continue
		}
		switch {
		case os.Args[i] == "--color" && len(os.Args) > i+1:
			snakeColor = os.Args[i+1]
			log.Println("Snake color set to", snakeColor)
		case os.Args[i] == "--port" && len(os.Args) > i+1:
			port = os.Args[i+1]
			log.Println("Port set to", port)
		case os.Args[i] == "--taunt" && len(os.Args) > i+1:
			snakeTaunt = os.Args[i+1]
			log.Println("Taunt set to", snakeTaunt)
		}
	}

	// Add filename into logging messages
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("Server running on port %s...\n", port)
	fmt.Printf("\nKeep this terminal focused to recieve input from the keyboard.\nUse the W A S D keys to move the snake.\n")

	// Run the HTTP server in a new goroutine
	go http.ListenAndServe(":"+port, nil)

	// Take input from the keyboard
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
		// If the provided move will result in the snake going back
		// on itself, don't add it to the queue
		switch {
		case char == 'w' && lastKey == 's':
			break
		case char == 'a' && lastKey == 'd':
			break
		case char == 's' && lastKey == 'w':
			break
		case char == 'd' && lastKey == 'a':
			break
		default:
			inputQueue.Enqueue(char)
		}
	}
}
