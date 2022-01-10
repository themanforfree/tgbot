package jwglxt

import (
	"fmt"
	"strings"

	"github.com/imroc/req"
)

type STU struct {
	username string
	password string
	client   *req.Req
}

func New(username, password string) *STU {
	return &STU{
		username: username,
		password: password,
		client:   req.New(),
	}
}

func (stu *STU) Login() error {
	res, err := stu.client.Get(Config.KeyUrl, Config.Header)
	if err != nil {
		return err
	}
	if res.Response().StatusCode != 200 {
		return fmt.Errorf("status code error: %d %s", res.Response().StatusCode, res.Response().Status)
	}
	json_data := map[string]interface{}{}
	err = res.ToJSON(&json_data)
	if err != nil {
		return err
	}
	n := json_data["modulus"].(string)
	e := json_data["exponent"].(string)
	rsa_pwd, err := getRsa(stu.password, n, e)
	if err != nil {
		return err
	}
	csrftoken, err := getCsrftoken()
	if err != nil {
		return err
	}
	param := req.Param{
		"csrftoken": csrftoken,
		"yhm":       stu.username,
		"mm":        rsa_pwd,
	}
	res, err = stu.client.Post(Config.LoginUrl, Config.Header, param)
	if err != nil {
		return err
	}
	if res.Response().StatusCode != 200 {
		return fmt.Errorf("status code error: %d %s", res.Response().StatusCode, res.Response().Status)
	}
	if strings.Contains(res.String(), "修改密码") {
		return nil
	} else if strings.Contains(res.String(), "不正确") {
		return fmt.Errorf("username or password is wrong")
	} else {
		return fmt.Errorf("unknown error")
	}
}

func (stu *STU) GetUserInfo() (UserInfo, error) {
	res, err := stu.client.Get(Config.UserInfoUrl, Config.Header)
	if err != nil {
		return UserInfo{}, err
	}
	if res.Response().StatusCode != 200 {
		return UserInfo{}, fmt.Errorf("status code error: %d %s", res.Response().StatusCode, res.Response().Status)
	}
	userinfo := UserInfo{}
	if err := res.ToJSON(&userinfo); err != nil {
		return UserInfo{}, fmt.Errorf("perhaps you are not login\njson error: %s", err)
	}
	return userinfo, nil
}

func (stu *STU) GetGrades(v ...interface{}) (Grade, error) {
	var xnm, xqm int
	if len(v) == 0 {
		xnm, xqm = getXnmXqm()
	} else if len(v) == 2 {
		xnm, xqm = v[0].(int), v[1].(int)
	} else {
		return Grade{}, fmt.Errorf("more than 2 params")
	}
	switch xqm {
	case 0:
		xqm = 0
	case 1:
		xqm = 3
	case 2:
		xqm = 12
	case 3:
		xqm = 16
	default:
		return Grade{}, fmt.Errorf("xqm error")
	}
	param := req.Param{
		"xnm":                    xnm,
		"_search":                "false",
		"queryModel.showCount":   "100",
		"queryModel.currentPage": "1",
		"queryModel.sortName":    "",
		"queryModel.sortOrder":   "asc",
		"time":                   "0",
	}
	if xqm == 0 {
		param["xqm"] = nil
	} else {
		param["xqm"] = xqm
	}
	res, err := stu.client.Get(Config.GradeUrl, Config.Header, param)
	if err != nil {
		return Grade{}, err
	}
	if res.Response().StatusCode != 200 {
		return Grade{}, fmt.Errorf("status code error: %d %s", res.Response().StatusCode, res.Response().Status)
	}
	grade := Grade{}
	if err := res.ToJSON(&grade); err != nil {
		return Grade{}, fmt.Errorf("perhaps you are not login\njson error: %s", err)
	}
	return grade, nil
}

func (stu *STU) GetScore(v ...interface{}) (Grade, error) {
	return stu.GetGrades(v...)
}

func (stu *STU) GetExam(v ...interface{}) (Exam, error) {
	var xnm, xqm int
	if len(v) == 0 {
		xnm, xqm = getXnmXqm()
	} else if len(v) == 2 {
		xnm, xqm = v[0].(int), v[1].(int)
	} else {
		return Exam{}, fmt.Errorf("more than 2 params")
	}
	switch xqm {
	case 0:
		xqm = 0
	case 1:
		xqm = 3
	case 2:
		xqm = 12
	case 3:
		xqm = 16
	default:
		return Exam{}, fmt.Errorf("xqm error")
	}
	param := req.Param{
		"xnm":                    xnm,
		"_search":                "false",
		"queryModel.showCount":   "100",
		"queryModel.currentPage": "1",
		"queryModel.sortName":    "",
		"queryModel.sortOrder":   "asc",
		"time":                   "0",
	}
	if xqm == 0 {
		param["xqm"] = nil
	} else {
		param["xqm"] = xqm
	}
	res, err := stu.client.Get(Config.ExamUrl, Config.Header, param)
	if err != nil {
		return Exam{}, err
	}
	if res.Response().StatusCode != 200 {
		return Exam{}, fmt.Errorf("status code error: %d %s", res.Response().StatusCode, res.Response().Status)
	}
	exam := Exam{}
	if err := res.ToJSON(&exam); err != nil {
		return Exam{}, fmt.Errorf("perhaps you are not login\njson error: %s", err)
	}
	return exam, nil
}
