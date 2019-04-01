package cstring

func String(str []byte) string {
	return string(str[:strLen(str)])
}

func strLen(str []byte) int {
	for i := 0; i < len(str); i++ {
		if str[i] == 0 {
			return i
		}
	}
	return len(str)
}
