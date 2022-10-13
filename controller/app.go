package controller

import (
	"CROSSROAD/models"
	"io"
	"net/http"
	"os"
	"unicode/utf8"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

// Index : Index Page
func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}

// Login : Login Page
func Login(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}

// ChangePassword : Change password page
func ChangePassword(c echo.Context) error {
	return c.Render(http.StatusOK, "changePassword", nil)
}

// ApplyCandidate : Apply candidate page
func ApplyCandidate(c echo.Context) error {
	groupName, studentNumber, studentName := SessionPop(c)
	groupInfo := models.GetGroupInfo(groupName)
	if groupInfo.ApplyAccess == 0 { // 후보자 등록 기간이 아닐 때
		return c.Redirect(http.StatusMovedPermanently, "/myInfo?err="+"후보자 등록 기간이 아닙니다.")
	}

	_, err := models.GetCandidateByGroupAndStudentNumber(groupName, studentNumber)
	if err == nil { // 이미 등록 했을 때
		return c.Redirect(http.StatusMovedPermanently, "/editCandidate")
	}

	return c.Render(http.StatusOK, "applyCandidate", map[string]interface{}{
		"groupName":   groupName,
		"studentName": studentName,
	})
}

// EditCandidate : Edit cadidate page
func EditCandidate(c echo.Context) error {
	groupName, studentNumber, studentName := SessionPop(c)
	groupInfo := models.GetGroupInfo(groupName)
	if groupInfo.ApplyAccess == 0 { // 후보자 등록 기간이 아닐 때
		return c.Redirect(http.StatusMovedPermanently, "/myInfo?err="+"후보자 등록 기간이 아닙니다.")
	}

	candidate, err := models.GetCandidateByGroupAndStudentNumber(groupName, studentNumber)
	if err != nil {
		return c.Redirect(http.StatusMovedPermanently, "/applyCandidate")
	}

	return c.Render(http.StatusOK, "editCandidate", map[string]interface{}{
		"groupName":   groupName,
		"studentName": studentName,
		"img1":        "/assets/file/" + groupName + "_" + studentNumber + ".png",
		"field":       candidate.Field,
		"oneLine":     candidate.OneLine,
		"pledge":      candidate.Pledge,
	})
}

// MyInfo : My Information page
func MyInfo(c echo.Context) error {
	groupName, _, _ := SessionPop(c)
	groupInfo := models.GetGroupInfo(groupName)
	if groupInfo.ResultAccess == 1 {
		return c.Render(http.StatusOK, "myInfo", map[string]interface{}{
			"applyAccess": "결과",
			"voteAccess":  "결과",
			"groupName":   groupName,
		})
	}
	applyAccess := "진행 중"
	voteAccess := "진행 중"
	if groupInfo.ApplyAccess == 0 {
		applyAccess = "중지"
	}
	if groupInfo.VoteAccess == 0 {
		voteAccess = "중지"
	}

	return c.Render(http.StatusOK, "myInfo", map[string]interface{}{
		"applyAccess": applyAccess,
		"voteAccess":  voteAccess,
		"groupName":   groupName,
	})
}

// CandidateMain : Candidate Main page
func CandidateMain(c echo.Context) error {
	groupName, _, _ := SessionPop(c)
	img := make([]string, 5)
	for i := 0; i <= 4; i = i + 1 {
		img[i] = "/assets/img/silh_blue.png"
	}
	name := make([]string, 5)
	studentNumber := make([]string, 5)
	oneLine := make([]string, 5)
	pledge := make([]string, 5)

	candidates := models.GetCandidatesByGroupAndField(groupName, "1")

	for index, candidate := range candidates {
		img[index] = "/assets/file/" + candidate.GroupName + "_" + candidate.StudentNumber + ".png"
		name[index] = candidate.StudentName
		studentNumber[index] = candidate.StudentNumber
		oneLine[index] = candidate.OneLine
		pledge[index] = candidate.Pledge
	}

	return c.Render(http.StatusOK, "main", map[string]interface{}{
		"groupName":      groupName,
		"img1":           img[0],
		"name1":          name[0],
		"studentNumber1": studentNumber[0],
		"oneLine1":       oneLine[0],
		"pledge1":        pledge[0],
		"img2":           img[1],
		"name2":          name[1],
		"studentNumber2": studentNumber[1],
		"oneLine2":       oneLine[1],
		"pledge2":        pledge[1],
		"img3":           img[2],
		"name3":          name[2],
		"studentNumber3": studentNumber[2],
		"oneLine3":       oneLine[2],
		"pledge3":        pledge[2],
		"img4":           img[3],
		"name4":          name[3],
		"studentNumber4": studentNumber[3],
		"oneLine4":       oneLine[3],
		"pledge4":        pledge[3],
		"img5":           img[4],
		"name5":          name[4],
		"studentNumber5": studentNumber[4],
		"oneLine5":       oneLine[4],
		"pledge5":        pledge[4],
	})
}

