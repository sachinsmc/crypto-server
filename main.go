package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/currency/:symbol", func(c *gin.Context) {
		path := c.Param("symbol")
		var host string
		if path != "all" {
			host = "https://api.hitbtc.com/api/2/public/ticker/" + path
			resp, err := http.Get(host)
			if err != nil {
				log.Fatalln(err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			var result map[string]interface{}
			json.Unmarshal([]byte(body), &result)
			c.JSON(200, result)
		} else {
			host := "https://api.hitbtc.com/api/2/public/symbol/"
			resp, err := http.Get(host)
			if err != nil {
				log.Fatalln(err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			var result []map[string]interface{}
			json.Unmarshal([]byte(body), &result)

			c.JSON(200, gin.H{"currencies": result})
		}
	})
	r.Run()
}
