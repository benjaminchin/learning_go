package main

import (
    "encoding/csv"
    "strconv"
    "fmt"
    "os"
)

type Comic struct {
    date string
    code string
    pub string
    title string
    cost string
}


func LoadCsv(output *os.File, list []Comic, filename string) error {
    file, err := os.Open(filename)

    if err != nil {
        return err
    }

    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return err
    }

    for _, record := range records {
        var c Comic
        UpdateComic(&c, record)
        AddComic(list, c)
    }

    fmt.Fprintf(output, "Number of comics: %d\n", len(list))
    return nil
}

func AddComic(list []Comic, c Comic) {
    list = append(list, c)
}

func UpdateComic(c *Comic, record []string) {
    c.date = record[0]
    c.code = record[1]
    c.pub = record[2]
    c.title = record[3]
    c.cost = record[4]
}

func FindComic(out *os.File, list []Comic, index int) {
    if index < len(list) {
        DisplayComic(out, list[index])
    }
}

func DisplayComic(out *os.File, c Comic) {
    fmt.Fprintf(out, "Date: %s\n", c.date)
    fmt.Fprintf(out, "Code: %s\n", c.code)
    fmt.Fprintf(out, "Publisher: %s\n", c.pub)
    fmt.Fprintf(out, "Title: %s\n", c.title)
    fmt.Fprintf(out, "Cost: %s\n", c.cost)
}

func DisplayComicHorizontal(out *os.File, c Comic) {
    fmt.Fprintf(out, "%s,%s,%s,%s,%s\n", c.date, c.code, c.pub, c.title, c.cost)
}

func DisplayList(out *os.File, list []Comic) {
    if len(list) == 0 {
        fmt.Fprintf(out, "List is currently empty.\n")
    }

    for i, c := range list {
        fmt.Fprintf(out, "Comic Number: %d\n", i)
        DisplayComic(out, c)
    }
}

func DisplayListHorizontal(out *os.File, list []Comic) {
    fmt.Fprintf(out, "DATE,CODE,PUBLISHER,TITLE,PRICE\n")

    for _, c := range list {
        DisplayComicHorizontal(out, c)
    }
}

func RemoveAt(list []Comic, index int) bool {
    if index <= len(list) {
        list = append(list[:index], list[index+1:]...)
        return true
    }
    return false
}

func BuyComic(list []Comic, cart []Comic, index int) {
    if index < len(list) {
        var c Comic = list[index]
        AddComic(cart, c)
    }
}

func Checkout(out *os.File, cart []Comic) {
    fmt.Fprintf(out, "Comics in Purchase List\n")
    subtotal := 0.0

    for i, c := range cart {
        if c.cost != "AR" {
            cost, err := strconv.ParseFloat(c.cost, 64)
            check(err)
            subtotal += cost
        }
        fmt.Fprintf(out, "Comic Number: %d\n", i + 1)
        DisplayComic(out, c)
    }

    tax := subtotal * 0.05
    total := subtotal + tax

    fmt.Fprintf(out, "Subtotal: %.2f\n", subtotal)
    fmt.Fprintf(out, "\tTax: %.2f\n", tax)
    fmt.Fprintf(out, "\tTax: %.2f\n", total)
}
