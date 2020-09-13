package search

import (
	"fmt"
	"net/http"

	"github.com/bountyeq/web/config"
	"github.com/bountyeq/web/model"
	"github.com/bountyeq/web/site"
	"github.com/gin-gonic/gin"
)

// List returns a list of creatures
func List(c *gin.Context) {
	var err error
	type Result struct {
		Name string
		URL  string
	}
	results := []Result{}
	search := c.Param("find")
	npcs := []model.NpcTypes{}
	err = config.ContentDB.Limit(10).Find(&npcs).Error
	if err != nil {
		c.Error(fmt.Errorf("find npcs: %w", err))
		return
	}
	for _, npc := range npcs {
		results = append(results, Result{
			Name: npc.Name,
			URL:  fmt.Sprintf("/bestiary/%d", npc.ID),
		})
	}
	site := site.Fetch()
	site.Title = "Search Results"

	c.HTML(http.StatusOK, "search_list.tmpl", gin.H{
		"Title":   "BountyEQ",
		"Results": results,
		"Site":    site,
		"Search":  search,
	})
}
