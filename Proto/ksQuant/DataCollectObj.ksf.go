// Package ksQuant comment
// This file was generated by ksf2go 1.3.20
// Generated from DataCollectObj.ksf
package ksQuant

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go.k8sf.cloud/go/KsfGo/ksf"
	m "go.k8sf.cloud/go/KsfGo/ksf/model"
	"go.k8sf.cloud/go/KsfGo/ksf/protocol/codec"
	"go.k8sf.cloud/go/KsfGo/ksf/protocol/kup"
	"go.k8sf.cloud/go/KsfGo/ksf/protocol/res/basef"
	"go.k8sf.cloud/go/KsfGo/ksf/protocol/res/requestf"
	"go.k8sf.cloud/go/KsfGo/ksf/util/current"
	"go.k8sf.cloud/go/KsfGo/ksf/util/tools"
	"go.k8sf.cloud/go/KsfGo/ksf/util/trace"
	"net"
	"unsafe"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = codec.FromInt8
	_ = unsafe.Pointer(nil)
	_ = bytes.ErrTooLarge
	_ = net.UDPConn{}
)

// DataCollectObj struct
type DataCollectObj struct {
	servant m.Servant
}

// HelloWorld is the proxy function for the method defined in the ksf file, with the context
func (obj *DataCollectObj) HelloWorld(opts ...map[string]string) (ret Error, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}
	ksfResp := new(requestf.ResponsePacket)
	ksfCtx := context.Background()

	err = obj.servant.KsfInvoke(ksfCtx, 0, "HelloWorld", buf.ToBytes(), statusMap, contextMap, ksfResp)
	if err != nil {
		return ret, err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
	err = ret.ReadBlock(readBuf, 0, true)
	if err != nil {
		return ret, err
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range ksfResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range ksfResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range ksfResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

// HelloWorldWithContext is the proxy function for the method defined in the ksf file, with the context
func (obj *DataCollectObj) HelloWorldWithContext(ksfCtx context.Context, opts ...map[string]string) (ret Error, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()

	traceData, ok := current.GetTraceData(ksfCtx)
	if ok && traceData.TraceCall {
		traceData.NewSpan()
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCS, uint(buf.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		ksf.Trace(traceData.GetTraceKey(trace.EstCS), trace.TraceAnnotationCS, ksf.GetClientConfig().ModuleName, obj.servant.Name(), "HelloWorld", 0, traceParam, "")
	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}

	ksfResp := new(requestf.ResponsePacket)
	err = obj.servant.KsfInvoke(ksfCtx, 0, "HelloWorld", buf.ToBytes(), statusMap, contextMap, ksfResp)
	if err != nil {
		return ret, err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
	err = ret.ReadBlock(readBuf, 0, true)
	if err != nil {
		return ret, err
	}

	if ok && traceData.TraceCall {
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCR, uint(readBuf.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value[""] = ret
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		ksf.Trace(traceData.GetTraceKey(trace.EstCR), trace.TraceAnnotationCR, ksf.GetClientConfig().ModuleName, obj.servant.Name(), "HelloWorld", int(ksfResp.IRet), traceParam, "")
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range ksfResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range ksfResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range ksfResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

// HelloWorldOneWayWithContext is the proxy function for the method defined in the ksf file, with the context
func (obj *DataCollectObj) HelloWorldOneWayWithContext(ksfCtx context.Context, opts ...map[string]string) (ret Error, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}

	ksfResp := new(requestf.ResponsePacket)
	err = obj.servant.KsfInvoke(ksfCtx, 1, "HelloWorld", buf.ToBytes(), statusMap, contextMap, ksfResp)
	if err != nil {
		return ret, err
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range ksfResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range ksfResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range ksfResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

// SetServant sets servant for the service.
func (obj *DataCollectObj) SetServant(servant m.Servant) {
	obj.servant = servant
}

// GetServant gets servant for the service.
func (obj *DataCollectObj) GetServant() (servant *m.Servant) {
	return &obj.servant
}

// SetOnConnectCallback
func (obj *DataCollectObj) SetOnConnectCallback(callback func(string)) {
	obj.servant.SetOnConnectCallback(callback)
}

// SetOnCloseCallback
func (obj *DataCollectObj) SetOnCloseCallback(callback func(string)) {
	obj.servant.SetOnCloseCallback(callback)
}

// SetKsfCallback
func (obj *DataCollectObj) SetKsfCallback(callback DataCollectObjKsfCallback) {
	var push DataCollectObjPushCallback
	push.Cb = callback
	obj.servant.SetKsfCallback(&push)
}

// SetPushCallback
func (obj *DataCollectObj) SetPushCallback(callback func([]byte)) {
	obj.servant.SetPushCallback(callback)
}
func (obj *DataCollectObj) req2Byte(rsp *requestf.ResponsePacket) []byte {
	req := requestf.RequestPacket{}
	req.IVersion = rsp.IVersion
	req.IRequestId = rsp.IRequestId
	req.IMessageType = rsp.IMessageType
	req.CPacketType = rsp.CPacketType
	req.Context = rsp.Context
	req.Status = rsp.Status
	req.SBuffer = rsp.SBuffer

	os := codec.NewBuffer()
	req.WriteTo(os)
	bs := os.ToBytes()
	sbuf := bytes.NewBuffer(nil)
	sbuf.Write(make([]byte, 4))
	sbuf.Write(bs)
	length := sbuf.Len()
	binary.BigEndian.PutUint32(sbuf.Bytes(), uint32(length))
	return sbuf.Bytes()
}

func (obj *DataCollectObj) rsp2Byte(rsp *requestf.ResponsePacket) []byte {
	if rsp.IVersion == basef.KUPVERSION {
		return obj.req2Byte(rsp)
	}
	os := codec.NewBuffer()
	rsp.WriteTo(os)
	bs := os.ToBytes()
	sbuf := bytes.NewBuffer(nil)
	sbuf.Write(make([]byte, 4))
	sbuf.Write(bs)
	length := sbuf.Len()
	binary.BigEndian.PutUint32(sbuf.Bytes(), uint32(length))
	return sbuf.Bytes()
}

// KsfPing
func (obj *DataCollectObj) KsfPing() {
	ksfCtx := context.Background()
	obj.servant.KsfPing(ksfCtx)
}

// KsfSetTimeout sets the timeout for the servant which is in ms.
func (obj *DataCollectObj) KsfSetTimeout(timeout int) {
	obj.servant.KsfSetTimeout(timeout)
}

// KsfSetVersion default as KSFVERSION,you can set JSONVERSION.
func (obj *DataCollectObj) KsfSetVersion(version int16) {
	obj.servant.KsfSetVersion(version)
}

// KsfSetProtocol sets the protocol for the servant.
func (obj *DataCollectObj) KsfSetProtocol(p m.Protocol) {
	obj.servant.KsfSetProtocol(p)
}

// AddServant adds servant  for the service.
func (obj *DataCollectObj) AddServant(imp DataCollectObjServant, servantObj string) {
	ksf.AddServant(obj, imp, servantObj)
}

// AddServantWithContext adds servant  for the service with context.
func (obj *DataCollectObj) AddServantWithContext(imp DataCollectObjServantWithContext, servantObj string) {
	ksf.AddServantWithContext(obj, imp, servantObj)
}

type DataCollectObjServant interface {
	HelloWorld() (ret Error, err error)
}
type DataCollectObjServantWithContext interface {
	HelloWorld(ksfCtx context.Context) (ret Error, err error)

	DoClose(ctx context.Context)
}

// Dispatch is used to call the server side implement for the method defined in the ksf file. withContext shows using context or not.
func (obj *DataCollectObj) Dispatch(ksfCtx context.Context, val interface{}, ksfReq *requestf.RequestPacket, ksfResp *requestf.ResponsePacket, withContext bool) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	readBuf := codec.NewReader(nil)
	buf := codec.NewBuffer()
	switch ksfReq.SFuncName {
	case "HelloWorld":

		traceData, ok := current.GetTraceData(ksfCtx)
		if ok && traceData.TraceCall {
			var traceParam string
			traceParamFlag := traceData.NeedTraceParam(trace.EstSR, uint(readBuf.Len()))
			if traceParamFlag == trace.EnpNormal {
				value := map[string]interface{}{}
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			ksf.Trace(traceData.GetTraceKey(trace.EstSR), trace.TraceAnnotationSR, ksf.GetClientConfig().ModuleName, ksfReq.SServantName, "HelloWorld", 0, traceParam, "")
		}

		var funRet Error
		if !withContext {
			imp := val.(DataCollectObjServant)
			funRet, err = imp.HelloWorld()
		} else {
			imp := val.(DataCollectObjServantWithContext)
			funRet, err = imp.HelloWorld(ksfCtx)
		}

		if err != nil {
			return err
		}

		if ksfReq.IVersion == basef.KSFVERSION {
			buf.Reset()

			err = funRet.WriteBlock(buf, 0)
			if err != nil {
				return err
			}

		} else if ksfReq.IVersion == basef.KUPVERSION {
			rspKup := kup.NewUniAttribute()

			err = funRet.WriteBlock(buf, 0)
			if err != nil {
				return err
			}

			rspKup.PutBuffer("", buf.ToBytes())
			rspKup.PutBuffer("ksf_ret", buf.ToBytes())

			buf.Reset()
			err = rspKup.Encode(buf)
			if err != nil {
				return err
			}
		} else if ksfReq.IVersion == basef.JSONVERSION {
			rspJson := map[string]interface{}{}
			rspJson["ksf_ret"] = funRet

			var rspByte []byte
			if rspByte, err = json.Marshal(rspJson); err != nil {
				return err
			}

			buf.Reset()
			err = buf.WriteSliceUint8(rspByte)
			if err != nil {
				return err
			}
		}

		if ok && traceData.TraceCall {
			var traceParam string
			traceParamFlag := traceData.NeedTraceParam(trace.EstSS, uint(buf.Len()))
			if traceParamFlag == trace.EnpNormal {
				value := map[string]interface{}{}
				value[""] = funRet
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			ksf.Trace(traceData.GetTraceKey(trace.EstSS), trace.TraceAnnotationSS, ksf.GetClientConfig().ModuleName, ksfReq.SServantName, "HelloWorld", 0, traceParam, "")
		}

	default:
		if ksfReq.SFuncName == "DoClose" {
			if withContext {
				imp := val.(DataCollectObjServantWithContext)
				imp.DoClose(ksfCtx)
			}
			return nil
		}
		return fmt.Errorf("func mismatch")
	}
	var statusMap map[string]string
	if status, ok := current.GetResponseStatus(ksfCtx); ok && status != nil {
		statusMap = status
	}
	var contextMap map[string]string
	if ctx, ok := current.GetResponseContext(ksfCtx); ok && ctx != nil {
		contextMap = ctx
	}
	*ksfResp = requestf.ResponsePacket{
		IVersion:     ksfReq.IVersion,
		CPacketType:  0,
		IRequestId:   ksfReq.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(buf.ToBytes()),
		Status:       statusMap,
		SResultDesc:  "",
		Context:      contextMap,
	}

	_ = readBuf
	_ = buf
	_ = length
	_ = have
	_ = ty
	return nil
}

type DataCollectObjKsfCallback interface {
	HelloWorld_Callback(ret *Error, opt ...map[string]string)
	HelloWorld_ExceptionCallback(err error)
}

// DataCollectObjPushCallback struct
type DataCollectObjPushCallback struct {
	Cb DataCollectObjKsfCallback
}

func (cb *DataCollectObjPushCallback) Ondispatch(ksfResp *requestf.ResponsePacket) {
	switch ksfResp.SResultDesc {
	case "HelloWorld":
		err := func() error {
			var (
				length int32
				have   bool
				ty     byte
			)
			var err error
			readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
			var ret Error
			err = ret.ReadBlock(readBuf, 0, true)
			if err != nil {
				return err
			}

			_ = length
			_ = have
			_ = ty
			if ksfResp.Context != nil {
				cb.Cb.HelloWorld_Callback(&ret, ksfResp.Context)
				return nil
			} else {
				cb.Cb.HelloWorld_Callback(&ret)
				return nil
			}
		}()
		if err != nil {
			cb.Cb.HelloWorld_ExceptionCallback(err)
		}
	}
}
func (obj *DataCollectObj) AsyncSendResponse_HelloWorld(ctx context.Context, ret Error, opt ...map[string]string) (err error) {

	conn, udpAddr, ok := current.GetRawConn(ctx)
	if !ok {
		return fmt.Errorf("connection not found")
	}
	buf := codec.NewBuffer()

	err = ret.WriteBlock(buf, 0)
	if err != nil {
		return err
	}

	resp := &requestf.ResponsePacket{
		SBuffer: tools.ByteToInt8(buf.ToBytes()),
	}
	resp.IVersion = basef.KSFVERSION
	if resp.Status == nil {
		resp.Status = make(map[string]string)
	}
	resp.Status["KSF_FUNC"] = "HelloWorld"
	resp.SResultDesc = "HelloWorld"
	if len(opt) > 0 {
		if opt[0] != nil {
			resp.Context = opt[0]
		}
	}
	rspData := obj.rsp2Byte(resp)
	if udpAddr != nil {
		udpConn, _ := conn.(*net.UDPConn)
		_, err = udpConn.WriteToUDP(rspData, udpAddr)
	} else {
		_, err = conn.Write(rspData)
	}
	return err
}
