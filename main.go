package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	//"github.com/bountyeq/web/dao"
	"github.com/bountyeq/web/bestiary"
	"github.com/bountyeq/web/character"
	"github.com/bountyeq/web/config"
	"github.com/bountyeq/web/site"
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/multitemplate"
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

	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Database))
	if err != nil {
		return fmt.Errorf("gorm open: %w", err)
	}
	config.DB = db
	db.LogMode(true)

	contentdb, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.ContentDatabase.Username, cfg.ContentDatabase.Password, cfg.ContentDatabase.Host, cfg.ContentDatabase.Database))
	if err != nil {
		return fmt.Errorf("gorm open contentDB: %w", err)
	}
	config.ContentDB = contentdb
	contentdb.LogMode(true)

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

	r.Static("assets", "assets")
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})
	//api.ConfigGinRouter(r)
	r.Use(gin.Recovery())
	//r.LoadHTMLGlob("templates/**/*")

	r.HTMLRender = loadTemplates("templates/")

	//root pages
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Title": "BountyEQ",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.tmpl", gin.H{
			"Title": "BountyEQ",
		})
	})

	forum := r.Group("/forum")
	forum.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "forum.tmpl", gin.H{
			"Title": "Forum",
			"Site":  site.Fetch(),
		})
	})

	characterGroup := r.Group("/character")
	characterGroup.GET("/", character.Get)
	bestiaryGroup := r.Group("/bestiary")
	bestiaryGroup.GET("/", bestiary.List)
	bestiaryGroup.GET("/:id", bestiary.List)

	databaseGroup := r.Group("/database")
	databaseGroup.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "database_list.tmpl", gin.H{
			"Title": "Databases",
		})
	})

	bountyGroup := r.Group("/bounty")
	bountyGroup.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "bounty_list.tmpl", gin.H{
			"Title": "Bounty Board",
		})
	})

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

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/content/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {

		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		fmt.Println("include", include, files)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
