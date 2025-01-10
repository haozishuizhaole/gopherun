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
)

func (i *GopherunFile) MkdirAll(path string) error {
	return i.MkdirAllWithMode(path, os.ModePerm)
}

func (i *GopherunFile) MkdirAllWithMode(path string, mode os.FileMode) error {
	return os.MkdirAll(path, mode)
}

func (i *GopherunFile) Remove(path string) error {
	return os.Remove(path)
}

func (i *GopherunFile) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (i *GopherunFile) IsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func (i *GopherunFile) GetAbsolutePath(path string) (string, error) {
	return filepath.Abs(path)
}

func (i *GopherunFile) GetPwd() (pwd string, err error) {
	return os.Getwd()
}
