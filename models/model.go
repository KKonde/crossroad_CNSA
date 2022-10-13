package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

const (
	// SQLConnectionString : MySQL Connection String
	SQLConnectionString = "root:root@tcp(localhost:3306)/crossroad?charset=utf8&parseTime=True"
	// SALT : SALT
	SALT = "WWtBQ2N<e%}_jqFkJ_[(}O}&C~dxCQU6.C4]wGiF*TLt*#E4<>g%?{#i2M.^YqVYuy<^WN<d@VbKd~.ZFJKNJ%^fZWz)IC^2.+/U`YiJadcD]Ce[NIKp%ioOR&q/nd4K%JMq!N9&AaMG@MnwRIG/]1+`8UdiHjw,I2YS`>q/92?%5*@e4}S%65SoWWSqK0Kq$O.k<4G2sGAz!gu?*.ON1QA@kK:EG/Yis8bq(2qaX?WAFsMQB2IMm;Zxov/N8bwt"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", SQLConnectionString)
	if err != nil {
		fmt.Println("error sql")
		panic(err)
	}
}

// ============================ 모델 ============================

// User model : 학생
type User struct {
	StudentNumber string `gorm:"type:VARCHAR(10); primary_key" json:"studentNumber"`
	Password      string `gorm:"type:VARCHAR(100)" json:"password"`
	StudentName   string `gorm:"type:VARCHAR(30)" json:"studentName"`
}

// TableName of User
func (c *User) TableName() string {
	return "users"
}

// Group model : 관리자
type Group struct {
	GroupName    string `gorm:"type:VARCHAR(50); primary_key" json:"groupName"`
	Password     string `gorm:"type:VARCHAR(100)" json:"password"`
	ApplyAccess  int    `gorm:"type:INT" json:"applyAccess"`
	VoteAccess   int    `gorm:"type:INT" json:"voteAccess"`
	ResultAccess int    `gorm:"type:INT" json:"resultAccess"`
}

// TableName of Group
func (c *Group) TableName() string {
	return "groups"
}

// Candidate model : 후보자
type Candidate struct {
	GroupName     string    `gorm:"type:VARCHAR(50); primary_key" json:"groupName"`
	Field         int       `gorm:"type:INT" json:"field"`
	StudentNumber string    `gorm:"type:VARCHAR(10); primary_key" json:"studentNumber"`
	StudentName   string    `gorm:"type:VARCHAR(30)" json:"studentName"`
	OneLine       string    `gorm:"type:VARCHAR(100)" json:"oneLine"`
	Pledge        string    `gorm:"type:VARCHAR(3000)" json:"pledge"`
	Vote          int       `gorm:"type:INT" json:"vote"`
	UploadTime    time.Time `gorm:"type:DATETIME" json:"uploadTime"`
}

// TableName of Candidate
func (c *Candidate) TableName() string {
	return "candidates"
}

// Member model : 멤버
type Member struct {
	GroupName     string `gorm:"type:VARCHAR(50); primary_key" json:"groupName"`
	StudentNumber string `gorm:"type:VARCHAR(10); primary_key" json:"studentNumber"`
	StudentName   string `gorm:"type:VARCHAR(30)" json:"studentName"`
	Vote          int    `gorm:"type:INT" json:"vote"`
}

// TableName of Member
func (c *Member) TableName() string {
	return "members"
}

// Join model : 가입 신청 멤버
type Join struct {
	GroupName     string `gorm:"type:VARCHAR(50); primary_key" json:"groupName"`
	StudentNumber string `gorm:"type:VARCHAR(10); primary_key" json:"studentNumber"`
	StudentName   string `gorm:"type:VARCHAR(30)" json:"studentName"`
}

// TableName of Join
func (c *Join) TableName() string {
	return "joins"
}

// Master model : 마스터
type Master struct {
	MasterID string `gorm:"type:VARCHAR(50); primary_key" json:"masterId"`
	Password string `gorm:"type:VARCHAR(100)" json:"password"`
}

// ============================ func ============================

// Login 학생 로그인 (성공 여부, 이름 or 원인)
func Login(studentNumber string, password string) (bool, string) {
	user := User{}
	err := db.Table("users").Where("student_number = ?", studentNumber).First(&user).Error

	if err != nil {
		return false, "studentNumber"
	}

	// SALT 적용 전 비밀번호 체크
	if password == user.Password {
		return true, user.StudentName
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password+SALT))
	if err != nil {
		return false, "password"
	}

	return true, user.StudentName
}

