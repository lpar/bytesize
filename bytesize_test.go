package bytesize

import (
	"strconv"
	"strings"
	"testing"
)

func TestFormatBytes(t *testing.T) {
	testCases := []struct {
		value     string
		base      int
		precision int
		result    string
	}{
		{"999", 10, 0, "999B"},
		{"999", 2, 0, "999B"},
		{"1000", 10, 0, "1KB"},
		{"1024", 10, 2, "1.02KB"},
		{"1024", 2, 1, "1.0KiB"},
		{"56000", 10, 2, "56.00KB"},     // 1 sec of 56K modem
		{"1024000", 2, 2, "1000.00KiB"}, // A "1MB" floppy disk is 1000KiB
		{"52428800", 2, 0, "50MiB"},     // A "50MB" file on a CD is 50MiB
		{"52428800", 10, 0, "52MB"},
		{"1300000000", 2, 2, "1.21GiB"},    // A 1.3GB file on a DVD is 1.3GB
		{"8589934592", 2, 2, "8.00GiB"},    // "8GB" of RAM is 8GiB
		{"300000000000", 10, 1, "300.0GB"}, // A 300GB hard disk is 300GB

		{"12345678920000", 10, 2, "12.35TB"}, // "8GB" of RAM is 8GiB
		{"12345678920000", 2, 2, "11.23TiB"}, // "8GB" of RAM is 8GiB

		{"30000000000000000", 10, 0, "30PB"},    // A year of Large Hadron Collider data
		{"30000000000000000", 2, 2, "26.65PiB"}, // A year of Large Hadron Collider data

		{"9223372036854775807", 10, 2, "9.22EB"}, // MaxInt64
		{"9223372036854775807", 2, 2, "8.00EiB"}, // MaxInt64
	}
	for _, test := range testCases {
		v64, err := strconv.ParseInt(test.value, 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		ts := FormatBytes(v64, test.base, test.precision)
		if ts != test.result {
			t.Errorf("got %s, expected %s (%d in base %d precision %d)", ts, test.result, v64,
				test.base, test.precision)
		}
	}
	e := FormatBytes(1, 3, 1)
	if !strings.HasPrefix(e, "%!") {
		t.Errorf("got err=nil expected error")
	}
	e = FormatBytes(1, 10, -1)
	if !strings.HasPrefix(e, "%!") {
		t.Errorf("got err=nil expected error")
	}
}

func TestSplit(t *testing.T) {
	x, y := split("10MB")
	if x != "10" || y != "MB" {
		t.Errorf("failed to split 10MB")
	}
	x, y = split("5 GiB")
	if x != "5" || y != "GiB" {
		t.Errorf("failed to split 5 GiB")
	}
}

func TestParse(t *testing.T) {
	testCases := []struct {
		Value  string
		Result int64
	}{
		{"1KB", 1000},
		{"1KiB", 1024},
		{"1.5GiB", 1536 * 1024 * 1024},
		{"6GiB", 6 * 1024 * 1024 * 1024},
		{"6 GiB", 6 * 1024 * 1024 * 1024},
	}
	for _, test := range testCases {
		x, err := ParseBytes(test.Value)
		if err != nil {
			t.Errorf("error: %v", err)
		}
		if x != test.Result {
			t.Errorf("failed to parse %s correctly", test.Value)
		}
	}
}

func TestErrors(t *testing.T) {
	_, err := ParseBytes("1ZB")
	if err == nil {
		t.Errorf("failed to flag range error")
	}
}
