// Package Base comment
// This file was generated by ksf2go 1.3.20
// Generated from PermUserAuth.ksf
package Base

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

// UserAuth struct
type UserAuth struct {
	servant m.Servant
}

// GetPermissionsForUser is the proxy function for the method defined in the ksf file, with the context
func (obj *UserAuth) GetPermissionsForUser(user string, domain string, permissions *[]UserPermission, opts ...map[string]string) (ret int32, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = buf.WriteString(user, 1)
	if err != nil {
		return ret, err
	}

	err = buf.WriteString(domain, 2)
	if err != nil {
		return ret, err
	}

	err = buf.WriteHead(codec.LIST, 3)
	if err != nil {
		return ret, err
	}

	err = buf.WriteInt32(int32(len(*permissions)), 0)
	if err != nil {
		return ret, err
	}

	for _, v := range *permissions {

		err = v.WriteBlock(buf, 0)
		if err != nil {
			return ret, err
		}

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

	err = obj.servant.KsfInvoke(ksfCtx, 0, "GetPermissionsForUser", buf.ToBytes(), statusMap, contextMap, ksfResp)
	if err != nil {
		return ret, err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
	err = readBuf.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	_, ty, err = readBuf.SkipToNoCheck(3, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = readBuf.ReadInt32(&length, 0, true)
		if err != nil {
			return ret, err
		}

		*permissions = make([]UserPermission, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = (*permissions)[i0].ReadBlock(readBuf, 0, false)
			if err != nil {
				return ret, err
			}

		}
	} else if ty == codec.SimpleList {
		err = fmt.Errorf("not support SimpleList type")
		if err != nil {
			return ret, err
		}

	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}

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

// GetPermissionsForUserWithContext is the proxy function for the method defined in the ksf file, with the context
func (obj *UserAuth) GetPermissionsForUserWithContext(ksfCtx context.Context, user string, domain string, permissions *[]UserPermission, opts ...map[string]string) (ret int32, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = buf.WriteString(user, 1)
	if err != nil {
		return ret, err
	}

	err = buf.WriteString(domain, 2)
	if err != nil {
		return ret, err
	}

	err = buf.WriteHead(codec.LIST, 3)
	if err != nil {
		return ret, err
	}

	err = buf.WriteInt32(int32(len(*permissions)), 0)
	if err != nil {
		return ret, err
	}

	for _, v := range *permissions {

		err = v.WriteBlock(buf, 0)
		if err != nil {
			return ret, err
		}

	}

	traceData, ok := current.GetTraceData(ksfCtx)
	if ok && traceData.TraceCall {
		traceData.NewSpan()
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCS, uint(buf.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value["user"] = user
			value["domain"] = domain
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		ksf.Trace(traceData.GetTraceKey(trace.EstCS), trace.TraceAnnotationCS, ksf.GetClientConfig().ModuleName, obj.servant.Name(), "GetPermissionsForUser", 0, traceParam, "")
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
	err = obj.servant.KsfInvoke(ksfCtx, 0, "GetPermissionsForUser", buf.ToBytes(), statusMap, contextMap, ksfResp)
	if err != nil {
		return ret, err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
	err = readBuf.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	_, ty, err = readBuf.SkipToNoCheck(3, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = readBuf.ReadInt32(&length, 0, true)
		if err != nil {
			return ret, err
		}

		*permissions = make([]UserPermission, length)
		for i1, e1 := int32(0), length; i1 < e1; i1++ {

			err = (*permissions)[i1].ReadBlock(readBuf, 0, false)
			if err != nil {
				return ret, err
			}

		}
	} else if ty == codec.SimpleList {
		err = fmt.Errorf("not support SimpleList type")
		if err != nil {
			return ret, err
		}

	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}

	}

	if ok && traceData.TraceCall {
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCR, uint(readBuf.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value[""] = ret
			value["permissions"] = *permissions
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		ksf.Trace(traceData.GetTraceKey(trace.EstCR), trace.TraceAnnotationCR, ksf.GetClientConfig().ModuleName, obj.servant.Name(), "GetPermissionsForUser", int(ksfResp.IRet), traceParam, "")
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

// GetPermissionsForUserOneWayWithContext is the proxy function for the method defined in the ksf file, with the context
func (obj *UserAuth) GetPermissionsForUserOneWayWithContext(ksfCtx context.Context, user string, domain string, permissions *[]UserPermission, opts ...map[string]string) (ret int32, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = buf.WriteString(user, 1)
	if err != nil {
		return ret, err
	}

	err = buf.WriteString(domain, 2)
	if err != nil {
		return ret, err
	}

	err = buf.WriteHead(codec.LIST, 3)
	if err != nil {
		return ret, err
	}

	err = buf.WriteInt32(int32(len(*permissions)), 0)
	if err != nil {
		return ret, err
	}

	for _, v := range *permissions {

		err = v.WriteBlock(buf, 0)
		if err != nil {
			return ret, err
		}

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
	err = obj.servant.KsfInvoke(ksfCtx, 1, "GetPermissionsForUser", buf.ToBytes(), statusMap, contextMap, ksfResp)
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
func (obj *UserAuth) CheckUserPermission(user string, domain string, object string, action string, opts ...map[string]string) (ret bool, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = buf.WriteString(user, 1)
	if err != nil {
		return ret, err
	}

	err = buf.WriteString(domain, 2)
	if err != nil {
		return ret, err
	}

	err = buf.WriteString(object, 3)
	if err != nil {
		return ret, err
	}

	err = buf.WriteString(action, 4)
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
	err = readBuf.ReadBool(&ret, 0, true)
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
func (obj *UserAuth) CheckUserPermissionWithContext(ksfCtx context.Context, user string, domain string, object string, action string, opts ...map[string]string) (ret bool, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = buf.WriteString(user, 1)
	if err != nil {
		return ret, err
	}

	err = buf.WriteString(domain, 2)
	if err != nil {
		return ret, err
	}

	err = buf.WriteString(object, 3)
	if err != nil {
		return ret, err
	}

	err = buf.WriteString(action, 4)
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
			value["user"] = user
			value["domain"] = domain
			value["object"] = object
			value["action"] = action
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
	err = readBuf.ReadBool(&ret, 0, true)
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
func (obj *UserAuth) CheckUserPermissionOneWayWithContext(ksfCtx context.Context, user string, domain string, object string, action string, opts ...map[string]string) (ret bool, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = buf.WriteString(user, 1)
	if err != nil {
		return ret, err
	}

	err = buf.WriteString(domain, 2)
	if err != nil {
		return ret, err
	}

	err = buf.WriteString(object, 3)
	if err != nil {
		return ret, err
	}

	err = buf.WriteString(action, 4)
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

// SetServant sets servant for the service.
func (obj *UserAuth) SetServant(servant m.Servant) {
	obj.servant = servant
}

// GetServant gets servant for the service.
func (obj *UserAuth) GetServant() (servant *m.Servant) {
	return &obj.servant
}

// SetOnConnectCallback
func (obj *UserAuth) SetOnConnectCallback(callback func(string)) {
	obj.servant.SetOnConnectCallback(callback)
}

// SetOnCloseCallback
func (obj *UserAuth) SetOnCloseCallback(callback func(string)) {
	obj.servant.SetOnCloseCallback(callback)
}

// SetKsfCallback
func (obj *UserAuth) SetKsfCallback(callback UserAuthKsfCallback) {
	var push UserAuthPushCallback
	push.Cb = callback
	obj.servant.SetKsfCallback(&push)
}

// SetPushCallback
func (obj *UserAuth) SetPushCallback(callback func([]byte)) {
	obj.servant.SetPushCallback(callback)
}
func (obj *UserAuth) req2Byte(rsp *requestf.ResponsePacket) []byte {
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

func (obj *UserAuth) rsp2Byte(rsp *requestf.ResponsePacket) []byte {
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
func (obj *UserAuth) KsfPing() {
	ksfCtx := context.Background()
	obj.servant.KsfPing(ksfCtx)
}

// KsfSetTimeout sets the timeout for the servant which is in ms.
func (obj *UserAuth) KsfSetTimeout(timeout int) {
	obj.servant.KsfSetTimeout(timeout)
}

// KsfSetVersion default as KSFVERSION,you can set JSONVERSION.
func (obj *UserAuth) KsfSetVersion(version int16) {
	obj.servant.KsfSetVersion(version)
}

// KsfSetProtocol sets the protocol for the servant.
func (obj *UserAuth) KsfSetProtocol(p m.Protocol) {
	obj.servant.KsfSetProtocol(p)
}

// AddServant adds servant  for the service.
func (obj *UserAuth) AddServant(imp UserAuthServant, servantObj string) {
	ksf.AddServant(obj, imp, servantObj)
}

// AddServantWithContext adds servant  for the service with context.
func (obj *UserAuth) AddServantWithContext(imp UserAuthServantWithContext, servantObj string) {
	ksf.AddServantWithContext(obj, imp, servantObj)
}

type UserAuthServant interface {
	GetPermissionsForUser(user string, domain string, permissions *[]UserPermission) (ret int32, err error)
	CheckUserPermission(user string, domain string, object string, action string) (ret bool, err error)
}
type UserAuthServantWithContext interface {
	GetPermissionsForUser(ksfCtx context.Context, user string, domain string, permissions *[]UserPermission) (ret int32, err error)
	CheckUserPermission(ksfCtx context.Context, user string, domain string, object string, action string) (ret bool, err error)

	DoClose(ctx context.Context)
}

// Dispatch is used to call the server side implement for the method defined in the ksf file. withContext shows using context or not.
func (obj *UserAuth) Dispatch(ksfCtx context.Context, val interface{}, ksfReq *requestf.RequestPacket, ksfResp *requestf.ResponsePacket, withContext bool) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	readBuf := codec.NewReader(tools.Int8ToByte(ksfReq.SBuffer))
	buf := codec.NewBuffer()
	switch ksfReq.SFuncName {
	case "GetPermissionsForUser":
		var user string
		var domain string
		var permissions []UserPermission
		permissions = make([]UserPermission, 0)

		if ksfReq.IVersion == basef.KSFVERSION {

			err = readBuf.ReadString(&user, 1, true)
			if err != nil {
				return err
			}

			err = readBuf.ReadString(&domain, 2, true)
			if err != nil {
				return err
			}

		} else if ksfReq.IVersion == basef.KUPVERSION {
			reqKup := kup.NewUniAttribute()
			reqKup.Decode(readBuf)

			var kupBuffer []byte

			reqKup.GetBuffer("user", &kupBuffer)
			readBuf.Reset(kupBuffer)
			err = readBuf.ReadString(&user, 0, true)
			if err != nil {
				return err
			}

			reqKup.GetBuffer("domain", &kupBuffer)
			readBuf.Reset(kupBuffer)
			err = readBuf.ReadString(&domain, 0, true)
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
				jsonStr, _ := json.Marshal(jsonData["user"])
				if err = json.Unmarshal(jsonStr, &user); err != nil {
					return err
				}
			}
			{
				jsonStr, _ := json.Marshal(jsonData["domain"])
				if err = json.Unmarshal(jsonStr, &domain); err != nil {
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
				value["user"] = user
				value["domain"] = domain
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			ksf.Trace(traceData.GetTraceKey(trace.EstSR), trace.TraceAnnotationSR, ksf.GetClientConfig().ModuleName, ksfReq.SServantName, "GetPermissionsForUser", 0, traceParam, "")
		}

		var funRet int32
		if !withContext {
			imp := val.(UserAuthServant)
			funRet, err = imp.GetPermissionsForUser(user, domain, &permissions)
		} else {
			imp := val.(UserAuthServantWithContext)
			funRet, err = imp.GetPermissionsForUser(ksfCtx, user, domain, &permissions)
		}

		if err != nil {
			return err
		}

		if ksfReq.IVersion == basef.KSFVERSION {
			buf.Reset()

			err = buf.WriteInt32(funRet, 0)
			if err != nil {
				return err
			}

			err = buf.WriteHead(codec.LIST, 3)
			if err != nil {
				return err
			}

			err = buf.WriteInt32(int32(len(permissions)), 0)
			if err != nil {
				return err
			}

			for _, v := range permissions {

				err = v.WriteBlock(buf, 0)
				if err != nil {
					return err
				}

			}

		} else if ksfReq.IVersion == basef.KUPVERSION {
			rspKup := kup.NewUniAttribute()

			err = buf.WriteInt32(funRet, 0)
			if err != nil {
				return err
			}

			rspKup.PutBuffer("", buf.ToBytes())
			rspKup.PutBuffer("ksf_ret", buf.ToBytes())

			buf.Reset()
			err = buf.WriteHead(codec.LIST, 0)
			if err != nil {
				return err
			}

			err = buf.WriteInt32(int32(len(permissions)), 0)
			if err != nil {
				return err
			}

			for _, v := range permissions {

				err = v.WriteBlock(buf, 0)
				if err != nil {
					return err
				}

			}
			rspKup.PutBuffer("permissions", buf.ToBytes())

			buf.Reset()
			err = rspKup.Encode(buf)
			if err != nil {
				return err
			}
		} else if ksfReq.IVersion == basef.JSONVERSION {
			rspJson := map[string]interface{}{}
			rspJson["ksf_ret"] = funRet
			rspJson["permissions"] = permissions

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
				value["permissions"] = permissions
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			ksf.Trace(traceData.GetTraceKey(trace.EstSS), trace.TraceAnnotationSS, ksf.GetClientConfig().ModuleName, ksfReq.SServantName, "GetPermissionsForUser", 0, traceParam, "")
		}

	case "CheckUserPermission":
		var user string
		var domain string
		var object string
		var action string

		if ksfReq.IVersion == basef.KSFVERSION {

			err = readBuf.ReadString(&user, 1, true)
			if err != nil {
				return err
			}

			err = readBuf.ReadString(&domain, 2, true)
			if err != nil {
				return err
			}

			err = readBuf.ReadString(&object, 3, true)
			if err != nil {
				return err
			}

			err = readBuf.ReadString(&action, 4, true)
			if err != nil {
				return err
			}

		} else if ksfReq.IVersion == basef.KUPVERSION {
			reqKup := kup.NewUniAttribute()
			reqKup.Decode(readBuf)

			var kupBuffer []byte

			reqKup.GetBuffer("user", &kupBuffer)
			readBuf.Reset(kupBuffer)
			err = readBuf.ReadString(&user, 0, true)
			if err != nil {
				return err
			}

			reqKup.GetBuffer("domain", &kupBuffer)
			readBuf.Reset(kupBuffer)
			err = readBuf.ReadString(&domain, 0, true)
			if err != nil {
				return err
			}

			reqKup.GetBuffer("object", &kupBuffer)
			readBuf.Reset(kupBuffer)
			err = readBuf.ReadString(&object, 0, true)
			if err != nil {
				return err
			}

			reqKup.GetBuffer("action", &kupBuffer)
			readBuf.Reset(kupBuffer)
			err = readBuf.ReadString(&action, 0, true)
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
				jsonStr, _ := json.Marshal(jsonData["user"])
				if err = json.Unmarshal(jsonStr, &user); err != nil {
					return err
				}
			}
			{
				jsonStr, _ := json.Marshal(jsonData["domain"])
				if err = json.Unmarshal(jsonStr, &domain); err != nil {
					return err
				}
			}
			{
				jsonStr, _ := json.Marshal(jsonData["object"])
				if err = json.Unmarshal(jsonStr, &object); err != nil {
					return err
				}
			}
			{
				jsonStr, _ := json.Marshal(jsonData["action"])
				if err = json.Unmarshal(jsonStr, &action); err != nil {
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
				value["user"] = user
				value["domain"] = domain
				value["object"] = object
				value["action"] = action
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			ksf.Trace(traceData.GetTraceKey(trace.EstSR), trace.TraceAnnotationSR, ksf.GetClientConfig().ModuleName, ksfReq.SServantName, "CheckUserPermission", 0, traceParam, "")
		}

		var funRet bool
		if !withContext {
			imp := val.(UserAuthServant)
			funRet, err = imp.CheckUserPermission(user, domain, object, action)
		} else {
			imp := val.(UserAuthServantWithContext)
			funRet, err = imp.CheckUserPermission(ksfCtx, user, domain, object, action)
		}

		if err != nil {
			return err
		}

		if ksfReq.IVersion == basef.KSFVERSION {
			buf.Reset()

			err = buf.WriteBool(funRet, 0)
			if err != nil {
				return err
			}

		} else if ksfReq.IVersion == basef.KUPVERSION {
			rspKup := kup.NewUniAttribute()

			err = buf.WriteBool(funRet, 0)
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
			ksf.Trace(traceData.GetTraceKey(trace.EstSS), trace.TraceAnnotationSS, ksf.GetClientConfig().ModuleName, ksfReq.SServantName, "CheckUserPermission", 0, traceParam, "")
		}

	default:
		if ksfReq.SFuncName == "DoClose" {
			if withContext {
				imp := val.(UserAuthServantWithContext)
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

type UserAuthKsfCallback interface {
	GetPermissionsForUser_Callback(ret *int32, permissions *[]UserPermission, opt ...map[string]string)
	GetPermissionsForUser_ExceptionCallback(err error)
	CheckUserPermission_Callback(ret *bool, opt ...map[string]string)
	CheckUserPermission_ExceptionCallback(err error)
}

// UserAuthPushCallback struct
type UserAuthPushCallback struct {
	Cb UserAuthKsfCallback
}

func (cb *UserAuthPushCallback) Ondispatch(ksfResp *requestf.ResponsePacket) {
	switch ksfResp.SResultDesc {
	case "GetPermissionsForUser":
		err := func() error {
			var (
				length int32
				have   bool
				ty     byte
			)
			var err error
			readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
			var ret int32
			err = readBuf.ReadInt32(&ret, 0, true)
			if err != nil {
				return err
			}

			var permissions []UserPermission

			_, ty, err = readBuf.SkipToNoCheck(3, true)
			if err != nil {
				return err
			}

			if ty == codec.LIST {
				err = readBuf.ReadInt32(&length, 0, true)
				if err != nil {
					return err
				}

				permissions = make([]UserPermission, length)
				for i2, e2 := int32(0), length; i2 < e2; i2++ {

					err = permissions[i2].ReadBlock(readBuf, 0, false)
					if err != nil {
						return err
					}

				}
			} else if ty == codec.SimpleList {
				err = fmt.Errorf("not support SimpleList type")
				if err != nil {
					return err
				}

			} else {
				err = fmt.Errorf("require vector, but not")
				if err != nil {
					return err
				}

			}

			_ = length
			_ = have
			_ = ty
			if ksfResp.Context != nil {
				cb.Cb.GetPermissionsForUser_Callback(&ret, &permissions, ksfResp.Context)
				return nil
			} else {
				cb.Cb.GetPermissionsForUser_Callback(&ret, &permissions)
				return nil
			}
		}()
		if err != nil {
			cb.Cb.GetPermissionsForUser_ExceptionCallback(err)
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
			var ret bool
			err = readBuf.ReadBool(&ret, 0, true)
			if err != nil {
				return err
			}

			_ = length
			_ = have
			_ = ty
			if ksfResp.Context != nil {
				cb.Cb.CheckUserPermission_Callback(&ret, ksfResp.Context)
				return nil
			} else {
				cb.Cb.CheckUserPermission_Callback(&ret)
				return nil
			}
		}()
		if err != nil {
			cb.Cb.CheckUserPermission_ExceptionCallback(err)
		}
	}
}
func (obj *UserAuth) AsyncSendResponse_GetPermissionsForUser(ctx context.Context, ret int32, permissions *[]UserPermission, opt ...map[string]string) (err error) {

	conn, udpAddr, ok := current.GetRawConn(ctx)
	if !ok {
		return fmt.Errorf("connection not found")
	}
	buf := codec.NewBuffer()

	err = buf.WriteInt32(ret, 0)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 3)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(*permissions)), 0)
	if err != nil {
		return err
	}

	for _, v := range *permissions {

		err = v.WriteBlock(buf, 0)
		if err != nil {
			return err
		}

	}
	resp := &requestf.ResponsePacket{
		SBuffer: tools.ByteToInt8(buf.ToBytes()),
	}
	resp.IVersion = basef.KSFVERSION
	if resp.Status == nil {
		resp.Status = make(map[string]string)
	}
	resp.Status["KSF_FUNC"] = "GetPermissionsForUser"
	resp.SResultDesc = "GetPermissionsForUser"
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
func (obj *UserAuth) AsyncSendResponse_CheckUserPermission(ctx context.Context, ret bool, opt ...map[string]string) (err error) {

	conn, udpAddr, ok := current.GetRawConn(ctx)
	if !ok {
		return fmt.Errorf("connection not found")
	}
	buf := codec.NewBuffer()

	err = buf.WriteBool(ret, 0)
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
