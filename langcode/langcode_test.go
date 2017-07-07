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
        { "crh-RU",     1, false },
        { "crh-RUS",    1, false },
        { "crh-TR",     0, false },
        { "crh-UA-43",  0, false },
        { "crh-UA-",   -1, true  },
    }

    for _, test := range tests {
        if id, err := getLangCodeId(&test.text); id != test.id || e2Bool(err) != test.err {
            t.Errorf("getLangCodeId(%q) == %d, error:%v; want %d, error:%v)",
                test.text, id, e2Bool(err), test.id, test.err)
        }
    }
}

func e2Bool(err error) bool {
    if err != nil {
        return true
    }
    return false
}
