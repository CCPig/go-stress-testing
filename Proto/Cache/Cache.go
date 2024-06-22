// Package Cache comment
// This file was generated by ksf2go 1.3.20
// Generated from Cache.ksf
package Cache

import (
	"fmt"

	"go.k8sf.cloud/go/KsfGo/ksf/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// const as define in ksf file
const (
	LEASE_TYPE_DURATION string = "duration"
	LEASE_TYPE_UTILTIME string = "utiltime"
	EVENT_TYPE_PUT      string = "PUT"
	EVENT_TYPE_DEL      string = "DELETE"
)

// Result struct implement
type Result struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func (st *Result) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *Result) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadInt32(&st.Code, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Msg, 1, false)
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
func (st *Result) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require Result, but not exist. tag %d", tag)
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
func (st *Result) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteInt32(st.Code, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Msg, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *Result) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// GetReq struct implement
type GetReq struct {
	K      string `json:"k"`
	Prefix string `json:"prefix"`
}

func (st *GetReq) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *GetReq) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.K, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Prefix, 1, false)
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
func (st *GetReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require GetReq, but not exist. tag %d", tag)
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
func (st *GetReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.K, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Prefix, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *GetReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// GetRsp struct implement
type GetRsp struct {
	Kv  map[string]string `json:"kv"`
	Ret Result            `json:"ret"`
}

func (st *GetRsp) ResetDefault() {
	st.Ret.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *GetRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.MAP, 0, false)
	if err != nil {
		return err
	}

	if have {
		err = readBuf.ReadInt32(&length, 0, true)
		if err != nil {
			return err
		}

		st.Kv = make(map[string]string)
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

			st.Kv[k0] = v0
		}
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
func (st *GetRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require GetRsp, but not exist. tag %d", tag)
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
func (st *GetRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteHead(codec.MAP, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Kv)), 0)
	if err != nil {
		return err
	}

	for k1, v1 := range st.Kv {

		err = buf.WriteString(k1, 0)
		if err != nil {
			return err
		}

		err = buf.WriteString(v1, 1)
		if err != nil {
			return err
		}

	}

	err = st.Ret.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *GetRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// DelReq struct implement
type DelReq struct {
	Ks     []string `json:"ks"`
	Prefix string   `json:"prefix"`
}

func (st *DelReq) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *DelReq) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	have, ty, err = readBuf.SkipToNoCheck(0, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			st.Ks = make([]string, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = readBuf.ReadString(&st.Ks[i0], 0, false)
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

	err = readBuf.ReadString(&st.Prefix, 1, false)
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
func (st *DelReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require DelReq, but not exist. tag %d", tag)
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
func (st *DelReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteHead(codec.LIST, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Ks)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Ks {

		err = buf.WriteString(v, 0)
		if err != nil {
			return err
		}

	}

	err = buf.WriteString(st.Prefix, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *DelReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// DelRsp struct implement
type DelRsp struct {
	Ks  []string `json:"ks"`
	Ret Result   `json:"ret"`
}

func (st *DelRsp) ResetDefault() {
	st.Ret.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *DelRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	have, ty, err = readBuf.SkipToNoCheck(0, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			st.Ks = make([]string, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = readBuf.ReadString(&st.Ks[i0], 0, false)
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
func (st *DelRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require DelRsp, but not exist. tag %d", tag)
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
func (st *DelRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteHead(codec.LIST, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Ks)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Ks {

		err = buf.WriteString(v, 0)
		if err != nil {
			return err
		}

	}

	err = st.Ret.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *DelRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// PutReq struct implement
type PutReq struct {
	Kvs map[string]string `json:"kvs"`
}

func (st *PutReq) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *PutReq) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.MAP, 0, false)
	if err != nil {
		return err
	}

	if have {
		err = readBuf.ReadInt32(&length, 0, true)
		if err != nil {
			return err
		}

		st.Kvs = make(map[string]string)
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

			st.Kvs[k0] = v0
		}
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *PutReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require PutReq, but not exist. tag %d", tag)
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
func (st *PutReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteHead(codec.MAP, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Kvs)), 0)
	if err != nil {
		return err
	}

	for k1, v1 := range st.Kvs {

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
func (st *PutReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// PutRsp struct implement
type PutRsp struct {
	Kvs map[string]string `json:"kvs"`
	Ret Result            `json:"ret"`
}

func (st *PutRsp) ResetDefault() {
	st.Ret.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *PutRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.MAP, 0, false)
	if err != nil {
		return err
	}

	if have {
		err = readBuf.ReadInt32(&length, 0, true)
		if err != nil {
			return err
		}

		st.Kvs = make(map[string]string)
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

			st.Kvs[k0] = v0
		}
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
func (st *PutRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require PutRsp, but not exist. tag %d", tag)
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
func (st *PutRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteHead(codec.MAP, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Kvs)), 0)
	if err != nil {
		return err
	}

	for k1, v1 := range st.Kvs {

		err = buf.WriteString(k1, 0)
		if err != nil {
			return err
		}

		err = buf.WriteString(v1, 1)
		if err != nil {
			return err
		}

	}

	err = st.Ret.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *PutRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// PutWithLeaseReq struct implement
type PutWithLeaseReq struct {
	Kvs           map[string]string `json:"kvs"`
	Lease_type    string            `json:"lease_type"`
	Expire_time   int64             `json:"expire_time"`
	Alive_seconds int64             `json:"alive_seconds"`
}

func (st *PutWithLeaseReq) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *PutWithLeaseReq) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.MAP, 0, false)
	if err != nil {
		return err
	}

	if have {
		err = readBuf.ReadInt32(&length, 0, true)
		if err != nil {
			return err
		}

		st.Kvs = make(map[string]string)
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

			st.Kvs[k0] = v0
		}
	}

	err = readBuf.ReadString(&st.Lease_type, 1, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.Expire_time, 2, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.Alive_seconds, 3, false)
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
func (st *PutWithLeaseReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require PutWithLeaseReq, but not exist. tag %d", tag)
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
func (st *PutWithLeaseReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteHead(codec.MAP, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Kvs)), 0)
	if err != nil {
		return err
	}

	for k1, v1 := range st.Kvs {

		err = buf.WriteString(k1, 0)
		if err != nil {
			return err
		}

		err = buf.WriteString(v1, 1)
		if err != nil {
			return err
		}

	}

	err = buf.WriteString(st.Lease_type, 1)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.Expire_time, 2)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.Alive_seconds, 3)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *PutWithLeaseReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// PutWithLeaseRsp struct implement
type PutWithLeaseRsp struct {
	Lease_id int64             `json:"lease_id"`
	Ttl      int64             `json:"ttl"`
	Kvs      map[string]string `json:"kvs"`
	Ret      Result            `json:"ret"`
}

func (st *PutWithLeaseRsp) ResetDefault() {
	st.Ret.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *PutWithLeaseRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadInt64(&st.Lease_id, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.Ttl, 1, false)
	if err != nil {
		return err
	}

	have, err = readBuf.SkipTo(codec.MAP, 2, false)
	if err != nil {
		return err
	}

	if have {
		err = readBuf.ReadInt32(&length, 0, true)
		if err != nil {
			return err
		}

		st.Kvs = make(map[string]string)
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

			st.Kvs[k0] = v0
		}
	}

	err = st.Ret.ReadBlock(readBuf, 3, false)
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
func (st *PutWithLeaseRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require PutWithLeaseRsp, but not exist. tag %d", tag)
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
func (st *PutWithLeaseRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteInt64(st.Lease_id, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.Ttl, 1)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.MAP, 2)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Kvs)), 0)
	if err != nil {
		return err
	}

	for k1, v1 := range st.Kvs {

		err = buf.WriteString(k1, 0)
		if err != nil {
			return err
		}

		err = buf.WriteString(v1, 1)
		if err != nil {
			return err
		}

	}

	err = st.Ret.WriteBlock(buf, 3)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *PutWithLeaseRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// WatchItem struct implement
type WatchItem struct {
	K           string `json:"k"`
	With_prefix bool   `json:"with_prefix"`
}

func (st *WatchItem) ResetDefault() {
	st.With_prefix = false
}

// ReadFrom reads  from readBuf and put into struct.
func (st *WatchItem) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.K, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadBool(&st.With_prefix, 1, false)
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
func (st *WatchItem) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require WatchItem, but not exist. tag %d", tag)
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
func (st *WatchItem) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.K, 0)
	if err != nil {
		return err
	}

	err = buf.WriteBool(st.With_prefix, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *WatchItem) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// WatchReq struct implement
type WatchReq struct {
	Items []WatchItem `json:"items"`
}

func (st *WatchReq) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *WatchReq) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	have, ty, err = readBuf.SkipToNoCheck(0, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			st.Items = make([]WatchItem, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = st.Items[i0].ReadBlock(readBuf, 0, false)
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
func (st *WatchReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require WatchReq, but not exist. tag %d", tag)
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
func (st *WatchReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteHead(codec.LIST, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Items)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Items {

		err = v.WriteBlock(buf, 0)
		if err != nil {
			return err
		}

	}

	return err
}

// WriteBlock encode struct
func (st *WatchReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// WatchRsp struct implement
type WatchRsp struct {
	Ret Result `json:"ret"`
}

func (st *WatchRsp) ResetDefault() {
	st.Ret.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *WatchRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = st.Ret.ReadBlock(readBuf, 0, false)
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
func (st *WatchRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require WatchRsp, but not exist. tag %d", tag)
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
func (st *WatchRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = st.Ret.WriteBlock(buf, 0)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *WatchRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// UnWatchReq struct implement
type UnWatchReq struct {
	Items []WatchItem `json:"items"`
}

func (st *UnWatchReq) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *UnWatchReq) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	have, ty, err = readBuf.SkipToNoCheck(0, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			st.Items = make([]WatchItem, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = st.Items[i0].ReadBlock(readBuf, 0, false)
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
func (st *UnWatchReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require UnWatchReq, but not exist. tag %d", tag)
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
func (st *UnWatchReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteHead(codec.LIST, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Items)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Items {

		err = v.WriteBlock(buf, 0)
		if err != nil {
			return err
		}

	}

	return err
}

// WriteBlock encode struct
func (st *UnWatchReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// UnWatchRsp struct implement
type UnWatchRsp struct {
	Ret Result `json:"ret"`
}

func (st *UnWatchRsp) ResetDefault() {
	st.Ret.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *UnWatchRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = st.Ret.ReadBlock(readBuf, 0, false)
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
func (st *UnWatchRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require UnWatchRsp, but not exist. tag %d", tag)
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
func (st *UnWatchRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = st.Ret.WriteBlock(buf, 0)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *UnWatchRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// Event struct implement
type Event struct {
	Type          string `json:"type"`
	K             string `json:"k"`
	V             string `json:"v"`
	CreateVersion int64  `json:"createVersion"`
	ModVersion    int64  `json:"modVersion"`
	Version       int64  `json:"version"`
	LeaseId       int64  `json:"leaseId"`
}

func (st *Event) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *Event) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Type, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.K, 1, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.V, 2, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.CreateVersion, 3, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.ModVersion, 4, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.Version, 5, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.LeaseId, 6, false)
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
func (st *Event) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require Event, but not exist. tag %d", tag)
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
func (st *Event) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Type, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.K, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.V, 2)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.CreateVersion, 3)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.ModVersion, 4)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.Version, 5)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.LeaseId, 6)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *Event) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// LeaderMark struct implement
type LeaderMark struct {
	ElectionKey string `json:"ElectionKey"`
	LeaderMark  string `json:"LeaderMark"`
}

func (st *LeaderMark) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *LeaderMark) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.ElectionKey, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.LeaderMark, 1, false)
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
func (st *LeaderMark) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require LeaderMark, but not exist. tag %d", tag)
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
func (st *LeaderMark) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.ElectionKey, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.LeaderMark, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *LeaderMark) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// LaunchElectionReq struct implement
type LaunchElectionReq struct {
	ElectionKey string `json:"ElectionKey"`
	MarkKey     string `json:"MarkKey"`
}

func (st *LaunchElectionReq) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *LaunchElectionReq) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.ElectionKey, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.MarkKey, 1, false)
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
func (st *LaunchElectionReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require LaunchElectionReq, but not exist. tag %d", tag)
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
func (st *LaunchElectionReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.ElectionKey, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.MarkKey, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *LaunchElectionReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// LaunchElectionRsp struct implement
type LaunchElectionRsp struct {
	Ret Result `json:"ret"`
}

func (st *LaunchElectionRsp) ResetDefault() {
	st.Ret.ResetDefault()
}

// ReadFrom reads  from readBuf and put into struct.
func (st *LaunchElectionRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = st.Ret.ReadBlock(readBuf, 0, false)
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
func (st *LaunchElectionRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require LaunchElectionRsp, but not exist. tag %d", tag)
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
func (st *LaunchElectionRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = st.Ret.WriteBlock(buf, 0)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *LaunchElectionRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
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
