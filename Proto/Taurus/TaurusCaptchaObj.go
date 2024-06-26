// Package Taurus comment
// This file was generated by ksf2go 1.3.21
// Generated from TaurusCaptchaObj.ksf
package Taurus

import (
	"fmt"

	"go.k8sf.cloud/go/KsfGo/ksf/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// GenerateCaptchaReq struct implement
type GenerateCaptchaReq struct {
	Captcha string `json:"captcha"`
	Length  int32  `json:"length"`
}

func (st *GenerateCaptchaReq) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *GenerateCaptchaReq) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Captcha, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Length, 1, false)
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
func (st *GenerateCaptchaReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require GenerateCaptchaReq, but not exist. tag %d", tag)
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
func (st *GenerateCaptchaReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Captcha, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Length, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *GenerateCaptchaReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// GenerateCaptchaRsp struct implement
type GenerateCaptchaRsp struct {
	Captcha_id string `json:"captcha_id"`
	Image      []int8 `json:"image"`
}

func (st *GenerateCaptchaRsp) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *GenerateCaptchaRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Captcha_id, 0, false)
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

			st.Image = make([]int8, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = readBuf.ReadInt8(&st.Image[i0], 0, false)
				if err != nil {
					return err
				}

			}
		} else if ty == codec.SimpleList {

			_, err = readBuf.SkipTo(codec.BYTE, 0, true)
			if err != nil {
				return err
			}

			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			err = readBuf.ReadSliceInt8(&st.Image, length, true)
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
func (st *GenerateCaptchaRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require GenerateCaptchaRsp, but not exist. tag %d", tag)
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
func (st *GenerateCaptchaRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Captcha_id, 0)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.SimpleList, 1)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.BYTE, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Image)), 0)
	if err != nil {
		return err
	}

	err = buf.WriteSliceInt8(st.Image)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *GenerateCaptchaRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// VerifyCaptchaReq struct implement
type VerifyCaptchaReq struct {
	Captcha_id string `json:"captcha_id"`
	Code       string `json:"code"`
}

func (st *VerifyCaptchaReq) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *VerifyCaptchaReq) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Captcha_id, 0, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Code, 1, true)
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
func (st *VerifyCaptchaReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require VerifyCaptchaReq, but not exist. tag %d", tag)
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
func (st *VerifyCaptchaReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Captcha_id, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Code, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *VerifyCaptchaReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// VerifyCaptchaRsp struct implement
type VerifyCaptchaRsp struct {
	Extra_params map[string]string `json:"extra_params"`
}

func (st *VerifyCaptchaRsp) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *VerifyCaptchaRsp) ReadFrom(readBuf *codec.Reader) error {
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
func (st *VerifyCaptchaRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require VerifyCaptchaRsp, but not exist. tag %d", tag)
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
func (st *VerifyCaptchaRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteHead(codec.MAP, 0)
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
func (st *VerifyCaptchaRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
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
