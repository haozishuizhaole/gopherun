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
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"io"
	"math/big"
	"testing"
)

type GopherunRandomTest struct {
	BaseTest
}

func TestGopherunRandomTest(t *testing.T) {
	suite.Run(t, new(GopherunRandomTest))
}

func (i *GopherunRandomTest) TestGopherunRandom_RandomString_case1() {
	defer func() {
		panicErr := recover()
		require.True(i.T(), panicErr != nil)
	}()
	randomString := Random.RandomString(1<<32, 16)
	require.True(i.T(), randomString == "")
}

func (i *GopherunRandomTest) TestGopherunRandom_RandomString_case2() {
	defer func() {
		panicErr := recover()
		require.True(i.T(), panicErr != nil)
	}()
	randomString := Random.RandomString(0, 0)
	require.True(i.T(), randomString == "")

}

func (i *GopherunRandomTest) TestGopherunRandom_RandomString_case3() {
	randomString := Random.RandomString(0, 16)
	require.True(i.T(), randomString != "")
	i.T().Logf("random string: %s", randomString)
}

func (i *GopherunRandomTest) TestGopherunRandom_RandomString_case4() {
	// mock
	mockrandInt := gomonkey.ApplyFunc(rand.Int, func(rand io.Reader, max *big.Int) (n *big.Int, err error) {
		return nil, errors.New("random error")
	})
	defer mockrandInt.Reset()

	// run
	randomString := Random.RandomString(0, 16)

	// assert
	require.True(i.T(), randomString != "")
	i.T().Logf("random string: %s", randomString)
}

func (i *GopherunRandomTest) TestGopherunRandom_RandomStringWithNumberAndLetter() {
	letter := Random.RandomStringWithNumberAndLetter(16)
	require.Regexp(i.T(), "^[a-zA-Z]+$", letter)
}

func (i *GopherunRandomTest) TestGopherunRandom_RandomInt64_case1() {
	randomInt64 := Random.RandomInt64(10, 5)
	require.True(i.T(), randomInt64 <= 10 && randomInt64 >= 5)

	randomInt64 = Random.RandomInt64(0, 1)
	require.True(i.T(), randomInt64 >= 0 && randomInt64 <= 1)

	randomInt64 = Random.RandomInt64(-1, 1)
	require.True(i.T(), randomInt64 >= -1 && randomInt64 <= 1)
}

func (i *GopherunRandomTest) TestGopherunRandom_RandomInt64_case2() {
	// mock
	mockrandInt := gomonkey.ApplyFunc(rand.Int, func(rand io.Reader, max *big.Int) (*big.Int, error) {
		return nil, errors.New("random error")
	})
	defer mockrandInt.Reset()

	// run
	randomInt64 := Random.RandomInt64(0, 1)

	// assert
	i.T().Logf("random int64: %v", randomInt64)
}
