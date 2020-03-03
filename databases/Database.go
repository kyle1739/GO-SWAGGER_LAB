package database

import(
	"github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/mysql"
)

type(

	todoModel struct{
		gorm.Model
		Title string `json:"title"`
		Completed int `json:"completed"`
	}
)

var Db *gorm.DB

func init() {
        var err error
        Db, err = gorm.Open("mysql", "root:123456@/todolist?charset=utf8&parseTime=True&loc=Local")
        if err != nil {
                panic("failed to connect database")
        }
        Db.AutoMigrate(&todoModel{})
}

