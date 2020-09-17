package framework

import (
	"database/sql"
	"fmt"
)

func CreateSearchRows(db *sql.DB, searchQuery , queryMin, queryMax, queryStore *string, queryLimit *int) (*sql.Rows, error) {

	var rows *sql.Rows

	exactSearchCondPtr := CreateSearchConditions(searchQuery, true)
	notExactSearchCondPtr := CreateSearchConditions(searchQuery, false)
	exactSearchCond := *exactSearchCondPtr
	notExactSearchCond := *notExactSearchCondPtr

	// create the query and add the search conditions if exists
	query := fmt.Sprintf(`
		select distinct * from (
			(
				select products.id, products.title as title, categories.title as category, products.price, stores.name, products.imageURL, products.affiliateURL, products.productURL, products.priority, products.commissionPercent, 1 as tNo
				from products 
				join categories on products.categoryID = categories.id 
				join stores on products.storeID = stores.id 
				where price >= ? and price <= ? and stores.name like ? %s
			)
			union
			(
				select products.id, products.title as title, categories.title as category, products.price, stores.name, products.imageURL, products.affiliateURL, products.productURL, products.priority, products.commissionPercent, 2 as tNo
				from products 
				join categories on products.categoryID = categories.id 
				join stores on products.storeID = stores.id 
				where price >= ? and price <= ? and stores.name like ? %s
			)
		) as products
		group by products.id
		order by tNo, products.priority DESC
		limit ?, 24 ;
		`, exactSearchCond, notExactSearchCond)

	rows, err := db.Query(query, *queryMin, *queryMax, *queryStore, *queryMin, *queryMax, *queryStore, *queryLimit)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

