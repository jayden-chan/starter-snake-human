# starter-snake-human
Takes input from the standard in to allow you to train your snakes

## Installation
### If you have [Go](https://golang.org/doc/install) installed:
```
$ go get github.com/jayden-chan/starter-snake-human
$ go install github.com/jayden-chan/starter-snake-human
```
### If you do not have Go installed:
Download the appropriate executable from `release` tab in GitHub and run it in your terminal.

#### Linux / Mac OS
From the directory containing the executable:
```
$ ./starter-snake-human
```

#### Windows
From the directory containing the executable:
```
> starter-snake-human.exe
```

## Usage
Use the W A S D keys to move the snake around. Press ESC to exit. You must focus the terminal window in order for the program to receive the keyboard input.

### Arguments
To change the color of the snake:
```
--color "#ffffff"
```

To change the taunt:
```
--taunt "i am snak"
```

To change the port:
```
--port "8080"
```

For example, the following will run a pink snake on port 8080 with taunt "i am snak"
```
./starter-snake-human --color "#ff00ff" --port "8080" --taunt "i am snak"
```

If you do not provide any command line arguments, sensible defaults will be chosen for you.
