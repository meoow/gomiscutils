package gomiscutils

import "errors"
//import "fmt"

type Reader interface {
	ReadString(byte) (string, error)
}

type ByteReader interface {
	ReadBytes(byte) ([]byte,error)
}

type ChunkReader interface {
	Read([]byte) (int, error)
	ByteReader
}

// func Readline(ireader fileReader) chan string
// readline := Readline(fileReader)
func Readline(ireader Reader) chan string {
	readline := make(chan string, 1)
	go func() {
		defer close(readline)
		for {
			line, err := ireader.ReadString('\n')
			if err != nil && line == "" {
				break
			}
			readline <- line
		}
	}()
	return readline
}

func ReadLineByte(ireader ByteReader) chan []byte {
	readline := make(chan []byte, 1)
	go func() {
		defer close(readline)
		for {
			line, err := ireader.ReadBytes('\n')
			if err != nil { break }
			readline <- line
		}
	}()
	return readline
}

func ReadChunkLineByte(ireader ChunkReader, byteSize int) (chan []byte, error){
	if byteSize == 0 {
		return nil, errors.New("byteSize must be greater than 0")
	}
	rclb := make(chan []byte, 1)
	go func() {
		defer close(rclb)
		for {
			buf := make([]byte, byteSize, byteSize + byteSize/10)
			n, e := ireader.Read(buf)
			//fmt.Println(n)
			if e!=nil {
				if n == 0 { break }
				if buf[n-1] != '\n' {
					rest, _ := ireader.ReadBytes('\n')
					rclb <- append(buf,rest...)
				} else {
					rclb <- buf
				}
			} else {
				if n == 0 { break }
				rclb <- buf[:n]
			}
		}
	}()
	return rclb, nil
}

func ReadChunkLine(ireader ChunkReader, byteSize int) (chan string,error) {
	rclb, err := ReadChunkLineByte(ireader, byteSize)
	if err != nil {
		return nil,err
	}
	rcl := make(chan string,1)
	go func() {
		defer close(rcl)
		for chunk := range rclb {
			rcl <- string(chunk)
		}
	}()
	return rcl, nil
}


