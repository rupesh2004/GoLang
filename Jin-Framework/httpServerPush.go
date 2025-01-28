package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

var Html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	router := gin.Default()

	// Serve static files from the "assets" directory
	router.Static("/assets", "./assets")
	router.SetHTMLTemplate(Html)

	router.GET("/", func(c *gin.Context) {
		// HTTP/2 server push
		if pusher := c.Writer.Pusher(); pusher != nil {
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		// Render the HTML template
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	// Start HTTPS server
	router.RunTLS(":3000", "./testdata/server.pem", "./testdata/server.key")
}
