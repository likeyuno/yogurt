package static

import (
	"database/sql"
	"time"
)

type Table struct {
	ID        string       `xorm:"<- notnull pk UUID 'id'"`
	CreatedAt time.Time    `xorm:"notnull created 'created_at'"`
	UpdatedAt time.Time    `xorm:"notnull updated 'updated_at'"`
	DeletedAt sql.NullTime `xorm:"deleted 'deleted_at'"`
}
