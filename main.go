package main

import (
    "fmt"
    "os"

    lc "./langcode"
    "./word"
)

func main() {
    args := os.Args

    lang, err := lc.GetId(&args[1])
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(lang, lc.Crh_RU)

    w := word.New(args[1], args[2])
    // fmt.Println(w)
    w.Trace("crh")
    w.Trace("crh_RU")
}
