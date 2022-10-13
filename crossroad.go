package main

import (
	"CROSSROAD/controller"
	"html/template"
	"io"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Template is a custom html/template renderer for Echo framework
type Template struct {
	templates *template.Template
}

// Render renders a template document
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.New("").Delims("[[", "]]").ParseFiles(
			"view/student/login.html", "view/student/changePassword.html",
			"view/student/myInfo.html", "view/student/main.html", "view/student/sub.html", "view/student/result.html",
			"view/student/applyCandidate.html", "view/student/editCandidate.html",
			"view/student/resultTest.html",

			"view/admin/login.html", "view/admin/signUp.html",
			"view/admin/myInfo.html", "view/admin/candidate.html",

			"view/master/login.html", "view/master/changePassword.html",
			"view/master/myInfo.html",
		)),
	}

	e := echo.New()

	// Set middlewares
	// Logger: loging all request and responses
	// Recover: Recover main thread if it fails
	e.Use(middleware.Logger(), middleware.Recover())

	// Session 설정
	store := session.NewCookieStore([]byte("secret"))
	e.Use(session.Sessions("CASESSION", store))
	e.Use(controller.CacheAPI)

	// Set template renderer
	// We uses standard golang template
	e.Renderer = t

	// Set static serve files
	e.Static("/assets", "static")

	e.GET("/", controller.Index)

	// ============== 학생 page ==============
	// Account
	e.GET("/login", controller.Login)
	e.GET("/changePassword", controller.ChangePassword, controller.AuthAPI)

	// Information
	e.GET("/myInfo", controller.MyInfo, controller.AuthAPI)

	// Candidate
	e.GET("/applyCandidate", controller.ApplyCandidate, controller.GroupAuthAPI)
	e.GET("/editCandidate", controller.EditCandidate, controller.GroupAuthAPI)
	e.GET("/candidate/main", controller.CandidateMain, controller.GroupAuthAPI)
	e.GET("/candidate/sub", controller.CandidateSub, controller.GroupAuthAPI)

	// Result
	e.GET("/result", controller.Result, controller.GroupAuthAPI)
	e.GET("/result2", controller.ResultTest, controller.GroupAuthAPI)

	// ============== 학생 API ==============
	// Account api
	e.POST("/api/login", controller.LoginAPI)                                       // Login api
	e.GET("/api/logout", controller.LogoutAPI)                                      // Logout api
	e.POST("/api/changePassword", controller.ChangePasswordAPI, controller.AuthAPI) // Change Password api
	e.GET("/api/getSession", controller.GetSessionAPI)                              // Get Session api

	// Candidate api
	e.POST("/api/applyCandidate", controller.ApplyCandidateAPI, controller.GroupAuthAPI)                            // Apply Candidate api
	e.POST("/api/editCandidate", controller.EditCandidateAPI, controller.GroupAuthAPI)                              // Edit Candidate api
	e.GET("/api/getCandidatesByGroupAndField", controller.GetCandidatesByGroupAndFieldAPI, controller.GroupAuthAPI) // Get Candidate (group, field) api
	e.GET("/api/getCandidatesByGroup", controller.GetCandidatesByGroupAPI, controller.GroupAuthAPI)                 // Get Candidate (group) api
	e.GET("/api/voteCandidate", controller.VoteCandidateAPI, controller.GroupAuthAPI)                               // Vote Candidate api

	// Group api
	e.GET("/api/getMyGroups", controller.GetMyGroupsAPI, controller.AuthAPI) // Get My Group api
	e.GET("/api/getGroups", controller.GetGroupsAPI, controller.AuthAPI)
	e.GET("/api/moveGroup", controller.MoveGroupAPI, controller.AuthAPI)
	e.GET("/api/leaveGroup", controller.LeaveGroupAPI, controller.AuthAPI)
	e.GET("/api/joinGroup", controller.JoinGroupAPI, controller.AuthAPI)

	// ============== 관리자 page ==============
	// Account
	e.GET("/admin/login", controller.AdminLogin)
	e.GET("/admin/signUp", controller.AdminSignUp)

	// Information
	e.GET("/admin/myInfo", controller.AdminMyInfo, controller.AdminAuthAPI)

	// ============== 관리자 API ==============
	// Account api
	e.POST("/admin/api/login", controller.AdminLoginAPI)
	e.GET("/admin/api/logout", controller.LogoutAPI)
	e.POST("/admin/api/signUp", controller.AdminSignUpAPI)

	// ============== 마스터 page ==============
	// Account
	e.GET("/master/login", controller.MasterLogin)

	// ============== 관리자 API ==============
	// Account api
	e.POST("/master/api/login", controller.MasterLoginAPI)
	e.GET("/master/api/logout", controller.LogoutAPI)

	e.Start(":8888")
}
