package sounds

import (
    "fmt"

    "unicode/utf8"
    "strings"
)

type soundMap map[string]string
type ruleFunc func(*Sounds) bool

type Sound struct {
    signs soundMap
    rules []ruleFunc
    property uint
}
type SoundsStore []Sound

type Sounds struct {
    property uint
    length int
    sounds []*Sound
    index []uint8
}

const (
    vowel               = 1 << iota                         // голосна
    consonant           = 1 << iota                         // приголосна
    sonorant            = consonant | (1 << iota)           // сонант (сонорний приголосний звук)
    voiced              = 1 << iota                         // дзвінка
    voiceless           = 1 << iota                         // глуха
    obstruent           = consonant | (1 << iota)           // шумна приголосна
    voicedObstruent     = consonant | voiced | obstruent    // дзвінка шумна приголосна
    voicelessObstruent  = consonant | voiceless | obstruent // глуха шумна приголосна

    hardVowel           = vowel | (1 << iota)               // тверда голосна
    softVowel           = vowel | (1 << iota)               // м'яка голосна
    // undefinedVowel      = vowel | (1 << iota)               // не визначена голосна (тверда чи м'яка) (перша голосна о у)

    hardness            = 1 << iota                         // звук має `Ъ`
    softness            = 1 << iota                         // звук має `Ь`
)
// Consonant согласные

// sonorant р, л, м, н, нъ
// звонкие шумные: б, в, д, з, ж, г, гъ, дж;        (voiced obstruent)
// глухие шумные: п, т, ф, с, ш, ч, к, ъ, х, ц, щ.  (voiceless obstruent)

// voicing              дзвінкі
// voiced consonants    дзвінкі приголосні

// voicelessness        глухі
// voiceless vowels     глухі голосні
// voiceless consonants глухі приголосні

// Vowel гласные
// твердые a, ı, o, u       (hard vowels)
// мягкие e, i, ö, ü        (soft vowels)
// // не работает в заимствованиях для звуков
// // работает для аффиксов
// karyer-lerimizden
// для аффиксов последний звук из корня ()
// берем последнее гласное (исключение несколько аффиксов soñ-u-nace)


func New(lang *string, text *string) (sounds *Sounds) {

    tail := strings.ToLower(*text)
    length := utf8.RuneCountInString(tail) + 2   //FIXIT: 2 провсяквипадок, як що виходимо за межі, то append(sounds) та append (index)
    position := 0

    sounds = &Sounds{0, 0, make([]*Sound, length), make([]uint8, length)}

    for utf8.RuneCountInString(tail) > 0 {
        needTrim := true

        for i := 0; i < len(soundsStore); i++ {
            if strings.HasPrefix(tail, soundsStore[i].signs[*lang]) {

                // Отложеная установка 
                //  - рассчитываем, что кол-во символов в soundsStore[i].sign идут от большего кол-ва к меньшему
                //  - при встрече с неопределенной твёрдостью/мягкостью или ещё что
                //    - возвращаем из checkRules() exception с ошибкой "нет зависимых данных" из первого правила и длинной звука
                //    - откусываем от tail длинну звка,
                //    - передаём звук, его размер и указатели на позицию (i, sounds.length) в отложенную функцию
                //    - после появления зависимостей выполняем поиск звука по soundsStore, повторяем выполнение правил
                //      - если остается хвост после отрезания от звука — exception
                //  - проверяем мягкость с конца слова
                //    - если встречаем `e`, это не означает, что слово мягкое, это могут быть твэрдые аффиксы -ğac-e/-qac-e
                //      - для проверки ищем следующее
                //    - если встречаем кирилическое у, о, — скорее таких не будет, так как они будут ждать зависимости,
                //      а sounds.sounds[i] == nil
                //    - если определить не удалось, то кидаем исключение и ждем завершения отложенных функций,
                //      от них тоже должны прийти исключения.
                //      Дальше нужно будет думать, как с этим бороться

                if !soundsStore[i].checkRules(sounds) {
                    continue
                }
                // fmt.Println(soundsStore[i].signs[*lang], tail, soundsStore[i].property)

                sounds.sounds[sounds.length] = &soundsStore[i]
                sounds.index[sounds.length] = uint8(position)
                position += utf8.RuneCountInString(soundsStore[i].signs[*lang])

                tail = strings.TrimPrefix(tail, soundsStore[i].signs[*lang])
                needTrim = false
                break
            }
        }
        if needTrim {  // exception
            fmt.Println(string([]rune(tail)[0]), tail, "BOO")

            // TODO: exception
            // TODO: посилання на SoundsStore на найближче через звичайний трансліт, або excepton
            // sounds.sounds[sounds.length] == nil
            sounds.index[sounds.length] = uint8(position)
            position += utf8.RuneCountInString(string([]rune(tail)[0]))

            tail = strings.TrimPrefix(tail, string([]rune(tail)[0]))
        }

        sounds.length++
    }

    sounds.setVowelHarmony(lang)

    return sounds
}

// твердые a, ı, o, u       (hard vowels)
// мягкие e, i, ö, ü        (soft vowels)

// если первое о или у то это не означает, что слово мягкое, нужно смотреть следующую согласную
// TRYIT

// Проблема: (кок, козь, кой) — а тут непонятно кроме "козь" по мягкому знаку
// по остальным нужно смотреть на аффикс, если он есть. 
// Таких корней очень мало. Можно посмотреть в скрипте транслитиратор как исключения.

func (sounds *Sounds) setVowelHarmony(lang *string) {
    fmt.Println(*sounds.Trace(lang))

    // fmt.Println(sounds.sounds[sounds.length - 1])
    for i := sounds.length - 1; i >= 0; i-- {
        if sounds.sounds[i] == nil || sounds.sounds[i].property & vowel != vowel {
            continue
        }

        fmt.Print(sounds.sounds[i].signs[*lang])
        if sounds.sounds[i].property & hardVowel == hardVowel {
            fmt.Print("_")
        }
        if sounds.sounds[i].property & softVowel == softVowel {
            fmt.Print("^")
        }
    }
    fmt.Println()
}

func (sound *Sound) checkRules(sounds *Sounds) bool {
    if len(sound.rules) == 0 {
        return true
    }

    for r := 0; r < len(sound.rules); r++ {
        if sound.rules[r](sounds) {
            // fmt.Println("*", sound.signs["crh"])
            return true
        }
    }

    return false
}

func (s *Sounds) Trace(inLang *string) *string {
    str := make([]string, s.length)

    for i := 0; i < s.length; i++ {
        if s.sounds[i] == nil {
            str[i] = "*"
            continue
        }

        str[i] = s.sounds[i].signs[*inLang]
    }

    result := strings.Join(str[:], "-")
    return &result
}
