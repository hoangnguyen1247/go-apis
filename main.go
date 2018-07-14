package go_apis

import (
	"context"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	log "github.com/sirupsen/logrus"

	"github.com/hoangnguyen1247/go-apis/controller/index"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Starting server.....")

	port := flag.String("port", "8080", "http serve port, default:8080")
	templatesPath := flag.String("template", "./template", "path of the template files")
	flag.Parse()

	r := gin.New()

	r.Use(gin.Recovery())

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	r.LoadHTMLGlob(*templatesPath + "/**/*")

	// handle 404
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "util/error_page.template", gin.H{
			"error_code": 404,
			"retry_link": "/",
		})
	})

	// setup routes for macrokiosk
	indexController, err := index.New()
	if err == nil {
		indexController.SetRoutes(r)
		//defer indexController.close()
	}

	// make grateful shutdown
	srv := &http.Server{
		Addr:    ":" + *port,
		Handler: r,
	}
	log.Println("Starting server ....")

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}