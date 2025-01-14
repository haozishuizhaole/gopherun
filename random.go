/*
 * Copyright (c) 2025 TootsCharlie
 * Gopherun is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 */

package gopherun

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// CharsetFlag 定义字符集标志类型
type CharsetFlag uint

// 定义字符集常量
const (
	CharsetLowercase CharsetFlag = 1 << iota // 小写字母
	CharsetUppercase                         // 大写字母
	CharsetNumbers                           // 数字
	CharsetSymbols                           // 符号

	CharsetLetter = CharsetLowercase + CharsetUppercase                                   // 大小写字母
	CharsetAll    = CharsetLowercase | CharsetUppercase | CharsetNumbers | CharsetSymbols // 全部字符集
)

// 定义字符集内容
var _charsets = map[CharsetFlag]string{
	CharsetLowercase: "abcdefghijklmnopqrstuvwxyz",
	CharsetUppercase: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	CharsetNumbers:   "0123456789",
	CharsetSymbols:   "!@#$%^&*()-_=+[]{}|;:',.<>?/`~",
}

// RandomString 根据指定的字符集和长度生成随机字符串
func (i GopherunRandom) RandomString(charsetsFlag CharsetFlag, length uint) (string, error) {
	if length == 0 {
		return "", errors.New("length must be greater than 0")
	}

	// 如果 charsetsFlag 为 0，使用所有字符集
	if charsetsFlag == 0 {
		charsetsFlag = CharsetAll
	}

	// 构建字符集池
	pool := ""
	for flag, chars := range _charsets {
		if charsetsFlag&flag != 0 {
			pool += chars
		}
	}

	poolLen := int64(len(pool))

	if poolLen == 0 {
		return "", errors.New("invalid charset configuration")
	}

	// 生成随机字符串
	result := make([]byte, length)
	for i := uint(0); i < length; i++ {
		idx, err := rand.Int(rand.Reader, big.NewInt(poolLen))
		if err != nil {
			return "", err
		}
		result[i] = pool[idx.Int64()]
	}

	return string(result), nil
}

// RandomStringWithNumberAndLetter 生成指定长度的字母+数字随机字符串
func (i GopherunRandom) RandomStringWithNumberAndLetter(length uint) (string, error) {
	return i.RandomString(CharsetLetter, length)
}

// RandomInt64 生成一个包含 min 和 max 边界的更安全的随机整数
func (i GopherunRandom) RandomInt64(min, max int64) (int64, error) {
	// 确保 min 小于 max
	if min > max {
		min, max = max, min
	}

	// 计算范围
	rangeValue := max - min + 1

	// 使用 crypto/rand 生成一个范围内的随机数
	num, err := rand.Int(rand.Reader, big.NewInt(rangeValue))
	if err != nil {
		return 0, err
	}

	return min + num.Int64(), nil
}
