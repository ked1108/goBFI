package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var code string = ""
var data [10000]byte
var dp uint16 = 0

var isLooping bool
var nestedLoops uint64
var loopStack []uint8

func opIncVal() {
	data[dp]++
}

func opDecVal() {
	data[dp]--
}

func opIncDp() {
	if int(dp) < len(data)-1 {
		dp++
	} else {
		fmt.Println("ERR BUFFER OVERFLOWED")
	}
}

func opDecDp() {
	if int(dp) > 1 {
		dp--
	} else {
		fmt.Println("ERR BUFFER UNDERFLOWED")
	}
}

func opIn() {
	var inp []byte
	_, err := fmt.Scanf("%s", &inp)
	if err != nil {
		return
	}
	data[dp] = inp[0]
	fmt.Printf("%c\n", inp[0])
}

func opOut() {
	fmt.Printf("%c\n", data[dp])
}

func loopTest(i int) {
	if isLooping {
		if code[i] == '[' {
			nestedLoops++
		} else if code[i] == ']' {
			if nestedLoops == 0 {
				isLooping = false
			} else {
				nestedLoops--
			}
		}
	}
}

func opLoops(i int) int {
	if data[dp] == 0 {
		idx := loopStack[len(loopStack)-1]
		if idx >= 0 {

		}
	}
	return i
}

func nop() {
	return
}

func cleaner(filename string) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0400)

	if err != nil {
		if err != nil {
			log.Fatal(err)
		}
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Getting rid of all the whitespace, its a waste anyways
		code = code + strings.Replace(scanner.Text(), " ", "", -1)
	}
}

func interpreter() {

	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '+':
			opIncVal()
		case '-':
			opDecVal()
		case '>':
			opIncDp()
		case '<':
			opDecDp()
		case ',':
			opIn()
		case '.':
			opOut()
		case '[':
			i = opLoops(i)
		default:
			nop()
		}
	}
	fmt.Println("")
	fmt.Println(data[:10])

}

func main() {
	if len(os.Args) >= 2 {
		filename := os.Args[1]
		fmt.Println("Executing File:\t" + filename)
		cleaner(filename)
		interpreter()
		fmt.Println("\n\nFile Finished Execution")
	} else {
		fmt.Println("No File Provided for Execution")
	}
}
