package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func getHTML() {
	req, _ := http.NewRequest("GET", "http://kjyy.ccnu.edu.cn/clientweb/xcus/ic2/Default.aspx", nil)
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	println(string(body))
}

var html = `
<html>

<body id="cas">
  <div id="container">
      <div id="content">
        <div class="title"></div>
        <div class="logo"><img id="img1" src="./images/main1-logo.png"/>
<div>
<div class="box" id="login" >
  <form id="fm1" action="/cas/login;jsessionid=2DFC56460A4E3EFC826ED792AD66050CjSme70?service=http://kjyy.ccnu.edu.cn/loginall.aspx?page=" method="post">
    <section class="row">
      <label for="username">用户名:</label>
          <input id="username" name="username" class="required" tabindex="1" accesskey="n" type="text" value="" size="25" autocomplete="off"/>
    </section>
    <section class="row">
      <label for="password">密　码:</label>
      <input id="password" name="password" class="required" tabindex="2" accesskey="p" type="password" value="" size="25" autocomplete="off"/>
    </section>
      <section class="row check"  >
      <label for="warn">请输入信息门户账号与密码。</label>
    </section>
    <section class="row btn-row">
      <input type="hidden" name="lt" value="LT-27747-0wdOBJE4SxphpKy0BnY2HngdfcayHF-account.ccnu.edu.cn" />
      <input type="hidden" name="execution" value="e1s1" />
      <input type="hidden" name="_eventId" value="submit" />
      <input class="btn-submit" name="submit" accesskey="l" value="登录" tabindex="4" type="submit" />
       <input class="btn-reset" accesskey="c" name="resetpass" type="button" tabindex="5" value="忘记密码" onclick="window.location='https://selfservice.ccnu.edu.cn/account/resetpass.jsf'" />
    </section>
  </form>
  <div id = "pc_qrcode" class="pc_sign_in_box">
    <div class="pc_sign_in_cont">
        <div class="pc_qr_code"></div>
        <div id="result">微信扫码登录</div>
    </div>
        <input type="hidden" id="servicestr" />
  </div>
</div><!-- logoform end -->
</div>
  </div> <!-- 31main-logo -->
</div>
    </div> <!-- END #container -->

  </body>

</html>`

func login() {
	req, _ := http.NewRequest(
		"GET",
		//"http://kjyy.ccnu.edu.cn/clientweb/xcus/ic2/Default.aspx",
		"https://account.ccnu.edu.cn/cas/login",
		nil,
	)
	client := &http.Client{}
	resp, _ := client.Do(req)
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	lt_value, _ := doc.Find("input[name='lt']").Attr("value")
	fmt.Println("lt_value = ", lt_value)
	execution_value, _ := doc.Find("input[name='execution']").Attr("value")
	fmt.Println("execution_value = ", execution_value)
	ll_cookie := resp.Header.Get("Set-Cookie")
	fmt.Println(ll_cookie)
	JSESSIONID, _, _ := strings.Cut(ll_cookie, ";")
	fmt.Println("JSESSIONID = ", JSESSIONID)
	entire_cookie := resp.Cookies()
	cookie_name := entire_cookie[0].Name
	fmt.Println("cookie_name = ", cookie_name)
	cookie_value := entire_cookie[0].Value
	fmt.Println("cookie_value = ", cookie_value)

	formData := url.Values{}
	formData.Set("username", "2024214274")
	formData.Set("password", "wasd+iop4444")
	formData.Set("lt", lt_value)
	formData.Set("execution", execution_value)
	formData.Set("_eventId", "submit")
	formData.Set("submit", "登录")

	newUrl := "https://account.ccnu.edu.cn/cas/login;jsessionid=" + cookie_value
	fmt.Println("newUrl = ", newUrl)

	req, _ = http.NewRequest(
		"POST",
		newUrl,
		strings.NewReader(formData.Encode()),
	)
	a := formData.Encode()
	fmt.Println("a = ", a)
	req.Header.Set("Cookie", JSESSIONID)
	resp, _ = client.Do(req)
	defer resp.Body.Close()

	final_cookie := resp.Cookies()
	fmt.Println(final_cookie)
}

func mkresv() {
	Turl := "http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=101699888&lab_id=&kind_id=&room_id=&type=dev&prop=&test_id=&term=&Vnumber=&classkind=&test_name=&start=2024-11-29+20%3A45&end=2024-11-29+22%3A00&start_time=2045&end_time=2200&up_file=&memo=&act=set_resv&_=1732876671234"
	req, _ := http.NewRequest(
		"GET",
		Turl,
		nil,
	)
	client := &http.Client{}
	req.Header.Add("Cookie", "ASP.NET_SessionId=ymbjh355mc20cdjgw0ocwg55")
	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	println(string(body))
}

