package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWrit	er, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}

	w.Write([]byte("Hello World!"))
}*/

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.Run(":8080")
}

func handler(c *gin.Context) {
	c.string(http.StatusOK, "Hello World!")
}
