// Package Alarm comment
// This file was generated by ksf2go 1.3.20
// Generated from Alarm.ksf
package Alarm

import (
	"fmt"

	"go.k8sf.cloud/go/KsfGo/ksf/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

type Platform int32

const (
	Platform_ALL       = 0
	Platform_DING_TALK = 1
	Platform_WE_CHAT   = 2
	Platform_MAIL      = 3
)

// DingTalkReq struct implement
type DingTalkReq struct {
	Url        string   `json:"url"`
	Secret     string   `json:"secret"`
	Content    string   `json:"content"`
	At         []string `json:"at"`
	Is_at_all  bool     `json:"is_at_all"`
	Group_name string   `json:"group_name"`
}

func (st *DingTalkReq) ResetDefault() {
	st.Is_at_all = false
}

// ReadFrom reads  from readBuf and put into struct.
func (st *DingTalkReq) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Url, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Secret, 1, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Content, 2, false)
	if err != nil {
		return err
	}

	have, ty, err = readBuf.SkipToNoCheck(3, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			st.At = make([]string, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = readBuf.ReadString(&st.At[i0], 0, false)
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

	err = readBuf.ReadBool(&st.Is_at_all, 4, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Group_name, 5, false)
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
func (st *DingTalkReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require DingTalkReq, but not exist. tag %d", tag)
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
func (st *DingTalkReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Url, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Secret, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Content, 2)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 3)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.At)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.At {

		err = buf.WriteString(v, 0)
		if err != nil {
			return err
		}

	}

	err = buf.WriteBool(st.Is_at_all, 4)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Group_name, 5)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *DingTalkReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// WeChatReq struct implement
type WeChatReq struct {
	Url        string `json:"url"`
	Group_name string `json:"group_name"`
	Content    string `json:"content"`
}

func (st *WeChatReq) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *WeChatReq) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Url, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Group_name, 1, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Content, 2, false)
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
func (st *WeChatReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require WeChatReq, but not exist. tag %d", tag)
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
func (st *WeChatReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Url, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Group_name, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Content, 2)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *WeChatReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// MailReq struct implement
type MailReq struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Content string   `json:"content"`
	Html    string   `json:"html"`
}

func (st *MailReq) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *MailReq) ReadFrom(readBuf *codec.Reader) error {
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

			st.To = make([]string, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = readBuf.ReadString(&st.To[i0], 0, false)
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

	err = readBuf.ReadString(&st.Subject, 1, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Content, 2, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Html, 3, false)
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
func (st *MailReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require MailReq, but not exist. tag %d", tag)
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
func (st *MailReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteHead(codec.LIST, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.To)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.To {

		err = buf.WriteString(v, 0)
		if err != nil {
			return err
		}

	}

	err = buf.WriteString(st.Subject, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Content, 2)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Html, 3)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *MailReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// AlarmReq struct implement
type AlarmReq struct {
	Platforms  []Platform    `json:"platforms"`
	Content    string        `json:"content"`
	Subject    string        `json:"subject"`
	Ding_talks []DingTalkReq `json:"ding_talks"`
	We_chats   []WeChatReq   `json:"we_chats"`
	Mails      []MailReq     `json:"mails"`
}

func (st *AlarmReq) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *AlarmReq) ReadFrom(readBuf *codec.Reader) error {
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

			st.Platforms = make([]Platform, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = readBuf.ReadInt32((*int32)(&st.Platforms[i0]), 0, false)
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

	err = readBuf.ReadString(&st.Content, 1, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Subject, 2, false)
	if err != nil {
		return err
	}

	have, ty, err = readBuf.SkipToNoCheck(3, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			st.Ding_talks = make([]DingTalkReq, length)
			for i1, e1 := int32(0), length; i1 < e1; i1++ {

				err = st.Ding_talks[i1].ReadBlock(readBuf, 0, false)
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

	have, ty, err = readBuf.SkipToNoCheck(4, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			st.We_chats = make([]WeChatReq, length)
			for i2, e2 := int32(0), length; i2 < e2; i2++ {

				err = st.We_chats[i2].ReadBlock(readBuf, 0, false)
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

	have, ty, err = readBuf.SkipToNoCheck(5, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			st.Mails = make([]MailReq, length)
			for i3, e3 := int32(0), length; i3 < e3; i3++ {

				err = st.Mails[i3].ReadBlock(readBuf, 0, false)
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
func (st *AlarmReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require AlarmReq, but not exist. tag %d", tag)
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
func (st *AlarmReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteHead(codec.LIST, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Platforms)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Platforms {

		err = buf.WriteInt32(int32(v), 0)
		if err != nil {
			return err
		}

	}

	err = buf.WriteString(st.Content, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Subject, 2)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 3)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Ding_talks)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Ding_talks {

		err = v.WriteBlock(buf, 0)
		if err != nil {
			return err
		}

	}

	err = buf.WriteHead(codec.LIST, 4)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.We_chats)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.We_chats {

		err = v.WriteBlock(buf, 0)
		if err != nil {
			return err
		}

	}

	err = buf.WriteHead(codec.LIST, 5)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Mails)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Mails {

		err = v.WriteBlock(buf, 0)
		if err != nil {
			return err
		}

	}

	return err
}

// WriteBlock encode struct
func (st *AlarmReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// AlarmRsp struct implement
type AlarmRsp struct {
	Code     int32    `json:"code"`
	Err_msgs []string `json:"err_msgs"`
}

func (st *AlarmRsp) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *AlarmRsp) ReadFrom(readBuf *codec.Reader) error {
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

			st.Err_msgs = make([]string, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = readBuf.ReadString(&st.Err_msgs[i0], 0, false)
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
func (st *AlarmRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require AlarmRsp, but not exist. tag %d", tag)
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
func (st *AlarmRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteInt32(st.Code, 0)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 1)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Err_msgs)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Err_msgs {

		err = buf.WriteString(v, 0)
		if err != nil {
			return err
		}

	}

	return err
}

// WriteBlock encode struct
func (st *AlarmRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
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
