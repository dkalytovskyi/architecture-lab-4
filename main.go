package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./engine"
)

type print struct {
	arg string
}

type palindrom struct {
	arg string
}

func (printComm *print) Execute(loop engine.Handler) {
	fmt.Println(printComm.arg)
}

func (palindrComm *palindrom) Execute(loop engine.Handler) {
	runes := []rune(palindrComm.arg)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	result := palindrComm.arg + string(runes)

	loop.Post(&print{arg: string(result)})
}

func parse(command string) engine.Command {
	args := strings.Fields(command)

	switch args[0] {
	case "print":
		return &print{arg: args[1]}
	case "palindrom":
		return &palindrom{arg: args[1]}
	default:
		return &print{arg: "Unknown command"}
	}
}

func main() {
	args := os.Args[1:]
	inputFile := "./commandsList.txt"

	if len(args) != 0 {
		inputFile = args[0]
	}

	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			eventLoop.Post(parse(scanner.Text()))
		}
	}

	eventLoop.AwaitFinish()
}
