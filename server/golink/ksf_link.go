// Package golink.ksf_link.go create by chaochen at 2023/10/24 下午6:26:00
package golink

import (
	"context"
	"github.com/google/uuid"
	"github.com/link1st/go-stress-testing/Proto/Taurus"
	"github.com/link1st/go-stress-testing/helper"
	"github.com/link1st/go-stress-testing/model"
	"sync"
	"time"
)

// Grpc grpc 接口请求
func Ksf(ctx context.Context, chanID uint64, ch chan<- *model.RequestResults, totalNumber uint64, wg *sync.WaitGroup,
	request *model.Request, prx interface{}, quick_token, refresh_token string) {
	defer func() {
		wg.Done()
	}()
	for i := uint64(0); i < totalNumber; i++ {
		ksfRequest(chanID, ch, i, request, prx, quick_token, refresh_token)
	}
	return
}

// grpcRequest 请求
func ksfRequest(chanID uint64, ch chan<- *model.RequestResults, i uint64, request *model.Request,
	prx interface{}, quick_token string, refresh_token string) {
	var (
		startTime = time.Now()
		isSucceed = false
		errCode   = model.HTTPOk
	)
	// 需要发送的数据
	switch request.FuncName {
	case "UserLogin":
		Auth := prx.(*Taurus.TaurusUserObj)
		if Auth == nil {
			errCode = model.RequestErr
		} else {
			// TODO::请求接口示例
			option := make(map[string]string, 2)
			option["obj"] = "Taurus.UserAuthServer.TaurusUserObj"
			option["mac_list"] = ""
			user_id, exist := request.Headers["user_id"]
			if !exist {
				panic("no user_id")
			}

			passwd, exist := request.Headers["passwd"]
			if !exist {
				panic("no passwd")
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
			_, err := Auth.UserLogin(req, rsp, option)
			// fmt.Printf("rsp:%+v", rsp)
			if err != nil {
				errCode = model.RequestErr
			} else {
				// 200 为成功
				isSucceed = true
			}
		}
	case "QryAsset":
		counter := prx.(*Taurus.CounterObj)
		if counter == nil {
			errCode = model.RequestErr
		} else {
			user_id, exist := request.Headers["user_id"]
			if !exist {
				panic("no user_id")
			}

			acc_id, exist := request.Headers["account_id"]
			if !exist {
				panic("no account_id")
			}
			option := make(map[string]string, 2)
			option["obj"] = "Trade.ATPProxyServer.TradeObj"
			option["mac_list"] = "7486E208DFE7"
			option["acc_id"] = acc_id
			option["quick_token"] = quick_token
			option["refresh_token"] = refresh_token
			req := &Taurus.QryAssetReq{
				User_id:      user_id,
				Account_id:   acc_id,
				Currency_id:  "",
				Account_ids:  nil,
				Start_time:   "",
				End_time:     "",
				Brow_index:   "",
				Record_count: 0,
				Request_id:   0,
				Extra_params: map[string]string{
					"OPT_STATION_CS": "PC;IIP=58.48.38.46;IPORT=44496;LIP=10.242.1.56;MAC=e0be035d9d9b;HD=0025_3881_22B4_494B;PCN=KS-SHA-LP220149;CPU=bfebfbff000a0634;PI=/dev/nvme0n1p2 468G;VOL=sysfs /sys@VMT;V1.3.0.69	",
				},
			}
			rsp := &Taurus.QryAssetRsp{}
			_, err := counter.QryAsset(req, rsp, option)
			//fmt.Printf("ret:%+v|rsp:%+v", ret, rsp)
			if err != nil {
				errCode = model.RequestErr
			} else {
				// 200 为成功
				isSucceed = true
			}
		}
	case "QryPosition":
		counter := prx.(*Taurus.CounterObj)
		if counter == nil {
			errCode = model.RequestErr
		} else {
			user_id, exist := request.Headers["user_id"]
			if !exist {
				panic("no user_id")
			}

			acc_id, exist := request.Headers["account_id"]
			if !exist {
				panic("no account_id")
			}
			option := make(map[string]string, 2)
			option["obj"] = "Trade.ATPProxyServer.TradeObj"
			option["mac_list"] = "7486E208DFE7"
			option["acc_id"] = acc_id
			option["quick_token"] = quick_token
			option["refresh_token"] = refresh_token
			req := &Taurus.QryPositionReq{
				User_id:      user_id,
				Account_id:   acc_id,
				Account_ids:  nil,
				Start_time:   "",
				End_time:     "",
				Brow_index:   "",
				Record_count: 0,
				Request_id:   0,
				Extra_params: map[string]string{
					"OPT_STATION_CS": "PC;IIP=58.48.38.46;IPORT=44496;LIP=10.242.1.56;MAC=e0be035d9d9b;HD=0025_3881_22B4_494B;PCN=KS-SHA-LP220149;CPU=bfebfbff000a0634;PI=/dev/nvme0n1p2 468G;VOL=sysfs /sys@VMT;V1.3.0.69	",
				},
			}
			rsp := &Taurus.QryPositionRsp{}
			_, err := counter.QryPosition(req, rsp, option)
			//fmt.Printf("ret:%+v|rsp:%+v", ret, rsp)
			if err != nil {
				errCode = model.RequestErr
			} else {
				// 200 为成功
				isSucceed = true
			}
		}
	case "InsertOrder":
		counter := prx.(*Taurus.CounterObj)
		if counter == nil {
			errCode = model.RequestErr
		} else {
			userId, exist := request.Headers["user_id"]
			if !exist {
				panic("no user_id")
			}

			accountId, exist := request.Headers["account_id"]
			if !exist {
				panic("no account_id")
			}

			uuid.EnableRandPool()
			uuid, _ := uuid.NewUUID()
			option := make(map[string]string, 2)
			option["obj"] = "Trade.SIMProxyServer.TradeObj"
			option["mac_list"] = ""
			option["acc_id"] = accountId
			option["quick_token"] = quick_token
			option["refresh_token"] = refresh_token
			req := &Taurus.InsertOrderReq{
				User_id:         userId,
				Account_id:      accountId,
				Instrument_type: "",
				Local_order_id:  uuid.String(),
				Instrument_id:   "sn2312.XSGE.CF",
				Order_side:      "F1",
				Volume:          1,
				Price_type:      "L",
				Price:           0,
				Task_id:         "",
				Order_way:       "",
				Request_id:      0,
				Extra_params: map[string]string{
					"hedge": "S",
				},
			}
			rsp := &Taurus.InsertOrderRsp{}
			_, err := counter.InsertOrder(req, rsp, option)
			// fmt.Printf("rsp:%+v", rsp)
			if err != nil {
				errCode = model.RequestErr
			} else {
				// 200 为成功
				isSucceed = true
			}
		}
	default:
		panic("funcName mismatch " + request.FuncName)
	}
	requestTime := uint64(helper.DiffNano(startTime))
	requestResults := &model.RequestResults{
		Time:      requestTime,
		IsSucceed: isSucceed,
		ErrCode:   errCode,
	}
	requestResults.SetID(chanID, i)
	ch <- requestResults
}
