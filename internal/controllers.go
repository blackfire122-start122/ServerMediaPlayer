package internal

import (
	. "SongServer/pkg"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func ListenAgainSongs(c *gin.Context) {
	loginUser, _ := CheckSessionUser(c.Request)

	if !loginUser {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	// ToDo: need create system for answer againSongs

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
	loginUser, _ := CheckSessionUser(c.Request)

	if !loginUser {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	// ToDo: need create system for answer top songs

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
	loginUser, _ := CheckSessionUser(c.Request)

	if !loginUser {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}

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
	loginUser, user := CheckSessionUser(c.Request)

	if !loginUser {
		c.Writer.WriteHeader(http.StatusUnauthorized)
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

func RegisterUser(c *gin.Context) {
	resp := make(map[string]string)

	var user UserRegister
	bodyBytes, _ := io.ReadAll(c.Request.Body)

	if err := json.Unmarshal(bodyBytes, &user); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Password == "" || user.Username == "" {
		resp["Register"] = "Not all field"

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := Sign(&user); err != nil {
		resp["Register"] = "Error create user"

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp["Register"] = "OK"
	c.JSON(http.StatusOK, resp)
}

func LoginUser(c *gin.Context) {
	resp := make(map[string]string)

	var user UserLogin
	bodyBytes, _ := io.ReadAll(c.Request.Body)

	if err := json.Unmarshal(bodyBytes, &user); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if Login(c.Writer, c.Request, &user) {
		resp["Login"] = "OK"
		c.JSON(http.StatusOK, resp)
	} else {
		resp["Login"] = "error login user"
		c.JSON(http.StatusForbidden, resp)
	}
}

func LogoutUser(c *gin.Context) {
	resp := make(map[string]string)

	if Logout(c.Writer, c.Request) {
		resp["Logout"] = "OK"
		c.JSON(http.StatusOK, resp)
	} else {
		resp["Logout"] = "error logout user"
		c.JSON(http.StatusInternalServerError, resp)
	}
}
