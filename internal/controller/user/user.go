package user

import (
	"context"
	"fmt"
	"zeekgo/api/user"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type User struct{}

// New function
func New() *User {
	return &User{}
}

func (u *User) AAdd(ctx context.Context, req *user.AddReq) (res *user.AddRes, err error) {
	r := g.RequestFromCtx(ctx)
	// name := r.Get("name", "23x")
	md := g.Model("student")
	type User struct {
		Id   int `json:"id"`
		Name string
		Age  string
		// openid   string
		// unionid  string
		// platform string
		// username string
		// mobile   string
		// password string
		// email    string
		// status   string
		// createAt *gtime.Time
		// updateAt *gtime.Time
	}
	var user *User

	err2 := md.Scan(&user)
	fmt.Println("www", user)
	if err2 == nil {
		fmt.Println("yes", user)
		r.Response.WriteJson(user)

	} else {
		fmt.Println("no")
		r.Response.Writeln(err2)

	}
	// err = gerror.New("erritc")
	// r.Response.Writeln(req)
	// r.Response.
	return
}

func (u *User) Db(req *ghttp.Request) {
	md := g.Model("user")
	user, err := md.One()
	if err != nil {
		req.Response.Writeln(user)

	} else {
		req.Response.Writeln("nil")

	}
	// r := g.RequestFromCtx(ctx)
	// name := r.Get("name", "23x")
	// fmt.Println("req", req)
	// res = &user.AddRes{
	// 	Name: "abkjdd",
	// 	Age:  12,
	// }
	// err = gerror.New("erritc")
	// r.Response.Writeln(req)
	// r.Response.
	return
}
