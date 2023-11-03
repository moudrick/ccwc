package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "unicode"
)

type CountingReader struct {
    Reader   *os.File
    ByteCount int
}

func (cr *CountingReader) Read(p []byte) (n int, err error) {
    n, err = cr.Reader.Read(p)
    cr.ByteCount += n
    return n, err
}

func CountNonBytes(scanner *bufio.Scanner) (int, int, int) {
    lines := 0
    words := 0
    characters := 0
    inWord := false

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
    return characters, lines, words
}

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

    countingReader := &CountingReader{Reader: file}
    scanner := bufio.NewScanner(countingReader)
    scanner.Split(bufio.ScanRunes)
    characters, lines, words := CountNonBytes(scanner)
    totalByteCount :=  countingReader.ByteCount

    if countLines {
        fmt.Printf("%7d", lines)
    }
    if countWords {
        fmt.Printf("%8d", words)
    }
    if countChars {
        fmt.Printf("%8d", characters)
    }
    if countBytes {
        fmt.Printf("%8d", totalByteCount)
    }
    fmt.Printf(" %s\n", filename)
}
