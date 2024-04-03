/*
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package code defines the error code for all module
package code

// RespCode represents the response code
type RespCode int

// Success 服务端返回码，200 表示成功，其他表示失败
const (
	Success RespCode = iota + 200
)

// 10000 - 19999 HTTP接口服务
// 20000 - 29999 IM系统
// ...

// errMsg 返回码对应具体的信息
var errMsg = map[RespCode]string{
	Success: "success",
}

// String 返回码对应具体的信息
func (rc RespCode) String() string {
	return errMsg[rc]
}
