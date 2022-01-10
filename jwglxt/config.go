package jwglxt

import "github.com/imroc/req"

var Config = struct {
	Header      req.Header
	BaseUrl     string
	LoginUrl    string
	KeyUrl      string
	UserInfoUrl string
	GradeUrl    string
	ExamUrl     string
}{
	Header: req.Header{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36",
		"Referer":    "https://jwglxt.haut.edu.cn",
	},
	BaseUrl:     "https://jwglxt.haut.edu.cn",
	LoginUrl:    "https://jwglxt.haut.edu.cn/jwglxt/xtgl/login_slogin.html",
	KeyUrl:      "https://jwglxt.haut.edu.cn/jwglxt/xtgl/login_getPublicKey.html",
	UserInfoUrl: "https://jwglxt.haut.edu.cn/jwglxt/xsxxxggl/xsxxwh_cxCkDgxsxx.html?gnmkdm=N100801",
	GradeUrl:    "https://jwglxt.haut.edu.cn/jwglxt/cjcx/cjcx_cxDgXscj.html?doType=query&gnmkdm=N305005",
	ExamUrl:     "https://jwglxt.haut.edu.cn/jwglxt/kwgl/kscx_cxXsksxxIndex.html?doType=query&gnmkdm=N358105",
}
