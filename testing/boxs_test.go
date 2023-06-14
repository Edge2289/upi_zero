package testing

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestRequest(t *testing.T) {
	client := &http.Client{}

	max := 10
	for i1 := 0; i1 < max; i1++ {
		go func() {
			c := strings.NewReader(`{
				"name": "----",
				"c_company_name": "三明佳妍贸易有限公司",
				"t_company_name": "济宁伊之美信息科技有限公司",
				"funds_type": "测试",
				"amount": "12.11",
				"remark": "就这？？"
			}`)

			req, _ := http.NewRequest("POST", "http://upi.test.com:81/api/fundManage/promptOrder/store", c)

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC91cGkudGVzdC5jb206ODFcL2FwaVwvYXV0aFwvbG9naW4iLCJpYXQiOjE2ODY2MjYxNTEsImV4cCI6MTY4NzA1ODE1MSwibmJmIjoxNjg2NjI2MTUxLCJqdGkiOiJiRDZPdXZUYnFoZ1B3VVBOIiwic3ViIjoyOSwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyJ9.N_AMeKFtRBW8i-LrrTGqzSlvWQ-kQ0cECWO70J8VMZw");
			res, _ := client.Do(req)
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			fmt.Println("client result : " , string(body))
		}()
	}
	for i := 0; i < max; i++ {
		go func() {
			c := strings.NewReader(`{
				"name": "啊啊啊啊",
				"c_company_name": "广州伊霓裳服饰有限公司",
				"t_company_name": "济宁伊之美信息科技有限公司",
				"funds_type": "测试",
				"amount": "12.11",
				"remark": "就这？？"
			}`)

			req, _ := http.NewRequest("POST", "http://upi.test.com:81/api/fundManage/promptOrder/store", c)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC91cGkudGVzdC5jb206ODFcL2FwaVwvYXV0aFwvbG9naW4iLCJpYXQiOjE2ODY2MjYxNTEsImV4cCI6MTY4NzA1ODE1MSwibmJmIjoxNjg2NjI2MTUxLCJqdGkiOiJiRDZPdXZUYnFoZ1B3VVBOIiwic3ViIjoyOSwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyJ9.N_AMeKFtRBW8i-LrrTGqzSlvWQ-kQ0cECWO70J8VMZw");
			res, _ := client.Do(req)
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			fmt.Println("client result : " , string(body))
		}()
	}

	time.Sleep(time.Second * 6)
}