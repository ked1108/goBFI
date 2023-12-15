package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var code = ""
var data [10000]byte
var dp uint16 = 0

var isLooping bool
var nestedLoops uint64
var loopStack []uint64

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
	if int(dp) >= 1 {
		dp--
	} else {
		fmt.Println("ERR BUFFER EMPTY")
	}
}

func opIn() {
	var inp []byte
	_, err := fmt.Scanf("%s", &inp)
	if err != nil {
		return
	}
	data[dp] = inp[0]
}

func opOut() {
	fmt.Printf("%c", data[dp])
}

func loopTest(i uint64) {
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

func opOpenLoops(i uint64) {
	if data[dp] == 0 {
		isLooping = true
	} else {
		loopStack = append(loopStack, i)
	}
}

func opCloseLoops(i *uint64) {
	if data[dp] != 0 {
		*i = loopStack[len(loopStack)-1]
	} else {
		if len(loopStack) > 0 {
			loopStack = loopStack[:len(loopStack)-1]
		}
	}
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
		// Getting rid of all the whitespace, it's a waste anyway
		code = code + strings.Replace(scanner.Text(), " ", "", -1)
	}
}

func interpreter() {
	var i uint64
	for i = 0; i < uint64(len(code)); i++ {
		loopTest(i)
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
			opOpenLoops(i)
		case ']':
			opCloseLoops(&i)
		default:
			nop()
		}
	}
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
