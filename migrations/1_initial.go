package migrations

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			fmt.Println("creating table base_model1...")
			_, err := db.Exec(`CREATE TABLE base_model1(user_id varchar(32) primary key, code int)`)
			return err
		}, func(db migrations.DB) error {
			fmt.Println("dropping table base_model1...")
			_, err := db.Exec(`DROP TABLE base_model1`)
			return err
		})
}