// CandidateMain2 asdf
func CandidateMain2(c echo.Context) error {
	groupName, _, _ := SessionPop(c)
	img := make([]string, 5)
	for i := 0; i <= 4; i = i + 1 {
		img[i] = "/assets/img/silh_blue.png"
	}
	name := make([]string, 5)
	oneLine := make([]string, 5)
	pledge := make([]string, 5)

	candidates := models.GetCandidatesByGroupAndField(groupName, "1")

	for index, candidate := range candidates {
		img[index] = "/assets/file/" + candidate.GroupName + "_" + candidate.StudentNumber + ".png"
		name[index] = candidate.StudentName
		oneLine[index] = candidate.OneLine
		pledge[index] = candidate.Pledge
	}

	return c.Render(http.StatusOK, "main2", map[string]interface{}{
		"groupName": groupName,
	})
}

// CandidateSub : Candidate Sub page
func CandidateSub(c echo.Context) error {
	groupName, _, _ := SessionPop(c)
	img := make([]string, 5)
	for i := 0; i <= 4; i = i + 1 {
		img[i] = "/assets/img/silh_blue.png"
	}
	name := make([]string, 5)
	studentNumber := make([]string, 5)
	oneLine := make([]string, 5)
	pledge := make([]string, 5)

	candidates := models.GetCandidatesByGroupAndField(groupName, "2")

	for index, candidate := range candidates {
		img[index] = "/assets/file/" + candidate.GroupName + "_" + candidate.StudentNumber + ".png"
		name[index] = candidate.StudentName
		studentNumber[index] = candidate.StudentNumber
		oneLine[index] = candidate.OneLine
		pledge[index] = candidate.Pledge
	}

	return c.Render(http.StatusOK, "sub", map[string]interface{}{
		"groupName":      groupName,
		"img1":           img[0],
		"name1":          name[0],
		"studentNumber1": studentNumber[0],
		"oneLine1":       oneLine[0],
		"pledge1":        pledge[0],
		"img2":           img[1],
		"name2":          name[1],
		"studentNumber2": studentNumber[1],
		"oneLine2":       oneLine[1],
		"pledge2":        pledge[1],
		"img3":           img[2],
		"name3":          name[2],
		"studentNumber3": studentNumber[2],
		"oneLine3":       oneLine[2],
		"pledge3":        pledge[2],
		"img4":           img[3],
		"name4":          name[3],
		"studentNumber4": studentNumber[3],
		"oneLine4":       oneLine[3],
		"pledge4":        pledge[3],
		"img5":           img[4],
		"name5":          name[4],
		"studentNumber5": studentNumber[4],
		"oneLine5":       oneLine[4],
		"pledge5":        pledge[4],
	})
}

// Result 최종 결과 페이지
func Result(c echo.Context) error {
	groupName, _, _ := SessionPop(c)
	groupInfo := models.GetGroupInfo(groupName)
	if groupInfo.ResultAccess == 0 { // 결과가 아직 안났을 때
		return c.Redirect(http.StatusMovedPermanently, "/myInfo?err="+"결과가 아직 나오지 않았습니다.")
	}

	winner := models.GetWinnerByGroup(groupName)

	return c.Render(http.StatusOK, "result", map[string]interface{}{
		"groupName":         groupName,
		"mainWinnerName":    winner[0].StudentName,
		"mainWinnerOneLine": winner[0].OneLine,
		"subWinnerName":     winner[1].StudentName,
		"subWinnerOneLine":  winner[1].OneLine,
	})
}

// ResultTest 테스트
func ResultTest(c echo.Context) error {
	return c.Render(http.StatusOK, "result2", nil)
}

// SessionPop 세션 값들 반환 (groupName, studentNumber, studentName)
func SessionPop(c echo.Context) (string, string, string) {
	session := session.Default(c)
	groupName := ""
	if session.Get("groupNameStudent") != nil {
		groupName = session.Get("groupNameStudent").(string)
	}
	return groupName, session.Get("studentNumber").(string), session.Get("studentName").(string)
}

