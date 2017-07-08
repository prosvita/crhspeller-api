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

    w := word.New(lang, args[2])

    w.Trace(lc.Crh)
    w.Trace(lc.Crh_RU)
}
