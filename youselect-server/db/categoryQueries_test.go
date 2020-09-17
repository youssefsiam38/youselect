package db

import (
	"os"
	// "github.com/youssefsiam38/youselect/framework"
	"reflect"
	"testing"

	"github.com/youssefsiam38/youselect/models"
)

func TestInsertCategory(t *testing.T) {

	inserted := models.Category{
		Title:    "testTitle1",
		ImageURL: "testUR1",
	}

	err := InsertCategory(&inserted)
	if err != nil {
		t.Error(err)
	}

	retrieved, err := SelectCategoryByTitle(&inserted.Title)
	if err != nil {
		t.Error(err)
	}

	if retrieved.Title != inserted.Title ||
		retrieved.ImageURL != inserted.ImageURL {
		t.Error("Inserted Category is not the retieved one")
	}
}

func TestSelectCategoryByID(t *testing.T) {

	inserted := models.Category{
		ID:       2,
		Title:    "testTitle2",
		ImageURL: "testURL2",
	}

	err := InsertCategory(&inserted)
	if err != nil {
		t.Error(err)
	}

	// test retrieving category by id only

	retrieved, err := SelectCategoryByID(&inserted.ID)
	if err != nil {
		t.Error(err)
	}

	if retrieved.Title != inserted.Title ||
		retrieved.ImageURL != inserted.ImageURL {
		t.Error("Inserted Category is not the retieved one")
	}

	// test retrieving category by title only
	inserted.Title = "testTitle2"

	retrieved, err = SelectCategoryByTitle(&inserted.Title)
	if err != nil {
		t.Error(err)
	}

	if retrieved.Title != inserted.Title ||
		retrieved.ImageURL != inserted.ImageURL {
		t.Error("Inserted Category is not the retieved one")
	}

	// test retrieving not existing category
	inserted.Title = "testTitle999"
	inserted.ID = 999

	retrieved, err = SelectCategoryByID(&inserted.ID)
	if err.Error() != "Not Found" {
		t.Error(err)
	}
}

func TestSelectAllCategories(t *testing.T) {
	FlushDatabase(os.Getenv("DB_NAME"))
	Setup()

	insertedCategories := []models.Category{
		models.Category{ID: 1, Title: "laptops", ImageURL: "1"},
		models.Category{ID: 2, Title: "mobiles", ImageURL: "2"},
		models.Category{ID: 3, Title: "pcs", ImageURL: "3"},
		models.Category{ID: 4, Title: "usbs", ImageURL: "4"},
	}

	for i := 0; i < len(insertedCategories); i++ {
		err := InsertCategory(&insertedCategories[i])
		if err != nil {
			t.Error(err)
		}
	}

	retrievedCategories, err := SelectAllCategories()
	if err != nil {
		t.Error(err)
	}

	valRetrievedCategories := *retrievedCategories // turn retrieved categories from pointer to value

	for i := 0; i < len(insertedCategories); i++ {

		// reflect the category struct to values to iterate over each field
		vInserted := reflect.ValueOf(insertedCategories[i])
		vRetrieved := reflect.ValueOf(valRetrievedCategories[i])

		for i := 0; i < vInserted.NumField(); i++ {
			if vInserted.Field(i).Interface() != vRetrieved.Field(i).Interface() {
				t.Error("Inserted Category is not the retieved one")
			}
		}
	}

}

func TestUpdateCategory(t *testing.T) {
	categ := models.Category{
		ID:       20,
		Title:    "testTitle20",
		ImageURL: "testURL20",
	}
	err := InsertCategory(&categ)
	if err != nil {
		t.Error(err)
	}

	beforeUpdate := models.Category{
		ID:       20,
		Title:    "testTitle20Updated",
		ImageURL: "testURL20",
	}

	err = UpdateCategory(20, &beforeUpdate)
	if err != nil {
		t.Error(err)
	}

	afterUpdate, err := SelectCategoryByID(&beforeUpdate.ID)
	if err != nil {
		t.Error(err)
	}

	// reflect the category struct to values to iterate over each field
	vBeforeUpdate := reflect.ValueOf(beforeUpdate)
	vAfterUpdate := reflect.ValueOf(*afterUpdate)

	for i := 0; i < vBeforeUpdate.NumField(); i++ {
		if vBeforeUpdate.Field(i).Interface() != vAfterUpdate.Field(i).Interface() {
			t.Error("Inserted Category is not the retieved one")
		}
	}
}

func TestDeleteCategory(t *testing.T) {

	categ := models.Category{
		ID:       55,
		Title:    "testTitle55",
		ImageURL: "testImageURL55",
	}

	err := InsertCategory(&categ)
	if err != nil {
		t.Error(err)
	}

	err = DeleteCategory(&categ.ID)
	if err != nil {
		t.Error(err)
	}

	_, err = SelectCategoryByTitle(&categ.Title)
	if err == nil {
		t.Error("the category must not be found")
	} else {
		if err.Error() != "Not Found" {
			t.Error("the category must not be found")
		} else {
		}
	}
}
