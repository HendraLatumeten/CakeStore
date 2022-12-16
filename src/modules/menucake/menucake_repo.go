package menucake

import (
	"database/sql"
	"log"

	"github.com/hendralatumeten/cakestore/src/database/models"
)

type menucake_repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *menucake_repo {
	return &menucake_repo{db: db}
}

func (r *menucake_repo) ListCake() (*models.MenuCakeAll, error) {
	var data models.MenuCakeAll

	rows, err := r.db.Query("SELECT title FROM menucake_tb ")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var row models.MenuCake

		err = rows.Scan(&row.Title)
		if err != nil {
			panic(err)
		}
		data = append(data, row)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &data, nil
}

func (r *menucake_repo) DetailCake(id int) (*models.MenuCakeAll, error) {
	var data models.MenuCakeAll

	rows, err := r.db.Query("SELECT id, title, description, rating, image FROM menucake_tb WHERE id= ?", id)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var row models.MenuCake

		err = rows.Scan(&row.ID, &row.Title, &row.Description, &row.Rating, &row.Image)
		if err != nil {
			panic(err)
		}
		data = append(data, row)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &data, nil
}

func (r *menucake_repo) Save(data *models.MenuCake) (*models.MenuCake, error) {
	var datas models.MenuCake

	stmt, err := r.db.Prepare("INSERT INTO menucake_tb (title , description, rating, image) VALUES (?, ?, ?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Title, data.Description, data.Rating, data.Image)
	if err != nil {
		panic(err.Error())
	}

	return &datas, nil
}
