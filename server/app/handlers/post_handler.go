package handlers

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/database"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/logger"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/models/apis"
	models "github.com/kaungmyathan22/golang-nextjs-blog/app/models/domain"
	"gorm.io/gorm"
)

type PostsHandlerImpl struct {
}

func NewPostsHandlerImpl() *PostsHandlerImpl {
	return &PostsHandlerImpl{}
}

func (handler *PostsHandlerImpl) CreatePost(c *gin.Context) {
	var payload apis.CreatePostPayload
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, apis.GetStatusBadRequestResponse("invalid body payload."))
		return
	}
	_, err := govalidator.ValidateStruct(payload)
	if err != nil {
		logger.Error("couldn't validate the payload..")
		c.JSON(http.StatusBadRequest, apis.GetStatusBadRequestResponse(err.Error()))
		return
	}
	rawUser, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, apis.UnauthorizedResponse)
		return
	}
	user := rawUser.(*models.User)
	logger.Info(strconv.Itoa(int(user.ID)))
	post := models.Post{Title: payload.Title, Content: payload.Content, User: *user}
	result := database.DB.Create(&post)
	if err := result.Error; err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, apis.InternalServerErrorResponse)
		return
	}
	c.JSON(http.StatusOK, apis.GetSuccessResponse(post))
}

func (handler *PostsHandlerImpl) UpdatePost(c *gin.Context) {
	var payload apis.UpdatePostPayload
	postId := c.Param("id")
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, apis.GetStatusBadRequestResponse("invalid body payload."))
		return
	}
	_, err := govalidator.ValidateStruct(payload)
	if err != nil {
		logger.Error("couldn't validate the payload..")
		c.JSON(http.StatusBadRequest, apis.GetStatusBadRequestResponse(err.Error()))
		return
	}
	rawUser, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, apis.UnauthorizedResponse)
		return
	}
	user := rawUser.(*models.User)
	var post *models.Post
	result := database.DB.Where("id = ? AND user_id = ?", postId, user.ID).First(&post)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, apis.GetNotFoundResponse(map[string]string{"message": "post with given id not found."}))
			return
		}
		c.JSON(http.StatusInternalServerError, apis.InternalServerErrorResponse)
		return
	}
	post.Title = payload.Title
	post.Content = payload.Content
	result = database.DB.Save(&post)
	if err := result.Error; err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, apis.InternalServerErrorResponse)
		return
	}
	c.JSON(http.StatusOK, apis.GetSuccessResponse(post))
}

func (handler *PostsHandlerImpl) DeletePost(c *gin.Context) {
	postId := c.Param("id")
	var post *models.Post
	result := database.DB.First(&post, postId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, apis.GetNotFoundResponse(map[string]string{"message": "post with given id not found."}))
			return
		}
		c.JSON(http.StatusInternalServerError, apis.InternalServerErrorResponse)
		return
	}
	rawUser, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, apis.UnauthorizedResponse)
		return
	}
	user := rawUser.(*models.User)
	if post.UserID != user.ID {
		c.JSON(http.StatusForbidden, apis.GetForbiddenResponse(gin.H{"error": "permission denied."}))
		return
	}
	result = database.DB.Delete(post)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, apis.InternalServerErrorResponse)
		return
	}

	c.JSON(http.StatusNoContent, apis.APIResponse{Message: "success", Status: http.StatusNoContent, Data: nil})
}

func (handler *PostsHandlerImpl) GetPosts(c *gin.Context) {
	var posts []models.Post
	var count int64
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	offset := (page - 1) * pageSize
	database.DB.Limit(pageSize).Offset(offset).Find(&posts)
	database.DB.Model(&models.Post{}).Count(&count)
	c.JSON(http.StatusOK, apis.GetStatusAcceptedResponse(gin.H{
		"totalPages": count / int64(pageSize),
		"totalItems": count,
		"page":       page,
		"pageSize":   pageSize,
		"posts":      posts,
	}))
}

func (handler *PostsHandlerImpl) GetPost(c *gin.Context) {
	postId := c.Param("id")
	var post *models.Post
	result := database.DB.First(&post, postId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, apis.GetNotFoundResponse(map[string]string{"message": "post with given id not found."}))
			return
		}
		c.JSON(http.StatusInternalServerError, apis.InternalServerErrorResponse)
		return
	}

	c.JSON(http.StatusOK, apis.GetSuccessResponse(post))
}
