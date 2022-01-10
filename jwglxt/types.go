package jwglxt

type UserInfo struct {
	StuId   string `json:"xh"`
	Name    string `json:"xm"`
	Class   string `json:"bh_id"`
	From    string `json:"jg"`
	Gender  string `json:"xbm"`
	College string `json:"jg_id"`
	Phone   string `json:"sjhm"`
	CardId  string `json:"zjhm"`
	Major   string `json:"zyh_id"`
}

type Grade struct {
	Items []struct {
		Name     string `json:"kcmc"`
		Score    string `json:"cj"`
		Teachers string `json:"jsxm"`
		ExamType string `json:"ksxz"`
		Credit   string `json:"xf"`
		Point    string `json:"jd"`
		// jxb_id     string
		// xnm        string
		// xqm        string
	} `json:"items"`
}

type Exam struct {
	Items []struct {
		Name     string `json:"kcmc"`
		Location string `json:"cdmc"`
		Time     string `json:"kssj"`
	} `json:"items"`
}
