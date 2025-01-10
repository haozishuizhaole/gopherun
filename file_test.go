/*
 *    Copyright (c) 2025 Chenzhihao
 *    Gopherun is licensed under Mulan PSL v2.
 *    You can use this software according to the terms and conditions of the Mulan PSL v2.
 *    You may obtain a copy of Mulan PSL v2 at:
 *             http://license.coscl.org.cn/MulanPSL2
 *    THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 *    See the Mulan PSL v2 for more details.
 */

package gopherun

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"os"
	"path/filepath"
	"testing"
)

type FileTest struct {
	BaseTest
}

func TestFileTest(t *testing.T) {
	suite.Run(t, new(FileTest))
}

func (f *FileTest) TestGopherunFile_MkdirAll() {
	path := filepath.Join(f.tempDir, "/test1/test2/test3/test4")

	err := File.MkdirAll(path)
	require.True(f.T(), err == nil, "MkdirAll err, ", path)

	require.DirExistsf(f.T(), path, "path not exist")
}

func (f *FileTest) TestGopherunFile_MkdirAllWithMode() {
	path := filepath.Join(f.tempDir, "/test1/test2/test3/test4")
	err := File.MkdirAllWithMode(path, 0755)
	require.True(f.T(), err == nil, "MkdirAllWithMode err, ", path, err)

	require.DirExistsf(f.T(), path, "path not exist")

	stat, err := os.Stat(path)
	require.True(f.T(), err == nil, "Stat err, ", path)
	require.True(f.T(), stat != nil, "Stat is nil")
	require.True(f.T(), stat.Mode().Perm() == 0755, "Mode should be 0755", stat.Mode())
}

func (f *FileTest) TestGopherunFile_Remove() {
	path := filepath.Join(f.tempDir, "/test1/test2/test3/test4")
	err := File.MkdirAllWithMode(path, 0755)
	require.True(f.T(), err == nil, "MkdirAllWithMode err, ", path, err)

	err = File.Remove(path)
	require.True(f.T(), err == nil, "Remove err, ", path)
	require.NoDirExistsf(f.T(), path, "path remove failed, has exist. %s", path)

	path = filepath.Join(f.tempDir, "/test1/test2/test3")
	_, err = os.Stat(path)
	require.Truef(f.T(), err == nil || os.IsExist(err), "parent dir not exist, %s", path)
}

func (f *FileTest) TestGopherunFile_RemoveAll() {
	path := filepath.Join(f.tempDir, "/test1/test2/test3/test4")
	err := File.MkdirAllWithMode(path, 0755)
	require.True(f.T(), err == nil, "MkdirAllWithMode err, ", path, err)

	path = filepath.Join(f.tempDir, "/test1")
	err = File.RemoveAll(path)
	require.Truef(f.T(), err == nil, "RemoveAll err, path = %s", path)
}

func (f *FileTest) TestGopherunFile_IsExists() {
	path := filepath.Join(f.tempDir, "/test1/test2/test3/test4")

	exists := File.IsExists(path)
	require.True(f.T(), !exists, "IsExists check failed")

	err := File.MkdirAllWithMode(path, 0755)
	require.Truef(f.T(), err == nil, "MkdirAllWithMode err, %s, %v", path, err)
	exists = File.IsExists(path)
	require.True(f.T(), exists, "IsExists check failed")
}

func (f *FileTest) TestGopherunFile_GetAbsolutePath() {
	err := os.Chdir(f.tempDir)
	require.Truef(f.T(), err == nil, "Chdir err, %s", f.tempDir)

	absolutePath, err := File.GetAbsolutePath(".")
	require.Truef(f.T(), err == nil, "GetAbsolutePath err, %s, %v", f.tempDir, err)

	// 某些操作系统（如：Mac）中，/var 是 /private/var 的符号链接，直接对比会导致case失败，实际上目录是一样的
	// 这里使用 filepath.EvalSymlinks 解析符号链接，确保路径一致。
	target, _ := filepath.EvalSymlinks(f.tempDir)
	source, _ := filepath.EvalSymlinks(absolutePath)

	require.Truef(f.T(), source != "" && target != "" && source == target, "GetAbsolutePath failed, %s, %s", source, target)
}

func (f *FileTest) TestGopherunFile_GetPwd() {
	err := os.Chdir(f.tempDir)
	require.Truef(f.T(), err == nil, "Chdir err, %s", f.tempDir)

	pwd, err := File.GetPwd()
	require.Truef(f.T(), err == nil, "GetPwd err, %s, %v", f.tempDir, err)

	// 某些操作系统（如：Mac）中，/var 是 /private/var 的符号链接，直接对比会导致case失败，实际上目录是一样的
	// 这里使用 filepath.EvalSymlinks 解析符号链接，确保路径一致。
	target, _ := filepath.EvalSymlinks(f.tempDir)
	source, _ := filepath.EvalSymlinks(pwd)

	require.Truef(f.T(), source != "" && target != "" && source == target, "GetAbsolutePath failed, %s, %s", source, target)
}
