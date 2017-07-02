package sounds

// import "fmt"

// Звук на початку слова
var wordInitially = func (s *Sounds) bool {
    if s.length == 0 {
        return true
    }

    return false
}

// TODO
// Звук на початку слова у «твердому» слові
var wordInitiallyInHard = func(s *Sounds) bool {return false}

// TODO
// Звук на початку слова у «м'якому» слові
var wordInitiallyInSoft = func(s *Sounds) bool {return false}

// TODO
// Це є перший голосний у «м'якому» слові
var firstVowelInSoft = func(s *Sounds) bool {return false}

// Після голосної
var afterVowel = func(s *Sounds) bool {
    if s.length == 0 {
        return false
    }

    if bool(s.sounds[s.length - 1].property & vowel == vowel) {
        return true
    }

    return false
}

// TODO
// Після голосної у «твердому» слові
var afterVowelInHard = func(s *Sounds) bool {return false}

// TODO
// Після голосної у «м'якому» слові
var afterVowelInSoft = func(s *Sounds) bool {return false}

// Після приголосної
var afterConsonant = func(s *Sounds) bool {
    if s.length == 0 {
        return false
    }

    if bool(s.sounds[s.length - 1].property & consonant == consonant) {
        return true
    }

    return false
}

// Після «ъ»
var afterHardness = func(s *Sounds) bool {
    if s.length == 0 {
        return false
    }

    if bool(s.sounds[s.length - 1].property & hardness == hardness) {
        return true
    }

    return false
}

// Після «ь»
var afterSoftness = func(s *Sounds) bool {
    if s.length == 0 {
        return false
    }

    if bool(s.sounds[s.length - 1].property & softness == softness) {
        return true
    }

    return false
}

// TODO
// Після «ь» у «твердому» слові
var afterSoftnessInHard = func(s *Sounds) bool {return false}

// TODO
// Після «ь» у «м'якому» слові
var afterSoftnessInSoft = func(s *Sounds) bool {return false}
