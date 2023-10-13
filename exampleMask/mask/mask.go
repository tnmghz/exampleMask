package mask

// mask маскирует ссылку
func Mask(n int, p []byte, s, m byte) string {
	for i := n; i < len(p); i++ {
		if p[i] == s {
			break
		}
		if p[i] != s {
			p[i] = m
		}
	}
	return string(p)
}
