package wpcom

import (
	"bytes"
	"strconv"
)

func hack(v interface{}) *softTypeHack {
	h := new(softTypeHack)
	h.value = v
	return h
}

type softTypeHack struct {
	value interface{}
}

func (h *softTypeHack) int() int64 {
	if h.value == nil {
		return 0
	}
	switch v := h.value.(type) {
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case string:
		f, e := strconv.ParseFloat(v, 64)
		if e != nil {
			return 0
		}
		frep := strconv.FormatFloat(f, 'f', 0, 64)
		i, e := strconv.ParseInt(frep, 10, 64)
		if e != nil {
			return 0
		}
		return i
	default:
		return 0
	}
}

func (h *softTypeHack) bool() bool {
	if h.value == nil {
		return false
	}
	switch v := h.value.(type) {
	case bool:
		return v
	case string:
		if v != "" {
			return true
		} else {
			return false
		}
	case []byte:
		if bytes.Equal(v, []byte("")) {
			return false
		} else {
			return true
		}
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		if v != 0 {
			return true
		} else {
			return false
		}
	}
	return false
}
