package main

import (
 "database/sql"
 "fmt"
 "log"
 "net/http"

 "github.com/gin-gonic/gin"
 _ "github.com/mattn/go-sqlite3"
)


type URL struct {
 ID     int      json:"id"
 Name   string   json:"name"
 URL    string   json:"url"
}

func main() {

 db, err := sql.Open("sqlite3", "urls.db")
 if err != nil {
  log.Fatal(err)
 }
 defer db.Close()

 
 createTable(db)

 
 router := gin.Default()

 router.POST("/urls", func(c *gin.Context) {
  var url URL
  if err := c.ShouldBindJSON(&url); err != nil {
   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
   return
  }


  insertURL(db, &url)

  c.JSON(http.StatusOK, gin.H{"message": "URL created successfully"})
 })

 router.GET("/urls/:id", func(c *gin.Context) {
  id := c.Param("id")

 
  url, err := getURLByID(db, id)
  if err != nil {
   c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
   return
  }

  c.JSON(http.StatusOK, url)
 })

 router.Run(":8080")
}


func createTable(db *sql.DB) {
 createTableSQL := `
  CREATE TABLE IF NOT EXISTS urls (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   name TEXT NOT NULL,
   url TEXT NOT NULL
  );
 `

 _, err := db.Exec(createTableSQL)
 if err != nil {
  log.Fatal(err)
 }
}


func insertURL(db *sql.DB, url *URL) {
 insertURLSQL := `
  INSERT INTO urls (name, url)
  VALUES (?, ?);
 `

 _, err := db.Exec(insertURLSQL, url.Name, url.URL)
 if err != nil {
  log.Fatal(err)
 }
}

// جستجوی رکورد با استفاده از شناسه
func getURLByID(db *sql.DB, id string) (*URL, error) {
 getURLByIDSQL := `
  SELECT id, name, url FROM urls
  WHERE id = ?;
 `

 var url URL
 err := db.QueryRow(getURLByIDSQL, id).Scan(&url.ID, &url.Name, &url.URL)
 if err != nil {
  return nil, err
 }

 return &url, nil
}