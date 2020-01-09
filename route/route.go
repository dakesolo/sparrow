package route

import (
	"main/app/index"
	"main/app/user"
	"main/unit"
	"net/http"
)

func InitRoute(mux *http.ServeMux) {
	mux.Handle("/user/result", &unit.Context{Action: user.Result})
	mux.Handle("/index/getIndex", &unit.Context{Action: index.GetIndex})
	mux.Handle("/index/getNav", &unit.Context{Action: index.GetNav})
	mux.Handle("/index/getBanner", &unit.Context{Action: index.GetBanner})
	mux.Handle("/index/hello", &unit.Context{Action: index.Hello})
}
