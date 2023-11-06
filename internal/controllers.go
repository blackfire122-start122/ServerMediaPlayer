package internal

import (
	. "SongServer/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListenAgainSongs(c *gin.Context) {
	var songs []Song

	if err := DB.Limit(20).Find(&songs).Error; err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	resp := make([]map[string]string, len(songs))

	for i, song := range songs {
		item := make(map[string]string)
		item["id"] = strconv.FormatUint(song.Id, 10)
		item["name"] = song.Name
		item["image"] = song.Image
		item["fileSong"] = song.File
		resp[i] = item
	}

	c.JSON(http.StatusOK, resp)
}

func GetImageSong(c *gin.Context) {
	c.File("media/songImage/" + c.Param("filename"))
}

func StreamSong(c *gin.Context) {
	fmt.Print("here")
	c.Writer.Header().Add("Content-Disposition", "inline; filename=audio.mp3")
	c.File("media/songs/" + c.Param("filename"))
}
