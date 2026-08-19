package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"players/services"
)

type PlayerHandler struct {
	Service *services.PlayerService
}

func NewPlayerHandler(service *services.PlayerService) *PlayerHandler {
	return &PlayerHandler{
		Service: service,
	}
}

func (h *PlayerHandler) GetAllPlayers(c *gin.Context) {
	page := 1
	limit := 20

	// Read page
	if value := c.Query("page"); value != "" {
		parsedPage, err := strconv.Atoi(value)

		if err != nil || parsedPage < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid page",
			})
			return
		}

		page = parsedPage
	}

	// Read limit
	if value := c.Query("limit"); value != "" {
		parsedLimit, err := strconv.Atoi(value)

		if err != nil || parsedLimit < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid limit",
			})
			return
		}

		limit = parsedLimit
	}

	// Read search name
	name := strings.TrimSpace(c.Query("name"))

	// ----------------------------------------------------
	// SEARCH
	// ----------------------------------------------------
	if name != "" {
		offset := (page - 1) * limit

		players, err := h.Service.SearchPlayers(
			c.Request.Context(),
			name,
			int32(limit),
			int32(offset),
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to search players",
			})
			return
		}

		total, err := h.Service.SearchPlayersCount(
			c.Request.Context(),
			name,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to get search players count",
			})
			return
		}

		totalPages := 1

		if total > 0 {
			totalPages = (int(total) + limit - 1) / limit
		}

		c.JSON(http.StatusOK, gin.H{
			"name":       name,
			"page":       page,
			"limit":      limit,
			"offset":     offset,
			"total":      total,
			"totalPages": totalPages,
			"players":    players,
		})

		return
	}

	// ----------------------------------------------------
	// NORMAL PAGINATION
	// ----------------------------------------------------

	offset := (page - 1) * limit

	players, err := h.Service.GetPlayers(
		c.Request.Context(),
		int32(limit),
		int32(offset),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get players",
		})
		return
	}

	total, err := h.Service.GetPlayersCount(
		c.Request.Context(),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get players count",
		})
		return
	}

	totalPages := 1

	if total > 0 {
		totalPages = (int(total) + limit - 1) / limit
	}

	c.JSON(http.StatusOK, gin.H{
		"page":       page,
		"limit":      limit,
		"offset":     offset,
		"total":      total,
		"totalPages": totalPages,
		"players":    players,
	})
}

func (h *PlayerHandler) GetPlayerByID(c *gin.Context) {
	playerID := c.Param("id")

	player, err := h.Service.GetPlayerByID(
		c.Request.Context(),
		playerID,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Player not found",
		})
		return
	}

	c.JSON(http.StatusOK, player)
}