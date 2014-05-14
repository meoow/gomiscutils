package gomiscutils

func TrimNewLineByte(b []byte) []byte {
	l := len(b)
	if l!=0 && b[l-1] == '\n' {
		l -= 1
		if l!=0 && b[l-1] == '\r' {
			l -= 1
		}
	}
	return b[:l]
}

func TrimRightWhiteSpaceByte(b []byte) []byte {
	l := len(b)
	for {
		if l == 0 { break }
		switch b[l-1] {
		case '\n','\r','\t',' ':
			l -= 1
		default:
			goto RETURN
		}
	}
	RETURN:
	return b[:l]
}

func TrimNewLine(s string) string {
	b := []byte(s)
	b = TrimNewLineByte(b)
	return string(b)
}

func TrimRightWhiteSpace(s string) string {
	b := []byte(s)
	b = TrimRightWhiteSpaceByte(b)
	return string(b)
}


