// Copyright © 2018 mritd <mritd1234@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package util

import (
	"github.com/mritd/idgen/metadata"
	"math/rand"
	"time"
)

// 随机单个中文字符
func GenOneChineseChars() string {
	return string(rune(RandInt(19968, 40869)))
}

// 指定长度随机中文字符
func GenFixedLengthChineseChars(length int) string {
	strRune := make([]rune, length)
	for i := range strRune {
		strRune[i] = rune(RandInt(19968, 40869))
	}
	return string(strRune)
}

// 指定范围随机中文字符
func GenRandomLengthChineseChars(start, end int) string {
	return GenFixedLengthChineseChars(RandInt(start, end))
}

// 随机单个复杂中文字符
func OneOddChar() string {
	rand.Seed(time.Now().UnixNano())
	return string([]rune(metadata.ODD_CHINESE_CHARS)[rand.Intn(len([]rune(metadata.ODD_CHINESE_CHARS)))])
}

// 随机英文字符
func RandStr(len int) string {
	rand.Seed(time.Now().UnixNano())
	data := make([]byte, len)
	var num int
	for i := 0; i < len; i++ {
		num = rand.Intn(57) + 65
		for {
			if num > 90 && num < 97 {
				num = rand.Intn(57) + 65
			} else {
				break
			}
		}
		data[i] = byte(num)
	}
	return string(data)
}

// 指定范围随机 int
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// 指定范围随机 int64
func RandInt64(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}