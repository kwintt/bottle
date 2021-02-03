package bot

import (
	"strconv"
	"strings"
)

func Color(s string) (uint, uint, uint) {
	hex := strings.TrimPrefix(s, "#")
	i, _ := strconv.ParseInt(hex, 16, 64)
	var (
		r = i >> 16
		g = (i >> 8) & 0xFF
		b = i & 0xFF
	)
	return uint(r), uint(g), uint(b)
}
