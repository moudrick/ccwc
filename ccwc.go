package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "unicode"
)

func main() {
    var file *os.File
    var err error
    var filename string
    var countChars, countWords, countLines bool

    flag.BoolVar(&countChars, "c", false, "Count characters")
    flag.BoolVar(&countWords, "w", false, "Count words")
    flag.BoolVar(&countLines, "l", false, "Count lines")
    flag.Parse()

    if !countChars && !countWords && !countLines {
        // If no flags provided, assume all flags are turned on
        countChars = true
        countWords = true
        countLines = true
    }

    if len(flag.Args()) > 0 {
        filename = flag.Args()[0]
        file, err = os.Open(filename)
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(2)
        }
        defer file.Close()
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

    if countLines {
        fmt.Printf("%7d", lines)
    }
    if countWords {
        fmt.Printf("%7d", words)
    }
    if countChars {
        fmt.Printf("%7d", characters)
    }
    fmt.Printf(" %s\n", filename)
}
