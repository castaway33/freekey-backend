// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Wallet is the golang structure of table u_wallet for DAO operations like Where/Data.
type Wallet struct {
	g.Meta       `orm:"table:u_wallet, do:true"`
	Id           interface{} //
	Uid          interface{} //
	Balance      interface{} //
	Pass         interface{} //
	PassErrCount interface{} // 密码输错次数
	Desc         interface{} //
	Status       interface{} // 金库状态 0 设置密码 1正常,2 锁定
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}