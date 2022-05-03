package merrors

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

func sprint(err error) string {
	str := fmt.Sprintf("%s\n", err)
	skip := 2
	for ; err != nil; err = errors.Unwrap(err) {
		if pc, file, line, ok := runtime.Caller(skip); ok {
			fn := runtime.FuncForPC(pc)
			strs := strings.Split(fn.Name(), ".")
			str += fmt.Sprintf("%s:%s:%v: %v\n", file, strs[len(strs)-1], line, err)
		}
		skip++
	}
	str = strings.TrimRight(str, "\n")
	return str
}
