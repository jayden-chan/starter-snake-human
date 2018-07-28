# starter-snake-human
Takes input from the standard in to allow you to train your snakes

## Installation
Download the appropriate executable from the `/bin` directory and run it from your terminal.

### Windows
From the directory containing the executable:
```
> starter-snake-human.exe
```

### Mac OS / Linux
From the directory containing the executable:
```
$ ./starter-snake-human
```

## Usage
Use the W A S D keys to move the snake around. Pres ESC to exit. You must focus the terminal window in order for the program to receive the keyboard input.

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
