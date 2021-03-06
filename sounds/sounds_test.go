package sounds

import (
    lc "../langcode"
    "testing"
)

type TestData []string

func TestNew(t *testing.T) {
    var tests = []TestData {

        { "evlerimizdendirlermi", "эвлеримиздендирлерми", "e-v-l-e-r-i-m-i-z-d-e-n-d-i-r-l-e-r-m-i", "э-в-л-е-р-и-м-и-з-д-е-н-д-и-р-л-е-р-м-и" },
        { "yazıcılarımızğa", "языджыларымызгъа", "ya-z-ı-c-ı-l-a-r-ı-m-ı-z-ğ-a", "я-з-ы-дж-ы-л-а-р-ы-м-ы-з-гъ-а" },

        { "böyle", "бойле", "b-ö-y-l-e", "б-о-й-л-е" }, // м'яке
        { "cümle", "джумле", "c-ü-m-l-e", "дж-у-м-л-е" }, // м'яке
        { "boldur", "болдур", "b-o-l-d-u-r", "б-о-л-д-у-р" }, // тверде
        { "tursat", "турсат", "t-u-r-s-a-t", "т-у-р-с-а-т" }, // тверде

        // { "yel", "ель", "ye-l-", "е-л-ь" },
        // { "el", "эль", "e-l-", "э-л-ь" },
        { "binaen", "бинаэн", "b-i-n-a-e-n", "б-и-н-а-э-н" },

        { "yekâne", "екяне", "ye-k-â-n-e", "е-к-я-н-е" },
        { "yekânedenlerniñ", "екянеденлернинъ", "ye-k-â-n-e-d-e-n-l-e-r-n-i-ñ", "е-к-я-н-е-д-е-н-л-е-р-н-и-нъ" },

        // { "söz", "сёз", "s-ö-z", "с-ё-з" },
        // { "örnek", "орьнек", "örnek", "орьнек" },
        // { "özgün", "озьгюн", "ö-z--g-ü-n", "о-з-ь-г-ю-н" },
        // { "körmek", "корьмек", "körmek", "корьмек" },
        // { "yol", "ёл", "yo-l", "ё-л" },
        { "yorğan", "ёргъан", "yo-r-ğ-a-n", "ё-р-гъ-а-н" },
        // { "afyon", "афьён", "afyon", "афьён" },
        { "yöneliş", "ёнелиш", "yö-n-e-l-i-ş", "ё-н-е-л-и-ш" },
        // { "tüz", "тюз", "t-ü-z", "т-ю-з" },
        // { "sürmek", "сюрмек", "s-ü-r-m-e-k", "с-ю-р-м-е-к" },
        // { "ürmet", "урьмет", "ürmet", "урьмет" },
        // { "üç", "учь", "üç", "учь" },

        { "mümkün", "мумкюн", "m-ü-m-k-ü-n", "м-у-м-к-ю-н" },
        // { "kün", "кунь", "kün", "кунь" },
        { "yutmaq", "ютмакъ", "yu-t-m-a-q", "ю-т-м-а-къ" },
        { "yuqarı", "юкъары", "yu-q-a-r-ı", "ю-къ-а-р-ы" },
        { "toyunda", "тоюнда", "t-o-yu-n-d-a", "т-о-ю-н-д-а" },
        { "qoyulmaq", "къоюлмакъ", "q-o-yu-l-m-a-q", "къ-о-ю-л-м-а-къ" },
        { "yüzüm", "юзюм", "yü-z-ü-m", "ю-з-ю-м" },
        // { "yürmek", "юрьмек", "yü-r--m-e-k", "ю-р-ь-м-е-к" },
        { "köyünde", "коюнде", "k-ö-yü-n-d-e", "к-о-ю-н-д-е" },
        // { "tüyüm", "тююм", "t-ü-yü-m", "т-ю-ю-м" },
        { "yahşı", "яхшы", "ya-h-ş-ı", "я-х-ш-ы" },
        // { "Yağya", "Ягъя", "ya-ğ-ya", "я-гъ-я" },
        { "qaya", "къая", "q-a-ya", "къ-а-я" },
        { "kâr", "кяр", "k-â-r", "к-я-р" },

        // виключення — тверде
        // { "tarih", "", "tarih", "" },
        // { "tarihta", "", "tarihta", "" },
        // { "tarihlar", "", "tarihlar", "" },
        // виключення — м'яке
        // { "teatr", "", "teatr", "" },
        // { "teatrde", "", "teatrde", "" },
        // { "teatrler", "", "teatrler", "" },

        // { "vaqıt", "", "vaqıt", "" }, // тверде
        // { "vaqıtta", "", "vaqıtta", "" },
        // { "vaqıtlar", "", "vaqıtlar", "" },
        // але зустрічаються
        // { "vaqıtte", "", "vaqıtte", "" },
        // { "vaqıtler", "", "vaqıtler", "" },

        // { "saat", "", "saat", "" }, // тверде
        // { "saatta", "", "saatta", "" },
        // { "saatlar", "", "saatlar", "" },
        // але зустрічаються
        // { "saatte", "", "saatte", "" },
        // { "saatler", "", "saatler", "" },

        // афікс -ki має тільки одну форму
        // { "Yaltadaki", "", "yaltadaki", "" },

        // "ъ" в складних словах в корінні
        // { "", "демиръёл", "", "демиръёл" },
        // { "yamyam", "ямъям", "ya-m--ya-m", "я-м-ъ-я-м" },

        // м'якість звуку л
        // { "sual", "суаль", "sual", "суаль" },
        // { "", "тиль", "", "тиль" },
        // { "", "сильмек", "", "сильмек" },
        // { "", "ольди", "", "ольди" },

        // роздільний знак в корінні після приголосних
        // { "dünya", "дюнья", "dünya", "дюнья" },
        // { "karyer", "карьер", "karyer", "карьер" },

        // у м'якому слові після твердого приголосного (ч, н, р, с, т, з) у першому твердому складі в корені
        // { "öz",     "озь",      "ö-z-",         "о-з-ь" },
        // { "kün",    "кунь",     "k-ü-n-",       "к-у-н-ь" },
        // { "yür",    "юрь",      "y-ü-r-",       "ю-р-ь" },
        // { "ürmet",  "урьмет",   "ü-r--m-e-t",   "у-р-ь-м-е-т" },
        // { "küç",    "кучь",     "k-ü-ç-",       "к-у-ч-ь" },
        // { "kün",    "кунь",     "k-ü-n-",       "к-у-н-ь" },
        // { "yür",    "юрь",      "y-ü-r-",       "ю-р-ь" },
        // { "köster", "косьтер",  "k-ö-s--t-e-r", "к-о-с-ь-т-е-р" },
        // { "köt",    "коть",     "k-ö-t-",       "к-о-т-ь" },
        // { "köz",    "козь",     "k-ö-z-",       "к-о-з-ь" },

        // виключення — м'яке ???
        // { "kök",    "кок",      "k-ö-k",        "к-о-к" },

        // { "aaa", "ааа", "a-a-a", "а-а-а" }, // Exception
        // { "bbb", "ббб", "b-b-b", "б-б-б" }, // Exception
        { "", "", "", "" }, // Ok
    }

    for _, test := range tests {
        test.crossAssert(t)
    }
}