// ============== API ==============

// CacheAPI 캐시 초기화
func CacheAPI(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Expires", "0")
		c.Response().Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
		c.Response().Header().Set("Pragma", "no-cache")

		return next(c)
	}
}

// AuthAPI 로그인 인증 middleware
func AuthAPI(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 로그인이 되어 있지 않으면 login page로 redirect
		session := session.Default(c)
		if session.Get("studentNumber") == nil {
			return c.Redirect(http.StatusMovedPermanently, "/login?err=로그인을 안하셨습니다.")
		}

		return next(c)
	}
}

// GroupAuthAPI 그룹 가입 인증 middleware
func GroupAuthAPI(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session := session.Default(c)
		if session.Get("studentNumber") == nil { // AuthAPI
			return c.Redirect(http.StatusMovedPermanently, "/login?err=로그인을 안하셨습니다.")
		}
		if session.Get("groupNameStudent") == nil { // 그룹 확인
			return c.Redirect(http.StatusMovedPermanently, "/myInfo?err=그룹을 정해주세요.")
		}

		return next(c)
	}
}

// LoginAPI 로그인 API
func LoginAPI(c echo.Context) error {
	isSuccessed, studentNameOrCause := models.Login(c.FormValue("studentNumber"), c.FormValue("password"))

	// Login 성공 시
	if isSuccessed {
		session := session.Default(c)
		session.Set("studentNumber", c.FormValue("studentNumber"))
		session.Set("studentName", studentNameOrCause)
		session.Save()

		// 비밀번호가 초기 세팅 값이면 비밀번호 변경 페이지로 이동
		if c.FormValue("studentNumber") == c.FormValue("password") {
			return c.Redirect(http.StatusMovedPermanently, "/changePassword")
		}

		return c.Redirect(http.StatusMovedPermanently, "/myInfo")
	}

	// Login 실패 시
	return c.Redirect(http.StatusMovedPermanently, "/login?err="+studentNameOrCause)
}

// LogoutAPI 로그아웃
func LogoutAPI(c echo.Context) error {
	// Session 초기화
	session := session.Default(c)
	session.Clear()
	session.Save()

	// 로그인 페이지로 빠이빠이
	return c.Render(http.StatusOK, "login", nil)
}

// ChangePasswordAPI 비밀번호 변경 API
func ChangePasswordAPI(c echo.Context) error {
	session := session.Default(c)

	isSuccessed, _ := models.Login(session.Get("studentNumber").(string), c.FormValue("password"))

	if !isSuccessed { // 비밀번호 오답
		return c.Redirect(http.StatusMovedPermanently, "/changePassword?err=비밀번호가 틀렸습니다.")
	}
	if c.FormValue("newPassword") == "" { // 새 비밀번호 공백
		return c.Redirect(http.StatusMovedPermanently, "/changePassword?err=바꿀 비밀번호를 설정해 주세요.")
	}
	if c.FormValue("newPassword") != c.FormValue("newPasswordCheck") { // 새 비밀번호 오답
		return c.Redirect(http.StatusMovedPermanently, "/changePassword?err=비밀번호 확인이 틀렸습니다.")
	}

	err := models.ChangePassword(session.Get("studentNumber").(string), c.FormValue("newPassword"))
	if err != nil {
		return c.Redirect(http.StatusMovedPermanently, "/changePassword?err=db에러 발생")
	}

	return c.Redirect(http.StatusMovedPermanently, "/myInfo")
}

// GetSessionAPI 세션 반환 API
func GetSessionAPI(c echo.Context) error {
	groupName, studentNumber, studentName := SessionPop(c)
	s := groupName + "__" + studentNumber + "__" + studentName

	return c.String(http.StatusOK, s)
}

