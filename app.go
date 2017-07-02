package main

import (
    // "fmt"
    "os"

    "./word"
)

func main() {
    args := os.Args

    w := word.New(args[1], args[2])
    // fmt.Println(w)
    w.Debug("crh")
    w.Debug("crh_RU")
}
