/*
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package gorm

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (

	// MaxLifetime 最大连接时间，防止使用无效的连接，单位：s
	MaxLifetime = 500

	// MaxIdleConns 空闲连接池中连接的最大数量
	MaxIdleConns = 10

	// MaxOpenConns 数据库连接的最大数量
	MaxOpenConns = 10

	// DefaultStringSize string 类型字段的默认长度
	DefaultStringSize = 256

	// DisableDatetimePrecision 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	DisableDatetimePrecision = true

	// DontSupportRenameIndex 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	DontSupportRenameIndex = true

	// DontSupportRenameColumn 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	DontSupportRenameColumn = true

	// SkipInitializeWithVersion 根据当前 MySQL 版本自动配置
	SkipInitializeWithVersion = false
)

// GetConnection 获取数据库连接
func GetConnection(dsn string) (*gorm.DB, error) {
	ctx := context.Background()
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		SkipInitializeWithVersion: SkipInitializeWithVersion,
		DefaultStringSize:         DefaultStringSize,
		DisableDatetimePrecision:  DisableDatetimePrecision,
		DontSupportRenameIndex:    DontSupportRenameIndex,
		DontSupportRenameColumn:   DontSupportRenameColumn,
	}), &gorm.Config{})

	if err != nil {
		logx.WithContext(ctx).Errorf("failed to open MySQL connection: %s", err)
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		logx.WithContext(ctx).Errorf("failed to get MySQL connection %s", err)
		return nil, err
	}

	sqlDb.SetMaxIdleConns(MaxIdleConns)
	sqlDb.SetMaxOpenConns(MaxOpenConns)
	sqlDb.SetConnMaxLifetime(MaxLifetime * time.Second)

	if err = sqlDb.Ping(); err != nil {
		logx.WithContext(ctx).Errorf("failed to ping MySQL connection", err)
		return nil, err
	}
	logx.WithContext(ctx).Info("success to connect MySQL")

	return db, err
}
