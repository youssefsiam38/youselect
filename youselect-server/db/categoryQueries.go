package db

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/youssefsiam38/youselect/models"
)

// InsertCategory in the database
func InsertCategory(c *models.Category) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	if c.Title == "" {
		return errors.New("Category must have title to be inserted in the database")
	}

	if c.ID != 0 {
		_, err = db.Exec(`
			insert into categories (id, title, imageURL) values (?, ?, ?)
		`, c.ID, c.Title, c.ImageURL)

	} else {
		_, err = db.Exec(`
			insert into categories (title, imageURL) values (?, ?)
		`, c.Title, c.ImageURL)

	}
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			duplicated := strings.Split(err.Error(), "'")[1]
			return errors.New(`duplicated, ` + duplicated)
		}
		return err
	}

	return nil
}

// SelectCategoryByID gets one Category from the database based on ID
func SelectCategoryByID(ID *uint) (*models.Category, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var row *sql.Row

	if *ID == 0 {
		return nil, errors.New("Must provide ID")
	}

	row = db.QueryRow(`
			select id, title, imageURL from categories where id=?;
		`, *ID)

	var category models.Category

	err = row.Scan(&category.ID, &category.Title, &category.ImageURL)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, errors.New("Not Found")
		}

		return nil, err
	}

	return &category, nil
}

// SelectCategoryByTitle gets one Category from the database based on title
func SelectCategoryByTitle(title *string) (*models.Category, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var row *sql.Row

	if *title == "" {
		return nil, errors.New("Must provide title")
	}

	row = db.QueryRow(`
			select id, title, imageURL from categories where title=?;
		`, *title)

	var newCateg models.Category

	err = row.Scan(&newCateg.ID, &newCateg.Title, &newCateg.ImageURL)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, errors.New("Not Found")
		}
		return nil, err
	}

	return &newCateg, nil
}

// SelectAllCategories in the database
func SelectAllCategories() (*[]models.Category, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select id, title, imageURL from categories order by title")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		var tempCateg models.Category
		rows.Scan(&tempCateg.ID, &tempCateg.Title, &tempCateg.ImageURL)
		categories = append(categories, tempCateg)
	}

	if len(categories) == 0 {
		return nil, errors.New("Not Found")
	}

	return &categories, nil
}

// UpdateCategory in the database
func UpdateCategory(oldID uint, new *models.Category) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	//must have id
	if oldID == 0 {
		return errors.New("ID cannot be 0")
	}

	_, err = db.Exec(`
		update categories
		set title=?, imageURL=?
		where id=? ;
	`, new.Title, new.ImageURL, oldID)

	if err != nil {
		return err
	}

	return nil
}

// DeleteCategory from database
func DeleteCategory(ID *uint) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	//must have id
	if *ID == 0 {
		return errors.New("ID cannot be 0")
	}

	_, err = db.Exec(`
		delete from categories where id=?
	`, *ID)

	if err != nil {
		return err
	}

	return nil
}
