package dvdrental

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
)

func logAndQuery(db *gorm.DB, query string, args ...interface{}) *sql.Rows {
	fmt.Println(query)
	res, err := db.Debug().Raw(query, args...).Rows()
	if err != nil {
		panic(err)
	}
	return res
}
