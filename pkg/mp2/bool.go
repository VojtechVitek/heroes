package mp2

type Bool uint8

func (v Bool) String() string {
	if v > 0 {
		return "true"
	}
	return "false"
}

func (v Bool) Bool() bool {
	if v > 0 {
		return true
	}
	return false
}
