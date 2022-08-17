package migrations

import (
	"fmt"
	"github.com/go-pg/migrations/v8"
)

func init() {

	//log.Println("migration init called")
	//migrations.MustRegisterTx(func(db migrations.DB) error {
	//	err := db.Model(&model.ProductModel{}).CreateTable(nil)
	//	if err != nil {
	//		log.Println(err)
	//		return err
	//	}
	//	return nil
	//})
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			fmt.Println("creating table product...")
			_, err := db.Exec(`CREATE TABLE product(name varchar(32) primary key, qty int)`)
			return err
		}, func(db migrations.DB) error {
			fmt.Println("dropping table product...")
			_, err := db.Exec(`DROP TABLE product`)
			return err
		})
}
