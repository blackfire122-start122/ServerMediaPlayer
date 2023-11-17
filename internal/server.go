package internal

import (
	"github.com/gin-gonic/gin"
)

func SetRouters(router *gin.Engine) {
	router.GET("api/getListenAgainSongs", ListenAgainSongs)
	router.GET("api/getTops", getTopsSongs)
	router.GET("api/songs/image/:filename", GetImageSong)
	router.GET("api/user/image/:filename", GetImageUser)
	router.GET("streamSong/:filename", StreamSong)
	router.GET("api/findSong/:findStr", FindSong)
	router.GET("ws/:roomId", handleConnections)
	router.GET("api/user/getUser", GetUser)
	router.POST("api/user/register", RegisterUser)
	router.POST("api/user/login", LoginUser)
	router.GET("api/user/logout", LogoutUser)
}
