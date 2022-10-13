package controller

import (
	"CROSSROAD/models"
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

// AdminLogin : Admin Login Page
func AdminLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "adminLogin", nil)
}

// AdminSignUp : Admin SignUp Page
func AdminSignUp(c echo.Context) error {
	return c.Render(http.StatusOK, "adminSignUp", nil)
}

// AdminMyInfo : Admin MyInfo Page
func AdminMyInfo(c echo.Context) error {
	return c.Render(http.StatusOK, "adminMyInfo", nil)
}

// ============== API ==============

// AdminAuthAPI 관리자 로그인 인증 middleware
func AdminAuthAPI(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 로그인이 되어 있지 않으면 login page로 redirect
		session := session.Default(c)
		if session.Get("groupName") == nil {
			return c.Redirect(http.StatusMovedPermanently, "/admin/login")
		}

		return next(c)
	}
}

// AdminLoginAPI 관리자 로그인 API
func AdminLoginAPI(c echo.Context) error {
	isSuccessed, cause := models.AdminLogin(c.FormValue("groupName"), c.FormValue("password"))

	// Login 성공 시
	if isSuccessed {
		// Session에 학번 저장
		session := session.Default(c)
		session.Set("groupName", c.FormValue("groupName"))
		session.Save()

		return c.Redirect(http.StatusMovedPermanently, "/admin/myInfo")
	}

	// Login 실패 시
	return c.Redirect(http.StatusMovedPermanently, "/admin/login?error="+cause)
}

// AdminSignUpAPI 관리자 회원가입 API
func AdminSignUpAPI(c echo.Context) error {
	if c.FormValue("password") != c.FormValue("passwordCheck") { // 비밀번호 확인과 비밀번혹 다를 때
		return c.Redirect(http.StatusMovedPermanently, "/admin/signUp?error=비밀번호가 서로 다릅니다.")
	}

	err := models.AdminSignUp(c.FormValue("groupName"), c.FormValue("password"))

	if err != nil { // 회원가입에서 오류 발생할 때
		return c.Redirect(http.StatusMovedPermanently, "/admin/signUp?error=이미 존재하는 그룹 이름입니다.")
	}

	// 로그인 페이지로 이동
	return c.Redirect(http.StatusMovedPermanently, "/admin/login?error=회원가입 성공하셨습니다.")
}
