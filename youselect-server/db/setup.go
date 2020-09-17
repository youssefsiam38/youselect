package db

import (
	"database/sql"
	"os"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/youssefsiam38/youselect/framework"
)

// Setup Creates DB If Not Exists
func Setup() {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+`:`+os.Getenv("MYSQL_PASS")+`@tcp(`+os.Getenv("MYSQL_HOST")+`:`+os.Getenv("MYSQL_PORT")+`)/`)
	defer db.Close()

	if err != nil {
		framework.Log(err)
	}

	

	err = createAdminsTable()
	if err != nil {
		framework.Log(err)
	}

	err = createStoresTable()
	if err != nil {
		framework.Log(err)
	}

	err = createCategoriesTable()
	if err != nil {
		framework.Log(err)
	}

	err = createProductsTable()
	if err != nil {
		framework.Log(err)
	}

}

// Connect to database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+`:`+os.Getenv("MYSQL_PASS")+`@tcp(`+os.Getenv("MYSQL_HOST")+`:`+os.Getenv("MYSQL_PORT")+`)/`+os.Getenv("DB_NAME"))
	if err != nil {
		framework.Log(err)
	}

	// reset the sql_mode to use not full group by and avoid error 1055
	defer db.Exec(`SET sql_mode=(SELECT REPLACE(@@sql_mode,'ONLY_FULL_GROUP_BY',''));`)

	duration, err := time.ParseDuration("60ms")
	if err != nil {
		return nil, err
	}

	defer db.SetConnMaxLifetime(duration)

	return db, nil
}

// createAdminsTable in the database
func createAdminsTable() error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(
		`create table if not exists admins (
			id int auto_increment not null unique, 
			email varchar(50)  not null unique,
			password varchar(255) not null,
			primary key(id)
		);
	`)
	if err != nil {
		return err
	}
	adminEmail, adminPass := os.Getenv("ADMIN_EMAIL"), os.Getenv("ADMIN_PASS")

	err = ForceInsertAdmin(&adminEmail, &adminPass)
	if err != nil {
		return err
	}
	
	return nil
}

// createStoresTable in the database
func createStoresTable() error {
	db, err := Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec(
		`create table if not exists stores (
			id int auto_increment not null unique,
			name varchar(255) not null unique,
			affiliateURL varchar(5000) not null,
			imageURL varchar(5000) not null,
			querySelector varchar(255) not null,
			primary key(id)
		);
	`)
	if err != nil {
		return err
	}

	return nil
}

// createCategoriesTable in the database
func createCategoriesTable() error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(
		`create table if not exists categories (
			id int auto_increment not null unique,
			title varchar(255) unique,
			imageURL varchar(5000) not null,
			primary key(id)
		);
	`)
	if err != nil {
		return err
	}

	return nil
}

// createProductsTable in the database
func createProductsTable() error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(
		`create table if not exists products (
			id int auto_increment not null unique,
			title varchar(500) not null,
			categoryID int not null,
			price decimal(11, 2) not null,
			storeID int not null,
			imageURL varchar(5000) not null,
			affiliateURL varchar(5000) not null,
			productURL varchar(5000) not null,
			priority decimal(11, 2) not null,
			commissionPercent decimal(4, 2) not null,
			primary key(id),
			foreign key(categoryID) references categories(id),
			foreign key(storeID) references stores(id)
		);
	`)
	if err != nil {
		return err
	}

	return nil
}

// FlushDatabase drops the databases
func FlushDatabase(DBname string) error {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+`:`+os.Getenv("MYSQL_PASS")+`@tcp(`+os.Getenv("MYSQL_HOST")+`:`+os.Getenv("MYSQL_PORT")+`)/`)
	if err != nil {
		return err
	}
	defer db.Close()

	err = framework.OneWordChecker(DBname)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`
	DROP DATABASE if exists ` + DBname + `;
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
