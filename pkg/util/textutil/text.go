package textutil

import "strconv"

func StringToPointer(value string) *string {
	return &value
}

func StringToUint64(value string, def uint64) uint64 {
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return def
	}

	return uint64(valueInt)
}
