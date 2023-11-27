package boot

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mooha76/Puy_me_kofe_API/config"
	"github.com/mooha76/Puy_me_kofe_API/global"
)

func test(cfg *config.Sysytem) {

	// Create a new Gin router
	// Create a new Gin router
	router := gin.Default()

	// Define your routes and handlers
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})

	// Run the server on port 8080
	router.Run(cfg.Port)
	global.MainLogger.Info("System started, listening " + cfg.Port + "...")

}

func Boot(cfg *config.Sysytem) {
	server := &http.Server{
		Addr: fmt.Sprintf(cfg.Port),
		//Handler: Router(),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("【【Kofe】】listen: %s\n", err)
		}
	}()
	logo(cfg.Port)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("【【Kofe】】start to exit...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("【【Kofe】】force exit：", err)
	}

	log.Println("【Kofe】 exit complete！")
}

func logo(port string) {
	fmt.Println("Welcome to Kofe API !")
	fmt.Println("Github: github.com/mooha76/Puy_me_kofe_API ")
	fmt.Println("Expecting Your Star!")
	fmt.Println("System started, listening " + port + "...")
	//global.MainLogger.Info("System started, listening " + port + "...")
}
