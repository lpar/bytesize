// Package human provides a function for formatting quantities of data as
// human-readable values such as 1MB or 2.2GB.
package bytesize

import "fmt"

// SI (base 10) units
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
// See https://en.wikipedia.org/wiki/Binary_prefix
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
			return fmt.Sprintf("%.*fYB", prec, b/YB)
		case b >= ZB:
			return fmt.Sprintf("%.*fZB", prec, b/ZB)
		case b >= EB:
			return fmt.Sprintf("%.*fEB", prec, b/EB)
		case b >= PB:
			return fmt.Sprintf("%.*fPB", prec, b/PB)
		case b >= TB:
			return fmt.Sprintf("%.*fTB", prec, b/TB)
		case b >= GB:
			return fmt.Sprintf("%.*fGB", prec, b/GB)
		case b >= MB:
			return fmt.Sprintf("%.*fMB", prec, b/MB)
		case b >= KB:
			return fmt.Sprintf("%.*fKB", prec, b/KB)
		}
		return fmt.Sprintf("%.*fB", prec, b)
	}
	switch {
	case b >= YiB:
		return fmt.Sprintf("%.*fYiB", prec, b/YiB)
	case b >= ZiB:
		return fmt.Sprintf("%.*fZiB", prec, b/ZiB)
	case b >= EiB:
		return fmt.Sprintf("%.*fEiB", prec, b/EiB)
	case b >= PiB:
		return fmt.Sprintf("%.*fPiB", prec, b/PiB)
	case b >= TiB:
		return fmt.Sprintf("%.*fTiB", prec, b/TiB)
	case b >= GiB:
		return fmt.Sprintf("%.*fGiB", prec, b/GiB)
	case b >= MiB:
		return fmt.Sprintf("%.*fMiB", prec, b/MiB)
	case b >= KiB:
		return fmt.Sprintf("%.*fKiB", prec, b/KiB)
	}
	return fmt.Sprintf("%.*fB", prec, b)
}
