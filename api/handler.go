package api

import "github.com/labstack/echo"
import "github.com/gtlservice/gutils/logger"
import "github.com/gtlservice/gutils/system"

import (
	"io/ioutil"
	"net/http"
	"time"
)

type UserData struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type UserResponse struct {
	UserId   string `json:"userid"`
	Name     string `json:"name"`
	CreateAt int64  `json:"createat"`
}

func (s *ApiServer) getCatalogHandleFunc(c *echo.Context) error {

	request := c.Request()
	defer request.Body.Close()
	buf, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "badrequest...")
	}
	logger.INFO("[#apil#] catalog handle func:%s path:%s buf:%s", request.Method, request.RequestURI, string(buf))
	return c.JSON(http.StatusOK, "calalog")
}

func (s *ApiServer) postUserHandleFunc(c *echo.Context) error {

	user := &UserData{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "userdata invalid.")
	}
	return c.JSON(http.StatusOK, &UserResponse{
		UserId:   system.MakeKey(true),
		Name:     user.Name,
		CreateAt: time.Now().UnixNano(),
	})
}
