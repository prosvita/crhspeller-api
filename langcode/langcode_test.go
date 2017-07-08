package langcode

import "testing"

func TestGetId(t *testing.T) {
    var tests = []struct {
        text string
        id int
        err bool
    }{
        { "",          -1, true  },
        { "crh",        0, false },
        { "en",        -1, true  },
        { "crh_RU",     1, false },
        { "crh-RU",    -1, true  },
        { "crh_ru",    -1, true  },
        { "crh_RUS",    1, false },
        { "crh_TR",     0, false },
        { "crh_UA-43",  0, false },
        { "crh_CYR",   -1, true  },
        { "crh_cyr",   -1, true  },
    }

    for _, test := range tests {
        if id, err := GetId(&test.text); id != test.id || e2Bool(err) != test.err {
            t.Errorf("GetId(%q) == %d, error:%v; want %d, error:%v)",
                test.text, id, e2Bool(err), test.id, test.err)
        }
    }
}

func TestLang(t *testing.T) {
    var tests = []struct {
        id int
        text string
        err bool
    }{
        { 0, "crh",    false },
        { 1, "crh_RU", false },
        { 2, "",       true  },
    }

    for _, test := range tests {
        if lang, err := Lang(test.id); *lang != test.text || e2Bool(err) != test.err {
            t.Errorf("Lang(%d) == %q, error:%v; want %q, error:%v)",
                test.id, *lang, e2Bool(err), test.text, test.err)
        }
    }
}

func e2Bool(err error) bool {
    if err != nil {
        return true
    }
    return false
}
