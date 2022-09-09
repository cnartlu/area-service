package strings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Pad(t *testing.T) {
	var (
		s1 string = "1"

		i1 int   = 1
		i2 int8  = 1
		i3 int16 = 1
		i4 int32 = 1
		i5 int64 = 1

		ui1 uint   = 1
		ui2 uint8  = 1
		ui3 uint16 = 1
		ui4 uint32 = 1
		ui5 uint64 = 1
	)
	for j := 0; j < 1; j++ {
		var (
			padLen = j + 3
			padStr = "0"
			plen   = len(padStr)
			modv   = (padLen - plen) / 2
			ls     = strings.Repeat(padStr, padLen-plen) + s1
			rs     = s1 + strings.Repeat(padStr, padLen-plen)
			bs     = strings.Repeat(padStr, modv) + s1 + strings.Repeat(padStr, padLen-modv-1)
		)
		// 补全左部字符串
		assert.Equal(t, Pad(i1, padLen, padStr, STR_PAD_LEFT), ls)
		assert.Equal(t, Pad(i2, padLen, padStr, STR_PAD_LEFT), ls)
		assert.Equal(t, Pad(i3, padLen, padStr, STR_PAD_LEFT), ls)
		assert.Equal(t, Pad(i4, padLen, padStr, STR_PAD_LEFT), ls)
		assert.Equal(t, Pad(i5, padLen, padStr, STR_PAD_LEFT), ls)

		assert.Equal(t, Pad(ui1, padLen, padStr, STR_PAD_LEFT), ls)
		assert.Equal(t, Pad(ui2, padLen, padStr, STR_PAD_LEFT), ls)
		assert.Equal(t, Pad(ui3, padLen, padStr, STR_PAD_LEFT), ls)
		assert.Equal(t, Pad(ui4, padLen, padStr, STR_PAD_LEFT), ls)
		assert.Equal(t, Pad(ui5, padLen, padStr, STR_PAD_LEFT), ls)

		assert.Equal(t, Pad(s1, padLen, padStr, STR_PAD_LEFT), ls)
		// 补全右部字符串
		assert.Equal(t, Pad(i1, padLen, padStr, STR_PAD_RIGHT), rs)
		assert.Equal(t, Pad(i2, padLen, padStr, STR_PAD_RIGHT), rs)
		assert.Equal(t, Pad(i3, padLen, padStr, STR_PAD_RIGHT), rs)
		assert.Equal(t, Pad(i4, padLen, padStr, STR_PAD_RIGHT), rs)
		assert.Equal(t, Pad(i5, padLen, padStr, STR_PAD_RIGHT), rs)

		assert.Equal(t, Pad(ui1, padLen, padStr, STR_PAD_RIGHT), rs)
		assert.Equal(t, Pad(ui2, padLen, padStr, STR_PAD_RIGHT), rs)
		assert.Equal(t, Pad(ui3, padLen, padStr, STR_PAD_RIGHT), rs)
		assert.Equal(t, Pad(ui4, padLen, padStr, STR_PAD_RIGHT), rs)
		assert.Equal(t, Pad(ui5, padLen, padStr, STR_PAD_RIGHT), rs)

		assert.Equal(t, Pad(s1, padLen, padStr, STR_PAD_RIGHT), rs)
		// 补全两边字符串
		assert.Equal(t, Pad(i1, padLen, padStr, STR_PAD_BOTH), bs)
		assert.Equal(t, Pad(i2, padLen, padStr, STR_PAD_BOTH), bs)
		assert.Equal(t, Pad(i3, padLen, padStr, STR_PAD_BOTH), bs)
		assert.Equal(t, Pad(i4, padLen, padStr, STR_PAD_BOTH), bs)
		assert.Equal(t, Pad(i5, padLen, padStr, STR_PAD_BOTH), bs)

		assert.Equal(t, Pad(ui1, padLen, padStr, STR_PAD_BOTH), bs)
		assert.Equal(t, Pad(ui2, padLen, padStr, STR_PAD_BOTH), bs)
		assert.Equal(t, Pad(ui3, padLen, padStr, STR_PAD_BOTH), bs)
		assert.Equal(t, Pad(ui4, padLen, padStr, STR_PAD_BOTH), bs)
		assert.Equal(t, Pad(ui5, padLen, padStr, STR_PAD_BOTH), bs)

		assert.Equal(t, Pad(s1, padLen, padStr, STR_PAD_BOTH), bs)
	}
}

