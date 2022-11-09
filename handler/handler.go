package handler

import (
	"fmt"
	"todolists/utility"
)

func Test(context utility.ReqRes) {
	fmt.Println(context.Req.URL)
}
