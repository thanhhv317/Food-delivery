package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type SQLModel struct {
	ID        int       `gorm:"column:id;"`
	CreatedAt time.Time `gorm:"column:created_at;"`
	UpdatedAt time.Time `gorm:"column:updated_at;"`
	Status    int       `gorm:"column:status;"`
}

type Note struct {
	SQLModel
	Name       string `gorm:"column:title;"`
	CategoryId int    `gorm:"column:category_id;"`
}

func (Note) TableName() string { return "notes" }

type NoteUpdate struct {
	Name       *string `gorm:"column:title;"`
	CategoryId *int    `gorm:"column:category_id;"`
}

func (NoteUpdate) TableName() string { return Note{}.TableName() }

func main() {
	fmt.Println("Hello world")

	dsn := "root:123456789@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug() // show query

	if err != nil {
		log.Fatal(err)
	}

	log.Println(db)

	// Create
	//n := &Note{
	//	Name: "Note 4",
	//}
	//
	//if err := db.Create(n).Error; err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println(n)

	// Get first

	var myNote Note

	if err := db.
		Where("id = ?", 2).
		First(&myNote).Error; err != nil {
		log.Fatal(err)
	}

	log.Println(myNote)

	// List
	var listNote []Note

	if err := db.Find(&listNote).Error; err != nil {
		log.Fatal(err)
	}

	log.Println(listNote)

	// update

	myNote.Name = "note updated"

	if err := db.Where("id = ?", 3).Updates(myNote).Error; err != nil {
		log.Fatal(err)
	}

	log.Println(myNote)

	// update name to ""

	emptyString := ""

	if err := db.Where("id = ?", 2).Updates(NoteUpdate{Name: &emptyString}).Error; err != nil {
		log.Fatal(err)
	}

	// Delete

	//if err := db.Table(Note{}.TableName()).Where("id = ?", 1).Delete(nil).Error; err != nil {
	//	log.Fatal(err)
	//}

}
