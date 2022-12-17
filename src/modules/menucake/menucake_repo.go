package menucake

import (
	"database/sql"
	"errors"
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

	rows, err := r.db.Query("SELECT * FROM menucake_tb ")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var row models.MenuCake

		err = rows.Scan(&row.ID, &row.Title, &row.Description, &row.Rating, &row.Image, &row.Created_at, &row.Updated_at)
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

	rows, err := r.db.Query("SELECT * FROM menucake_tb WHERE id= ?", id)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var row models.MenuCake

		err = rows.Scan(&row.ID, &row.Title, &row.Description, &row.Rating, &row.Image, &row.Created_at, &row.Updated_at)
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

func (r *menucake_repo) SaveCake(data *models.MenuCake) (*models.MenuCake, error) {
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

func (r *menucake_repo) DeleteCake(id int) (*models.MenuCake, error) {
	var data models.MenuCake

	res, err := r.db.Exec("DELETE FROM menucake_tb WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	CheckID, err := res.RowsAffected()
	if err != nil {
		panic(err.Error())
	} else if CheckID == 0 {
		return nil, errors.New("ID not Found!")
	}

	return &data, nil
}

func (r *menucake_repo) UpdateCake(id string, data *models.MenuCake) (*models.MenuCake, error) {
	var datas models.MenuCake

	stmt, err := r.db.Prepare("UPDATE menucake_tb SET title = ?, description = ?, rating = ?, image= ?  WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	res, err := stmt.Exec(data.Title, data.Description, data.Rating, data.Image, id)
	if err != nil {
		panic(err.Error())
	}

	CheckID, err := res.RowsAffected()
	if err != nil {
		panic(err.Error())
	} else if CheckID == 0 {
		return nil, errors.New("ID not Found!")
	}

	return &datas, nil
}
