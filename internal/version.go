/*
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package internal defines the internal constants
package internal

import "fmt"

// Version 参数
var (
	CommitID  = "" // git commit hash
	BuildTime = "" // when to build
	Version   = "" // what version
)

// VersionInfo show version data
//
//	@Description:
//	@return string
func VersionInfo() string {
	return fmt.Sprintf("\nCurrent version: %s\nCommit hash: %s\nBuild time: %s\n\n", Version, CommitID, BuildTime)
}
