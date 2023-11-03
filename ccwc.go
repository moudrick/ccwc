package main

import (
    "bufio"
    "fmt"
    "os"
    "unicode"
)

func main() {
    var file *os.File
    var err error
	var filename string 

    if len(os.Args) > 1 {
        filename = os.Args[1]
        if filename != "" {
            file, err = os.Open(filename)
            if err != nil {
                fmt.Println("Error:", err)
                os.Exit(2)
            }
            defer file.Close()
        }
    } else {
        file = os.Stdin
		filename = ""
    }

    lines := 0
    words := 0
    characters := 0
    inWord := false

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanRunes)

    for scanner.Scan() {
        char := scanner.Text()
        characters++

        if char == "\n" {
            lines++
        }

        if unicode.IsSpace([]rune(char)[0]) {
            inWord = false
        } else if !inWord {
            inWord = true
            words++
        }
    }

    fmt.Printf("%7d %7d %7d %s\n", lines, words, characters, filename)
}
