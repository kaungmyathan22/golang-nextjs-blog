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
)

type PostsHandlerImpl struct {
}

func NewPostsHandlerImpl() *PostsHandlerImpl {
	return &PostsHandlerImpl{}
}

func (handler *PostsHandlerImpl) CreatePost(c *gin.Context) {
	logger.Info("CreatePost")
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
func (handler *PostsHandlerImpl) UpdatePost(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "UpdatePost"})
}
func (handler *PostsHandlerImpl) DeletePost(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "DeletePost"})
}
func (handler *PostsHandlerImpl) GetPosts(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "GetPosts"})
}
func (handler *PostsHandlerImpl) GetPost(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "GetPost"})
}
