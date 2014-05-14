package gomiscutils

import "path/filepath"
import "strings"

func SplitExt(s string) (name, ext string) {
	fname := filepath.Base(s)
	ext = filepath.Ext(fname)
	name = strings.TrimRight(fname, ext)
	return
}
