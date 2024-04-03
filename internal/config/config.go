/*
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MysqlDNS  string
	RedisConf RedisConf
}

type RedisConf struct {
	Host string
	Port int32
	Pass string
}
