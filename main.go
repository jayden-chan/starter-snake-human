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

	for i, _ := range os.Args {
		if i == 0 {
			continue
		}
		if os.Args[i] == "--color" && len(os.Args) > i+1 {
			snakeColor = os.Args[i+1]
			log.Println("Snake color set to", snakeColor)
		}
	}

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
