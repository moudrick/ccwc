package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "unicode"
	"unicode/utf8"
)

func main() {
    var file *os.File
    var err error
    var filename string
    var countBytes, countWords, countLines, countChars bool

    flag.BoolVar(&countBytes, "c", false, "Count bytes")
    flag.BoolVar(&countWords, "w", false, "Count words")
    flag.BoolVar(&countLines, "l", false, "Count lines")
	flag.BoolVar(&countChars, "m", false, "Count characters")
    flag.Parse()

    if !countBytes && !countWords && !countLines && !countChars {
        // If no flags provided, assume three well known flags are turned on
        countBytes = true
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
    bytes := 0
    inWord := false	

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanRunes)

    for scanner.Scan() {
        char := scanner.Text()
        characters++

        byteCount := utf8.RuneCountInString(char)
        bytes += byteCount

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
	if countBytes {
        fmt.Printf("%7d", bytes)
    }
    if countChars {
        fmt.Printf("%7d", characters)
    }
    fmt.Printf(" %s\n", filename)
}
