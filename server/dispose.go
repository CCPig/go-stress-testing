// Package server 压测启动
package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/link1st/go-stress-testing/Proto/Taurus"
	"github.com/link1st/go-stress-testing/model"
	"github.com/link1st/go-stress-testing/server/client"
	com "github.com/link1st/go-stress-testing/server/comm"
	"github.com/link1st/go-stress-testing/server/golink"
	"github.com/link1st/go-stress-testing/server/statistics"
	"github.com/link1st/go-stress-testing/server/verify"
	"go.k8sf.cloud/go/KsfGo/ksf"
	"go.k8sf.cloud/go/KsfGo/ksf/util/endpoint"
	"strings"
	"sync"
	"time"
)

type Token struct {
	sync.RWMutex
	quick_token   string
	refresh_token string
}

func (t *Token) GetQuickToken() string {
	t.RLock()
	defer t.RUnlock()
	return t.quick_token
}

func (t *Token) GetRefreshToken() string {
	t.RLock()
	defer t.RUnlock()
	return t.refresh_token
}

const (
	connectionMode = 1 // 1:顺序建立长链接 2:并发建立长链接
)

var token = &Token{}

// init 注册验证器
func init() {

	// http
	model.RegisterVerifyHTTP("statusCode", verify.HTTPStatusCode)
	model.RegisterVerifyHTTP("json", verify.HTTPJson)

	// webSocket
	model.RegisterVerifyWebSocket("json", verify.WebSocketJSON)
}

