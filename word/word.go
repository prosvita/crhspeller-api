package word

import (
   "fmt"

    S "../sounds"
//    "unicode/utf8"
//    "strings"
)

// Как ищем:
//   - формируем из разных словарей (корней и исключений) дерево корней разбитое на звуки
//   - разбираем слово на звуки, ошибки складываем в стек e.conficts, привязав их к позиции в слове
//   - ищем начало по дереву корней
//     - в root добавляем ссылку на последнюю ноду корня в дереве (чтоб не копировать ничего)
//     - в tail slice от sounds
//     - отсекаем из стека с ошибками (e.conficts) корень — должно исключить ошибки на именах собственных и заимствованиях
//     - разбираем tail на аффиксы (в сопряжении с корнем)
//   - Действие:
//     1. возвращаем ошибку
//     2. возвращаем в заданой кодировке слово собранное по правилам из корня и аффиксов
//     3. находим для некорректного слова ближайшие верные варианты

type Word struct {
    lang string
    word string
    sounds *S.Sounds
    // root = ???
    // tail = ???
    // capitalMask = []??? // nil by default
}

func New(lang string, text string) *Word {

    sounds := S.New(&lang, &text)
    word := Word{lang, text, sounds}

    // fmt.Println(sounds)
    return &word
}

func (w *Word) Debug(inLang string) {
    fmt.Printf("%s: %s => %s: %s\n", w.lang, w.word, inLang, (*w.sounds.Debug(&inLang)))
}
