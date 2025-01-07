/*
 *    Copyright (c) 2025 Middling
 *    Gopherun is licensed under Mulan PSL v2.
 *    You can use this software according to the terms and conditions of the Mulan PSL v2.
 *    You may obtain a copy of Mulan PSL v2 at:
 *             http://license.coscl.org.cn/MulanPSL2
 *    THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 *    See the Mulan PSL v2 for more details.
 */

package gopherun

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGopherunJSON_Encode_case1(t *testing.T) {
	// mock
	var i GopherunJSON

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := &User{
		Name: "zhangsan",
		Age:  12,
	}

	// run
	bytes, err := i.Encode(user)

	// assert
	assert.True(t, err == nil)
	assert.True(t, bytes != nil)
	assert.True(t, string(bytes) == "{\"name\":\"zhangsan\",\"age\":12}")
}

func TestGopherunJSON_EncodeToJSONStr_case1(t *testing.T) {
	// mock
	var i GopherunJSON

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := &User{
		Name: "zhangsan",
		Age:  12,
	}

	// run
	jsonStr, err := i.EncodeToJSONStr(user)

	// assert
	assert.True(t, err == nil)
	assert.True(t, jsonStr == "{\"name\":\"zhangsan\",\"age\":12}")
}

func TestGopherunJSON_EncodeToJSONStr_case2(t *testing.T) {
	// mock
	var i GopherunJSON

	type Node struct {
		Value string
		Next  *Node
	}

	node1 := &Node{Value: "Node1"}
	node2 := &Node{Value: "Node2"}
	node1.Next = node2
	node2.Next = node1 // 循环引用

	// run
	jsonStr, err := i.EncodeToJSONStr(node1)

	// assert
	assert.True(t, err != nil)
	assert.True(t, jsonStr == "")
}

func TestGopherunJSON_Decode_case1(t *testing.T) {
	// mock
	var i GopherunJSON

	jsonStr := "{\"name\":\"zhangsan\",\"age\":12}"

	// run

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := &User{}

	err := i.Decode([]byte(jsonStr), user)

	// assert
	assert.True(t, err == nil)
	assert.True(t, reflect.DeepEqual(user, &User{Name: "zhangsan", Age: 12}))
}

func TestGopherunJSON_Decode_case2(t *testing.T) {
	// mock
	var i GopherunJSON

	jsonStr := "{\"name\":\"zhangsan\",\"age\":12}sdsdsdsd" // invalid json str

	// run

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := &User{}

	err := i.Decode([]byte(jsonStr), user)

	// assert
	assert.True(t, err != nil)
	assert.True(t, reflect.DeepEqual(user, &User{}))
}

func TestGopherunJSON_Decode_case3(t *testing.T) {
	// mock
	var i GopherunJSON

	jsonStr := "{\"name\":\"zhangsan\",\"age\":\"abc\"}" // invalid age type

	// run

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := &User{}

	err := i.Decode([]byte(jsonStr), user)

	// assert
	assert.True(t, err != nil)
	assert.True(t, reflect.DeepEqual(user, &User{Name: "zhangsan"})) // json反序列化部分失败，注意此时不要使用该结构
}

func TestGopherunJSON_DecodeByJSONStr_case1(t *testing.T) {
	// mock
	var i GopherunJSON

	jsonStr := "{\"name\":\"zhangsan\",\"age\":12}"

	// run

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := &User{}

	err := i.DecodeByJSONStr(jsonStr, user)

	// assert
	assert.True(t, err == nil)
	assert.True(t, reflect.DeepEqual(user, &User{Name: "zhangsan", Age: 12}))
}