func main() {

	newUrl := "http://kjyy.ccnu.edu.cn/clientweb/xcus/ic2/Default.aspx"
	req1, _ := http.NewRequest("GET", newUrl, nil)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res1, _ := client.Do(req1)
	c := res1.Header.Get("Set-Cookie")
	fmt.Println("c = ", c)
	c1, _, _ := strings.Cut(c, ";")
	fmt.Println("c1 = ", c1)

	req, _ := http.NewRequest(
		"GET",
		"https://account.ccnu.edu.cn/cas/login",
		nil,
	)
	client = &http.Client{}
	res, _ := client.Do(req)
	//body, _ := io.ReadAll(res.Body)
	//println("body = ", string(body))

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	lt_value := doc.Find("input[name=lt]").AttrOr("value", "")
	fmt.Println("lt_value = ", lt_value)
	execution_value := doc.Find("input[name=execution]").AttrOr("value", "")
	fmt.Println("execution_value = ", execution_value)

	cookie1 := res.Header.Get("Set-Cookie")
	fmt.Println("cookie1 = ", cookie1)
	JSESSIONID, _, _ := strings.Cut(cookie1, ";")
	fmt.Println("JSESSIONID = ", JSESSIONID)

	data := url.Values{}
	data.Set("username", "2024214274")
	data.Set("password", "wasd+iop4444")
	data.Set("lt", lt_value)
	data.Set("execution", execution_value)
	data.Set("_eventId", "submit")
	data.Set("submit", "登录")

	newUrl = fmt.Sprintf("https://account.ccnu.edu.cn/cas/login;" + JSESSIONID)
	fmt.Println("newUrl = ", newUrl)

	req, _ = http.NewRequest(
		"POST",
		"https://account.ccnu.edu.cn/cas/login;"+JSESSIONID,
		bytes.NewBufferString(data.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", JSESSIONID)

	res, _ = client.Do(req)
	defer res.Body.Close()
	//body, _ := io.ReadAll(res.Body)
	//println("body = ", string(body))
	cc := res.Cookies()
	fmt.Println("cc = ", cc)
	cookie3 := res.Cookies()[1]
	fmt.Println("cookie3 = ", cookie3)
	cookie4, _, _ := strings.Cut(cookie3.String(), ";")
	fmt.Println("cookie4 = ", cookie4)

	req, _ = http.NewRequest(
		"GET",
		"http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/data/searchAccount.aspx?type=logonname&ReservaApply=ReservaApply&term=2024214240&_=1732899726120",
		nil,
	)
	req.Header.Set("Cookie", cookie4)
	req.Header.Set("Cookie", c1)
	res, _ = client.Do(req)
	body, _ := io.ReadAll(res.Body)

	fmt.Println("body = ", string(body))
}

//mk resv
/*
http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?
dialogid=&
dev_id=101699888& //seat_id
lab_id=&
kind_id=&
room_id=&
type=dev&
prop=&
test_id=&
term=&
Vnumber=&
classkind=&
test_name=&
start=2024-11-29+20%3A45& //settime
end=2024-11-29+22%3A00&  //settime
start_time=2045&		//settime
end_time=2200&		//settime
up_file=&
memo=&
act=set_resv&
_=1732611395713 //timestamp
*/

//mv resv
/*
http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?
act=del_resv&
id=	159555958&		//resv_id
_=1732611395714 //timestamp

*/

//req, _ := http.NewRequest(
//	"GET",
//	//"http://kjyy.ccnu.edu.cn/clientweb/xcus/ic2/Default.aspx",
//	"https://account.ccnu.edu.cn/cas/login",
//	nil,
//)
//client := &http.Client{}
//resp, _ := client.Do(req)
//doc, _ := goquery.NewDocumentFromReader(resp.Body)
//lt_value, _ := doc.Find("input[name='lt']").Attr("value")
//fmt.Println("lt_value = ", lt_value)
//execution_value, _ := doc.Find("input[name='execution']").Attr("value")
//fmt.Println("execution_value = ", execution_value)
//ll_cookie := resp.Header.Get("Set-Cookie")
//fmt.Println(ll_cookie)
//JSESSIONID, _, _ := strings.Cut(ll_cookie, ";")
//fmt.Println("JSESSIONID = ", JSESSIONID)
//entire_cookie := resp.Cookies()
//cookie_name := entire_cookie[0].Name
//fmt.Println("cookie_name = ", cookie_name)
//cookie_value := entire_cookie[0].Value
//fmt.Println("cookie_value = ", cookie_value)
//
//formData := url.Values{}
//formData.Set("username", "2024214274")
//formData.Set("password", "wasd+iop4444")
//formData.Set("lt", lt_value)
//formData.Set("execution", execution_value)
//formData.Set("_eventId", "submit")
//formData.Set("submit", "登录")

//newUrl := "https://account.ccnu.edu.cn/cas/login;jsessionid=" + cookie_value
//fmt.Println("newUrl = ", newUrl)
//
//req, _ = http.NewRequest(
//	"POST",
//	newUrl,
//	strings.NewReader(formData.Encode()),
//)
//a := formData.Encode()
//fmt.Println("a = ", a)
//req.Header.Set("Cookie", JSESSIONID)
//resp, _ = client.Do(req)
//defer resp.Body.Close()
//
//final_cookie := resp.Cookies()
//fmt.Println(final_cookie)

//newUrl := "http://kjyy.ccnu.edu.cn/clientweb/xcus/ic2/Default.aspx"
//req, _ := http.NewRequest("GET", newUrl, nil)
//client := &http.Client{
//	CheckRedirect: func(req *http.Request, via []*http.Request) error {
//		return http.ErrUseLastResponse
//	},
//}
//res, _ := client.Do(req)
//c := res.Header.Get("Set-Cookie")
//fmt.Println("c = ", c)
//time := time.Now().UnixMilli()
//urll := "http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/data/searchAccount.aspx?type=logonname&ReservaApply=ReservaApply&term=2024214274&_=" + fmt.Sprint(time)
//req, _ = http.NewRequest("GET", urll, nil)
//cookie, _, _ := strings.Cut(c, ";")
//req.Header.Add("Cookie", cookie)
//resp, _ := client.Do(req)
//body, _ := io.ReadAll(resp.Body)
//defer resp.Body.Close()
//println(string(body))
