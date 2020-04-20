//basenames removes directory components and a .suffix
// e.g, a => a, a.go => a, a/b/c.go => c.go, a/b.c.go => b.c

package basename

import "strings"

func basename(s string)string{
	//Discard the last '/' and everything before 
	for i := len(s) - 1; i >= 0; i--{
		if s[i] == '/'{
			s = s[i+1:]
			break
		}
	}
	// preserve everything before the last '.'
	for i := len(s) - 1; i >= 0; i--{
		if s[i] == '.'{
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string{
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0{
		s = s[:dot]
	}
	return s
}
