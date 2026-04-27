package handlers

import (
	"task9/internal/db"

	"github.com/gin-gonic/gin"
)

func CreatePerson(c *gin.Context) {
	var req struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		BirthYear int    `json:"birth_year"`
		GroupID   int    `json:"group_id"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, err)
		return
	}

	_, err := db.DB.Exec(`
		INSERT INTO people(first_name,last_name,birth_year,group_id)
		VALUES ($1,$2,$3,$4)
	`, req.FirstName, req.LastName, req.BirthYear, req.GroupID)

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.Status(201)
}

func UpdatePerson(c *gin.Context) {
	id := c.Param("id")

	var req struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		BirthYear int    `json:"birth_year"`
		GroupID   int    `json:"group_id"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, err)
		return
	}

	_, err := db.DB.Exec(`
		UPDATE people
		SET first_name=$1, last_name=$2, birth_year=$3, group_id=$4
		WHERE id=$5
	`, req.FirstName, req.LastName, req.BirthYear, req.GroupID, id)

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.Status(200)
}

func DeletePerson(c *gin.Context) {
	id := c.Param("id")

	_, err := db.DB.Exec(`DELETE FROM people WHERE id=$1`, id)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.Status(204)
}

func GetPeopleByGroup(c *gin.Context) {
	groupID := c.Param("id")
	mode := c.Query("mode")

	var query string

	if mode == "all" {
		query = `
		WITH RECURSIVE tree AS (
			SELECT id FROM groups WHERE id=$1
			UNION ALL
			SELECT g.id FROM groups g
			INNER JOIN tree t ON g.parent_id = t.id
		)
		SELECT p.id, p.first_name, p.last_name, p.birth_year, p.group_id
		FROM people p
		WHERE p.group_id IN (SELECT id FROM tree)
		`
	} else {
		query = `
		SELECT id, first_name, last_name, birth_year, group_id
		FROM people WHERE group_id=$1
		`
	}

	rows, err := db.DB.Query(query, groupID)
	if err != nil {
		c.JSON(500, err)
		return
	}
	defer rows.Close()

	var result []map[string]any

	for rows.Next() {
		var id, birthYear, groupID int
		var fn, ln string

		rows.Scan(&id, &fn, &ln, &birthYear, &groupID)

		result = append(result, gin.H{
			"id":         id,
			"first_name": fn,
			"last_name":  ln,
			"birth_year": birthYear,
			"group_id":   groupID,
		})
	}

	c.JSON(200, result)
}
