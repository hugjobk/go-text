package text_test

import (
	"testing"
	"time"

	"github.com/hugjobk/go-text"
)

func TestUnmarshalString(t *testing.T) {
	var v string
	if err := text.Unmarshal("abc", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalBool(t *testing.T) {
	var v bool
	if err := text.Unmarshal("true", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalInt(t *testing.T) {
	var v int
	if err := text.Unmarshal("123", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalInt8(t *testing.T) {
	var v int8
	if err := text.Unmarshal("123", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalInt16(t *testing.T) {
	var v int16
	if err := text.Unmarshal("123", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalInt32(t *testing.T) {
	var v int32
	if err := text.Unmarshal("123", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalInt64(t *testing.T) {
	var v int64
	if err := text.Unmarshal("123", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalUint(t *testing.T) {
	var v uint
	if err := text.Unmarshal("123", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalUint8(t *testing.T) {
	var v uint8
	if err := text.Unmarshal("123", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalUint16(t *testing.T) {
	var v uint16
	if err := text.Unmarshal("123", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalUint32(t *testing.T) {
	var v uint32
	if err := text.Unmarshal("123", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalUint64(t *testing.T) {
	var v uint64
	if err := text.Unmarshal("123", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalFloat32(t *testing.T) {
	var v float32
	if err := text.Unmarshal("1.23", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalFloat64(t *testing.T) {
	var v float64
	if err := text.Unmarshal("1.23", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalComplex64(t *testing.T) {
	var v complex64
	if err := text.Unmarshal("10+11i", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalComplex128(t *testing.T) {
	var v complex64
	if err := text.Unmarshal("10+11i", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalDuration(t *testing.T) {
	var v time.Duration
	if err := text.Unmarshal("1h20m30s", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalSlice(t *testing.T) {
	var v []int
	if err := text.Unmarshal("1,2,3", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestUnmarshalMap(t *testing.T) {
	var v map[int]string
	if err := text.Unmarshal("1:a,2:b,3:c", &v); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}
