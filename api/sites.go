package api

import (
	"database/sql"
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

	data, err := s.store.GetSite(ctx, req.KEY)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	info := db.CountSearchParams{
		SiteID:    data.ID,
		IpAddress: ctx.ClientIP(),
		Browser:   ctx.GetHeader("User-Agent"),
	}
	err = s.store.CountSearch(ctx, info)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := generateSiteResponse{
		ID:        data.ID,
		Url:       data.LongURL,
		ShortCode: data.Key,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

/* GENERATE SITE */
type generateSiteRequest struct {
	URL string `json:"url" binding:"required"`
}
type generateSiteResponse struct {
	ID        int       `json:"id"`
	Url       string    `json:"url"`
	ShortCode string    `json:"shortCode"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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
		LONG_URL:   req.URL,
		SALT:       salt,
		EXPIRATION: time.Now().Add(24 * time.Hour),
	}

	data, err := s.store.GenerateSite(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response := generateSiteResponse{
		ID:        data.ID,
		Url:       data.LongURL,
		ShortCode: key,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (s *Server) getSiteStats(ctx *gin.Context) {

}
