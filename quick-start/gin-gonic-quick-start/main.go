package main

//https://gin-gonic.com/docs/examples/graceful-restart-or-stop/

import (
	"fmt"
	"time"
	"sync"
	"io/ioutil"
	"log"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sleepForever(wg* sync.WaitGroup) {
	defer wg.Done()
	for  {
		time.Sleep(time.Second)
		fmt.Println("Sleep")
	}
}

func checkByHttpClient() {
	var http_client = &http.Client{Timeout: 2*time.Second}

	resp, err := http_client.Get("http://127.0.0.1:8080/ping")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	} 

	sb := string(body)
	fmt.Println("Get Response:" + sb)

}

func main() {

	var wg sync.WaitGroup	
	wg.Add(1)
	srv := startGinServer(&wg)	
	time.Sleep(time.Second)
	checkByHttpClient()
	
	wg.Wait()
	
	if err := srv.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		log.Printf("HTTP server Shutdown: %v", err)
	}

}

func startGinServer(wg *sync.WaitGroup) *http.Server {
	defer wg.Done()
    router := gin.New()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	fmt.Println("starting server")
	//r.Run("127.0.0.1:8080") 
	srv := &http.Server{
		Addr : "127.0.0.1:8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %\n", err)
		}
	} ()

	return srv
	
}


