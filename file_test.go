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
	"errors"
	"github.com/agiledragon/gomonkey/v2"
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

func (f *FileTest) TestGopherunFile_Size() {
	err := os.Chdir(f.tempDir)
	require.Truef(f.T(), err == nil, "Chdir err, %s", f.tempDir)

	err = File.WriteFileSafer("student.txt", []byte("zhangsan"), os.ModePerm)
	require.Truef(f.T(), err == nil, "WriteFileSafer err, %s", f.tempDir)
	require.FileExists(f.T(), "student.txt", "student.txt file not exists")

	size, err := File.Size("student.txt")
	require.Truef(f.T(), err == nil, "WriteFileSafer err, %s", f.tempDir)
	require.True(f.T(), size > 0)
	f.T().Logf("student.txt file size: %d", size)

	size, err = File.Size("student.txt.back")
	require.True(f.T(), err != nil)
}

func (f *FileTest) TestGopherunFile_IsDir_case1() {
	err := os.Chdir(f.tempDir)
	require.Truef(f.T(), err == nil, "Chdir err, %s", f.tempDir)

	isDir := File.IsDir("./student/")
	require.True(f.T(), isDir == false)

	isDir = File.IsDir("./")
	require.True(f.T(), isDir == true)

	err = File.WriteFileSafer("student.txt", []byte("zhangsan"), os.ModePerm)
	require.Truef(f.T(), err == nil, "WriteFileSafer err, %s", f.tempDir)
	require.FileExists(f.T(), "student.txt", "student.txt file not exists")
	isDir = File.IsDir("student.txt")
	require.True(f.T(), isDir == false)

	err = File.MkdirAll("./student/")
	require.Truef(f.T(), err == nil, "WriteFileSafer err, %s", f.tempDir)
	require.DirExists(f.T(), "./student/", "student.txt file not exists")
	isDir = File.IsDir("./student/")
	require.True(f.T(), isDir == true)
}

func (f *FileTest) TestGopherunFile_IsDir_case2() {
	// mock
	mockosLstat := gomonkey.ApplyFunc(os.Lstat, func(name string) (os.FileInfo, error) {
		return nil, errors.New("mock err")
	})
	defer mockosLstat.Reset()

	// run
	isDir := File.IsDir("./")

	// assert
	require.True(f.T(), isDir == false)
}

func (f *FileTest) TestGopherunFile_WriteFileSafer_case1() {
	err := os.Chdir(f.tempDir)
	require.Truef(f.T(), err == nil, "Chdir err, %s", f.tempDir)

	err = File.WriteFileSafer("student.txt", []byte("zhangsan"), os.ModePerm)
	require.Truef(f.T(), err == nil, "WriteFileSafer err, %s", f.tempDir)
	require.FileExists(f.T(), "student.txt", "student.txt file not exists")
}

func (f *FileTest) TestGopherunFile_WriteFileSafer_case2() {
	// mock
	mockosOpenFile := gomonkey.ApplyFunc(os.OpenFile, func(name string, flag int, perm os.FileMode) (*os.File, error) {
		return nil, errors.New("mock err")
	})
	defer mockosOpenFile.Reset()

	// run
	err := os.Chdir(f.tempDir)
	require.Truef(f.T(), err == nil, "Chdir err, %s", f.tempDir)
	err = File.WriteFileSafer("student.txt", []byte("zhangsan"), os.ModePerm)

	// assert
	require.Truef(f.T(), err != nil, "WriteFileSafer err, %s", f.tempDir)
	require.NoFileExists(f.T(), "student.txt")
}

func (f *FileTest) TestGopherunFile_WriteFileSafer_case3() {
	// mock
	mockFile := &os.File{}
	mockWrite := gomonkey.ApplyMethod(mockFile, "Write", func(_ *os.File, b []byte) (n int, err error) {
		return 0, errors.New("mock err")
	})
	defer mockWrite.Reset()

	mockosOpenFile := gomonkey.ApplyFunc(os.OpenFile, func(name string, flag int, perm os.FileMode) (*os.File, error) {
		return mockFile, nil
	})
	defer mockosOpenFile.Reset()

	// run
	err := os.Chdir(f.tempDir)
	require.Truef(f.T(), err == nil, "Chdir err, %s", f.tempDir)
	err = File.WriteFileSafer("student.txt", []byte("zhangsan"), os.ModePerm)

	// assert
	require.Truef(f.T(), err != nil, "WriteFileSafer err, %s", f.tempDir)
	require.NoFileExists(f.T(), "student.txt")
}

