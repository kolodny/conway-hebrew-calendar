package conway

import (
	"fmt"
	"time"
)

type quality int

const (
	abundant  = 1
	regular   = 0
	deficient = -1
)

type HebrewYear struct {
	Y        int // the year
	s        quality
	leapYear bool
}

func (h *HebrewYear) Length() int {
	// ref: p. 4
	days := 354 + int(h.s)
	if h.leapYear {
		days += 30
	}
	return days
}

// String implements stringer.String.
func (h HebrewYear) String() string {
	return fmt.Sprintf("%d", h.Y)
}

func (h *HebrewYear) monthLength(m HebrewMonth) int {
	switch m {
	case Nissan, Sivan, Av, Tishrei, Shevat, Adar_I:
		return 30
	case Iyar, Tamuz, Elul, Tevet, Adar_II, Adar:
		return 29
	case Marcheshvan:
		if h.s == 1 { // ref: p. 4
			return 30
		}
		return 29
	case Kislev:
		if h.s == -1 { // ref: p. 4
			return 29
		}
		return 30
	default:
		panic(fmt.Sprint("Invalid month:", m))
	}
}

type HebrewMonth int

const (
	Nissan HebrewMonth = 3 + iota
	Iyar
	Sivan
	Tamuz
	Av
	Elul
	Tishrei
	Marcheshvan
	Kislev
	Tevet
	Shevat
	Adar_I
	Adar_II
	Adar
)

func (m HebrewMonth) num() int {
	if m < Tishrei {
		return int(m)
	} else if m == Adar {
		return int(Adar_I) - 1
	}
	return int(m) - 1 // Tishrei and Elul are both height 8
}

func (m HebrewMonth) String() string {
	switch m {
	case Tishrei:
		return "Tishrei"
	case Marcheshvan:
		return "Marcheshvan"
	case Kislev:
		return "Kislev"
	case Tevet:
		return "Tevet"
	case Shevat:
		return "Shevat"
	case Adar:
		return "Adar"
	case Adar_I:
		return "Adar_I"
	case Adar_II:
		return "Adar_II"
	case Nissan:
		return "Nissan"
	case Iyar:
		return "Iyar"
	case Sivan:
		return "Sivan"
	case Tamuz:
		return "Tamuz"
	case Av:
		return "Av"
	case Elul:
		return "Elul"
	}
	panic(fmt.Sprintf("No known Hebrew month %d", m))
}

type HebrewDate struct {
	Y HebrewYear
	D int
	M HebrewMonth
}

func NewHebrewDate(t time.Time) HebrewDate {
	return ToHebrewDate(t)
}

// String implements stringer.String.
func (h HebrewDate) String() string {
	return fmt.Sprintf("%d %s %s", h.D, h.M, h.Y)
}

// height gives the "height" of the date, per Conway.
func (h HebrewDate) height() int {
	return h.D + int(h.M)
}

// equals compares two HebrewDates for equality
func (h HebrewDate) Equals(d HebrewDate) bool {
	return h.D == d.D && h.M == d.M && h.Y.Y == d.Y.Y
}
