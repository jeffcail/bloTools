package bloTools

import "strings"

func CompactStr(args ...string) string {
	var builder strings.Builder
	for i := range args {
		builder.WriteString(args[i])
	}
	return builder.String()
}
