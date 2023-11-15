package internal

import (
	. "SongServer/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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

func getTopsSongs(c *gin.Context) {
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

func GetImageUser(c *gin.Context) {
	c.File("media/users/" + c.Param("filename"))
}

func StreamSong(c *gin.Context) {
	c.Writer.Header().Add("Content-Disposition", "inline; filename=audio.mp3")
	c.File("media/songs/" + c.Param("filename"))
}

func FindSong(c *gin.Context) {
	var songs []Song

	if err := DB.Where("lower(Name) LIKE ?", "%"+strings.ToLower(c.Param("findStr"))+"%").Limit(20).Find(&songs).Error; err != nil {
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

func GetUser(c *gin.Context) {
	var user User

	if err := DB.First(&user).Error; err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	resp := make(map[string]string)
	resp["id"] = strconv.FormatUint(user.Id, 10)
	resp["username"] = user.Username
	resp["image"] = user.Image
	resp["email"] = user.Email
	resp["phone"] = user.Phone

	c.JSON(http.StatusOK, resp)
}