func Benchmark_Pad(b *testing.B) {
	var (
		i1 int   = 1
		i2 int8  = 1
		i3 int16 = 1
		i4 int32 = 1
		i5 int64 = 1

		ui1 uint   = 1
		ui2 uint8  = 1
		ui3 uint16 = 1
		ui4 uint32 = 1
		ui5 uint64 = 1

		s1 string = "1"
	)
	for i := 0; i < b.N; i++ {
		for j := 0; j < 2; j++ {
			var (
				padLen = j + 3
				padStr = "0"
				plen   = len(padStr)
				modv   = plen / 2
				ls     = strings.Repeat(padStr, padLen-plen) + s1
				rs     = s1 + strings.Repeat(padStr, padLen-plen)
				bs     = strings.Repeat(padStr, modv) + s1 + strings.Repeat(padStr, padLen-modv)
			)
			// 补全左部字符串
			assert.Equal(b, Pad(i1, padLen, padStr, STR_PAD_LEFT), ls)
			assert.Equal(b, Pad(i2, padLen, padStr, STR_PAD_LEFT), ls)
			assert.Equal(b, Pad(i3, padLen, padStr, STR_PAD_LEFT), ls)
			assert.Equal(b, Pad(i4, padLen, padStr, STR_PAD_LEFT), ls)
			assert.Equal(b, Pad(i5, padLen, padStr, STR_PAD_LEFT), ls)

			assert.Equal(b, Pad(ui1, padLen, padStr, STR_PAD_LEFT), ls)
			assert.Equal(b, Pad(ui2, padLen, padStr, STR_PAD_LEFT), ls)
			assert.Equal(b, Pad(ui3, padLen, padStr, STR_PAD_LEFT), ls)
			assert.Equal(b, Pad(ui4, padLen, padStr, STR_PAD_LEFT), ls)
			assert.Equal(b, Pad(ui5, padLen, padStr, STR_PAD_LEFT), ls)

			assert.Equal(b, Pad(s1, padLen, padStr, STR_PAD_LEFT), ls)
			// 补全右部字符串
			assert.Equal(b, Pad(i1, padLen, padStr, STR_PAD_RIGHT), rs)
			assert.Equal(b, Pad(i2, padLen, padStr, STR_PAD_RIGHT), rs)
			assert.Equal(b, Pad(i3, padLen, padStr, STR_PAD_RIGHT), rs)
			assert.Equal(b, Pad(i4, padLen, padStr, STR_PAD_RIGHT), rs)
			assert.Equal(b, Pad(i5, padLen, padStr, STR_PAD_RIGHT), rs)

			assert.Equal(b, Pad(ui1, padLen, padStr, STR_PAD_RIGHT), rs)
			assert.Equal(b, Pad(ui2, padLen, padStr, STR_PAD_RIGHT), rs)
			assert.Equal(b, Pad(ui3, padLen, padStr, STR_PAD_RIGHT), rs)
			assert.Equal(b, Pad(ui4, padLen, padStr, STR_PAD_RIGHT), rs)
			assert.Equal(b, Pad(ui5, padLen, padStr, STR_PAD_RIGHT), rs)

			assert.Equal(b, Pad(s1, padLen, padStr, STR_PAD_RIGHT), rs)
			// 补全两边字符串
			assert.Equal(b, Pad(i1, padLen, padStr, STR_PAD_BOTH), bs)
			assert.Equal(b, Pad(i2, padLen, padStr, STR_PAD_BOTH), bs)
			assert.Equal(b, Pad(i3, padLen, padStr, STR_PAD_BOTH), bs)
			assert.Equal(b, Pad(i4, padLen, padStr, STR_PAD_BOTH), bs)
			assert.Equal(b, Pad(i5, padLen, padStr, STR_PAD_BOTH), bs)

			assert.Equal(b, Pad(ui1, padLen, padStr, STR_PAD_BOTH), bs)
			assert.Equal(b, Pad(ui2, padLen, padStr, STR_PAD_BOTH), bs)
			assert.Equal(b, Pad(ui3, padLen, padStr, STR_PAD_BOTH), bs)
			assert.Equal(b, Pad(ui4, padLen, padStr, STR_PAD_BOTH), bs)
			assert.Equal(b, Pad(ui5, padLen, padStr, STR_PAD_BOTH), bs)

			assert.Equal(b, Pad(s1, padLen, padStr, STR_PAD_BOTH), bs)
		}
	}
}
