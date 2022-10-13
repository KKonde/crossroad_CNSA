package controller

import (
	"CROSSROAD/models"
	"fmt"
	"net/http"
	"os"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

// MasterLogin Master Login Page
func MasterLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "masterLogin", nil)
}

func MasterMyInfo(c echo.Context) error {
	return c.Render(http.StatusOK, "masterMyInfo", nil)
}

// ============== API ==============
// MasterAuthAPI 관리자 로그인 인증 middleware
func MasterAuthAPI(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 로그인이 되어 있지 않으면 login page로 redirect
		session := session.Default(c)
		if session.Get("master") == nil {
			return c.Redirect(http.StatusMovedPermanently, "/?err=권한을 넘보지 마세요.")
		}

		return next(c)
	}
}

// MasterLoginAPI 마스터 로그인 API
func MasterLoginAPI(c echo.Context) error {
	isSuccessed := models.MasterLogin(c.FormValue("master"), c.FormValue("password"))

	// Login 성공 시
	if isSuccessed {
		// Session에 학번 저장
		session := session.Default(c)
		session.Set("master", c.FormValue("master"))
		session.Save()

		return c.Redirect(http.StatusMovedPermanently, "/master/myInfo")
	}

	// Login 실패 시
	return c.Redirect(http.StatusMovedPermanently, "/?err=권한을 넘보지 마세요.")
}

