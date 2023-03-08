package util

func Max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int64) int64 {
	if a > b {
		return b
	} else {
		return a
	}
}

func GetString(s string) *string {
	return &s
}