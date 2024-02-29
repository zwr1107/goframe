// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goframe/internal/dao/internal"
)

// internalUserInfoDao is internal type for wrapping internal DAO implements.
type internalUserInfoDao = *internal.UserInfoDao

// userInfoDao is the data access object for table user_info.
// You can define custom methods on it to extend its functionality as you wish.
type userInfoDao struct {
	internalUserInfoDao
}

var (
	// UserInfo is globally public accessible object for table user_info operations.
	UserInfo = userInfoDao{
		internal.NewUserInfoDao(),
	}
)

// Fill with you ideas below.
