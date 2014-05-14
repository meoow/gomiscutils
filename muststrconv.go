package gomiscutils

import "strconv"
import "log"
import "os"

var logger = log.New(os.Stderr, "", 0)

func Die(err error) {
	if err != nil {
		logger.Fatal(err)
	}
}

func MustParseBool(s string) (value bool) {
	value, err := strconv.ParseBool(s)
	Die(err)
	return
}

func MustParseFloat(s string, bitSize int) (f float64) {
	f, err := strconv.ParseFloat(s, bitSize)
	Die(err)
	return
}

func MustParseInt(s string, base int, bitSize int) (i int64) {
	i, err := strconv.ParseInt(s, base, bitSize)
	Die(err)
	return
}

func MustParseUint(s string, base int, bitSize int) (u uint64) {
	u, err := strconv.ParseUint(s, base, bitSize)
	Die(err)
	return
}
