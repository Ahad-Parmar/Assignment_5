package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Truck struct {
	Id      int    `json:id`
	DriverName    string `json:drivername`
	CleanerName    string `json:cleanername`
	TruckNo int    `json:truckno`
	
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "order_db"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(ctx *gin.Context) {
		//render only file,  full name with extension is must
		db := dbConn()
		selDB, err := db.Query("SELECT * FROM truck ORDER BY id DESC")
		if err != nil {
			panic(err.Error())
		}
		truck := Truck{}
		res := []truck{}
		for selDB.Next() {
			var id, truckno int
			var drivername, cleanername string
			err = selDB.Scan(&id, &drivername, &cleanername, &truckno)
			if err != nil {
				panic(err.Error())
			}
			truck.Id = id
			truck.DriverName = drivername
			truck.CleanerName = cleanername
			truck.TruckNo = truckno
			res = append(res, truck)
		}
		//var a = "hello words"
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Home Page!!", "a": res})
	})

	r.GET("/add", func(ctx *gin.Context) {
		//render only file, full name with extension is must
		ctx.HTML(http.StatusOK, "add.html", gin.H{"title": "Add truck!!"})
	})

	r.POST("/insert", func(ctx *gin.Context) {
		//render only file, full name with extension is must
		var drivername, cleanername string
		var truckno int
		drivername = ctx.Request.FormValue("drivername")
		cleanername = ctx.Request.FormValue("cleanername")
		truckno = ctx.Request.FormValue("truckno")
		

		db := dbConn()
		insForm, err := db.Prepare("INSERT INTO player (drivername, cleanername, truckno) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, role, matches, age)
		// ctx.HTML(http.StatusOK, "updated.html", gin.H{"title": "Player"})

		selDB, err := db.Query("SELECT * FROM truck ORDER BY id DESC")
		if err != nil {
			panic(err.Error())
		}
		truck := Truck{}
		res := []Truck{}
		for selDB.Next() {
			var id, truckno int
			var drivername, cleanername string
			err = selDB.Scan(&id, &drivername, &cleanername, &truckno)
			if err != nil {
				panic(err.Error())
			}
			truck.Id = id
			truck.DriverName = drivername
			truck.CleanerName = cleanername
			truck.TruckNo = truckno
			res = append(res, truck)
		}
		//var a = "hello words"
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Home Page!!", "a": res})
	})

	r.Run(":8080")

}
