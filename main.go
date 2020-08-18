package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	//"github.com/bountyeq/web/dao"
	"github.com/bountyeq/web/character"
	"github.com/bountyeq/web/config"
	"github.com/bountyeq/web/npc"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog"
	"github.com/xackery/log"
)

var (
	rxURL = regexp.MustCompile(`^/regexp\d*`)
)

func main() {
	log := log.New()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	ctx, cancel := context.WithCancel(context.Background())
	err := run(ctx, cancel)
	cancel()
	if err != nil {
		log.Error().Err(err).Msg("main exited with errors")
	}
	log.Info().Msg("main exited cleanly")
}

func run(ctx context.Context, cancel context.CancelFunc) error {
	log := log.New()

	db, err := gorm.Open("mysql", "eqemu:eqemupass@tcp(127.0.0.1:3306)/eqemu")
	if err != nil {
		return fmt.Errorf("gorm open: %w", err)
	}
	config.DB = db
	db.LogMode(true)
	//dao.DB = db

	/*dao.Logger = func(ctx context.Context, sql string) {
		fmt.Printf("SQL: %s\n", sql)
	}*/

	c := make(chan os.Signal, 1)
	r := gin.Default()
	//gin.DefaultWriter = log
	r.Use(logger.SetLogger(logger.Config{
		Logger:         &log,
		UTC:            true,
		SkipPath:       []string{"/skip"},
		SkipPathRegexp: rxURL,
	}))
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})
	//api.ConfigGinRouter(r)
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Title": "BountyEQ",
		})
	})

	forum := r.Group("/forum")
	forum.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "forum/index.tmpl", gin.H{
			"Title": "BountyEQ",
		})
	})
	characterGroup := r.Group("/character")
	characterGroup.GET("/", character.Get)
	npcGroup := r.Group("/npc")
	npcGroup.GET("/", npc.Get)

	/*authorized := r.Group("/account")
	authorized.Use(AuthRequired()) {
		authorauthorized.POST("/login", loginEndpoint)
	}*/
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	s := <-c
	fmt.Println("Got signal:", s)
	cancel()

	return nil
}
