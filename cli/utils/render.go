package utils

import (
	"fmt"
	"strings"
)

func format(str string, kwargs map[string]string) string {
	args, i := make([]string, len(kwargs)*2), 0
	for k, v := range kwargs {
		args[i] = "{" + k + "}"
		args[i+1] = fmt.Sprint(v)
		i += 2
	}
	return strings.NewReplacer(args...).Replace(str)
}

func Render(str string, kwargs map[string]string) {
	styledString := format(str+"{RESET}", styles)
	fmt.Println(format(styledString, kwargs))
}
