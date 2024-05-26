package handlers

import "github.com/gin-gonic/gin"

type PostsHandlerImpl struct {
}

func NewPostsHandlerImpl() *PostsHandlerImpl {
	return &PostsHandlerImpl{}
}

func CreatePost(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "CreatePost"})
}
func UpdatePost(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "UpdatePost"})
}
func DeletePost(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "DeletePost"})
}
func GetPosts(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "GetPosts"})
}
func GetPost(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "GetPost"})
}
