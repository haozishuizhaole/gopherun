/*
 *    Copyright (c) 2025 TootsCharlie
 *    Gopherun is licensed under Mulan PSL v2.
 *    You can use this software according to the terms and conditions of the Mulan PSL v2.
 *    You may obtain a copy of Mulan PSL v2 at:
 *             http://license.coscl.org.cn/MulanPSL2
 *    THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 *    See the Mulan PSL v2 for more details.
 */

package gopherun

import (
	"github.com/stretchr/testify/suite"
)

type BaseTest struct {
	suite.Suite
	tempDir string
}

func (b *BaseTest) SetupSuite() {
	b.T().Log("setting up Suite")
}

func (b *BaseTest) TearDownSuite() {
	b.T().Log("tearing down Suite")
}

func (b *BaseTest) SetupTest() {
	b.T().Logf("setting up Test: %s", b.T().Name())
}

func (b *BaseTest) TearDownTest() {
	b.T().Logf("tearing down Test: %s", b.T().Name())
}
