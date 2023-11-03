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

    if len(os.Args) > 1 {
        filename := os.Args[1]
        file, err = os.Open(filename)
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(2)
        }
        defer file.Close()
    } else {
        file = os.Stdin
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

    fmt.Printf("Lines: %d\nWords: %d\nCharacters: %d\n", lines, words, characters)
}
