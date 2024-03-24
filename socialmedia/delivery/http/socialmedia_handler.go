package delivery

import (
	"mygram-api/domain"
	"mygram-api/helpers"
	"mygram-api/socialmedia/delivery/http/middleware"
	"mygram-api/socialmedia/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type socialMediaHandler struct {
	socialMediaUseCase domain.SocialMediaUseCase
}

func NewSocialMediaHandler(routers *gin.Engine, socialMediaUseCase domain.SocialMediaUseCase) {
	handler := &socialMediaHandler{socialMediaUseCase}

	router := routers.Group("/socialmedias")
	{
		router.Use(middleware.Authentication())
		router.GET("", handler.Fetch)
		router.POST("", handler.Store)
		router.PUT("/:socialMediaId", middleware.Authorization(handler.socialMediaUseCase), handler.Update)
		router.DELETE("/:socialMediaId", middleware.Authorization(handler.socialMediaUseCase), handler.Delete)
	}
}

// Fetch godoc
// @Summary    	Fetch all social media
// @Description	Get all social media with authentication user
// @Tags        socialmedias
// @Accept      json
// @Produce     json
// @Success     200	{object}	utils.ResponseDataFetchedSocialMedia
// @Failure     400	{object}	utils.ResponseMessage
// @Failure     401	{object}	utils.ResponseMessage
// @Security    Bearer
// @Router      /socialmedias	[get]
func (handler *socialMediaHandler) Fetch(ctx *gin.Context) {
	var (
		socialMedias []domain.SocialMedia
		err          error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = handler.socialMediaUseCase.Fetch(ctx.Request.Context(), &socialMedias, userID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data: utils.FetchedSocialMedia{
			SocialMedias: socialMedias,
		},
	})
}

// Store godoc
// @Summary    	Add a social media
// @Description	Create and store a social media with authentication user
// @Tags        socialmedias
// @Accept      json
// @Produce     json
// @Param       json	body			utils.AddSocialMedia true  "Add Social Media"
// @Success     201		{object}  utils.ResponseDataAddedSocialMedia
// @Failure     400		{object}	utils.ResponseMessage
// @Failure     401		{object}	utils.ResponseMessage
// @Security    Bearer
// @Router      /socialmedias		[post]
func (handler *socialMediaHandler) Store(ctx *gin.Context) {
	var (
		socialMedia domain.SocialMedia
		err         error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	socialMedia.UserID = userID

	if err = handler.socialMediaUseCase.Store(ctx.Request.Context(), &socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, helpers.ResponseData{
		Status: "success",
		Data: utils.AddedSocialMedia{
			ID:             socialMedia.ID,
			UserID:         socialMedia.UserID,
			Name:           socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
			CreatedAt:      socialMedia.CreatedAt,
		},
	})
}

// Update godoc
// @Summary     Update a social media
// @Description	Update a social media by id with authentication user
// @Tags        socialmedias
// @Accept      json
// @Produce     json
// @Param       id		path      string	true	"SocialMedia ID"
// @Param				json	body			utils.UpdateSocialMedia	true	"Update Social Media"
// @Success     200		{object}	utils.ResponseDataUpdatedSocialMedia
// @Failure     400		{object}	utils.ResponseMessage
// @Failure     401		{object}	utils.ResponseMessage
// @Failure     404		{object}	utils.ResponseMessage
// @Security    Bearer
// @Router      /socialmedias/{id} [put]
func (handler *socialMediaHandler) Update(ctx *gin.Context) {
	var (
		socialMedia domain.SocialMedia
		err         error
	)

	socialMediaID := ctx.Param("socialMediaId")
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	updatedSocialMedia := domain.SocialMedia{
		UserID:         userID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
	}

	if socialMedia, err = handler.socialMediaUseCase.Update(ctx.Request.Context(), updatedSocialMedia, socialMediaID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, utils.UpdatedSocialMedia{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID:         socialMedia.UserID,
		UpdatedAt:      socialMedia.UpdatedAt,
	})
}

// Delete godoc
// @Summary     Delete a social media
// @Description	Delete a social media by id with authentication user
// @Tags        socialmedias
// @Accept      json
// @Produce     json
// @Param       id   path     string  true  "SocialMedia ID"
// @Success     200  {object}	utils.ResponseMessageDeletedSocialMedia
// @Failure     400  {object}	utils.ResponseMessage
// @Failure     401  {object}	utils.ResponseMessage
// @Failure     404  {object}	utils.ResponseMessage
// @Security    Bearer
// @Router      /socialmedias/{id} [delete]
func (handler *socialMediaHandler) Delete(ctx *gin.Context) {
	socialMediaID := ctx.Param("socialMediaId")

	if err := handler.socialMediaUseCase.Delete(ctx.Request.Context(), socialMediaID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "your social media has been successfully deleted",
	})
}
