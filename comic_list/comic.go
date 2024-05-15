package main

import (
    "encoding/csv"
    "os"
)

type Comic struct {
    date string
    code string
    pub string
    title string
    cost string
}

type Comic_List struct {
    list []Comic
    size int
    count int
}

func LoadCsv(filename string) ([][]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    return records, nil
}
