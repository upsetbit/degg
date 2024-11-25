package capital

import (
	"errors"
)

type (
	Capital string
)

const (
	_enumName = "Capital"

	AUSTRALIA Capital = "Canberra"
	BRAZIL    Capital = "Brasilia"
	CANADA    Capital = "Ottawa"
	CHINA     Capital = "Beijing"
	FRANCE    Capital = "Paris"
	GERMANY   Capital = "Berlin"
	JAPAN     Capital = "Tokyo"
	RUSSIA    Capital = "Moscow"
	UK        Capital = "London"
	USA       Capital = "Washington"

	_unknown Capital = ""
)

var (
	InvalidCapitalErr = errors.New("invalid Capital, must be one of [AUSTRALIA, BRAZIL, CANADA, CHINA, FRANCE, GERMANY, JAPAN, RUSSIA, UK, USA]")
)

func From(c string) (Capital, error) {
	switch c {
	case "Canberra":
		return AUSTRALIA, nil
	case "Brasilia":
		return BRAZIL, nil
	case "Ottawa":
		return CANADA, nil
	case "Beijing":
		return CHINA, nil
	case "Paris":
		return FRANCE, nil
	case "Berlin":
		return GERMANY, nil
	case "Tokyo":
		return JAPAN, nil
	case "Moscow":
		return RUSSIA, nil
	case "London":
		return UK, nil
	case "Washington":
		return USA, nil
	default:
		return _unknown, InvalidCapitalErr
	}
}

func FromName(c string) (Capital, error) {
	switch c {
	case "AUSTRALIA":
		return AUSTRALIA, nil
	case "BRAZIL":
		return BRAZIL, nil
	case "CANADA":
		return CANADA, nil
	case "CHINA":
		return CHINA, nil
	case "FRANCE":
		return FRANCE, nil
	case "GERMANY":
		return GERMANY, nil
	case "JAPAN":
		return JAPAN, nil
	case "RUSSIA":
		return RUSSIA, nil
	case "UK":
		return UK, nil
	case "USA":
		return USA, nil
	default:
		return _unknown, InvalidCapitalErr
	}
}

func (c Capital) String() string {
	return string(c)
}

func (c Capital) Int() int {
	switch c {
	case AUSTRALIA:
		return 0
	case BRAZIL:
		return 1
	case CANADA:
		return 2
	case CHINA:
		return 3
	case FRANCE:
		return 4
	case GERMANY:
		return 5
	case JAPAN:
		return 6
	case RUSSIA:
		return 7
	case UK:
		return 8
	case USA:
		return 9
	default:
		return -1
	}
}

func (c Capital) Code() string {
	return _enumName + "." + c.String()
}

func (c Capital) Repr() string {
	return _enumName + "(" + c.String() + ")"
}
