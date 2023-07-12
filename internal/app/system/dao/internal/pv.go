// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PvDao is the data access object for table pv.
type PvDao struct {
	table   string    // table is the underlying table name of the DAO.
	group   string    // group is the database configuration group name of current DAO.
	columns PvColumns // columns contains all the column names of Table for convenient usage.
}

// PvColumns defines and stores column names for table pv.
type PvColumns struct {
	Id        string //
	Url       string //
	Event     string //
	Ip        string //
	CreatedAt string //
	UpdatedAt string //
}

// pvColumns holds the columns for table pv.
var pvColumns = PvColumns{
	Id:        "id",
	Url:       "url",
	Event:     "event",
	Ip:        "ip",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewPvDao creates and returns a new DAO object for table data access.
func NewPvDao() *PvDao {
	return &PvDao{
		group:   "default",
		table:   "pv",
		columns: pvColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PvDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PvDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PvDao) Columns() PvColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PvDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PvDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PvDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}