// Package Taurus comment
// This file was generated by ksf2go 1.3.21
// Generated from XBond.ksf
package Taurus

import (
	"fmt"

	"go.k8sf.cloud/go/KsfGo/ksf/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// NewTask struct implement
type NewTask struct {
	Request_id string    `json:"request_id"`
	Task_param TaskParam `json:"task_param"`
}

func (st *NewTask) ResetDefault() {
	st.Task_param.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *NewTask) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	err = st.Task_param.ReadBlock(readBuf, 1, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *NewTask) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require NewTask, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *NewTask) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = st.Task_param.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *NewTask) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// NewTaskRsp struct implement
type NewTaskRsp struct {
	Request_id string `json:"request_id"`
	Ret        Result `json:"ret"`
}

func (st *NewTaskRsp) ResetDefault() {
	st.Ret.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *NewTaskRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	err = st.Ret.ReadBlock(readBuf, 1, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *NewTaskRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require NewTaskRsp, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *NewTaskRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = st.Ret.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *NewTaskRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// ControlTask struct implement
type ControlTask struct {
	Request_id   string            `json:"request_id"`
	Task_id      string            `json:"task_id"`
	Command_type string            `json:"command_type"`
	Extra_params map[string]string `json:"extra_params"`
}

func (st *ControlTask) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *ControlTask) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Task_id, 1, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Command_type, 2, false)
	if err != nil {
		return err
	}

	have, err = readBuf.SkipTo(codec.MAP, 3, false)
	if err != nil {
		return err
	}

	if have {
		err = readBuf.ReadInt32(&length, 0, true)
		if err != nil {
			return err
		}

		st.Extra_params = make(map[string]string)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {
			var k0 string
			var v0 string

			err = readBuf.ReadString(&k0, 0, false)
			if err != nil {
				return err
			}

			err = readBuf.ReadString(&v0, 1, false)
			if err != nil {
				return err
			}

			st.Extra_params[k0] = v0
		}
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *ControlTask) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require ControlTask, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *ControlTask) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Task_id, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Command_type, 2)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.MAP, 3)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Extra_params)), 0)
	if err != nil {
		return err
	}

	for k1, v1 := range st.Extra_params {

		err = buf.WriteString(k1, 0)
		if err != nil {
			return err
		}

		err = buf.WriteString(v1, 1)
		if err != nil {
			return err
		}

	}

	return err
}

// WriteBlock encode struct
func (st *ControlTask) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// ControlTaskRsp struct implement
type ControlTaskRsp struct {
	Request_id string `json:"request_id"`
	Ret        Result `json:"ret"`
}

func (st *ControlTaskRsp) ResetDefault() {
	st.Ret.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *ControlTaskRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	err = st.Ret.ReadBlock(readBuf, 1, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *ControlTaskRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require ControlTaskRsp, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *ControlTaskRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = st.Ret.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *ControlTaskRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// NewChildOrder struct implement
type NewChildOrder struct {
	Request_id       string         `json:"request_id"`
	Insert_order_req InsertOrderReq `json:"insert_order_req"`
}

func (st *NewChildOrder) ResetDefault() {
	st.Insert_order_req.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *NewChildOrder) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	err = st.Insert_order_req.ReadBlock(readBuf, 1, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *NewChildOrder) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require NewChildOrder, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *NewChildOrder) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = st.Insert_order_req.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *NewChildOrder) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// NewChildOrderRsp struct implement
type NewChildOrderRsp struct {
	Request_id       string         `json:"request_id"`
	Ret              Result         `json:"ret"`
	Insert_order_rsp InsertOrderRsp `json:"insert_order_rsp"`
}

func (st *NewChildOrderRsp) ResetDefault() {
	st.Ret.ResetDefault()
	st.Insert_order_rsp.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *NewChildOrderRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	err = st.Ret.ReadBlock(readBuf, 1, false)
	if err != nil {
		return err
	}

	err = st.Insert_order_rsp.ReadBlock(readBuf, 2, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *NewChildOrderRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require NewChildOrderRsp, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *NewChildOrderRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = st.Ret.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	err = st.Insert_order_rsp.WriteBlock(buf, 2)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *NewChildOrderRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// CancelChildOrder struct implement
type CancelChildOrder struct {
	Request_id       string         `json:"request_id"`
	Cancel_order_req CancelOrderReq `json:"cancel_order_req"`
}

func (st *CancelChildOrder) ResetDefault() {
	st.Cancel_order_req.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *CancelChildOrder) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	err = st.Cancel_order_req.ReadBlock(readBuf, 1, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *CancelChildOrder) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require CancelChildOrder, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *CancelChildOrder) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = st.Cancel_order_req.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *CancelChildOrder) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// CancelChildOrderRsp struct implement
type CancelChildOrderRsp struct {
	Request_id       string         `json:"request_id"`
	Ret              Result         `json:"ret"`
	Cancel_order_rsp CancelOrderRsp `json:"cancel_order_rsp"`
}

