package main

import (
    "bufio"
    "fmt"
    "os"
    "unicode"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run wc.go <filename>")
        os.Exit(1)
    }

    filename := os.Args[1]

    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(2)
    }
    defer file.Close()

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