func (f *FileTest) TestGopherunFile_WriteFileSafer_case4() {
	// mock
	mockFile := &os.File{}
	mockWrite := gomonkey.ApplyMethod(mockFile, "Write", func(_ *os.File, b []byte) (n int, err error) {
		return 0, nil
	})
	defer mockWrite.Reset()

	mockSync := gomonkey.ApplyMethod(mockFile, "Sync", func(_ *os.File) (err error) {
		return errors.New("mock err")
	})
	defer mockSync.Reset()

	mockosOpenFile := gomonkey.ApplyFunc(os.OpenFile, func(name string, flag int, perm os.FileMode) (*os.File, error) {
		return mockFile, nil
	})
	defer mockosOpenFile.Reset()

	// run
	err := os.Chdir(f.tempDir)
	require.Truef(f.T(), err == nil, "Chdir err, %s", f.tempDir)
	err = File.WriteFileSafer("student.txt", []byte("zhangsan"), os.ModePerm)

	// assert
	require.Truef(f.T(), err != nil, "WriteFileSafer err, %s", f.tempDir)
	require.NoFileExists(f.T(), "student.txt")
}

func (f *FileTest) TestGopherunFile_WriteFileSafer_case5() {
	// mock
	mockFile := &os.File{}
	mockWrite := gomonkey.ApplyMethod(mockFile, "Write", func(_ *os.File, b []byte) (n int, err error) {
		return 0, nil
	})
	defer mockWrite.Reset()

	mockSync := gomonkey.ApplyMethod(mockFile, "Sync", func(_ *os.File) (err error) {
		return nil
	})
	defer mockSync.Reset()

	mockClose := gomonkey.ApplyMethod(mockFile, "Close", func(_ *os.File) (err error) {
		return errors.New("mock err")
	})
	defer mockClose.Reset()

	mockosOpenFile := gomonkey.ApplyFunc(os.OpenFile, func(name string, flag int, perm os.FileMode) (*os.File, error) {
		return mockFile, nil
	})
	defer mockosOpenFile.Reset()

	// run
	err := os.Chdir(f.tempDir)
	require.Truef(f.T(), err == nil, "Chdir err, %s", f.tempDir)
	err = File.WriteFileSafer("student.txt", []byte("zhangsan"), os.ModePerm)

	// assert
	require.Truef(f.T(), err != nil, "WriteFileSafer err, %s", f.tempDir)
	require.NoFileExists(f.T(), "student.txt")
}

func (f *FileTest) TestGopherunFile_WriteFileSafer_case6() {
	// mock
	mockFile := &os.File{}
	mockWrite := gomonkey.ApplyMethod(mockFile, "Write", func(_ *os.File, b []byte) (n int, err error) {
		return 0, nil
	})
	defer mockWrite.Reset()

	mockSync := gomonkey.ApplyMethod(mockFile, "Sync", func(_ *os.File) (err error) {
		return nil
	})
	defer mockSync.Reset()

	mockClose := gomonkey.ApplyMethod(mockFile, "Close", func(_ *os.File) (err error) {
		return nil
	})
	defer mockClose.Reset()

	mockName := gomonkey.ApplyMethod(mockFile, "Name", func(_ *os.File) string {
		return "mockName"
	})
	defer mockName.Reset()

	mockChmod := gomonkey.ApplyFunc(os.Chmod, func(path string, mode os.FileMode) error {
		return errors.New("mock err")
	})
	defer mockChmod.Reset()

	mockosOpenFile := gomonkey.ApplyFunc(os.OpenFile, func(name string, flag int, perm os.FileMode) (*os.File, error) {
		return mockFile, nil
	})
	defer mockosOpenFile.Reset()

	// run
	err := os.Chdir(f.tempDir)
	require.Truef(f.T(), err == nil, "Chdir err, %s", f.tempDir)
	err = File.WriteFileSafer("student.txt", []byte("zhangsan"), os.ModePerm)

	// assert
	require.Truef(f.T(), err != nil, "WriteFileSafer err, %s", f.tempDir)
	require.NoFileExists(f.T(), "student.txt")
}

func (f *FileTest) TestGopherunFile_WriteFileSafer_case7() {
	// mock
	mockosRename := gomonkey.ApplyFunc(os.Rename, func(oldpath, newpath string) error {
		f.T().Logf("oldpath: %s, newpath: %s", oldpath, newpath)
		return errors.New("mock err")
	})
	defer mockosRename.Reset()

	// run
	err := os.Chdir(f.tempDir)
	require.Truef(f.T(), err == nil, "Chdir err, %s", f.tempDir)
	err = File.WriteFileSafer("student.txt", []byte("zhangsan"), os.ModePerm)

	// assert
	require.Truef(f.T(), err != nil, "WriteFileSafer err, %s", f.tempDir)
	require.NoFileExists(f.T(), "student.txt")
}

func (f *FileTest) TestGopherunFile_WriteFileSafer_case8() {
	// mock
	mockosRename := gomonkey.ApplyFunc(os.Rename, func(oldpath, newpath string) error {
		f.T().Logf("oldpath: %s, newpath: %s", oldpath, newpath)
		return errors.New("mock err. access is denied")
	})
	defer mockosRename.Reset()

	// run
	err := os.Chdir(f.tempDir)
	require.Truef(f.T(), err == nil, "Chdir err, %s", f.tempDir)
	err = File.WriteFileSafer("student.txt", []byte("zhangsan"), os.ModePerm)

	// assert
	require.Truef(f.T(), err != nil, "WriteFileSafer err, %s", f.tempDir)
	require.NoFileExists(f.T(), "student.txt")
}
