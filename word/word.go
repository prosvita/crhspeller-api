package word

import "fmt"
import (
    lc "../langcode"
    s "../sounds"
)

type Word struct {
    lang int
    word string
    sounds *s.Sounds
    // root = ???
    // tail = ???
    // capitalMask = []??? // nil by default
}

func New(lang int, text string) *Word {

    sounds := s.New(lang, &text)
    word := Word{lang, text, sounds}

    return &word
}

func (w *Word) Trace(inLang int) {
    fmt.Printf("%s: %s => %s: %s\n", *lc.GetLang(w.lang), w.word, *lc.GetLang(inLang), (*w.sounds.Trace(inLang)))
}
