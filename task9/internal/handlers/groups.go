package handlers

import (
	"task9/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func CreateGroup(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		ParentID *int   `json:"parent_id"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, err)
		return
	}

	_, err := db.DB.Exec(
		"INSERT INTO groups(name, parent_id) VALUES ($1,$2)",
		req.Name, req.ParentID,
	)

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.Status(201)
}

func UpdateGroup(c *gin.Context) {
	id := c.Param("id")

	var req struct {
		Name     string `json:"name"`
		ParentID *int   `json:"parent_id"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, err)
		return
	}

	_, err := db.DB.Exec(`
		UPDATE groups
		SET name=$1, parent_id=$2
		WHERE id=$3
	`, req.Name, req.ParentID, id)

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.Status(200)
}

func DeleteGroup(c *gin.Context) {
	id := c.Param("id")

	_, err := db.DB.Exec(`DELETE FROM groups WHERE id=$1`, id)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.Status(204)
}

func GetGroups(c *gin.Context) {
	rows, err := db.DB.Query(`SELECT id, name FROM groups`)
	if err != nil {
		c.JSON(500, err)
		return
	}
	defer rows.Close()

	type GroupResp struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		DirectCount int    `json:"direct_count"`
		TotalCount  int    `json:"total_count"`
	}

	var result []GroupResp

	for rows.Next() {
		var g GroupResp
		rows.Scan(&g.ID, &g.Name)

		// 1. direct count (только в этой группе)
		db.DB.QueryRow(`
			SELECT COUNT(*) FROM people WHERE group_id=$1
		`, g.ID).Scan(&g.DirectCount)

		// 2. total count (с детьми)
		childIDs, err := getAllChildGroups(g.ID)
		if err != nil {
			c.JSON(500, err)
			return
		}

		if len(childIDs) > 0 {
			db.DB.QueryRow(`
				SELECT COUNT(*) FROM people WHERE group_id = ANY($1)
			`, pq.Array(childIDs)).Scan(&g.TotalCount)
		}

		result = append(result, g)
	}

	c.JSON(200, result)
}

func getAllChildGroups(groupID int) ([]int, error) {
	rows, err := db.DB.Query(`
		WITH RECURSIVE tree AS (
			SELECT id FROM groups WHERE id=$1
			UNION ALL
			SELECT g.id FROM groups g
			INNER JOIN tree t ON g.parent_id = t.id
		)
		SELECT id FROM tree
	`, groupID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int

	for rows.Next() {
		var id int
		rows.Scan(&id)
		ids = append(ids, id)
	}

	return ids, nil
}
