package utl

import "strings"

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// GetIP reads an IP from a string (this doesn't look to me like it is
// correctly written)
func GetIP(s string) (ip, port string) {
	ip = reverse(strings.Join(strings.Split(reverse(s), ":")[1:], ":"))
	port = reverse(strings.Split(reverse(s), ":")[0])
	return
}
