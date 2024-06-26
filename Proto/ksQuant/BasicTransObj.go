// Package ksQuant comment
// This file was generated by ksf2go 1.3.20
// Generated from BasicTransObj.ksf
package ksQuant

import (
	"fmt"

	"go.k8sf.cloud/go/KsfGo/ksf/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// LastUpdateInfoReq struct implement
type LastUpdateInfoReq struct {
	Data_source string `json:"data_source"`
}

func (st *LastUpdateInfoReq) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *LastUpdateInfoReq) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Data_source, 0, false)
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
func (st *LastUpdateInfoReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require LastUpdateInfoReq, but not exist. tag %d", tag)
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
func (st *LastUpdateInfoReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Data_source, 0)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *LastUpdateInfoReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// LastUpdateInfoRsp struct implement
type LastUpdateInfoRsp struct {
	Data_source string `json:"data_source"`
	Timestamp   int64  `json:"timestamp"`
}

func (st *LastUpdateInfoRsp) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *LastUpdateInfoRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Data_source, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.Timestamp, 1, false)
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
func (st *LastUpdateInfoRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require LastUpdateInfoRsp, but not exist. tag %d", tag)
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
func (st *LastUpdateInfoRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Data_source, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.Timestamp, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *LastUpdateInfoRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
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
