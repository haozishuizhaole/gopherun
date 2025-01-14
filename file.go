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
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (i GopherunFile) MkdirAll(path string) error {
	return i.MkdirAllWithMode(path, os.ModePerm)
}

func (i GopherunFile) MkdirAllWithMode(path string, mode os.FileMode) error {
	return os.MkdirAll(path, mode)
}

func (i GopherunFile) Remove(path string) error {
	return os.Remove(path)
}

func (i GopherunFile) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (i GopherunFile) IsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func (i GopherunFile) GetAbsolutePath(path string) (string, error) {
	return filepath.Abs(path)
}

func (i GopherunFile) GetPwd() (pwd string, err error) {
	return os.Getwd()
}

func (i GopherunFile) Size(path string) (int64, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return -1, err
	}
	return fileInfo.Size(), nil
}
func (i GopherunFile) IsDir(path string) bool {
	fileInfo, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// WriteFileSafer 将数据先写入临时文件，成功后自动重命名为指定文件名。
func (i GopherunFile) WriteFileSafer(writePath string, data []byte, perm os.FileMode) (err error) {
	// credits: https://github.com/88250/gulu/blob/master/file.go
	dir, name := filepath.Split(writePath)

	// 临时文件全限定路径
	tmp := filepath.Join(dir, name+Random.RandomStringWithNumberAndLetter(10)+".tmp")

	// 创建临时文件
	f, err := os.OpenFile(tmp, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if nil != err {
		return
	}

	// 写入数据
	if _, err = f.Write(data); nil != err {
		return
	}

	if err = f.Sync(); nil != err {
		return
	}

	if err = f.Close(); nil != err {
		return
	}

	// 修改临时文件mod
	if err = os.Chmod(f.Name(), perm); nil != err {
		return
	}

	// 重命名
	for retryCount := 0; retryCount < 3; retryCount++ {
		err = os.Rename(f.Name(), writePath) // Windows 上重命名是非原子的
		if nil == err {
			_ = i.Remove(f.Name())
			return
		}

		if errMsg := strings.ToLower(err.Error()); strings.Contains(errMsg, "access is denied") || strings.Contains(errMsg, "used by another process") { // 文件可能是被锁定
			time.Sleep(200 * time.Millisecond)
			continue
		}
		break
	}
	return
}
