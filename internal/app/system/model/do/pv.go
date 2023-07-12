// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Pv is the golang structure of table pv for DAO operations like Where/Data.
type Pv struct {
	g.Meta    `orm:"table:pv, do:true"`
	Id        interface{} //
	Url       interface{} //
	Event     interface{} //
	Ip        interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}