func (test TestData) crossAssert(t *testing.T) {
    var matrix = []struct {
        sourceLang, targetLang int
        textIndex, wantIndex int
    }{
        { lc.Crh,    lc.Crh,    0, 2 },
        { lc.Crh,    lc.Crh_RU, 0, 3 },
        { lc.Crh_RU, lc.Crh,    1, 2 },
        { lc.Crh_RU, lc.Crh_RU, 1, 3 },
    }
    var sounds *Sounds

    for i, m := range matrix {
        if i == 0 || m.sourceLang != matrix[i - 1].sourceLang {
            sounds = New(m.sourceLang, &test[m.textIndex])
        }

        got := *sounds.Trace(m.targetLang)
        if got != test[m.wantIndex] {
            t.Errorf("New(%q, %q)[%q] == %q; want %q",
                *lc.GetLang(m.sourceLang), test[m.textIndex], *lc.GetLang(m.targetLang),
                got, test[m.wantIndex])
        }
    }
}

func TestInsert(t *testing.T) {
    var tests = []struct {
        position int
        text, want string
        err bool
    }{
        { 0, "", "x", false },
        { 1, "", "", true },
        { 0, "ana", "x-a-n-a", false },
        { 2, "ana", "a-n-x-a", false },
        { 3, "ana", "a-n-a-x", false },
        { 4, "ana", "a-n-a", true },
        { -1, "ana", "a-n-a", true },
        { 2, "yahşı", "ya-h-x-ş-ı", false },
        { 4, "yahşı", "ya-h-ş-ı-x", false },
    }
    var dump = Sound { Signs{"x", "х"}, []ruleFunc{}, Consonant }

    for _, test := range tests {
        sounds := New(lc.Crh, &test.text)
        err := sounds.Insert(test.position, &dump)
        got := *sounds.Trace(lc.Crh)
        if got != test.want || e2Bool(err) != test.err {
            t.Errorf("(%q).InsertTo(%v, %q) => %q, error:%v; want %q, error:%v",
                test.text, test.position, "x", got, e2Bool(err), test.want, test.err)
        }
    }
}

func e2Bool(err error) bool {
    if err != nil {
        return true
    }
    return false
}
