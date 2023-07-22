package wr

import "io"

func CustomWriter(wr io.Writer, args ...interface{}) {
	for _, arg := range args {
		if cur, ok := arg.(string); ok {
			wr.Write([]byte(cur))
		}
	}
}
