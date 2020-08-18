package character

import (
	"fmt"
	"net/http"

	"github.com/bountyeq/web/config"
	"github.com/bountyeq/web/model"
	"github.com/gin-gonic/gin"
)

// Get returns a get request
func Get(c *gin.Context) {
	characters := []model.CharacterData{}
	err := config.DB.Find(&characters).Error
	if err != nil {
		c.Error(fmt.Errorf("find characterData: %w", err))
		return
	}
	c.HTML(http.StatusOK, "character/index.tmpl", gin.H{
		"Title":      "BountyEQ",
		"Characters": characters,
	})
}
