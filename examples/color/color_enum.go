package color

import (
	"errors"
)

type (
	Color string
)

const (
	_enumName = "Color"

	RED   Color = "RED"
	BLUE  Color = "BLUE"
	GREEN Color = "GREEN"
	WHITE Color = "WHITE"
	BLACK Color = "BLACK"

	_unknown Color = ""
)

var (
	InvalidColorErr = errors.New("invalid Color, must be one of [RED, BLUE, GREEN, WHITE, BLACK]")
)

func From(c string) (Color, error) {
	switch c {
	case "RED":
		return RED, nil
	case "BLUE":
		return BLUE, nil
	case "GREEN":
		return GREEN, nil
	case "WHITE":
		return WHITE, nil
	case "BLACK":
		return BLACK, nil
	default:
		return _unknown, InvalidColorErr
	}
}

func FromName(c string) (Color, error) {
	switch c {
	case "RED":
		return RED, nil
	case "BLUE":
		return BLUE, nil
	case "GREEN":
		return GREEN, nil
	case "WHITE":
		return WHITE, nil
	case "BLACK":
		return BLACK, nil
	default:
		return _unknown, InvalidColorErr
	}
}

func (c Color) String() string {
	return string(c)
}

func (c Color) Int() int {
	switch c {
	case RED:
		return 0
	case BLUE:
		return 1
	case GREEN:
		return 2
	case WHITE:
		return 3
	case BLACK:
		return 4
	default:
		return -1
	}
}

func (c Color) Code() string {
	return _enumName + "." + c.String()
}

func (c Color) Repr() string {
	return _enumName + "(" + c.String() + ")"
}
