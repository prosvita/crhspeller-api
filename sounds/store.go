package sounds

var dumpVowel = Sound { Signs{"?", "?"}, []ruleFunc{}, Vowel }

var soundsStore = SoundsStore {
// TODO: відсортувати по лічильникам використання, чи зробити динамічне сортування

    { Signs{"c",  "дж"}, []ruleFunc{}, Consonant },
    { Signs{"ğ",  "гъ"}, []ruleFunc{}, Consonant | Hardness },
    { Signs{"ñ",  "нъ"}, []ruleFunc{}, Consonant | Hardness },
    { Signs{"q",  "къ"}, []ruleFunc{}, Consonant | Hardness },
    { Signs{"ts", "ц" }, []ruleFunc{}, Consonant },
    { Signs{"şç", "щ" }, []ruleFunc{}, Consonant },

    { Signs{"ye", "е"}, []ruleFunc{setFirstVowel, wordInitially, afterVowel, afterSoftness}, SoftVowel },
    { Signs{"e",  "е"}, []ruleFunc{setFirstVowel, afterConsonant}, SoftVowel },
    { Signs{"e",  "э"}, []ruleFunc{setFirstVowel, eq}, SoftVowel },

    { Signs{"ya", "я"}, []ruleFunc{setFirstVowel, wordInitially, afterVowel, afterSoftness}, HardVowel },
    { Signs{"â",  "я"}, []ruleFunc{setFirstVowel, afterConsonant}, HardVowel },

    { Signs{"yo", "ё"}, []ruleFunc{setFirstVowel, wordInitiallyInHard, afterVowel, afterHardness, afterSoftness}, HardVowel },
    { Signs{"yö", "ё"}, []ruleFunc{setFirstVowel, wordInitiallyInSoft}, SoftVowel },
// TRYIT: незрозуміло, що робити, як-що жодне правило не спрацює, то буде y+o|y+ö, а з ё?
    { Signs{"ö",  "о"}, []ruleFunc{setFirstVowel, firstVowelInSoft}, SoftVowel },
    { Signs{"ö",  "ё"}, []ruleFunc{setFirstVowel, afterConsonant}, SoftVowel },
    { Signs{"o",  "о"}, []ruleFunc{setFirstVowel, eq}, HardVowel },

    { Signs{"yu", "ю"}, []ruleFunc{setFirstVowel, wordInitiallyInHard, afterVowelInHard, afterSoftnessInHard}, HardVowel },
    { Signs{"yü", "ю"}, []ruleFunc{setFirstVowel, wordInitiallyInSoft, afterVowelInSoft, afterSoftnessInSoft}, SoftVowel },
    { Signs{"ü",  "у"}, []ruleFunc{setFirstVowel, firstVowelInSoft}, SoftVowel },
    { Signs{"ü",  "ю"}, []ruleFunc{setFirstVowel, afterConsonant}, SoftVowel },
    { Signs{"u",  "у"}, []ruleFunc{setFirstVowel, eq}, HardVowel },

    { Signs{"a",  "а"}, []ruleFunc{setFirstVowel, eq}, HardVowel },
    { Signs{"b",  "б"}, []ruleFunc{}, Consonant },
    { Signs{"ç",  "ч"}, []ruleFunc{}, Consonant },
    { Signs{"d",  "д"}, []ruleFunc{}, Consonant },
    { Signs{"f",  "ф"}, []ruleFunc{}, Consonant },
    { Signs{"g",  "г"}, []ruleFunc{}, Consonant },
    { Signs{"h",  "х"}, []ruleFunc{}, Consonant },
    { Signs{"ı",  "ы"}, []ruleFunc{setFirstVowel, eq}, HardVowel },
    { Signs{"i",  "и"}, []ruleFunc{setFirstVowel, eq}, SoftVowel },
    { Signs{"j",  "ж"}, []ruleFunc{}, Consonant },
    { Signs{"k",  "к"}, []ruleFunc{}, Consonant },
    { Signs{"l",  "л"}, []ruleFunc{}, Consonant },
    { Signs{"m",  "м"}, []ruleFunc{}, Consonant },
    { Signs{"n",  "н"}, []ruleFunc{}, Consonant },
    { Signs{"p",  "п"}, []ruleFunc{}, Consonant },
    { Signs{"r",  "р"}, []ruleFunc{}, Consonant },
    { Signs{"s",  "с"}, []ruleFunc{}, Consonant },
    { Signs{"ş",  "ш"}, []ruleFunc{}, Consonant },
    { Signs{"t",  "т"}, []ruleFunc{}, Consonant },
    { Signs{"v",  "в"}, []ruleFunc{}, Consonant },
    { Signs{"y",  "й"}, []ruleFunc{}, Consonant },
    { Signs{"z",  "з"}, []ruleFunc{}, Consonant },

// TRYIT: додати правило яке визначає, за якими буквами не можуть йти "ъ" чи "ь"
// FIXIT: навчити визначати латинці, за яким звуком має йти "ь"
    { Signs{"",   "ъ"}, []ruleFunc{}, Hardness },
    { Signs{"",   "ь"}, []ruleFunc{}, Softness },
}
