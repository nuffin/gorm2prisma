package main

import (
	"os"

	"github.com/nuffin/gorm2prisma/lib"
)

// type User struct {
// 	ID        uint      `gorm:"primaryKey"`
// 	UserName  string    `gorm:"column:user_name;size:100"`
// 	Email     string    `gorm:"unique"`
// 	Age       int       `gorm:"column:user_age"`
// 	Active    bool      `gorm:"column:is_active"`
// 	CreatedAt time.Time `gorm:"column:created_at"`
// 	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;"`
// }

// type UserDocument struct {
// 	ID   uint `gorm:"primaryKey"`
// 	User User `gorm:"foreignKey:UserID"`
// }

// User model
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"uniqueIndex;size:100;not null"`
	Profile   Profile
	Posts     []Post    // One-to-Many relationship (User has many Posts)
	Addresses []Address `gorm:"many2many:user_addresses"` // Many-to-Many relationship
}

func (User) TableName() string {
	return "user"
}

// Profile model (One-to-One relationship with User)
type Profile struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `gorm:"uniqueIndex"` // Foreign key to User
	Bio    string `gorm:"size:255"`
}

// Post model (Belongs to User)
type Post struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"size:200;not null"`
	Body   string `gorm:"type:text;not null"`
	UserID uint   `gorm:"not null"` // Foreign key to User
}

// Address model (Many-to-Many relationship)
type Address struct {
	ID      uint   `gorm:"primaryKey"`
	Street  string `gorm:"size:100;not null"`
	City    string `gorm:"size:50;not null"`
	ZipCode string `gorm:"size:20;not null"`
	Users   []User `gorm:"many2many:user_addresses"` // Many-to-Many relationship
}

func main() {
	// Generate the Prisma schema
	// schema := generatePrismaSchema(User{}, Profile{}, Post{}, Address{})

	// Instantiate PrismaSchemaGenerator
	generator := lib.PrismaSchemaGenerator{}
	// Generate the Prisma schema
	schema := generator.Generate(User{}, Profile{}, Post{}, Address{})

	// Write to schema.prisma file
	file, err := os.Create("schema.prisma")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(schema)
	if err != nil {
		panic(err)
	}
}
