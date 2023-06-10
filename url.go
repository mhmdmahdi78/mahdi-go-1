package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type URL struct {
	
	gorm.Model
	LongURL  string `gorm:"unique_index"`
	ShortURL string `gorm:"unique_index"`
}

func main() {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=youruser dbname=yourdb password=ypurpasswprd  sslmode=disable")
	if err != nil {
		log.Fatal("خطا در اتصال پایگاه داده:", err)
	}
	defer db.Close()

	db.AutoMigrate(&URL{})

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "سرویس کوتاه کننده URL"})
	})

	router.POST("/shorten", func(c *gin.Context) {
		
		longURL := c.PostForm("url")

		ShortURL := ShortURL(LongURL)

		c.JSON(http.StatusOK, gin.H{"shortURL": ShortURL})
	})
}
func ShortenURL(LongURL string) string {
	h := md5.New()
	h.Write([]byte(LongURL))
	shortenURL := hex.EncodeToString(h.sum(nil))[:8]

	url := URL{LongURL: longURL, ShortURL: shortenURL}
	db.Create(&url)

	return shortenURL
}
