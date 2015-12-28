package bytesize

import "testing"

type FmtTest struct {
	Value     int64
	Base      int
	Precision int
	Result    string
}

var FmtTests = []FmtTest{
	FmtTest{1000, 10, 0, "1KB"},
	FmtTest{1024, 10, 2, "1.02KB"},
	FmtTest{1024, 2, 1, "1.0KiB"},
	FmtTest{56000, 10, 2, "56.00KB"},        // 1 sec of 56K modem
	FmtTest{1024000, 2, 2, "1000.00KiB"},    // A "1MB" floppy disk is 1000KiB
	FmtTest{52428800, 2, 0, "50MiB"},        // A "50MB" file on a CD is 50MiB
	FmtTest{1300000000, 2, 2, "1.21GiB"},    // A 1.3GB file on a DVD is 1.3GB
	FmtTest{300000000000, 10, 1, "300.0GB"}, // A 300GB hard disk is 300GB
	FmtTest{8589934592, 2, 2, "8.00GiB"},    // "8GB" of RAM is 8GiB
}

type ParseTest struct {
	Value  string
	Result int64
}

var ParseTests = []ParseTest{
	ParseTest{"1KB", 1000},
	ParseTest{"1KiB", 1024},
	ParseTest{"6GiB", 6 * 1024 * 1024 * 1024},
}

func TestFormatBytes(t *testing.T) {
	for _, test := range FmtTests {
		ts := FormatBytes(test.Value, test.Base, test.Precision)
		if ts != test.Result {
			t.Errorf("%d base %d precision %d should be %s, was %s", test.Value,
				test.Base, test.Precision, test.Result, ts)
		}
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
	for _, test := range ParseTests {
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
