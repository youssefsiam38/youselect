package db

import (
	"os"
	"encoding/json"
	"errors"
	"github.com/youssefsiam38/youselect/models"
	"github.com/youssefsiam38/youselect/framework"
	"strings"
)

// InsertAdmin inserts admin in the database and hash his password
func InsertAdmin(email, password *string) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	encryptedPassword, err := framework.GenerateFromPassword(password)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		insert into admins (email, password)
		values (?, ?);
	`, *email, *encryptedPassword)
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			duplicated := strings.Split(err.Error(), "'")[1]
			return errors.New(`duplicated, ` + duplicated)
		}
		return err
	}

	return nil
}

// InsertAdmin inserts admin in the database and hash his password
func ForceInsertAdmin(email, password *string) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	encryptedPassword, err := framework.GenerateFromPassword(password)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		insert into admins (email, password)
		values (?, ?);
	`, *email, *encryptedPassword)
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			_, err = db.Exec(`
				update admins
				set email = ?, password = ?
				where email = ? ; 
			`, *email, *encryptedPassword, *email)
			if err != nil {
				return err
			}
				return nil
		
		}
		return err
	}

	return nil
}


// CheckAdmin verifies the admin
func CheckAdmin(a *models.Admin) (bool, *string, error) {
	db, err := Connect()
	if err != nil {
		return false, nil, err
	}
	defer db.Close()

	var newAdmin models.Admin

	row := db.QueryRow(`
		select id, email, password from admins where email=? ;
	`, a.Email)

	err = row.Scan(&newAdmin.ID, &newAdmin.Email, &newAdmin.Password)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return false, nil, errors.New("Not Found")
		}

		return false, nil, err
	}

	err = framework.CompareHashandPassword(newAdmin.Password, a.Password)
	if err != nil {
		return false, nil, nil // this time return no err because this time the err is that passwords didn't match
	}

	a.ID = newAdmin.ID

	adminJSON, err := json.Marshal(a)
	if err != nil {
		return false, nil, err
	}
	adminJSONStr := string(adminJSON)
	framework.Log(adminJSONStr)
	token, err := framework.GenerateFromPassword(&adminJSONStr)
	if err != nil {
		return false, nil, err
	}

	framework.Log(*token)

	return true, token, nil

}
// CheckAdminByToken is to check the token that the encryption of the FULL ADMIN STRUCT
func CheckAdminByToken(token *string) error {

	var databaseAdmin models.Admin

	db, err := Connect()
	if err != nil {
		return nil
	}
	defer db.Close()

	rows, err := db.Query(`
		select * from admins
		limit 10;
	`)
	defer rows.Close()


	for rows.Next() {
		rows.Scan(&databaseAdmin.ID, &databaseAdmin.Email, &databaseAdmin.Password) // hashed
		err = framework.CompareHashandPassword(databaseAdmin.Password, os.Getenv("ADMIN_PASS"))
		if err == nil {
			databaseAdmin.Password = os.Getenv("ADMIN_PASS")
		} else {
			return err
		}

		databaseAdminJSON, err := json.Marshal(databaseAdmin)
		if err != nil {
			return err
		}

		err = framework.CompareHashandPassword(*token, string(databaseAdminJSON))

		if err == nil {
			return nil
		}
	}

	return errors.New("Not Found")

}