// Dispose 处理函数
func Dispose(ctx context.Context, concurrency, totalNumber uint64, request *model.Request) {
	// 设置接收数据缓存
	ch := make(chan *model.RequestResults, 1000)
	var (
		wg          sync.WaitGroup // 发送数据完成
		wgReceiving sync.WaitGroup // 数据处理完成
	)
	wgReceiving.Add(1)
	statistics.TotalTimes = concurrency * totalNumber
	if strings.HasSuffix(request.URL, "InsertOrder") {
		statistics.Optimize = false
	}
	go statistics.ReceivingResults(concurrency, ch, &wgReceiving, request)

	vec := strings.Split(request.URL, "/")
	if len(vec) < 2 {
		panic("ksf url error")
	}
	funcName := vec[len(vec)-1]
	request.FuncName = funcName

	ep := endpoint.Parse(vec[0])

	comm := ksf.NewCommunicator()
	Auth := &Taurus.TaurusUserObj{}
	counter := &Taurus.CounterObj{}
	bypass := &Taurus.BypassObj{}

	RefreshToken := func(user_id, quick_token, refresh_token string) error {
		option := make(map[string]string, 2)
		option["obj"] = "Taurus.UserAuthServer.TaurusUserObj"
		option["mac_list"] = ""
		req := &Taurus.RefreshTokenReq{
			Refresh_token: refresh_token,
			Quick_token:   quick_token,
			Client: Taurus.ClientInfo{
				User_id:      user_id,
				Channel:      "",
				Guid:         "",
				Xua:          "",
				Imei:         "",
				Macs:         nil,
				Hosts:        nil,
				Extra_params: nil,
			},
		}
		rsp := &Taurus.RefreshTokenRsp{}
		ret, err := Auth.RefreshToken(req, rsp, option)
		if err != nil {
			return err
		}
		if ret.Code < 0 {
			return errors.New(fmt.Sprintf("%+v", ret))
		}
		func() {
			token.Lock()
			defer token.Unlock()
			token.quick_token = rsp.Quick_token
		}()
		return nil
	}
	UserLogin := func() (err error) {
		user_id, exist := request.Headers["user_id"]
		if !exist {
			panic("no user_id")
		}

		rawPasswd, exist := request.Headers["passwd"]
		if !exist {
			panic("no passwd")
		}

		comm.StringToProxy("Taurus.AccessServer.TaurusUserObj@"+ep.String(), Auth)
		//comm.StringToProxy("Taurus.AccessServer.AccessObj@"+ep.String(), bypass)
		//comm.StringToProxy("Taurus.AccessServer.TradeObj@"+ep.String(), counter)
		bypass.SetServant(*Auth.GetServant())
		counter.SetServant(*Auth.GetServant())

		Auth.KsfSetTimeout(20000)
		bypass.KsfSetTimeout(20000)
		counter.KsfSetTimeout(20000)
		option := make(map[string]string, 2)
		option["obj"] = "Taurus.UserAuthServer.TaurusUserObj"
		option["mac_list"] = ""
		var passwd string
		passwd, err = com.DESUserPasswd(rawPasswd)
		if err != nil {
			panic(err)
		}
		req := &Taurus.UserLoginReq{
			User_id: user_id,
			Passwd:  passwd,
			Client: Taurus.ClientInfo{
				User_id:      user_id,
				Channel:      "",
				Guid:         "",
				Xua:          "",
				Imei:         "",
				Macs:         nil,
				Hosts:        nil,
				Extra_params: nil,
			},
			Extends: nil,
		}
		rsp := &Taurus.UserLoginRsp{}
		login, err := Auth.UserLogin(req, rsp, option)
		if err != nil {
			return err
		}
		if login.Code < 0 {
			return errors.New(fmt.Sprintf("%+v", login))
		}
		func() {
			token.Lock()
			defer token.Unlock()
			token.quick_token = rsp.Quick_token
			token.refresh_token = rsp.Refresh_token
		}()

		go func() {
			//ticker := time.NewTicker(time.Minute * 3)
			ticker := time.NewTicker(time.Second * 1)
			select {
			case <-ticker.C:
				RefreshToken(user_id, token.GetQuickToken(), token.GetRefreshToken())
			}
		}()
		return nil
	}

	accountLogin := func() (err error) {
		acc_id, exist := request.Headers["account_id"]
		if !exist {
			panic("no user_id")
		}

		option := make(map[string]string, 2)
		option["obj"] = "Taurus.UserAuthServer.TaurusUserObj"
		option["mac_list"] = ""
		if err != nil {
			panic(err)
		}
		req := &Taurus.AccountLoginReq{
			Account_id: acc_id,
			Pwd:        "",
			Request_id: 0,
			Client:     Taurus.ClientInfo{},
			Extra_params: map[string]string{
				"OPT_STATION_CS": "PC;IIP=58.48.38.46;IPORT=44496;LIP=10.242.1.56;MAC=e0be035d9d9b;HD=0025_3881_22B4_494B;PCN=KS-SHA-LP220149;CPU=bfebfbff000a0634;PI=/dev/nvme0n1p2 468G;VOL=sysfs /sys@VMT;V1.3.0.69	",
				"quick_token":    token.GetQuickToken(),
			},
		}
		rsp := &Taurus.AccountLoginRsp{}
		login, err := Auth.AccountLogin(req, rsp, option)
		if err != nil {
			return err
		}
		if login.Code < 0 {
			return errors.New(fmt.Sprintf("%+v", login))
		}
		return nil
	}

	if err := UserLogin(); err != nil {
		panic(err)
	}

	switch funcName {
	case "QryAsset":
		fallthrough
	case "QryPosition":
		fallthrough
	case "QryOrder":
		fallthrough
	case "InsertOrder":
		if err := accountLogin(); err != nil {
			panic(err)
		}
	}

	for i := uint64(0); i < concurrency; i++ {
		wg.Add(1)
		switch request.Form {
		case model.FormTypeHTTP:
			go golink.HTTP(ctx, i, ch, totalNumber, &wg, request)
		case model.FormTypeWebSocket:
			switch connectionMode {
			case 1:
				// 连接以后再启动协程
				ws := client.NewWebSocket(request.URL)
				ws.SetHeader(request.Headers)
				err := ws.GetConn()
				if err != nil {
					fmt.Println("连接失败:", i, err)
					continue
				}
				go golink.WebSocket(ctx, i, ch, totalNumber, &wg, request, ws)
			case 2:
				// 并发建立长链接
				go func(i uint64) {
					// 连接以后再启动协程
					ws := client.NewWebSocket(request.URL)
					ws.SetHeader(request.Headers)
					err := ws.GetConn()
					if err != nil {
						fmt.Println("连接失败:", i, err)
						return
					}
					golink.WebSocket(ctx, i, ch, totalNumber, &wg, request, ws)
				}(i)
				// 注意:时间间隔太短会出现连接失败的报错 默认连接时长:20毫秒(公网连接)
				time.Sleep(5 * time.Millisecond)
			default:
				data := fmt.Sprintf("不支持的类型:%d", connectionMode)
				panic(data)
			}
		case model.FormTypeGRPC:
			// 连接以后再启动协程
			ws := client.NewGrpcSocket(request.URL)
			err := ws.Link()
			if err != nil {
				fmt.Println("连接失败:", i, err)
				continue
			}
			go golink.Grpc(ctx, i, ch, totalNumber, &wg, request, ws)
		case model.FormTypeRadius:
			// Radius use udp, does not a connection
			go golink.Radius(ctx, i, ch, totalNumber, &wg, request)

		case model.FormTypeKsf:
			switch funcName {
			case "UserLogin":
				go golink.Ksf(ctx, i, ch, totalNumber, &wg, request, Auth, token.GetQuickToken(), token.GetRefreshToken())
			case "QryAsset":
				fallthrough
			case "QryPosition":
				go golink.Ksf(ctx, i, ch, totalNumber, &wg, request, counter, token.GetQuickToken(), token.GetRefreshToken())
			case "QryOrder":
				fallthrough
			case "QryTrade":
				fallthrough
			case "QueryTaskStatus":
				go golink.Ksf(ctx, i, ch, totalNumber, &wg, request, bypass, token.GetQuickToken(), token.GetRefreshToken())
			case "InsertOrder":
				go golink.Ksf(ctx, i, ch, totalNumber, &wg, request, counter, token.GetQuickToken(), token.GetRefreshToken())
			default:
				panic("ksf url error")
			}

		default:
			// 类型不支持
			wg.Done()
		}
	}
	// 等待所有的数据都发送完成
	wg.Wait()
	// 延时1毫秒 确保数据都处理完成了
	time.Sleep(1 * time.Millisecond)
	close(ch)
	// 数据全部处理完成了
	wgReceiving.Wait()
	return
}
