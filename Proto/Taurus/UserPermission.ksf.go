// Package Taurus comment
// This file was generated by ksf2go 1.3.21
// Generated from TaurusPermissionObj.ksf
package Taurus

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

// UserPermission struct
type UserPermission struct {
	servant m.Servant
}

// QryUserPermission is the proxy function for the method defined in the ksf file, with the context
func (obj *UserPermission) QryUserPermission(req *QryUserPermissionReq, rsp *QryUserPermissionRsp, opts ...map[string]string) (ret Result, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return ret, err
	}

	err = (*rsp).WriteBlock(buf, 2)
	if err != nil {
		return ret, err
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
	ksfCtx := context.Background()

	err = obj.servant.KsfInvoke(ksfCtx, 0, "QryUserPermission", buf.ToBytes(), statusMap, contextMap, ksfResp)
	if err != nil {
		return ret, err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
	err = ret.ReadBlock(readBuf, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*rsp).ReadBlock(readBuf, 2, true)
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

// QryUserPermissionWithContext is the proxy function for the method defined in the ksf file, with the context
func (obj *UserPermission) QryUserPermissionWithContext(ksfCtx context.Context, req *QryUserPermissionReq, rsp *QryUserPermissionRsp, opts ...map[string]string) (ret Result, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return ret, err
	}

	err = (*rsp).WriteBlock(buf, 2)
	if err != nil {
		return ret, err
	}

	traceData, ok := current.GetTraceData(ksfCtx)
	if ok && traceData.TraceCall {
		traceData.NewSpan()
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCS, uint(buf.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value["req"] = req
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		ksf.Trace(traceData.GetTraceKey(trace.EstCS), trace.TraceAnnotationCS, ksf.GetClientConfig().ModuleName, obj.servant.Name(), "QryUserPermission", 0, traceParam, "")
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
	err = obj.servant.KsfInvoke(ksfCtx, 0, "QryUserPermission", buf.ToBytes(), statusMap, contextMap, ksfResp)
	if err != nil {
		return ret, err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
	err = ret.ReadBlock(readBuf, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*rsp).ReadBlock(readBuf, 2, true)
	if err != nil {
		return ret, err
	}

	if ok && traceData.TraceCall {
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCR, uint(readBuf.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value[""] = ret
			value["rsp"] = *rsp
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		ksf.Trace(traceData.GetTraceKey(trace.EstCR), trace.TraceAnnotationCR, ksf.GetClientConfig().ModuleName, obj.servant.Name(), "QryUserPermission", int(ksfResp.IRet), traceParam, "")
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

// QryUserPermissionOneWayWithContext is the proxy function for the method defined in the ksf file, with the context
func (obj *UserPermission) QryUserPermissionOneWayWithContext(ksfCtx context.Context, req *QryUserPermissionReq, rsp *QryUserPermissionRsp, opts ...map[string]string) (ret Result, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return ret, err
	}

	err = (*rsp).WriteBlock(buf, 2)
	if err != nil {
		return ret, err
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
	err = obj.servant.KsfInvoke(ksfCtx, 1, "QryUserPermission", buf.ToBytes(), statusMap, contextMap, ksfResp)
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

// CheckUserPermission is the proxy function for the method defined in the ksf file, with the context
func (obj *UserPermission) CheckUserPermission(req *CheckUserPermissionReq, rsp *CheckUserPermissionRsp, opts ...map[string]string) (ret Result, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return ret, err
	}

	err = (*rsp).WriteBlock(buf, 2)
	if err != nil {
		return ret, err
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
	ksfCtx := context.Background()

	err = obj.servant.KsfInvoke(ksfCtx, 0, "CheckUserPermission", buf.ToBytes(), statusMap, contextMap, ksfResp)
	if err != nil {
		return ret, err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
	err = ret.ReadBlock(readBuf, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*rsp).ReadBlock(readBuf, 2, true)
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

// CheckUserPermissionWithContext is the proxy function for the method defined in the ksf file, with the context
func (obj *UserPermission) CheckUserPermissionWithContext(ksfCtx context.Context, req *CheckUserPermissionReq, rsp *CheckUserPermissionRsp, opts ...map[string]string) (ret Result, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return ret, err
	}

	err = (*rsp).WriteBlock(buf, 2)
	if err != nil {
		return ret, err
	}

	traceData, ok := current.GetTraceData(ksfCtx)
	if ok && traceData.TraceCall {
		traceData.NewSpan()
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCS, uint(buf.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value["req"] = req
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		ksf.Trace(traceData.GetTraceKey(trace.EstCS), trace.TraceAnnotationCS, ksf.GetClientConfig().ModuleName, obj.servant.Name(), "CheckUserPermission", 0, traceParam, "")
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
	err = obj.servant.KsfInvoke(ksfCtx, 0, "CheckUserPermission", buf.ToBytes(), statusMap, contextMap, ksfResp)
	if err != nil {
		return ret, err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
	err = ret.ReadBlock(readBuf, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*rsp).ReadBlock(readBuf, 2, true)
	if err != nil {
		return ret, err
	}

	if ok && traceData.TraceCall {
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCR, uint(readBuf.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value[""] = ret
			value["rsp"] = *rsp
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		ksf.Trace(traceData.GetTraceKey(trace.EstCR), trace.TraceAnnotationCR, ksf.GetClientConfig().ModuleName, obj.servant.Name(), "CheckUserPermission", int(ksfResp.IRet), traceParam, "")
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

// CheckUserPermissionOneWayWithContext is the proxy function for the method defined in the ksf file, with the context
func (obj *UserPermission) CheckUserPermissionOneWayWithContext(ksfCtx context.Context, req *CheckUserPermissionReq, rsp *CheckUserPermissionRsp, opts ...map[string]string) (ret Result, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return ret, err
	}

	err = (*rsp).WriteBlock(buf, 2)
	if err != nil {
		return ret, err
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
	err = obj.servant.KsfInvoke(ksfCtx, 1, "CheckUserPermission", buf.ToBytes(), statusMap, contextMap, ksfResp)
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

// CheckUserConfig is the proxy function for the method defined in the ksf file, with the context
func (obj *UserPermission) CheckUserConfig(req *CheckUserConfigReq, rsp *CheckUserConfigRsp, opts ...map[string]string) (ret Result, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return ret, err
	}

	err = (*rsp).WriteBlock(buf, 2)
	if err != nil {
		return ret, err
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
	ksfCtx := context.Background()

	err = obj.servant.KsfInvoke(ksfCtx, 0, "CheckUserConfig", buf.ToBytes(), statusMap, contextMap, ksfResp)
	if err != nil {
		return ret, err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
	err = ret.ReadBlock(readBuf, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*rsp).ReadBlock(readBuf, 2, true)
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

// CheckUserConfigWithContext is the proxy function for the method defined in the ksf file, with the context
func (obj *UserPermission) CheckUserConfigWithContext(ksfCtx context.Context, req *CheckUserConfigReq, rsp *CheckUserConfigRsp, opts ...map[string]string) (ret Result, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return ret, err
	}

	err = (*rsp).WriteBlock(buf, 2)
	if err != nil {
		return ret, err
	}

	traceData, ok := current.GetTraceData(ksfCtx)
	if ok && traceData.TraceCall {
		traceData.NewSpan()
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCS, uint(buf.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value["req"] = req
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		ksf.Trace(traceData.GetTraceKey(trace.EstCS), trace.TraceAnnotationCS, ksf.GetClientConfig().ModuleName, obj.servant.Name(), "CheckUserConfig", 0, traceParam, "")
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
	err = obj.servant.KsfInvoke(ksfCtx, 0, "CheckUserConfig", buf.ToBytes(), statusMap, contextMap, ksfResp)
	if err != nil {
		return ret, err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
	err = ret.ReadBlock(readBuf, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*rsp).ReadBlock(readBuf, 2, true)
	if err != nil {
		return ret, err
	}

	if ok && traceData.TraceCall {
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCR, uint(readBuf.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value[""] = ret
			value["rsp"] = *rsp
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		ksf.Trace(traceData.GetTraceKey(trace.EstCR), trace.TraceAnnotationCR, ksf.GetClientConfig().ModuleName, obj.servant.Name(), "CheckUserConfig", int(ksfResp.IRet), traceParam, "")
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

// CheckUserConfigOneWayWithContext is the proxy function for the method defined in the ksf file, with the context
func (obj *UserPermission) CheckUserConfigOneWayWithContext(ksfCtx context.Context, req *CheckUserConfigReq, rsp *CheckUserConfigRsp, opts ...map[string]string) (ret Result, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return ret, err
	}

	err = (*rsp).WriteBlock(buf, 2)
	if err != nil {
		return ret, err
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
	err = obj.servant.KsfInvoke(ksfCtx, 1, "CheckUserConfig", buf.ToBytes(), statusMap, contextMap, ksfResp)
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
func (obj *UserPermission) SetServant(servant m.Servant) {
	obj.servant = servant
}

// GetServant gets servant for the service.
func (obj *UserPermission) GetServant() (servant *m.Servant) {
	return &obj.servant
}

// SetOnConnectCallback
func (obj *UserPermission) SetOnConnectCallback(callback func(string)) {
	obj.servant.SetOnConnectCallback(callback)
}

// SetOnCloseCallback
func (obj *UserPermission) SetOnCloseCallback(callback func(string)) {
	obj.servant.SetOnCloseCallback(callback)
}

// SetKsfCallback
func (obj *UserPermission) SetKsfCallback(callback UserPermissionKsfCallback) {
	var push UserPermissionPushCallback
	push.Cb = callback
	obj.servant.SetKsfCallback(&push)
}

// SetPushCallback
func (obj *UserPermission) SetPushCallback(callback func([]byte)) {
	obj.servant.SetPushCallback(callback)
}
func (obj *UserPermission) req2Byte(rsp *requestf.ResponsePacket) []byte {
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

func (obj *UserPermission) rsp2Byte(rsp *requestf.ResponsePacket) []byte {
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
func (obj *UserPermission) KsfPing() {
	ksfCtx := context.Background()
	obj.servant.KsfPing(ksfCtx)
}

// KsfPingWithContext
func (obj *UserPermission) KsfPingWithContext(ksfCtx context.Context) {
	obj.servant.KsfPing(ksfCtx)
}

// KsfSetTimeout sets the timeout for the servant which is in ms.
func (obj *UserPermission) KsfSetTimeout(timeout int) {
	obj.servant.KsfSetTimeout(timeout)
}

// KsfSetVersion default as KSFVERSION,you can set JSONVERSION.
func (obj *UserPermission) KsfSetVersion(version int16) {
	obj.servant.KsfSetVersion(version)
}

// KsfSetProtocol sets the protocol for the servant.
func (obj *UserPermission) KsfSetProtocol(p m.Protocol) {
	obj.servant.KsfSetProtocol(p)
}

// AddServant adds servant  for the service.
func (obj *UserPermission) AddServant(imp UserPermissionServant, servantObj string) {
	ksf.AddServant(obj, imp, servantObj)
}

// AddServantWithContext adds servant  for the service with context.
func (obj *UserPermission) AddServantWithContext(imp UserPermissionServantWithContext, servantObj string) {
	ksf.AddServantWithContext(obj, imp, servantObj)
}

type UserPermissionServant interface {
	QryUserPermission(req *QryUserPermissionReq, rsp *QryUserPermissionRsp) (ret Result, err error)
	CheckUserPermission(req *CheckUserPermissionReq, rsp *CheckUserPermissionRsp) (ret Result, err error)
	CheckUserConfig(req *CheckUserConfigReq, rsp *CheckUserConfigRsp) (ret Result, err error)
}
type UserPermissionServantWithContext interface {
	QryUserPermission(ksfCtx context.Context, req *QryUserPermissionReq, rsp *QryUserPermissionRsp) (ret Result, err error)
	CheckUserPermission(ksfCtx context.Context, req *CheckUserPermissionReq, rsp *CheckUserPermissionRsp) (ret Result, err error)
	CheckUserConfig(ksfCtx context.Context, req *CheckUserConfigReq, rsp *CheckUserConfigRsp) (ret Result, err error)

	DoClose(ctx context.Context)
}

// Dispatch is used to call the server side implement for the method defined in the ksf file. withContext shows using context or not.
func (obj *UserPermission) Dispatch(ksfCtx context.Context, val interface{}, ksfReq *requestf.RequestPacket, ksfResp *requestf.ResponsePacket, withContext bool) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	readBuf := codec.NewReader(tools.Int8ToByte(ksfReq.SBuffer))
	buf := codec.NewBuffer()
	switch ksfReq.SFuncName {
	case "QryUserPermission":
		var req QryUserPermissionReq
		var rsp QryUserPermissionRsp

		if ksfReq.IVersion == basef.KSFVERSION {

			err = req.ReadBlock(readBuf, 1, true)
			if err != nil {
				return err
			}

		} else if ksfReq.IVersion == basef.KUPVERSION {
			reqKup := kup.NewUniAttribute()
			reqKup.Decode(readBuf)

			var kupBuffer []byte

			reqKup.GetBuffer("req", &kupBuffer)
			readBuf.Reset(kupBuffer)
			err = req.ReadBlock(readBuf, 0, true)
			if err != nil {
				return err
			}

		} else if ksfReq.IVersion == basef.JSONVERSION {
			var jsonData map[string]interface{}
			decoder := json.NewDecoder(bytes.NewReader(readBuf.ToBytes()))
			decoder.UseNumber()
			err = decoder.Decode(&jsonData)
			if err != nil {
				return fmt.Errorf("decode reqpacket failed, error: %+v", err)
			}
			{
				jsonStr, _ := json.Marshal(jsonData["req"])
				req.ResetDefault()
				if err = json.Unmarshal(jsonStr, &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("decode reqpacket fail, error version: %d", ksfReq.IVersion)
			return err
		}

		traceData, ok := current.GetTraceData(ksfCtx)
		if ok && traceData.TraceCall {
			var traceParam string
			traceParamFlag := traceData.NeedTraceParam(trace.EstSR, uint(readBuf.Len()))
			if traceParamFlag == trace.EnpNormal {
				value := map[string]interface{}{}
				value["req"] = req
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			ksf.Trace(traceData.GetTraceKey(trace.EstSR), trace.TraceAnnotationSR, ksf.GetClientConfig().ModuleName, ksfReq.SServantName, "QryUserPermission", 0, traceParam, "")
		}

		var funRet Result
		if !withContext {
			imp := val.(UserPermissionServant)
			funRet, err = imp.QryUserPermission(&req, &rsp)
		} else {
			imp := val.(UserPermissionServantWithContext)
			funRet, err = imp.QryUserPermission(ksfCtx, &req, &rsp)
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

			err = rsp.WriteBlock(buf, 2)
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
			err = rsp.WriteBlock(buf, 0)
			if err != nil {
				return err
			}

			rspKup.PutBuffer("rsp", buf.ToBytes())

			buf.Reset()
			err = rspKup.Encode(buf)
			if err != nil {
				return err
			}
		} else if ksfReq.IVersion == basef.JSONVERSION {
			rspJson := map[string]interface{}{}
			rspJson["ksf_ret"] = funRet
			rspJson["rsp"] = rsp

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
				value["rsp"] = rsp
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			ksf.Trace(traceData.GetTraceKey(trace.EstSS), trace.TraceAnnotationSS, ksf.GetClientConfig().ModuleName, ksfReq.SServantName, "QryUserPermission", 0, traceParam, "")
		}

	case "CheckUserPermission":
		var req CheckUserPermissionReq
		var rsp CheckUserPermissionRsp

		if ksfReq.IVersion == basef.KSFVERSION {

			err = req.ReadBlock(readBuf, 1, true)
			if err != nil {
				return err
			}

		} else if ksfReq.IVersion == basef.KUPVERSION {
			reqKup := kup.NewUniAttribute()
			reqKup.Decode(readBuf)

			var kupBuffer []byte

			reqKup.GetBuffer("req", &kupBuffer)
			readBuf.Reset(kupBuffer)
			err = req.ReadBlock(readBuf, 0, true)
			if err != nil {
				return err
			}

		} else if ksfReq.IVersion == basef.JSONVERSION {
			var jsonData map[string]interface{}
			decoder := json.NewDecoder(bytes.NewReader(readBuf.ToBytes()))
			decoder.UseNumber()
			err = decoder.Decode(&jsonData)
			if err != nil {
				return fmt.Errorf("decode reqpacket failed, error: %+v", err)
			}
			{
				jsonStr, _ := json.Marshal(jsonData["req"])
				req.ResetDefault()
				if err = json.Unmarshal(jsonStr, &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("decode reqpacket fail, error version: %d", ksfReq.IVersion)
			return err
		}

		traceData, ok := current.GetTraceData(ksfCtx)
		if ok && traceData.TraceCall {
			var traceParam string
			traceParamFlag := traceData.NeedTraceParam(trace.EstSR, uint(readBuf.Len()))
			if traceParamFlag == trace.EnpNormal {
				value := map[string]interface{}{}
				value["req"] = req
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			ksf.Trace(traceData.GetTraceKey(trace.EstSR), trace.TraceAnnotationSR, ksf.GetClientConfig().ModuleName, ksfReq.SServantName, "CheckUserPermission", 0, traceParam, "")
		}

		var funRet Result
		if !withContext {
			imp := val.(UserPermissionServant)
			funRet, err = imp.CheckUserPermission(&req, &rsp)
		} else {
			imp := val.(UserPermissionServantWithContext)
			funRet, err = imp.CheckUserPermission(ksfCtx, &req, &rsp)
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

			err = rsp.WriteBlock(buf, 2)
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
			err = rsp.WriteBlock(buf, 0)
			if err != nil {
				return err
			}

			rspKup.PutBuffer("rsp", buf.ToBytes())

			buf.Reset()
			err = rspKup.Encode(buf)
			if err != nil {
				return err
			}
		} else if ksfReq.IVersion == basef.JSONVERSION {
			rspJson := map[string]interface{}{}
			rspJson["ksf_ret"] = funRet
			rspJson["rsp"] = rsp

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
				value["rsp"] = rsp
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			ksf.Trace(traceData.GetTraceKey(trace.EstSS), trace.TraceAnnotationSS, ksf.GetClientConfig().ModuleName, ksfReq.SServantName, "CheckUserPermission", 0, traceParam, "")
		}

	case "CheckUserConfig":
		var req CheckUserConfigReq
		var rsp CheckUserConfigRsp

		if ksfReq.IVersion == basef.KSFVERSION {

			err = req.ReadBlock(readBuf, 1, true)
			if err != nil {
				return err
			}

		} else if ksfReq.IVersion == basef.KUPVERSION {
			reqKup := kup.NewUniAttribute()
			reqKup.Decode(readBuf)

			var kupBuffer []byte

			reqKup.GetBuffer("req", &kupBuffer)
			readBuf.Reset(kupBuffer)
			err = req.ReadBlock(readBuf, 0, true)
			if err != nil {
				return err
			}

		} else if ksfReq.IVersion == basef.JSONVERSION {
			var jsonData map[string]interface{}
			decoder := json.NewDecoder(bytes.NewReader(readBuf.ToBytes()))
			decoder.UseNumber()
			err = decoder.Decode(&jsonData)
			if err != nil {
				return fmt.Errorf("decode reqpacket failed, error: %+v", err)
			}
			{
				jsonStr, _ := json.Marshal(jsonData["req"])
				req.ResetDefault()
				if err = json.Unmarshal(jsonStr, &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("decode reqpacket fail, error version: %d", ksfReq.IVersion)
			return err
		}

		traceData, ok := current.GetTraceData(ksfCtx)
		if ok && traceData.TraceCall {
			var traceParam string
			traceParamFlag := traceData.NeedTraceParam(trace.EstSR, uint(readBuf.Len()))
			if traceParamFlag == trace.EnpNormal {
				value := map[string]interface{}{}
				value["req"] = req
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			ksf.Trace(traceData.GetTraceKey(trace.EstSR), trace.TraceAnnotationSR, ksf.GetClientConfig().ModuleName, ksfReq.SServantName, "CheckUserConfig", 0, traceParam, "")
		}

		var funRet Result
		if !withContext {
			imp := val.(UserPermissionServant)
			funRet, err = imp.CheckUserConfig(&req, &rsp)
		} else {
			imp := val.(UserPermissionServantWithContext)
			funRet, err = imp.CheckUserConfig(ksfCtx, &req, &rsp)
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

			err = rsp.WriteBlock(buf, 2)
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
			err = rsp.WriteBlock(buf, 0)
			if err != nil {
				return err
			}

			rspKup.PutBuffer("rsp", buf.ToBytes())

			buf.Reset()
			err = rspKup.Encode(buf)
			if err != nil {
				return err
			}
		} else if ksfReq.IVersion == basef.JSONVERSION {
			rspJson := map[string]interface{}{}
			rspJson["ksf_ret"] = funRet
			rspJson["rsp"] = rsp

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
				value["rsp"] = rsp
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			ksf.Trace(traceData.GetTraceKey(trace.EstSS), trace.TraceAnnotationSS, ksf.GetClientConfig().ModuleName, ksfReq.SServantName, "CheckUserConfig", 0, traceParam, "")
		}

	default:
		if ksfReq.SFuncName == "DoClose" {
			if withContext {
				imp := val.(UserPermissionServantWithContext)
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

type UserPermissionKsfCallback interface {
	QryUserPermission_Callback(ret *Result, rsp *QryUserPermissionRsp, opt ...map[string]string)
	QryUserPermission_ExceptionCallback(err error)
	CheckUserPermission_Callback(ret *Result, rsp *CheckUserPermissionRsp, opt ...map[string]string)
	CheckUserPermission_ExceptionCallback(err error)
	CheckUserConfig_Callback(ret *Result, rsp *CheckUserConfigRsp, opt ...map[string]string)
	CheckUserConfig_ExceptionCallback(err error)
}

// UserPermissionPushCallback struct
type UserPermissionPushCallback struct {
	Cb UserPermissionKsfCallback
}

func (cb *UserPermissionPushCallback) Ondispatch(ksfResp *requestf.ResponsePacket) {
	switch ksfResp.SResultDesc {
	case "QryUserPermission":
		err := func() error {
			var (
				length int32
				have   bool
				ty     byte
			)
			var err error
			readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
			var ret Result
			err = ret.ReadBlock(readBuf, 0, true)
			if err != nil {
				return err
			}

			var rsp QryUserPermissionRsp

			err = rsp.ReadBlock(readBuf, 2, true)
			if err != nil {
				return err
			}

			_ = length
			_ = have
			_ = ty
			if ksfResp.Context != nil {
				cb.Cb.QryUserPermission_Callback(&ret, &rsp, ksfResp.Context)
				return nil
			} else {
				cb.Cb.QryUserPermission_Callback(&ret, &rsp)
				return nil
			}
		}()
		if err != nil {
			cb.Cb.QryUserPermission_ExceptionCallback(err)
		}
	case "CheckUserPermission":
		err := func() error {
			var (
				length int32
				have   bool
				ty     byte
			)
			var err error
			readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
			var ret Result
			err = ret.ReadBlock(readBuf, 0, true)
			if err != nil {
				return err
			}

			var rsp CheckUserPermissionRsp

			err = rsp.ReadBlock(readBuf, 2, true)
			if err != nil {
				return err
			}

			_ = length
			_ = have
			_ = ty
			if ksfResp.Context != nil {
				cb.Cb.CheckUserPermission_Callback(&ret, &rsp, ksfResp.Context)
				return nil
			} else {
				cb.Cb.CheckUserPermission_Callback(&ret, &rsp)
				return nil
			}
		}()
		if err != nil {
			cb.Cb.CheckUserPermission_ExceptionCallback(err)
		}
	case "CheckUserConfig":
		err := func() error {
			var (
				length int32
				have   bool
				ty     byte
			)
			var err error
			readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
			var ret Result
			err = ret.ReadBlock(readBuf, 0, true)
			if err != nil {
				return err
			}

			var rsp CheckUserConfigRsp

			err = rsp.ReadBlock(readBuf, 2, true)
			if err != nil {
				return err
			}

			_ = length
			_ = have
			_ = ty
			if ksfResp.Context != nil {
				cb.Cb.CheckUserConfig_Callback(&ret, &rsp, ksfResp.Context)
				return nil
			} else {
				cb.Cb.CheckUserConfig_Callback(&ret, &rsp)
				return nil
			}
		}()
		if err != nil {
			cb.Cb.CheckUserConfig_ExceptionCallback(err)
		}
	}
}
func (obj *UserPermission) AsyncSendResponse_QryUserPermission(ctx context.Context, ret Result, rsp *QryUserPermissionRsp, opt ...map[string]string) (err error) {

	conn, udpAddr, ok := current.GetRawConn(ctx)
	if !ok {
		return fmt.Errorf("connection not found")
	}
	buf := codec.NewBuffer()

	err = ret.WriteBlock(buf, 0)
	if err != nil {
		return err
	}

	err = (*rsp).WriteBlock(buf, 2)
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
	resp.Status["KSF_FUNC"] = "QryUserPermission"
	resp.SResultDesc = "QryUserPermission"
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
func (obj *UserPermission) AsyncSendResponse_CheckUserPermission(ctx context.Context, ret Result, rsp *CheckUserPermissionRsp, opt ...map[string]string) (err error) {

	conn, udpAddr, ok := current.GetRawConn(ctx)
	if !ok {
		return fmt.Errorf("connection not found")
	}
	buf := codec.NewBuffer()

	err = ret.WriteBlock(buf, 0)
	if err != nil {
		return err
	}

	err = (*rsp).WriteBlock(buf, 2)
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
	resp.Status["KSF_FUNC"] = "CheckUserPermission"
	resp.SResultDesc = "CheckUserPermission"
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
func (obj *UserPermission) AsyncSendResponse_CheckUserConfig(ctx context.Context, ret Result, rsp *CheckUserConfigRsp, opt ...map[string]string) (err error) {

	conn, udpAddr, ok := current.GetRawConn(ctx)
	if !ok {
		return fmt.Errorf("connection not found")
	}
	buf := codec.NewBuffer()

	err = ret.WriteBlock(buf, 0)
	if err != nil {
		return err
	}

	err = (*rsp).WriteBlock(buf, 2)
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
	resp.Status["KSF_FUNC"] = "CheckUserConfig"
	resp.SResultDesc = "CheckUserConfig"
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
