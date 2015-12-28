package bytesize

import "testing"

type TestCase struct {
	Value     int64
	Base      int
	Precision int
	Result    string
}

var Tests = []TestCase{
	TestCase{1000, 10, 0, "1KB"},
	TestCase{1024, 10, 2, "1.02KB"},
	TestCase{1024, 2, 1, "1.0KiB"},
	TestCase{56000, 10, 2, "56.00KB"},        // 1 sec of 56K modem
	TestCase{1024000, 2, 2, "1000.00KiB"},    // A "1MB" floppy disk is 1000KiB
	TestCase{52428800, 2, 0, "50MiB"},        // A "50MB" file on a CD is 50MiB
	TestCase{1300000000, 2, 2, "1.21GiB"},    // A 1.3GB file on a DVD is 1.3GB
	TestCase{300000000000, 10, 1, "300.0GB"}, // A 300GB hard disk is 300GB
	TestCase{8589934592, 2, 2, "8.00GiB"},    // "8GB" of RAM is 8GiB
}

func TestFormatBytes(t *testing.T) {
	for _, test := range Tests {
		ts := FormatBytes(test.Value, test.Base, test.Precision)
		if ts != test.Result {
			t.Errorf("%d base %d precision %d should be %s, was %s", test.Value,
				test.Base, test.Precision, test.Result, ts)
		}
	}
}
