package db

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/youssefsiam38/youselect/models"
)

// InsertStore in the database
func InsertStore(s *models.Store) error {

	db, err := Connect()
	if err != nil {
		return err
	}

	if s.ID != 0 {
		_, err = db.Exec(`
			insert into stores (id, name, affiliateURL, imageURL, querySelector)
			values (?, ?, ?, ?, ?);
		`, s.ID, s.Name, s.AffiliateURL, s.ImageURL, s.QuerySelector)
	} else {
		_, err = db.Exec(`
			insert into stores (name, affiliateURL, imageURL, querySelector)
			values (?, ?, ?, ?);
		`, s.Name, s.AffiliateURL, s.ImageURL, s.QuerySelector)
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

// SelectStoreByID gets one Store from the database based on provided ID
func SelectStoreByID(ID *uint) (*models.Store, error) {
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
			select id, name, affiliateURL, imageURL, querySelector from stores where id=?;
		`, *ID)

	var store models.Store

	err = row.Scan(&store.ID, &store.Name, &store.AffiliateURL, &store.ImageURL, &store.QuerySelector)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, errors.New("Not Found")
		}

		return nil, err
	}

	return &store, nil
}

// SelectStoreByName gets one Store from the database based on provided name
func SelectStoreByName(name *string) (*models.Store, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var row *sql.Row

	if *name == "" {
		return nil, errors.New("Must provide name")
	}

	row = db.QueryRow(`
			select id, name, affiliateURL, imageURL, querySelector from stores where name=?;
		`, *name)

	var store models.Store

	err = row.Scan(&store.ID, &store.Name, &store.AffiliateURL, &store.ImageURL, &store.QuerySelector)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, errors.New("Not Found")
		}

		return nil, err
	}

	return &store, nil
}

// SelectAllStores in the database
func SelectAllStores() (*[]models.Store, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select id, name, affiliateURL, imageURL, querySelector from stores order by name ;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stores []models.Store

	for rows.Next() {
		var tempStore models.Store
		rows.Scan(&tempStore.ID, &tempStore.Name, &tempStore.AffiliateURL, &tempStore.ImageURL, &tempStore.QuerySelector)
		stores = append(stores, tempStore)
	}

	if len(stores) == 0 {
		return nil, errors.New("Not Found")
	}

	return &stores, nil
}

// UpdateStore in the database
func UpdateStore(oldID uint, new *models.Store) error {
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
		update stores
		set name=?, affiliateURL=?, imageURL=?, querySelector=?
		where id=? ;
	`, new.Name, new.AffiliateURL, new.ImageURL, new.QuerySelector, oldID)

	if err != nil {
		return err
	}

	return nil
}

// DeleteStore from database
func DeleteStore(ID *uint) error {
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
		delete from stores where id=? ;
	`, *ID)

	if err != nil {
		return err
	}

	return nil
}
