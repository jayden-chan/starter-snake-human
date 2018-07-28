package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type StartRes struct {
	Color string `json:"color,omitempty"`
	Taunt string `json:"taunt,omitempty"`
}

type MoveRes struct {
	Move string `json:"move"`
}

// Start is called when a request to /start is received,
// which will happen once at the beginning of each game.
// It will contain information about the match including
// the width and height of the board.
func Start(res http.ResponseWriter, req *http.Request) {
	log.Printf("Start req recieved")
	inputQueue.Clear()
	lastKey = 'w'

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(StartRes{
		// Color: "#84D5E2",
		Color: "#222222",
		Taunt: "NOT A HUMAN...",
	})

	res.Write([]byte("\n"))
}

// Move is called on each game tick. We must send our
// move response here.
func Move(res http.ResponseWriter, req *http.Request) {
	move := MoveRes{}
	var key rune
	if inputQueue.Size() > 0 {
		key = inputQueue.Dequeue()
		lastKey = key
	} else {
		key = lastKey
	}
	switch key {
	case 'w':
		move = MoveRes{Move: "up"}
	case 'a':
		move = MoveRes{Move: "left"}
	case 's':
		move = MoveRes{Move: "down"}
	case 'd':
		move = MoveRes{Move: "right"}
	}
	HTTPRes(res, move)
}

// End is only called at the end of the game and is
// essentially useless I think..
func End(res http.ResponseWriter, req *http.Request) {
	log.Printf("End req recieved")
}

// HTTPRes sends an HTTP response to the given ResponseWriter
func HTTPRes(res http.ResponseWriter, obj interface{}) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(obj)
	res.Write([]byte("\n"))
}
