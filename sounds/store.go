package sounds

var soundsStore = SoundsStore {
// TODO: навесить на звуки счетчики, прогнать большой текс, отсортировать по популярности

    { soundMap{"crh":"c",  "crh_RU":"дж"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"ğ",  "crh_RU":"гъ"}, []ruleFunc{}, consonant | hardness },
    { soundMap{"crh":"ñ",  "crh_RU":"нъ"}, []ruleFunc{}, consonant | hardness },
    { soundMap{"crh":"q",  "crh_RU":"къ"}, []ruleFunc{}, consonant | hardness },
    { soundMap{"crh":"ts", "crh_RU":"ц" }, []ruleFunc{}, consonant },
    { soundMap{"crh":"şç", "crh_RU":"щ" }, []ruleFunc{}, consonant },

    { soundMap{"crh":"ye", "crh_RU":"е"}, []ruleFunc{wordInitially, afterVowel, afterSoftness}, softVowel },
    { soundMap{"crh":"e",  "crh_RU":"е"}, []ruleFunc{afterConsonant}, softVowel },
    { soundMap{"crh":"e",  "crh_RU":"э"}, []ruleFunc{}, softVowel },

    { soundMap{"crh":"ya", "crh_RU":"я"}, []ruleFunc{wordInitially, afterVowel, afterSoftness}, hardVowel },
    { soundMap{"crh":"â",  "crh_RU":"я"}, []ruleFunc{afterConsonant}, hardVowel },

    { soundMap{"crh":"yo", "crh_RU":"ё"}, []ruleFunc{wordInitiallyInHard, afterVowel, afterHardness, afterSoftness}, hardVowel },
    { soundMap{"crh":"yö", "crh_RU":"ё"}, []ruleFunc{wordInitiallyInSoft}, softVowel },

//TRYIT: тут может нужно будет "ё" и "о" местами поменять
    { soundMap{"crh":"ö",  "crh_RU":"ё"}, []ruleFunc{afterConsonant}, softVowel },
    { soundMap{"crh":"ö",  "crh_RU":"о"}, []ruleFunc{firstVowelInSoft}, softVowel },
    { soundMap{"crh":"o",  "crh_RU":"о"}, []ruleFunc{}, hardVowel },

    { soundMap{"crh":"yu", "crh_RU":"ю"}, []ruleFunc{wordInitiallyInHard, afterVowelInHard, afterSoftnessInHard}, hardVowel },
    { soundMap{"crh":"yü", "crh_RU":"ю"}, []ruleFunc{wordInitiallyInSoft, afterVowelInSoft, afterSoftnessInSoft}, softVowel },
    { soundMap{"crh":"ü",  "crh_RU":"ю"}, []ruleFunc{afterConsonant}, softVowel },
    { soundMap{"crh":"ü",  "crh_RU":"у"}, []ruleFunc{firstVowelInSoft}, softVowel },
    { soundMap{"crh":"u",  "crh_RU":"у"}, []ruleFunc{}, hardVowel },

    { soundMap{"crh":"a",  "crh_RU":"а"}, []ruleFunc{}, hardVowel },
    { soundMap{"crh":"b",  "crh_RU":"б"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"ç",  "crh_RU":"ч"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"d",  "crh_RU":"д"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"f",  "crh_RU":"ф"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"g",  "crh_RU":"г"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"h",  "crh_RU":"х"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"ı",  "crh_RU":"ы"}, []ruleFunc{}, hardVowel },
    { soundMap{"crh":"i",  "crh_RU":"и"}, []ruleFunc{}, softVowel },
    { soundMap{"crh":"j",  "crh_RU":"ж"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"k",  "crh_RU":"к"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"l",  "crh_RU":"л"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"m",  "crh_RU":"м"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"n",  "crh_RU":"н"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"p",  "crh_RU":"п"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"r",  "crh_RU":"р"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"s",  "crh_RU":"с"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"ş",  "crh_RU":"ш"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"t",  "crh_RU":"т"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"v",  "crh_RU":"в"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"y",  "crh_RU":"й"}, []ruleFunc{}, consonant },
    { soundMap{"crh":"z",  "crh_RU":"з"}, []ruleFunc{}, consonant },

//TRYIT: добавить правило, по которомуму определять, после каких букв не может идти "ъ" или "ь"
//FIXIT: научится определять в латинице, после каких звуков должен идти "ь"
    { soundMap{"crh":"",   "crh_RU":"ъ"}, []ruleFunc{}, hardness },
    { soundMap{"crh":"",   "crh_RU":"ь"}, []ruleFunc{}, softness },
}
