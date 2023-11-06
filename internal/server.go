package internal

import (
	"github.com/gin-gonic/gin"
)

func SetRouters(router *gin.Engine) {
	router.GET("api/getListenAgainSongs", ListenAgainSongs)
	router.GET("api/songs/image/:filename", GetImageSong)
	router.GET("streamSong/:filename", StreamSong)
}