// ChangePassword 학생 비밀번호 변경
func ChangePassword(studentNumber string, newPassword string) error {
	user := User{}
	db.Table("users").Where("student_number = ?", studentNumber).First(&user)

	bytes, _ := bcrypt.GenerateFromPassword([]byte(newPassword+SALT), bcrypt.DefaultCost)
	user.Password = string(bytes)

	return db.Save(&user).Error
}

// ApplyCandidate 후보자 등록
func ApplyCandidate(groupName string, field string, studentNumber string, studentName string, oneLine string, pledge string) error {
	fieldi, _ := strconv.Atoi(field)

	mount := GetCandidateMountByGroupAndField(groupName, field)

	if mount >= 5 {
		return errors.New("후보자 수 초과하였습니다.")
	}

	return db.Create(&Candidate{
		GroupName:     groupName,
		Field:         fieldi,
		StudentNumber: studentNumber,
		StudentName:   studentName,
		OneLine:       oneLine,
		Pledge:        pledge,
		UploadTime:    time.Now(),
	}).Error
}

// EditCandidate 후보자 수정
func EditCandidate(groupName string, field string, studentNumber string, oneLine string, pledge string) error {
	fieldi, _ := strconv.Atoi(field)

	candidate := Candidate{}
	db.Table("candidates").Where("group_name = ? AND student_number = ?", groupName, studentNumber).First(&candidate)

	candidate.Field = fieldi
	candidate.OneLine = oneLine
	candidate.Pledge = pledge

	return db.Save(&candidate).Error
}

// VoteCandidate 투표
func VoteCandidate(groupName string, studentNumber string, field string, studentNumberCandidate string) string {
	fieldi, _ := strconv.Atoi(field)

	member := Member{}
	candidate := Candidate{}
	db.Table("members").Where("group_name = ? AND student_number = ?", groupName, studentNumber).Find(&member)
	err := db.Table("candidates").Where("group_name = ? AND student_number = ?", groupName, studentNumberCandidate).First(&candidate).Error

	if err != nil {
		return "잘못된 후보입니다."
	}

	if fieldi == 1 && member.Vote%2 == 0 {
		member.Vote = member.Vote + 1
		candidate.Vote = candidate.Vote + 1
	} else if fieldi == 2 && member.Vote/2 == 0 {
		member.Vote = member.Vote + 2
		candidate.Vote = candidate.Vote + 1
	} else {
		return "이미 투표 하였습니다."
	}

	err = db.Save(&member).Error
	if err != nil {
		return err.Error()
	}

	err = db.Save(&candidate).Error
	if err != nil {
		member.Vote = member.Vote - 1
		db.Save(&member)

		return err.Error()
	}

	return "투표 성공하였습니다."
}

// VoteAbstain 투표 기권
func VoteAbstain(groupName string, studentNumber string, field string) string {
	fieldi, _ := strconv.Atoi(field)

	member := Member{}
	db.Table("members").Where("group_name = ? AND student_number = ?", groupName, studentNumber).Find(&member)

	if fieldi == 1 && member.Vote%2 == 0 {
		member.Vote = member.Vote + 1
	} else if fieldi == 2 && member.Vote/2 == 0 {
		member.Vote = member.Vote + 2
	} else {
		return "이미 투표 하였습니다."
	}

	err := db.Save(&member).Error
	if err != nil {
		return err.Error()
	}

	return "투표 성공하였습니다."
}

// GetCandidateByGroupAndStudentNumber 학번으로 후보자 정보 반환
func GetCandidateByGroupAndStudentNumber(groupName string, studentNumber string) (Candidate, error) {
	candidate := Candidate{}
	err := db.Table("candidates").Where("group_name = ? AND student_number = ?", groupName, studentNumber).Find(&candidate).Error
	return candidate, err
}

// GetCandidatesByGroupAndField 분야로 후보자 정보 반환
func GetCandidatesByGroupAndField(groupName string, field string) []Candidate {
	fieldi, _ := strconv.Atoi(field)

	candidates := []Candidate{}
	db.Table("candidates").Where("group_name = ? AND field = ?", groupName, fieldi).Find(&candidates)

	return candidates
}

// GetCandidateMountByGroupAndField 분야로 후보자 수 반환
func GetCandidateMountByGroupAndField(groupName string, field string) int {
	fieldi, _ := strconv.Atoi(field)
	var cnt int

	db.Table("candidates").Where("group_name = ? AND field = ?", groupName, fieldi).Count(&cnt)
	return cnt
}

