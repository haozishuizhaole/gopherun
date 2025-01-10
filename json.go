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

import "encoding/json"

func (i *GopherunJSON) Encode(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func (i *GopherunJSON) EncodeToJSONStr(obj interface{}) (jsonStr string, err error) {
	bytes, err := i.Encode(obj)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (i *GopherunJSON) Decode(bytes []byte, obj interface{}) error {
	return json.Unmarshal(bytes, obj)
}

func (i *GopherunJSON) DecodeByJSONStr(jsonStr string, obj interface{}) error {
	return i.Decode([]byte(jsonStr), obj)
}
