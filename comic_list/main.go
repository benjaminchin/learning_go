package main

import (
    "fmt"
    "os"
    "strings"
    "bufio"
    "strconv"
)

const initSize = 10

func main() {
    args := os.Args[1:]

    if len(args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: comics <input-file> <output-file>\n")
        os.Exit(1)
    }

    fmt.Println(args)

    // Open input and output files for read and write
    output, err := os.Open(args[1])
    check(err)
    defer output.Close()

    input, err := os.Open(args[0])
    check(err)
    defer input.Close()

    list := make([]Comic, initSize)

    cart := make([]Comic, initSize)

    scanner := bufio.NewScanner(input)

    for scanner.Scan() {
        line := scanner.Text()

        fields := strings.Fields(line)

        command := fields[0]

        if command == "load" {
            // Get arg for load command
            csv := fields[1]

            fmt.Fprintf(output, "Command: %s %s\n", command, csv)

            // Run command
            err := LoadCsv(output, list, csv)
            check(err)
        } else if command == "buy" {
            buy, err := strconv.ParseInt(fields[1], 10, 32)
            check(err)

            fmt.Fprintf(output, "Command: %s %d\n", command, buy)

            BuyComic(list, cart, int(buy))
            fmt.Fprintf(output, "Comic #%d added to purchase list\n", buy)
        } else if command == "checkout" {
            fmt.Fprintf(output, "Command: %s\n", command)

            Checkout(output, cart)
        } else if command == "display" {
            fmt.Fprintf(output, "Command: %s\n", command)

            DisplayList(output, list)
        } else if command == "save" {
            filename := args[1]

            fmt.Fprintf(output, "Command: %s %s\n", command, filename)
            save, err := os.Open(filename)
            defer save.Close()
            check(err)

            DisplayListHorizontal(save, list)
        } else if command == "clear" {
            fmt.Fprintf(output, "Command: %s\n", command)
            list = make([]Comic, initSize)
        } else if command == "find" {
            search, err := strconv.ParseInt(args[1], 10, 32)
            check(err)

            fmt.Fprintf(output, "Command: %s %d\n", command, search)

            FindComic(output, list, int(search))
        } else if command == "remove" {
            remove, err := strconv.ParseInt(args[1], 10, 32)
            check(err)

            fmt.Fprintf(output, "Command %s %d\n", command, remove)

            RemoveAt(list, int(remove))
        }
    }
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