// GetCandidateByGroup 그룹으로 후보자 정보 반환
func GetCandidateByGroup(groupName string) []Candidate {
	candidates := []Candidate{}
	db.Table("candidates").Where("group_name = ?", groupName).Find(&candidates)

	return candidates
}

// GetWinnerByGroupAndField 당선자 반환 (소속, 분야)
func GetWinnerByGroupAndField(groupName string, field string) Candidate {
	fieldi, _ := strconv.Atoi(field)
	noWinner := Candidate{
		GroupName:     groupName,
		Field:         1,
		StudentNumber: "??????",
		StudentName:   "???",
		OneLine:       "당선자가 없습니다.",
		Pledge:        "?",
		Vote:          0,
		UploadTime:    time.Now(),
	}
	winner := noWinner
	candidates := []Candidate{}

	db.Table("candidates").Where("group_name = ? AND field = ?", groupName, fieldi).Find(&candidates)
	for _, candidate := range candidates {
		if winner.Vote < candidate.Vote { // 더 높은 득표수
			winner = candidate
		} else if winner.Vote == candidate.Vote { // 득표수가 같으면 초기화
			winner = noWinner
			winner.Vote = candidate.Vote
		}
	}
	return winner
}

// GetWinnerByGroup 그룹으로 당선자 반환
func GetWinnerByGroup(groupName string) []Candidate {
	mainWinner := GetWinnerByGroupAndField(groupName, "1")
	subWinner := GetWinnerByGroupAndField(groupName, "2")

	return []Candidate{mainWinner, subWinner}
}

// GetGroupInfo 그룹 정보 반환
func GetGroupInfo(groupName string) Group {
	group := Group{}
	db.Table("groups").Where("group_name = ?", groupName).First(&group)

	return group
}

// GetMyGroups 자신의 그룹 반환
func GetMyGroups(studentNumber string) []Member {
	members := []Member{}
	db.Table("members").Where("student_number = ?", studentNumber).Find(&members)

	return members
}

// GetGroups 모든 그룹 반환
func GetGroups() []Group {
	groups := []Group{}
	db.Table("groups").Find(&groups)

	return groups
}

// CheckGroupMember 그룹 멤버인지 확인
func CheckGroupMember(groupName string, studentNumber string) error {
	member := Member{}
	return db.Table("members").Where("group_name = ? AND student_number =?", groupName, studentNumber).First(&member).Error
}

// DeleteGroupMember 그룹 멤버 삭제
func DeleteGroupMember(groupName string, studentNumber string) error {
	return db.Table("members").Where("group_name = ? AND student_number =?", groupName, studentNumber).Delete(Member{}).Error
}

// JoinGroup 그룹 가입 신청
func JoinGroup(groupName string, studentNumber string, studentName string) error {
	return db.Create(&Join{
		GroupName:     groupName,
		StudentNumber: studentNumber,
		StudentName:   studentName,
	}).Error
}

// CheckJoinGroup 그룹 가입 신청 했는지 확인
func CheckJoinGroup(groupName string, studentNumber string) error {
	join := Join{}
	return db.Table("joins").Where("group_name = ? AND student_number =?", groupName, studentNumber).First(&join).Error
}

// ============== Admin ==============

// AdminLogin 관리자 로그인 (성공 여부, 원인)
func AdminLogin(groupName string, password string) (bool, string) {
	group := Group{}
	err := db.Table("groups").Where("group_name = ?", groupName).First(&group).Error

	if err != nil {
		return false, "없는 그룹입니다."
	}

	// SALT 적용 전 비밀번호 체크
	if password == group.Password {
		return true, ""
	}

	err = bcrypt.CompareHashAndPassword([]byte(group.Password), []byte(password+SALT))
	if err != nil {
		return false, "비밀번호가 틀렸습니다."
	}

	return true, ""
}

// AdminSignUp 관리자 회원가입 (에러)
func AdminSignUp(groupName string, password string) error {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password+SALT), bcrypt.DefaultCost)

	return db.Create(&Group{
		GroupName: groupName,
		Password:  string(bytes),
	}).Error
}

// ============== Master ==============

// MasterLogin 마스터 로그인
func MasterLogin(masterID string, password string) bool {
	master := Master{}
	err := db.Table("masters").Where("master = ?", masterID).First(&master).Error

	if err != nil {
		return false
	}

	// SALT 적용 전 비밀번호 체크
	if password == master.Password {
		return true
	}

	err = bcrypt.CompareHashAndPassword([]byte(master.Password), []byte(password+SALT))
	if err != nil {
		return false
	}

	return true
}

