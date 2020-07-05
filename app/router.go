package app

import (
	"github.com/gin-gonic/gin"
	"goweb/app/handlers/blog"
	"goweb/pkg/hot"
	resp "goweb/pkg/response"
	"time"
)

// 注册路由
func registerRoute(r *gin.Engine) {
	// 跨域处理
	r.Use(Cors())

	// 根目录
	r.GET("/", func(c *gin.Context) {
		resp.String(c, 200, time.Now().Format(hot.GetTimeCommonFormat()))
	})

	// blog
	blogRouter(r.Group("blog"))
}

// 博客路由
func blogRouter(r *gin.RouterGroup) {
	// 分类列表
	r.GET("ls", blog.Ls)
	// 文章列表
	r.GET("posts", blog.Posts)
	// 文章内容
	r.GET("posts.detail", blog.PostsDetail)
}
