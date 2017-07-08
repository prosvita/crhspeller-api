package sounds

import "fmt"
import (
    lc "../langcode"
    "unicode/utf8"
    "strings"
)

type Signs [lc.Size]string
type ruleFunc func(*Sounds, int) (bool, bool)

type Sound struct {
    signs Signs
    rules []ruleFunc
    property uint
}
type SoundsStore []Sound

type Sounds struct {
    property uint
    length int
    firstVowel int
    sounds []*Sound
    index []uint8
}

const (
    Vowel               = 1 << iota                         // голосна
    HardVowel           = Vowel | (1 << iota)               // тверда голосна
    SoftVowel           = Vowel | (1 << iota)               // м'яка голосна

    Consonant           = 1 << iota                         // приголосна

    // sonorant            = Consonant | (1 << iota)           // сонант (сонорний приголосний звук)
    // voiced              = 1 << iota                         // дзвінка
    // voiceless           = 1 << iota                         // глуха
    // obstruent           = Consonant | (1 << iota)           // шумна приголосна
    // voicedObstruent     = Consonant | voiced | obstruent    // дзвінка шумна приголосна
    // voicelessObstruent  = Consonant | voiceless | obstruent // глуха шумна приголосна

    Hardness            = 1 << iota                         // звук має `Ъ`
    Softness            = 1 << iota                         // звук має `Ь`
)

func New(lang int, text *string) (sounds *Sounds) {

    tail := strings.ToLower(*text)
    length := utf8.RuneCountInString(tail) + 2   //FIXIT: 2 провсяквипадок, як що виходимо за межі, то append(sounds) та append (index)
    position := 0

    sounds = &Sounds{0, 0, -1, make([]*Sound, length), make([]uint8, length)}

    for utf8.RuneCountInString(tail) > 0 {
        needTrim := true

        for i := 0; i < len(soundsStore); i++ {
            if strings.HasPrefix(tail, soundsStore[i].signs[lang]) {

                caught, compute := soundsStore[i].checkRules(sounds, sounds.length)
                if !caught && compute { // жодне правило не спрацювало
                    continue
                }

                if compute {
                    sounds.sounds[sounds.length] = &soundsStore[i]
                } else {
                    sounds.sounds[sounds.length] = &dumpVowel
                    defer sounds.resetSound(sounds.length, i, lang)
                }
                sounds.index[sounds.length] = uint8(position)
                position += utf8.RuneCountInString(soundsStore[i].signs[lang])

                tail = strings.TrimPrefix(tail, soundsStore[i].signs[lang])
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

func (sounds *Sounds) resetSound(index, start, lang int) {
    needTrim := true
    tail := soundsStore[start].signs[lang]

    for i := start; i < len(soundsStore); i++ {
        if strings.HasPrefix(tail, soundsStore[i].signs[lang]) {

            caught, compute := soundsStore[i].checkRules(sounds, index)
// fmt.Printf("%v> BOO %v %v %v\n",index , caught, compute, soundsStore[i].signs[lang])
            if !caught && compute { // жодне правило не спрацювало
                continue
            }

            if compute {
                sounds.sounds[index] = &soundsStore[i]
            } else {
                //FIXIT: потрібно з'ясувати, у яких випадках це можливе, і всі ці випадки відобразити у rules
                fmt.Printf("%v> EXCEPTION not compute\n", index)
                return
            }

            needTrim = false
            break
        }
    }
    if needTrim {  // exception
        fmt.Printf("%v> EXCEPTION needTrim\n", index)
        //FIXIT: такої ситуації не має статись, але як що виникне, потрібно розсунути масив
        return
    }
}

// Шукаємо з кінця голосні
//   - a, â, ı, o, u — hard vowels
//   - e, i, ö, ü — soft vowels
// За ними встановлюємо гармонію слова
// FIXIT: як визначити гармонію у словах з одним голосним в корні, який замінили на dumpVowel:
//  - подивитись Softness у sounds[i+2] (козь)
//  - для crh o, u — завжди hard
//  - для crh ö, ü — завжди soft
//  - що ще?
func (sounds *Sounds) setVowelHarmony(  lang int) {

    for i := sounds.length - 1; i >= 0; i-- {
        // FIXIT: що робити з dumpVowel?
        if sounds.sounds[i] == nil || sounds.sounds[i] == &dumpVowel || sounds.sounds[i].property & Vowel != Vowel {
            continue
        }

        // FIXIT: `e` не дає гарантії, що слово м'яке, потрібно перевірити на тверді афікси -ğace/-qace
        if sounds.sounds[i].property & HardVowel == HardVowel {
            sounds.property |= HardVowel
            // FIXIT: виключення
            break
        }
        if sounds.sounds[i].property & SoftVowel == SoftVowel {
            sounds.property |= SoftVowel
            // FIXIT: виключення
            break
        }
    }
}

func (sound *Sound) checkRules(sounds *Sounds, index int) (bool, bool) {
    if len(sound.rules) == 0 {
        return true, true
    }

    for r := 0; r < len(sound.rules); r++ {
        caught, compute := sound.rules[r](sounds, index)
        if caught || !compute {
            return caught, compute
        }
    }

    return false, true
}

func (s *Sounds) Trace(inLang int) *string {
    str := make([]string, s.length)

    for i := 0; i < s.length; i++ {
        if s.sounds[i] == nil {
            str[i] = "*"
            continue
        }

        str[i] = s.sounds[i].signs[inLang]
    }

    result := strings.Join(str[:], "-")
    return &result
}
