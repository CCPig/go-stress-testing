package comm

import (
	"github.com/link1st/go-stress-testing/Proto/Taurus"
	"strings"
)

func Get(code int, lang ...int) string {
	if len(lang) == 0 {
		return Taurus.GetErrFormat(code, Taurus.ZH_CN)
	} else {
		return Taurus.GetErrFormat(code, lang[0])
	}
}

func ES(code int, param ...string) Taurus.Result {
	msg := Get(code)
	msg = strings.ReplaceAll(msg, "{}", "")
	msgSlice := strings.Split(msg, "[")
	placeHoldersCnt := len(msgSlice) - 1

	buildMsg := ""
	if len(msg) == 0 && len(param) == 1 {
		buildMsg = param[0]
		return Taurus.Result{
			Code: int32(code),
			Msg:  buildMsg,
		}
	}
	for i := 0; i < len(msgSlice); i++ {
		buildMsg += msgSlice[i]
		if i != len(msgSlice)-1 {
			buildMsg += "["
		}

		if i < len(param) && i < placeHoldersCnt {
			buildMsg += param[i]
		}
	}

	return Taurus.Result{
		Code: int32(code),
		Msg:  buildMsg,
	}
}
