package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	db "github.com/TonyGLL/url-shortener/db/sql"
	"github.com/TonyGLL/url-shortener/util"
	"github.com/gin-gonic/gin"
)

/* GET SITE */
type getSiteRequest struct {
	KEY string `uri:"key" binding:"required"`
}

// Get Site	godoc
// @Summary Get site
// @Description Get site by ID
// @Tags Sites
// @Accept json
// @Produce application/json
// @Param			key	path		getSiteRequest		true	"Site KEY"
// @in header
// @name Authorization
// @Success 200 {object} string
// @Failure		400			{string}	gin.H	"StatusBadRequest"
// @Failure		404			{string}	gin.H	"StatusNotFound"
// @Failure		500			{string}	gin.H	"StatusInternalServerError"
// @Router /sites/{id} [get]
func (s *Server) getSite(ctx *gin.Context) {
	var req getSiteRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	site, err := s.store.GetSite(ctx, req.KEY)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusPermanentRedirect, site)
}

/* GENERATE SITE */
type generateSiteRequest struct {
	SITE string `json:"site" binding:"required"`
}

func (s *Server) generateSite(ctx *gin.Context) {
	var req generateSiteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	salt := time.Now().UnixNano()

	key := util.EncryptAndConvertToBase62(12345, salt, s.config.Secret)

	arg := db.GenerateSiteParams{
		KEY:        key,
		LONG_URL:   req.SITE,
		SALT:       salt,
		EXPIRATION: time.Now().Add(24 * time.Hour),
	}

	err := s.store.GenerateSite(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	data := fmt.Sprintf("%s/%s", s.config.URL, key)

	ctx.JSON(http.StatusOK, gin.H{"data": data})
}
