package langcode

import (
    "strings"
    "regexp"
    "errors"
)

const (
    Crh     = 0
    Crh_RU  = 1

    Size    = 2
)

var (
    empty   = ""
    crh     = "crh"
    crh_RU  = "crh_RU"
)

func GetId(name *string) (int, error) {
    reCrhAny := regexp.MustCompile(`^crh_[A-Z]{2,3}(-[0-9A-Z]+)?$`)

    switch {
        case *name == "":
            return -1, errors.New("langcode: language is not set")
        case strings.Compare(*name, "crh") == 0:
            return Crh, nil
        case strings.Compare(*name, "crh_RU") == 0:
            return Crh_RU, nil
        case strings.Compare(*name, "crh_RUS") == 0:
            return Crh_RU, nil
        case strings.Compare(*name, "crh_CYR") == 0:
        case reCrhAny.MatchString(*name):
            return Crh, nil
    }

    return -1, errors.New("langcode: language code is not supported")
}

func Lang(id int) (result *string, err error) {
    result = GetLang(id)
    err = nil

    if *result == empty {
        err = errors.New("langcode: language is not supported")
    }
    return
}

func GetLang(id int) (*string) {
    switch {
        case id == Crh:
            return &crh
        case id == Crh_RU:
            return &crh_RU
    }
    return &empty
}
