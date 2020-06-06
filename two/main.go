package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {

	db, err := gorm.Open("mysql",  os.Getenv("DB_USER")+ ":"+ os.Getenv("DB_PASS") +"@/" + os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local")
	if(err != nil){
		panic("Error al conectar con la base de datos")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	//db.Create(&Product{Code: "Sargento Robot", Price: 1000})

	//db.Create(&Product{Code: "Hola Mundo", Price: 2000})

	var product Product
	db.First(&product, 5)
	fmt.Println(product) // Hola Mundo

	var product2 Product
	db.First(&product2, "code = ? AND id = ?", "Sargento Robot", 1)
	fmt.Println(product2)


	db.Model(&product).Update("Price", 50000)

	db.Delete(&product)
}