// Package human provides a function for formatting and parsing quantities of
// data as human-readable values such as 1MB or 2.2GB.
package bytesize

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

// SI (base 10) units
// See http://physics.nist.gov/cuu/Units/prefixes.html
const (
	KB float64 = 1000
	MB         = 1000 * KB
	GB         = 1000 * MB
	TB         = 1000 * GB
	PB         = 1000 * TB
	EB         = 1000 * PB
	ZB         = 1000 * EB
	YB         = 1000 * ZB
)

// IEC (base 2) units
// See http://physics.nist.gov/cuu/Units/binary.html
// and https://en.wikipedia.org/wiki/Binary_prefix
const (
	_           = iota
	KiB float64 = 1 << (10 * iota)
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

var multipliers = map[string]float64{
	"KB": KB, "MB": MB, "GB": GB, "TB": TB,
	"PB": PB, "EB": EB, "ZB": ZB, "YB": YB,
	"KiB": KiB, "MiB": MiB, "GiB": GiB, "TiB": TiB,
	"PiB": PiB, "EiB": EiB, "ZiB": ZiB, "YiB": YiB,
}

// FormatBytes formats the specified number of bytes in human-readable format,
// in either base 2 (IEC) or base 10 (SI) units as specified by the base
// parameter, with the specified number of digits of precision.
//
// e.g.
//     FormatBytes(1024*1024, 2, 2) => "1.00MiB"
//     FormatBytes(2000000000, 10, 2) => "2.00GB"
//
// If an invalid base or precision is given, a Sprintf-style error string
// is returned.
func FormatBytes(bytes int64, base int, prec int) string {
	if base != 10 && base != 2 {
		return "%!(BADBASE)"
	}
	b := float64(bytes)
	if base == 10 {
		switch {
		case b >= YB:
			return fmt.Sprintf("%.*fYB", prec, float64(b/YB))
		case b >= ZB:
			return fmt.Sprintf("%.*fZB", prec, float64(b/ZB))
		case b >= EB:
			return fmt.Sprintf("%.*fEB", prec, float64(b/EB))
		case b >= PB:
			return fmt.Sprintf("%.*fPB", prec, float64(b/PB))
		case b >= TB:
			return fmt.Sprintf("%.*fTB", prec, float64(b/TB))
		case b >= GB:
			return fmt.Sprintf("%.*fGB", prec, float64(b/GB))
		case b >= MB:
			return fmt.Sprintf("%.*fMB", prec, float64(b/MB))
		case b >= KB:
			return fmt.Sprintf("%.*fKB", prec, float64(b/KB))
		}
		return fmt.Sprintf("%.*fB", prec, float64(b))
	}
	switch {
	case b >= YiB:
		return fmt.Sprintf("%.*fYiB", prec, float64(b/YiB))
	case b >= ZiB:
		return fmt.Sprintf("%.*fZiB", prec, float64(b/ZiB))
	case b >= EiB:
		return fmt.Sprintf("%.*fEiB", prec, float64(b/EiB))
	case b >= PiB:
		return fmt.Sprintf("%.*fPiB", prec, float64(b/PiB))
	case b >= TiB:
		return fmt.Sprintf("%.*fTiB", prec, float64(b/TiB))
	case b >= GiB:
		return fmt.Sprintf("%.*fGiB", prec, float64(b/GiB))
	case b >= MiB:
		return fmt.Sprintf("%.*fMiB", prec, float64(b/MiB))
	case b >= KiB:
		return fmt.Sprintf("%.*fKiB", prec, float64(b/KiB))
	}
	return fmt.Sprintf("%.*fB", prec, float64(b))
}

func split(s string) (string, string) {
	s1 := make([]rune, 0, len(s))
	s2 := make([]rune, 0, len(s))
	for _, r := range s {
		if (unicode.IsDigit(r) || r == '.') && len(s2) == 0 {
			s1 = append(s1, r)
		} else {
			if r != ' ' {
				s2 = append(s2, r)
			}
		}
	}
	return string(s1), string(s2)
}

// ParseBytesFloat parses a human-readable quantity of bytes, and returns the
// raw number of bytes as a float64.
//
// Whitespace is allowed between the number and the units. The 'B' for bytes
// is required to be upper case; 'b' in lower case would be bits.
func ParseBytesFloat(s string) (float64, error) {
	b, u := split(s)
	fb, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return fb, err
	}
	if u == "" {
		return fb, fmt.Errorf("no units found")
	}
	m, ok := multipliers[u]
	if !ok {
		return fb, fmt.Errorf("unrecognized units %s", u)
	}
	return m * fb, nil
}

// ParseBytesFloat parses a human-readable quantity of bytes, and returns the
// raw number of bytes as an int64. If the value is too large for an int64,
// an error value is returned.
//
// Whitespace is allowed between the number and the units. The 'B' for bytes
// is required to be upper case; 'b' in lower case would be bits.
func ParseBytes(s string) (int64, error) {
	fv, err := ParseBytesFloat(s)
	if err != nil {
		return int64(fv), err
	}
	if fv > math.MaxInt64 {
		return 0, fmt.Errorf("value too large for int64")
	}
	return int64(fv), nil
}
