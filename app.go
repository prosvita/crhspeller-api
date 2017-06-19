package main

import (
    "fmt"
    "os"
    "unicode/utf8"
    "strings"
)

type soundMap map[string]string
type ruleFunc func(string, *Word) bool

type Sounds []struct {
    text soundMap
    rules []ruleFunc
}

type Word struct {
    lang string
    word string
    tail []string
    pTail []*soundMap
}

var sounds = Sounds {
// TODO: навесить на звуки счетчики, прогнать большой текс, отсортировать по популярности
    { soundMap{"crh":"c",  "crh_RU":"дж"}, []ruleFunc{} },
    { soundMap{"crh":"ğ",  "crh_RU":"гъ"}, []ruleFunc{} },
    { soundMap{"crh":"ñ",  "crh_RU":"нъ"}, []ruleFunc{} },
    { soundMap{"crh":"q",  "crh_RU":"къ"}, []ruleFunc{} },
    { soundMap{"crh":"la", "crh_RU":"ла"}, []ruleFunc{} }, // ? можно проверить на конвертации
    { soundMap{"crh":"lâ", "crh_RU":"ля"}, []ruleFunc{} }, // ?
    { soundMap{"crh":"ts", "crh_RU":"ц" }, []ruleFunc{} },
    { soundMap{"crh":"şç", "crh_RU":"щ" }, []ruleFunc{} },

    { soundMap{"crh":"ye", "crh_RU":"е"}, []ruleFunc{eq} }, // в начале слова, после гласных и ь в латинице соответствует ye (ye на початку слова, після голосної або ь)
    { soundMap{"crh":"e",  "crh_RU":"е"}, []ruleFunc{eq} }, // e <-> е після приголосної
    { soundMap{"crh":"e",  "crh_RU":"э"}, []ruleFunc{eq} },
    { soundMap{"crh":"ya", "crh_RU":"я"}, []ruleFunc{eq} }, // â після приголосної; ya на початку слова, після голосної або ь
    { soundMap{"crh":"â",  "crh_RU":"я"}, []ruleFunc{eq} }, // <-> після приголосної

    { soundMap{"crh":"yo", "crh_RU":"ё"}, []ruleFunc{eq} }, // yo на початку слова у «твердих» словах; після голосної, ь чи ъ
    { soundMap{"crh":"yö", "crh_RU":"ё"}, []ruleFunc{eq} }, // yö на початку слова у «м'яких» словах
    { soundMap{"crh":"ö",  "crh_RU":"ё"}, []ruleFunc{eq} }, // ё <-> ö після приголосної
    { soundMap{"crh":"ö",  "crh_RU":"о"}, []ruleFunc{eq} }, // ö якщо о — перша голосна у «м'якому» слові
    { soundMap{"crh":"o",  "crh_RU":"о"}, []ruleFunc{eq} },

    { soundMap{"crh":"yu", "crh_RU":"ю"}, []ruleFunc{eq} }, // yu на початку слова, після голосної або ь у «твердих» словах
    { soundMap{"crh":"yü", "crh_RU":"ю"}, []ruleFunc{eq} }, // yü на початку слова, після голосної або ь у «м'яких» словах
    { soundMap{"crh":"ü",  "crh_RU":"ю"}, []ruleFunc{eq} }, // ü <-> ю після приголосної
    { soundMap{"crh":"ü",  "crh_RU":"у"}, []ruleFunc{eq} }, // ü якщо у — перша голосна у «м'якому» слові
    { soundMap{"crh":"u",  "crh_RU":"у"}, []ruleFunc{eq} },

    { soundMap{"crh":"a",  "crh_RU":"а"}, []ruleFunc{} },
    { soundMap{"crh":"b",  "crh_RU":"б"}, []ruleFunc{} },
    { soundMap{"crh":"ç",  "crh_RU":"ч"}, []ruleFunc{} },
    { soundMap{"crh":"d",  "crh_RU":"д"}, []ruleFunc{} },
    { soundMap{"crh":"f",  "crh_RU":"ф"}, []ruleFunc{} },
    { soundMap{"crh":"g",  "crh_RU":"г"}, []ruleFunc{} },
    { soundMap{"crh":"h",  "crh_RU":"х"}, []ruleFunc{} },
    { soundMap{"crh":"ı",  "crh_RU":"ы"}, []ruleFunc{} },
    { soundMap{"crh":"i",  "crh_RU":"и"}, []ruleFunc{} },
    { soundMap{"crh":"j",  "crh_RU":"ж"}, []ruleFunc{} },
    { soundMap{"crh":"k",  "crh_RU":"к"}, []ruleFunc{} },
    { soundMap{"crh":"l",  "crh_RU":"л"}, []ruleFunc{} },
    { soundMap{"crh":"m",  "crh_RU":"м"}, []ruleFunc{} },
    { soundMap{"crh":"n",  "crh_RU":"н"}, []ruleFunc{} },
    { soundMap{"crh":"p",  "crh_RU":"п"}, []ruleFunc{} },
    { soundMap{"crh":"r",  "crh_RU":"р"}, []ruleFunc{} },
    { soundMap{"crh":"s",  "crh_RU":"с"}, []ruleFunc{} },
    { soundMap{"crh":"ş",  "crh_RU":"ш"}, []ruleFunc{} },
    { soundMap{"crh":"t",  "crh_RU":"т"}, []ruleFunc{} },
    { soundMap{"crh":"v",  "crh_RU":"в"}, []ruleFunc{} },
    { soundMap{"crh":"y",  "crh_RU":"й"}, []ruleFunc{} },
    { soundMap{"crh":"z",  "crh_RU":"з"}, []ruleFunc{} },
}

var eq = func(sound string, word *Word) bool {return true}


func NewWord(lang string, text string) *Word {

    word := Word{lang, text, []string{}, []*soundMap{}}
    tail := strings.ToLower(text)

    for utf8.RuneCountInString(tail) > 0 {
        needTrim := true

        for i := 0; i < len(sounds); i++ {
            if strings.HasPrefix(tail, sounds[i].text[lang]) {
                verified := true
                for r := 0; r < len(sounds[i].rules); r++ {
                    if !sounds[i].rules[r](sounds[i].text[lang], &word) {
                        verified = false
                        break
                    }
                }
                if !verified {
                    continue
                }

                word.pTail = append(word.pTail, &sounds[i].text)

                word.tail = append(word.tail, sounds[i].text[lang])
                tail = strings.TrimPrefix(tail, sounds[i].text[lang])
                needTrim = false

                fmt.Println(sounds[i].text[lang], tail)
                break
            }
        }
        if needTrim {
            fmt.Println(string([]rune(tail)[0]), tail, "BOO")
            // TODO: посилання на Sounds на найближче через звичайний трансліт, або excepton
            word.pTail = append(word.pTail, nil)

            word.tail = append(word.tail, string([]rune(tail)[0]))
            tail = strings.TrimPrefix(tail, string([]rune(tail)[0]))
        }
    }

    return &word
}

func main() {
    args := os.Args

    word := NewWord(args[1], args[2])
    fmt.Println(word)
}
