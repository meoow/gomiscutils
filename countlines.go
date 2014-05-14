package gomiscutils

import "log"
import "io"
import "bytes"
import "os"
//import "fmt"

var textBufSize uint = 1024 * 1024 // read 1 mega bytes
var err error
var nbytes int
//var sep rune = '\n'
var sep []byte = []byte{'\n'}

func Countlines(f *os.File) uint64 {
	defer f.Seek(0,0)
	_, err = f.Seek(0,0)
	if err != nil { log.Fatal(err) }

	var textBuf []byte = make([]byte,textBufSize,textBufSize)
	var count int
	var loopcount uint64
	var lastNBytes int

	for {
		nbytes, err = f.Read(textBuf)
		if err != io.EOF {
			loopcount++
			count += bytes.Count(textBuf[0:nbytes], sep)
			lastNBytes = nbytes
		} else if err == io.EOF {
			break
		} else {
			log.Fatal(err)
		}
	}

	if loopcount > 0 {
		if textBuf[lastNBytes-1] == '\n' {
			return uint64(count)
		} else { return uint64(count+1) }
	}

	return 0
}

//func runeCount(r []rune, c rune) uint64 {
//	var count uint64
//	for _,v:= range r {
//		if v == c { count++ }
//	}
//	return count
//}
