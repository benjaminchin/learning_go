package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]

    if len(args) != 3 {
        fmt.Fprintf(os.Stderr, "Usage: comics <input-file> <output-file>\n")
        os.Exit(1)
    }

    fmt.Println(args)

    // Open input and output files for read and write
//    output, err := os.Open(args[2])
//    check(err)
//
//    input, err := os.Open(args[1])
//    check(err)
//
    // Create Comic_Lists for all loaded comics, and bought comics
    records, err := LoadCsv(args[0])
    check(err)

    for _, record := range records {
        fmt.Println(record)
    }
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
