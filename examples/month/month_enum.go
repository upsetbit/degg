package month

import (
	"errors"
)

type (
	Month int
)

const (
	_enumName = "Month"

	JANUARY Month = iota
	FEBRUARY
	MARCH
	APRIL
	MAY
	JUNE
	JULY
	AUGUST
	SEPTEMBER
	OCTOBER
	NOVEMBER
	DECEMBER

	_unknown Month = -1
)

var (
	InvalidMonthErr = errors.New("invalid Month, must be one of [JANUARY, FEBRUARY, MARCH, APRIL, MAY, JUNE, JULY, AUGUST, SEPTEMBER, OCTOBER, NOVEMBER, DECEMBER]")
)

func From(m int) (Month, error) {
	switch m {
	case 0:
		return JANUARY, nil
	case 1:
		return FEBRUARY, nil
	case 2:
		return MARCH, nil
	case 3:
		return APRIL, nil
	case 4:
		return MAY, nil
	case 5:
		return JUNE, nil
	case 6:
		return JULY, nil
	case 7:
		return AUGUST, nil
	case 8:
		return SEPTEMBER, nil
	case 9:
		return OCTOBER, nil
	case 10:
		return NOVEMBER, nil
	case 11:
		return DECEMBER, nil
	default:
		return _unknown, InvalidMonthErr
	}
}

func FromName(m string) (Month, error) {
	switch m {
	case "JANUARY":
		return JANUARY, nil
	case "FEBRUARY":
		return FEBRUARY, nil
	case "MARCH":
		return MARCH, nil
	case "APRIL":
		return APRIL, nil
	case "MAY":
		return MAY, nil
	case "JUNE":
		return JUNE, nil
	case "JULY":
		return JULY, nil
	case "AUGUST":
		return AUGUST, nil
	case "SEPTEMBER":
		return SEPTEMBER, nil
	case "OCTOBER":
		return OCTOBER, nil
	case "NOVEMBER":
		return NOVEMBER, nil
	case "DECEMBER":
		return DECEMBER, nil
	default:
		return _unknown, InvalidMonthErr
	}
}

func (m Month) String() string {
	switch m {
	case JANUARY:
		return "JANUARY"
	case FEBRUARY:
		return "FEBRUARY"
	case MARCH:
		return "MARCH"
	case APRIL:
		return "APRIL"
	case MAY:
		return "MAY"
	case JUNE:
		return "JUNE"
	case JULY:
		return "JULY"
	case AUGUST:
		return "AUGUST"
	case SEPTEMBER:
		return "SEPTEMBER"
	case OCTOBER:
		return "OCTOBER"
	case NOVEMBER:
		return "NOVEMBER"
	case DECEMBER:
		return "DECEMBER"
	default:
		return ""
	}
}

func (m Month) Int() int {
	return int(m)
}

func (m Month) Code() string {
	return _enumName + "." + m.String()
}

func (m Month) Repr() string {
	return _enumName + "(" + m.String() + ")"
}
