package npc

import (
	"fmt"
	"net/http"

	"github.com/bountyeq/web/config"
	"github.com/bountyeq/web/model"
	"github.com/gin-gonic/gin"
)

// Get returns a get request
func Get(c *gin.Context) {
	npcs := []model.NpcTypes{}
	err := config.DB.Limit(10).Find(&npcs).Error
	if err != nil {
		c.Error(fmt.Errorf("find NpcTypes: %w", err))
		return
	}
	c.HTML(http.StatusOK, "npc/index.tmpl", gin.H{
		"Title": "BountyEQ",
		"Npcs":  npcs,
	})
}
