package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
)

type Data struct {
    People_id int
    Address   string 
    Birth string 
    Gender string
    Job		string
    Status 	int
}

func main() {
  db, err :=  gorm.Open("postgres", "host=localhost user=postgres dbname=go sslmode=disable password=tekanenter")
  if err != nil {
    panic("failed to connect database")
  }
  	defer db.Close()
    
    //create table
    if err := db.CreateTable(&Data{}).Error; err != nil {
    	fmt.Sprintf("No error should happen when create table, but got %+v", err)
    }

	// Seeding tables:
	var users []Data = []Data{
		Data{People_id: 1, Address: "Jl. Lapan", Birth: "1980-09-09", Gender: "l", Job: "Staff", Status: 1},
		Data{People_id: 2, Address: "Jl. Sembilan", Birth: "1980-09-09", Gender: "l", Job: "Staff", Status: 1},
		Data{People_id: 3, Address: "Jl. Sepuluh", Birth: "1980-09-09", Gender: "l", Job: "Staff", Status: 1},
		Data{People_id: 4, Address: "Jl. Sebelas", Birth: "1980-09-09", Gender: "l", Job: "Staff", Status: 1},
		Data{People_id: 5, Address: "Jl. DuaBelas", Birth: "1980-09-09", Gender: "l", Job: "Staff", Status: 1},
		Data{People_id: 6, Address: "Jl. TigaBelas", Birth: "1980-09-09", Gender: "l", Job: "Staff", Status: 1},
		Data{People_id: 7, Address: "Jl. EmpatBelas", Birth: "1980-09-09", Gender: "l", Job: "Staff", Status: 1},
		Data{People_id: 8, Address: "Jl. LimaBelas", Birth: "1980-09-09", Gender: "l", Job: "Staff", Status: 1},
		Data{People_id: 9, Address: "Jl. EnamBelas", Birth: "1980-09-09", Gender: "l", Job: "Staff", Status: 1},
		Data{People_id: 10, Address: "Jl. TujuhBelas", Birth: "1980-09-09", Gender: "l", Job: "Staff", Status: 1},
		Data{People_id: 11, Address: "Jl. LapanBelas", Birth: "1980-09-09", Gender: "l", Job: "Staff", Status: 1},
		Data{People_id: 12, Address: "Jl. SembilanBelas", Birth: "1980-09-09", Gender: "l", Job: "Staff", Status: 1},
		Data{People_id: 13, Address: "Jl. Sepuluh", Birth: "1980-09-09", Gender: "l", Job: "Staff", Status: 1},
	}

	for _, user := range users {
	  db.Create(&user)
	}

}