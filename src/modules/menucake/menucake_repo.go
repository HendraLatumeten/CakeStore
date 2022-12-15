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

func (r *menucake_repo) FindAll() (*models.MenuCake, error) {
	var data models.MenuCake

	rows, err := r.db.Query("SELECT * FROM menucake_tb ")

	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var description string
		var rating int
		var image string
		var created_at string
		var updated_at string

		err = rows.Scan(&id, &title, &description, &rating, &image, &created_at, &updated_at)
		if err != nil {
			log.Fatal(err)
		}

		data.ID = id
		data.Title = title
		data.Description = description
		data.Rating = rating
		data.Image = image
		data.Created_at = created_at
		data.Updated_at = updated_at
	}
	// To just check if any error was occured during iteration.
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &data, nil
}
