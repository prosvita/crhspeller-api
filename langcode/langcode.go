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

func GetId(name *string) (int, error) {
    reCrhAny := regexp.MustCompile(`^crh-[A-Z]{2,3}(-[0-9A-Z]+)?$`)

    switch {
        case *name == "":
            return -1, errors.New("langcode: language is not set")
        case strings.Compare(*name, "crh") == 0:
            return Crh, nil
        case strings.Compare(*name, "crh-RU") == 0:
            return Crh_RU, nil
        case strings.Compare(*name, "crh-RUS") == 0:
            return Crh_RU, nil
        case reCrhAny.MatchString(*name):
            return Crh, nil
    }

    return -1, errors.New("langcode: language code is not supported")
}
