package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AddReq struct {
	g.Meta `path:"/users" method:"get"` // key和value之间不能有空格

	// g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
	UserName string `p:"name" d:"李四"`
	Password string
	Age      int
}

type AddRes struct {
	Name string
	Age  int
}

// type IUser interface {
// 	Add(ctx context.Context, req *AddReq) (res *AddResp, err error)
// }
