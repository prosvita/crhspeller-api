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
    wordType uint
    sounds []*Sound
    index []uint8
    length int
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
    undefinedVowel      = vowel | (1 << iota)               // не визначена голосна (тверда чи м'яка) (перша голосна о у)

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


func New(lang *string, text *string) *Sounds {

    tail := strings.ToLower(*text)
    length := utf8.RuneCountInString(tail) + 2   // 2 провсяквипадок
    sounds := Sounds{0, make([]*Sound, length), make([]uint8, length), 0}
    position := 0

    for utf8.RuneCountInString(tail) > 0 {
        needTrim := true

        for i := 0; i < len(soundsStore); i++ {
            if strings.HasPrefix(tail, soundsStore[i].signs[*lang]) {

                if !soundsStore[i].checkRules(&sounds) {
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
            // sounds.sounds[sounds.length] = nil
            sounds.index[sounds.length] = uint8(position)
            position += utf8.RuneCountInString(string([]rune(tail)[0]))

            tail = strings.TrimPrefix(tail, string([]rune(tail)[0]))
        }

        sounds.length++
    }

    return &sounds
}

func (sound *Sound) checkRules(sounds *Sounds) bool {
    if len(sound.rules) == 0 {
        return true
    }

    for r := 0; r < len(sound.rules); r++ {
        if sound.rules[r](sounds) {
fmt.Println("*", sound.signs["crh"])
            return true
        }
    }

    return false
}

func (s *Sounds) Debug(inLang *string) *string {
    str := make([]string, s.length)

    for i := 0; i < s.length; i++ {
        str[i] = s.sounds[i].signs[*inLang]
    }

    result := strings.Join(str[:], "-")
    return &result
}
