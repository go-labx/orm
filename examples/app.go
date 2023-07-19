package main

import (
	"fmt"

	"github.com/go-labx/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int    `db:"id" pk:"true" auto:"true"`
	Name      string `db:"name"`
	Email     string `db:"email" unique:"true"`
	Password  string `db:"password"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

func (u *User) TableName() string {
	return "user"
}

func main() {
	DB, err := orm.NewDB(orm.NewDataSource(
		orm.SetUser("root"),
		orm.SetPassword("12345678"),
		orm.SetDBName("orm"),
	))
	if err != nil {
		panic(err)
	}
	defer func(DB *orm.DB) {
		err := DB.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(DB)

	session := DB.NewSession()
	result, err := session.Raw("SELECT * FROM user").Exec()
	fmt.Println(result, err)
}
