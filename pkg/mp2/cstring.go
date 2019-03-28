package mp2

func nullTerminatedString(nullTerminatedString []byte) string {
	return string(nullTerminatedString[:nullTerminatedStringLen(nullTerminatedString)])
}

func nullTerminatedStringLen(nullTerminatedString []byte) int {
	for i := 0; i < len(nullTerminatedString); i++ {
		if nullTerminatedString[i] == 0 {
			return i
		}
	}
	return len(nullTerminatedString)
}