// ImageUploadAPI 이미지 업로드 API
func ImageUploadAPI(c echo.Context) error {
	groupName, studentNumber, _ := SessionPop(c)
	filename := groupName + "_" + studentNumber + ".png"
	face, err := c.FormFile("face")

	// face가 비어있으면
	if err != nil {
		src, err := os.Open("static/img/silh_blue.png")
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create("static/file/" + filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		return nil
	}

	src, err := face.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create("static/file/" + filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

// ApplyCandidateAPI 후보자 등록 API
func ApplyCandidateAPI(c echo.Context) error {
	groupName, studentNumber, studentName := SessionPop(c)
	groupInfo := models.GetGroupInfo(groupName)
	if groupInfo.ApplyAccess == 0 { // 후보자 등록 기간이 아닐 때
		return c.Redirect(http.StatusMovedPermanently, "/myInfo?err="+"후보자 등록 기간이 아닙니다.")
	}

	if c.FormValue("field") == "0" { // 분야 선택 하지 않은 경우
		return c.Redirect(http.StatusMovedPermanently, "/applyCandidate?err=분야를 설정해주세요.")
	}
	if utf8.RuneCountInString(c.FormValue("oneLine")) < 2 { // 한줄 소개 글자 수 확인
		return c.Redirect(http.StatusMovedPermanently, "/applyCandidate?err=한줄 소개가 너무 짧습니다.")
	}
	if utf8.RuneCountInString(c.FormValue("pledge")) < 2 { // 공약 글자 수 확인
		return c.Redirect(http.StatusMovedPermanently, "/applyCandidate?err=공약이 너무 짧습니다.")
	}

	ImageUploadAPI(c)
	err := models.ApplyCandidate(groupName, c.FormValue("field"), studentNumber, studentName, c.FormValue("oneLine"), c.FormValue("pledge"))

	if err != nil {
		if err.Error()[:10] == "Error 1062" { // 이미 신청한 경우
			models.EditCandidate(groupName, c.FormValue("field"), studentNumber, c.FormValue("oneLine"), c.FormValue("pledge"))

			return c.Redirect(http.StatusMovedPermanently, "/candidate/main")
		}
		return c.Redirect(http.StatusMovedPermanently, "/applyCandidate?err="+err.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/candidate/main")
}

// EditCandidateAPI 후보자 수정 API
func EditCandidateAPI(c echo.Context) error {
	groupName, studentNumber, studentName := SessionPop(c)
	groupInfo := models.GetGroupInfo(groupName)
	if groupInfo.ApplyAccess == 0 { // 후보자 등록 기간이 아닐 때
		return c.Redirect(http.StatusMovedPermanently, "/myInfo?err="+"후보자 등록 기간이 아닙니다.")
	}

	if c.FormValue("field") == "0" { // 분야 선택 하지 않은 경우
		return c.Redirect(http.StatusMovedPermanently, "/editCandidate?err=분야를 설정해주세요.")
	}
	if utf8.RuneCountInString(c.FormValue("oneLine")) < 2 { // 한줄 소개 글자 수 확인
		return c.Redirect(http.StatusMovedPermanently, "/editCandidate?err=한줄 소개가 너무 짧습니다.")
	}
	if utf8.RuneCountInString(c.FormValue("pledge")) < 2 { // 공약 글자 수 확인
		return c.Redirect(http.StatusMovedPermanently, "/editCandidate?err=공약이 너무 짧습니다.")
	}

	if c.FormValue("token") == "1" {
		ImageUploadAPI(c)
	}
	err := models.EditCandidate(groupName, c.FormValue("field"), studentNumber, c.FormValue("oneLine"), c.FormValue("pledge"))

	if err != nil {
		if err.Error()[:10] == "Error 1364" { // 신청하지 않은 경우
			models.ApplyCandidate(groupName, c.FormValue("field"), studentNumber, studentName, c.FormValue("oneLine"), c.FormValue("pledge"))

			return c.Redirect(http.StatusMovedPermanently, "/candidate/main")
		}
		return c.Redirect(http.StatusMovedPermanently, "/editCandidate?err="+err.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/candidate/main")
}

// VoteCandidateAPI 투표하기 API
func VoteCandidateAPI(c echo.Context) error {
	groupName, studentNumber, _ := SessionPop(c)
	groupInfo := models.GetGroupInfo(groupName)
	if groupInfo.VoteAccess == 0 { // 투표 기간이 아닐 때
		return c.Redirect(http.StatusMovedPermanently, "/candidate/main?err=투표기간이 아닙니다.")
	}
	err := models.VoteCandidate(groupName, studentNumber, c.QueryParam("field"), c.QueryParam("studentNumberCandidate"))

	return c.Redirect(http.StatusMovedPermanently, "/candidate/main?err="+err)
}

// GetCandidatesByGroupAndFieldAPI 그룹, 분야단위 후보자 불러오기 API
func GetCandidatesByGroupAndFieldAPI(c echo.Context) error {
	groupName, _, _ := SessionPop(c)

	return c.JSON(http.StatusOK, models.GetCandidatesByGroupAndField(groupName, c.QueryParam("field")))
}

// GetCandidatesByGroupAPI 그룹단위 후보자 불러오기 API
func GetCandidatesByGroupAPI(c echo.Context) error {
	groupName, _, _ := SessionPop(c)

	return c.JSON(http.StatusOK, models.GetCandidateByGroup(groupName))
}

// GetMyGroupsAPI 자신이 속한 그룹 불러오기 API
func GetMyGroupsAPI(c echo.Context) error {
	_, studentNumber, _ := SessionPop(c)

	return c.JSON(http.StatusOK, models.GetMyGroups(studentNumber))
}

// GetGroupsAPI 모든 그룹 불러오기 API
func GetGroupsAPI(c echo.Context) error {
	_, studentNumber, _ := SessionPop(c)

	groups := models.GetGroups()
	groups2 := models.GetMyGroups(studentNumber)
	rgroups := []models.Group{}

	for _, group := range groups {
		key := 0
		for _, group2 := range groups2 {
			if group.GroupName == group2.GroupName { // 이미 가입된 그룹이면
				key = 1
				break
			}
		}
		if key == 1 { // 이미 가입된 그룹이면 rgroups에 포함 X
			continue
		}
		rgroups = append(rgroups, group) // 가입되지 않은 그룹이면 rgroups에 포함
	}

	return c.JSON(http.StatusOK, rgroups)
}

// MoveGroupAPI 그룹 이동 API
func MoveGroupAPI(c echo.Context) error {
	_, studentNumber, _ := SessionPop(c)
	err := models.CheckGroupMember(c.QueryParam("groupName"), studentNumber)
	if err != nil {
		return c.Redirect(http.StatusMovedPermanently, "/myInfo?err="+c.QueryParam("groupName")+"에 가입되어 있지 않습니다.")
	}
	session := session.Default(c)
	session.Set("groupNameStudent", c.QueryParam("groupName"))
	session.Save()

	return c.Redirect(http.StatusMovedPermanently, "/myInfo?err="+c.QueryParam("groupName")+"으로 이동했습니다.")
}

// LeaveGroupAPI 그룹 탈퇴 API / query(groupName)
func LeaveGroupAPI(c echo.Context) error {
	_, studentNumber, _ := SessionPop(c)
	groupName := c.QueryParam("groupName")
	groupInfo := models.GetGroupInfo(groupName)
	if groupInfo.VoteAccess == 1 || groupInfo.ResultAccess == 1 { // 투표 기간 혹은 결과 기간 이면
		return c.Redirect(http.StatusMovedPermanently, "/myInfo?err=특정 기간에는 탈퇴에 관리자의 허가가 필요합니다.")
	}

	err := models.DeleteGroupMember(groupName, studentNumber)
	if err != nil {
		return c.Redirect(http.StatusMovedPermanently, "/myInfo?err="+groupName+"에 탈퇴 실패했습니다.")
	}

	session := session.Default(c)
	if session.Get("groupnameStudent") == groupName { // 현재 속해있는 그룹 탈퇴 시 세션 초기화
		session.Set("groupNameStudent", nil)
		session.Save()
	}

	return c.Redirect(http.StatusMovedPermanently, "/myInfo?err="+groupName+"에 탈퇴했습니다.")
}

// JoinGroupAPI 그룹 가입 신청 API
func JoinGroupAPI(c echo.Context) error {
	_, studentNumber, studentName := SessionPop(c)
	groupName := c.QueryParam("groupName")
	err := models.CheckGroupMember(groupName, studentNumber)
	if err == nil { // 그룹 멤버인지 확인
		return c.Redirect(http.StatusMovedPermanently, "/myInfo?err="+c.QueryParam("groupName")+"에 이미 가입되어 있습니다.")
	}

	err = models.CheckJoinGroup(groupName, studentNumber)
	if err == nil { // 그룹 가입 신청했는지 확인
		return c.Redirect(http.StatusMovedPermanently, "/myInfo?err="+c.QueryParam("groupName")+"에 이미 가입 신청되어 있습니다.")
	}

	err = models.JoinGroup(groupName, studentNumber, studentName)
	if err != nil { // 기타 에러
		return c.Redirect(http.StatusMovedPermanently, "/myInfo?err="+err.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/myInfo?err=가입 신청했습니다.")
}