func (st *CancelChildOrderRsp) ResetDefault() {
	st.Ret.ResetDefault()
	st.Cancel_order_rsp.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *CancelChildOrderRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	err = st.Ret.ReadBlock(readBuf, 1, false)
	if err != nil {
		return err
	}

	err = st.Cancel_order_rsp.ReadBlock(readBuf, 2, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *CancelChildOrderRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require CancelChildOrderRsp, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *CancelChildOrderRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = st.Ret.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	err = st.Cancel_order_rsp.WriteBlock(buf, 2)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *CancelChildOrderRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// ChildOrderUpdate struct implement
type ChildOrderUpdate struct {
	Request_id string    `json:"request_id"`
	Order_info OrderInfo `json:"order_info"`
}

func (st *ChildOrderUpdate) ResetDefault() {
	st.Order_info.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *ChildOrderUpdate) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	err = st.Order_info.ReadBlock(readBuf, 1, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *ChildOrderUpdate) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require ChildOrderUpdate, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *ChildOrderUpdate) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = st.Order_info.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *ChildOrderUpdate) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// ChildOrderUpdateRsp struct implement
type ChildOrderUpdateRsp struct {
	Request_id string `json:"request_id"`
	Ret        Result `json:"ret"`
}

func (st *ChildOrderUpdateRsp) ResetDefault() {
	st.Ret.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *ChildOrderUpdateRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	err = st.Ret.ReadBlock(readBuf, 1, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *ChildOrderUpdateRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require ChildOrderUpdateRsp, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *ChildOrderUpdateRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = st.Ret.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *ChildOrderUpdateRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// ChildTradeUpdate struct implement
type ChildTradeUpdate struct {
	Request_id string    `json:"request_id"`
	Trade_info TradeInfo `json:"trade_info"`
}

func (st *ChildTradeUpdate) ResetDefault() {
	st.Trade_info.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *ChildTradeUpdate) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	err = st.Trade_info.ReadBlock(readBuf, 1, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *ChildTradeUpdate) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require ChildTradeUpdate, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *ChildTradeUpdate) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = st.Trade_info.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *ChildTradeUpdate) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// ChildTradeUpdateRsp struct implement
type ChildTradeUpdateRsp struct {
	Request_id string `json:"request_id"`
	Ret        Result `json:"ret"`
}

func (st *ChildTradeUpdateRsp) ResetDefault() {
	st.Ret.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *ChildTradeUpdateRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	err = st.Ret.ReadBlock(readBuf, 1, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *ChildTradeUpdateRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require ChildTradeUpdateRsp, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *ChildTradeUpdateRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = st.Ret.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *ChildTradeUpdateRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// TaskStatusUpdate struct implement
type TaskStatusUpdate struct {
	Request_id  string     `json:"request_id"`
	Task_status TaskStatus `json:"task_status"`
}

func (st *TaskStatusUpdate) ResetDefault() {
	st.Task_status.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *TaskStatusUpdate) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	err = st.Task_status.ReadBlock(readBuf, 1, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *TaskStatusUpdate) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require TaskStatusUpdate, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *TaskStatusUpdate) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = st.Task_status.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *TaskStatusUpdate) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// TaskStatusUpdateRsp struct implement
type TaskStatusUpdateRsp struct {
	Request_id string `json:"request_id"`
	Ret        Result `json:"ret"`
}

func (st *TaskStatusUpdateRsp) ResetDefault() {
	st.Ret.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *TaskStatusUpdateRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	err = st.Ret.ReadBlock(readBuf, 1, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *TaskStatusUpdateRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require TaskStatusUpdateRsp, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *TaskStatusUpdateRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = st.Ret.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *TaskStatusUpdateRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// QryTaskStatuses struct implement
type QryTaskStatuses struct {
	Request_id string   `json:"request_id"`
	Task_id    []string `json:"task_id"`
}

func (st *QryTaskStatuses) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *QryTaskStatuses) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	have, ty, err = readBuf.SkipToNoCheck(1, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			st.Task_id = make([]string, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = readBuf.ReadString(&st.Task_id[i0], 0, false)
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
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *QryTaskStatuses) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require QryTaskStatuses, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *QryTaskStatuses) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 1)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Task_id)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Task_id {

		err = buf.WriteString(v, 0)
		if err != nil {
			return err
		}

	}

	return err
}

// WriteBlock encode struct
func (st *QryTaskStatuses) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// QryTaskStatusesRsp struct implement
type QryTaskStatusesRsp struct {
	Request_id  string       `json:"request_id"`
	Ret         Result       `json:"ret"`
	Task_status []TaskStatus `json:"task_status"`
}

func (st *QryTaskStatusesRsp) ResetDefault() {
	st.Ret.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *QryTaskStatusesRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	err = st.Ret.ReadBlock(readBuf, 1, false)
	if err != nil {
		return err
	}

	have, ty, err = readBuf.SkipToNoCheck(2, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			st.Task_status = make([]TaskStatus, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = st.Task_status[i0].ReadBlock(readBuf, 0, false)
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
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *QryTaskStatusesRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require QryTaskStatusesRsp, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *QryTaskStatusesRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = st.Ret.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 2)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Task_status)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Task_status {

		err = v.WriteBlock(buf, 0)
		if err != nil {
			return err
		}

	}

	return err
}

// WriteBlock encode struct
func (st *QryTaskStatusesRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}
