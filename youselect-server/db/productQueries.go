package db

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"

	"github.com/youssefsiam38/youselect/framework"

	"github.com/youssefsiam38/youselect/models"
)

// InsertProduct in the database
func InsertProduct(p *models.Product) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	if p.Title == "" {
		return errors.New("Product must have title to be inserted in the database")
	}

	err = p.CalcPriority()
	if err != nil {
		return err
	}

	if p.ID != 0 {
		_, err = db.Exec(`
		insert into products (id, title, categoryID, price, storeID, imageURL, affiliateURL, productURL, priority, commissionPercent) 
		values (?, ?, (select id from categories where title=?), ?, (select id from stores where name=?), ?, ?, ?, ?, ?);
		`, p.ID, p.Title, p.Category, p.Price, p.Store, p.ImageURL, p.AffiliateURL, p.ProductURL, p.Priority, p.CommissionPercent)

	} else {
		_, err = db.Exec(`
		insert into products (title, categoryID, price, storeID, imageURL, affiliateURL, productURL, priority, commissionPercent) 
		values (?, (select id from categories where title=?), ?, (select id from stores where name=?), ?, ?, ?, ?, ?);
		`, p.Title, p.Category, p.Price, p.Store, p.ImageURL, p.AffiliateURL, p.ProductURL, p.Priority, p.CommissionPercent)

	}
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			duplicated := strings.Split(err.Error(), "'")[1]
			return errors.New(`duplicated, ` + duplicated)
		}
		return err
	}

	return nil
}

// SelectProductByID gets one product from the database based on provided ID
func SelectProductByID(ID *uint) (*models.Product, error) {
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
		select products.id, products.title, categories.title, products.price, stores.name, products.imageURL, products.affiliateURL, products.productURL, products.priority, products.commissionPercent 
		from products 
		join categories on products.categoryID = categories.id 
		join stores on products.storeID = stores.id 
		where products.id=?;
	`, *ID)

	var p models.Product

	err = row.Scan(&p.ID, &p.Title, &p.Category, &p.Price, &p.Store, &p.ImageURL, &p.AffiliateURL, &p.ProductURL, &p.Priority, &p.CommissionPercent)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, errors.New("Not Found")
		}

		return nil, err
	}

	return &p, nil
}

// SelectAllProducts in the database
func SelectAllProducts() (*[]models.Product, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		select products.id, products.title, categories.title, products.price, stores.name, products.imageURL, products.affiliateURL, products.productURL, products.priority, products.commissionPercent 
		from products 
		join categories on products.categoryID = categories.id 
		join stores on products.storeID = stores.id 
		order by products.priority;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product

		rows.Scan(&p.ID, &p.Title, &p.Category, &p.Price, &p.Store, &p.ImageURL, &p.AffiliateURL, &p.ProductURL, &p.Priority, &p.CommissionPercent)

		products = append(products, p)
	}

	if len(products) == 0 {
		return nil, errors.New("Not Found")
	}
	return &products, nil
}

// SelectProducts from the database with pagination
//
// withSearch till the database that this query include search
//
// exact till the database that the resulted product title must include all the words in the search, in other words it ask to use and not or
//
// page is the pagination
//
// min and max is like (min < price < max)
//
// store for specifing store
//
// search is only useful if withSearch is true
func SelectProducts(withSearch, exact bool, page, min, max, store, search *string) (*[]models.Product, error) {

	queryMin := float64(0)
	queryMax := float64(999999999.99)
	queryStore := "%"
	queryLimit := 0

	if *min != "" {
		min, err := strconv.ParseFloat(*min, 64)
		if err != nil {
			return nil, err
		}
		queryMin = min
	}

	if *max != "" {
		max, err := strconv.ParseFloat(*max, 64)
		if err != nil {
			return nil, err
		}
		queryMax = max
	}

	if *store != "" {
		queryStore = *store
	}

	if *page != "" {
		page, err := strconv.Atoi(*page)
		if err != nil {
			return nil, err
		}
		queryLimit = (page * 24) - 24
	}

	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var rows *sql.Rows

	if withSearch {
		queryMinStr := framework.FloatToStr(queryMin)
		queryMaxStr := framework.FloatToStr(queryMax)
		rows, err = framework.CreateSearchRows(db, search, &queryMinStr, &queryMaxStr, &queryStore, &queryLimit)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

	} else {
		// create the query and add the search conditions if exists
		query := `
			select products.id, products.title, categories.title, products.price, stores.name, products.imageURL, products.affiliateURL, products.productURL, products.priority, products.commissionPercent 
			from products 
			join categories on products.categoryID = categories.id 
			join stores on products.storeID = stores.id 
			where price >= ? and price <= ? and stores.name like ?
			order by products.priority DESC 
			limit ?, 24 ;
		`
		rows, err = db.Query(query, framework.FloatToStr(queryMin), framework.FloatToStr(queryMax), queryStore, queryLimit)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
	}
	
	var products []models.Product

	if withSearch {
		for rows.Next() {
			var p models.Product

			rows.Scan(&p.ID, &p.Title, &p.Category, &p.Price, &p.Store, &p.ImageURL, &p.AffiliateURL, &p.ProductURL, &p.Priority, &p.CommissionPercent, nil)

			products = append(products, p)
		}
	} else {
		for rows.Next() {
			var p models.Product

			rows.Scan(&p.ID, &p.Title, &p.Category, &p.Price, &p.Store, &p.ImageURL, &p.AffiliateURL, &p.ProductURL, &p.Priority, &p.CommissionPercent)

			products = append(products, p)
		}

	}

	if len(products) == 0 {
		return nil, errors.New("Not Found")
	}
	return &products, nil
}

// UpdateProduct in the database
func UpdateProduct(oldID uint, new *models.Product) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	//must have id
	if oldID == 0 {
		return errors.New("ID cannot be 0")
	}

	err = new.CalcPriority()
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		update products
		set title=?, categoryID=(select id from categories where title=?), price=?, storeID=(select id from stores where name=?), imageURL=?, affiliateURL=?, productURL=?, priority=?, commissionPercent=?
		where id=? ;

	`, new.Title, new.Category, new.Price, new.Store, new.ImageURL, new.AffiliateURL, new.ProductURL, new.Priority, new.CommissionPercent, oldID)

	if err != nil {
		return err
	}

	return nil
}

// DeleteProduct from database
func DeleteProduct(ID *uint) error {
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
		delete from products where id=?
	`, *ID)

	if err != nil {
		return err
	}

	return nil
}
