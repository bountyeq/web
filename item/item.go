package item

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/bountyeq/web/config"
	"github.com/bountyeq/web/model"
	"github.com/bountyeq/web/sanitize"
	"github.com/bountyeq/web/site"
	"github.com/gin-gonic/gin"
)

// List returns a list of creatures
func List(c *gin.Context) {
	var err error
	search, ok := c.Get("search")
	fmt.Println("search", search)
	if ok {
		fmt.Println("search", search)
	}
	id := c.Param("id")
	if len(id) > 0 {
		err = getItem(c, id)
		if err != nil {
			c.Error(err)
			return
		}
		return
	}
	page := 1
	pageStr, ok := c.GetQuery("page")
	if ok {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
	}
	if page < 1 {
		page = 1
	}
	limit := 10

	items := []model.Items{}
	err = config.ContentDB.Limit(limit).Offset(limit*page - 1).Find(&items).Error
	if err != nil {
		c.Error(fmt.Errorf("find Items: %w", err))
		return
	}
	site := site.Fetch()
	site.Title = "Item List"
	previousDisplay := fmt.Sprintf("%d", page-1)
	previous := fmt.Sprintf("/item?page=%s", previousDisplay)
	if page == 1 {
		previous = ""
		previousDisplay = ""
	}
	site.Page.Generate("/item", "Start", previous, previousDisplay, fmt.Sprintf("%d", page), fmt.Sprintf("/item?page=%d", page+1), fmt.Sprintf("%d", page+1))
	c.HTML(http.StatusOK, "item_list.tmpl", gin.H{
		"Title": "BountyEQ",
		"Items": items,
		"Site":  site,
	})
}

func getItem(c *gin.Context, strID string) error {
	description := ""
	id, err := strconv.Atoi(strID)
	if err != nil {
		return fmt.Errorf("parse id %s: %w", strID, err)
	}
	if id == 0 {
		return fmt.Errorf("invalid id %s", strID)
	}

	item := &model.Items{}

	config.ContentDB.Where("id = ?", id).First(item)
	if err != nil {
		return fmt.Errorf("could not find ITEM ID %d: %w", id, err)
	}

	description += fmt.Sprintf("<p>%s is an item that I'll add more intel later about.</p>", item.Name)

	site := site.Fetch()
	site.Title = item.Name
	c.HTML(http.StatusOK, "item_view.tmpl", gin.H{
		"Title":       sanitize.CleanName(item.Name),
		"Item":        item,
		"Description": template.HTML(description),
		"Site":        site,
	})
	return nil
}
