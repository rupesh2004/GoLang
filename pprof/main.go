package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"time"

	pprofMiddleware "github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func busyWork() {
	for {
		fmt.Println("Processing request...")
		time.Sleep(10000 * time.Millisecond)

		
	start := time.Now()
	for time.Since(start) < 30*time.Second {
		for i := 0; i < 1000000; i++ {
			_ = i * i
		}
	}
	}
}

func profileCPU() {
	f, err := os.Create("cpu_profile.prof")
	if err != nil {
		log.Fatal("Error creating CPU profile file:", err)
	}
	defer f.Close()

	log.Println("Starting CPU profiling...")
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("Error starting CPU profiling:", err)
	}
	defer pprof.StopCPUProfile()

	start := time.Now()
	for time.Since(start) < 30*time.Second {
		for i := 0; i < 1000000; i++ {
			_ = i * i
		}
	}

	log.Println("CPU profiling completed")
}


func main() {
	r := gin.Default()
	pprofMiddleware.Register(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	go busyWork()
	go profileCPU()



	log.Println("Starting Gin server on :8080 with pprof enabled")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
