package sounds

// Запам'ятовуємо індекс першої голосної
var setFirstVowel = func(s *Sounds, index int) (bool, bool) {
    if s.firstVowel == -1 {
        s.firstVowel = index
    }

    return false, true
}

// Додаємо до звуків після setFirstVowel, де немає інших правил
var eq = func(s *Sounds, index int) (bool, bool) {
    return true, true
}

// Звук на початку слова
var wordInitially = func (s *Sounds, index int) (bool, bool) {
    if index == 0 {
        return true, true
    }

    return false, true
}

// Звук на початку слова у «твердому» слові
var wordInitiallyInHard = func(s *Sounds, index int) (bool, bool) {
    if index != 0 {
        return false, true
    }

    if s.property & Vowel != Vowel {
        return false, false
    }

    if s.property & HardVowel == HardVowel {
        return true, true
    }

    return false, true
}

// Звук на початку слова у «м'якому» слові
var wordInitiallyInSoft = func(s *Sounds, index int) (bool, bool) {
    if index != 0 {
        return false, true
    }

    if s.property & Vowel != Vowel {
        return false, false
    }

    if s.property & SoftVowel == SoftVowel {
        return true, true
    }

    return false, true
}

// Це є перший голосний у «м'якому» слові
var firstVowelInSoft = func(s *Sounds, index int) (bool, bool) {
    if s.firstVowel != index {
        return false, true
    }

    if s.property & Vowel != Vowel {
        return false, false
    }

    if s.property & SoftVowel == SoftVowel {
        return true, true
    }

    return false, true
}

// Після голосної
var afterVowel = func(s *Sounds, index int) (bool, bool) {
    if index == 0 || s.sounds[index - 1].property & Vowel != Vowel {
        return false, true
    }

    return true, true
}

// Після голосної у «твердому» слові
var afterVowelInHard = func(s *Sounds, index int) (bool, bool) {
    if index == 0 || s.sounds[index - 1].property & Vowel != Vowel {
        return false, true
    }

    if s.property & Vowel != Vowel {
        return false, false
    }

    if s.property & HardVowel == HardVowel {
        return true, true
    }

    return false, true
}

// Після голосної у «м'якому» слові
var afterVowelInSoft = func(s *Sounds, index int) (bool, bool) {
    if index == 0 || s.sounds[index - 1].property & Vowel != Vowel {
        return false, true
    }

    if s.property & Vowel != Vowel {
        return false, false
    }

    if s.property & SoftVowel == SoftVowel {
        return true, true
    }

    return false, true
}

// Після приголосної
var afterConsonant = func(s *Sounds, index int) (bool, bool) {
    if index == 0 || s.sounds[index - 1].property & Consonant != Consonant {
        return false, true
    }

    return true, true
}

// Після «ъ»
var afterHardness = func(s *Sounds, index int) (bool, bool) {
    if index == 0 || s.sounds[index - 1].property & Hardness != Hardness {
        return false, true
    }

    return true, true
}

// Після «ь»
var afterSoftness = func(s *Sounds, index int) (bool, bool) {
    if index == 0 || s.sounds[index - 1].property & Softness != Softness {
        return false, true
    }

    return true, true
}

// Після «ь» у «твердому» слові
var afterSoftnessInHard = func(s *Sounds, index int) (bool, bool) {
    if index == 0 || s.sounds[index - 1].property & Softness != Softness {
        return false, true
    }

    if s.property & Vowel != Vowel {
        return false, false
    }

    if s.property & HardVowel == HardVowel {
        return true, true
    }

    return false, true
}

// Після «ь» у «м'якому» слові
var afterSoftnessInSoft = func(s *Sounds, index int) (bool, bool) {
    if index == 0 || s.sounds[index - 1].property & Softness != Softness {
        return false, true
    }

    if s.property & Vowel != Vowel {
        return false, false
    }

    if s.property & SoftVowel == SoftVowel {
        return true, true
    }

    return false, true
}
