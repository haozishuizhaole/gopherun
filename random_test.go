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
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GopherunRandomTest struct {
	BaseTest
}

func TestGopherunRandomTest(t *testing.T) {
	suite.Run(t, new(GopherunRandomTest))
}

func (i *GopherunRandomTest) TestGopherunRandom_RandomString() {
	randomString, err := Random.RandomString(1<<32, 16)
	require.True(i.T(), err != nil)
	require.True(i.T(), randomString == "")

	randomString, err = Random.RandomString(0, 0)
	require.True(i.T(), err != nil)
	require.True(i.T(), randomString == "")

	randomString, err = Random.RandomString(0, 16)
	require.True(i.T(), err == nil)
	require.True(i.T(), randomString != "")
	i.T().Logf("random string: %s", randomString)
}

func (i *GopherunRandomTest) TestGopherunRandom_RandomStringWithNumberAndLetter() {
	letter, err := Random.RandomStringWithNumberAndLetter(16)
	require.True(i.T(), err == nil)
	require.Regexp(i.T(), "^[a-zA-Z]+$", letter)
}

func (i *GopherunRandomTest) TestGopherunRandom_RandomInt64() {
	randomInt64, err := Random.RandomInt64(10, 5)
	require.True(i.T(), err == nil)
	require.True(i.T(), randomInt64 <= 10 && randomInt64 >= 5)

	randomInt64, err = Random.RandomInt64(0, 1)
	require.True(i.T(), err == nil)
	require.True(i.T(), randomInt64 >= 0 && randomInt64 <= 1)

	randomInt64, err = Random.RandomInt64(-1, 1)
	require.True(i.T(), err == nil)
	require.True(i.T(), randomInt64 >= -1 && randomInt64 <= 1)
}
