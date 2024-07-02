// Package ksFactor comment
// This file was generated by ksf2go 1.3.20
// Generated from FactorData.ksf
package ksFactor

import (
	"fmt"

	"go.k8sf.cloud/go/KsfGo/ksf/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// Meta struct implement
type Meta struct {
	Code         string   `json:"code"`
	Name         string   `json:"name"`
	Factor_type  string   `json:"factor_type"`
	Source_type  string   `json:"source_type"`
	Description  string   `json:"description"`
	Src_table    string   `json:"src_table"`
	Src_section  string   `json:"src_section"`
	Formula      string   `json:"formula"`
	Formula_desc string   `json:"formula_desc"`
	Version      int64    `json:"version"`
	Enabled      bool     `json:"enabled"`
	Metainfo     string   `json:"metainfo"`
	Fill_mode    string   `json:"fill_mode"`
	Support_mode int32    `json:"support_mode"`
	Indi_formula string   `json:"indi_formula"`
	Depends      []string `json:"depends"`
	Etype        string   `json:"etype"`
	Foreign_key  string   `json:"foreign_key"`
	Dtype        string   `json:"dtype"`
	Offset       int32    `json:"offset"`
	Foreign_name string   `json:"foreign_name"`
	Access       int32    `json:"access"`
	Db_id        int32    `json:"db_id"`
	Remarks      string   `json:"remarks"`
	Modified     int64    `json:"modified"`
}

func (st *Meta) ResetDefault() {
	st.Enabled = true
}

// ReadFrom reads  from readBuf and put into struct.
func (st *Meta) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Code, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Name, 1, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Factor_type, 2, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Source_type, 3, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Description, 4, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Src_table, 5, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Src_section, 6, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Formula, 7, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Formula_desc, 8, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.Version, 9, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadBool(&st.Enabled, 10, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Metainfo, 11, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Fill_mode, 12, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Support_mode, 13, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Indi_formula, 14, false)
	if err != nil {
		return err
	}

	have, ty, err = readBuf.SkipToNoCheck(15, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			st.Depends = make([]string, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = readBuf.ReadString(&st.Depends[i0], 0, false)
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

	err = readBuf.ReadString(&st.Etype, 16, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Foreign_key, 17, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Dtype, 18, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Offset, 19, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Foreign_name, 20, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Access, 21, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Db_id, 22, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Remarks, 23, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.Modified, 24, false)
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
func (st *Meta) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require Meta, but not exist. tag %d", tag)
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
func (st *Meta) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Code, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Name, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Factor_type, 2)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Source_type, 3)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Description, 4)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Src_table, 5)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Src_section, 6)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Formula, 7)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Formula_desc, 8)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.Version, 9)
	if err != nil {
		return err
	}

	err = buf.WriteBool(st.Enabled, 10)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Metainfo, 11)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Fill_mode, 12)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Support_mode, 13)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Indi_formula, 14)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 15)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Depends)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Depends {

		err = buf.WriteString(v, 0)
		if err != nil {
			return err
		}

	}

	err = buf.WriteString(st.Etype, 16)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Foreign_key, 17)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Dtype, 18)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Offset, 19)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Foreign_name, 20)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Access, 21)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Db_id, 22)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Remarks, 23)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.Modified, 24)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *Meta) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// Field struct implement
type Field struct {
	Index      int32  `json:"index"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	Alias      string `json:"alias"`
	Handler    string `json:"handler"`
	Size       int32  `json:"size"`
	Enum_table int32  `json:"enum_table"`
}

func (st *Field) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *Field) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadInt32(&st.Index, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Type, 1, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Name, 2, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Alias, 3, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Handler, 4, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Size, 5, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Enum_table, 6, false)
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
func (st *Field) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require Field, but not exist. tag %d", tag)
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
func (st *Field) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteInt32(st.Index, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Type, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Name, 2)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Alias, 3)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Handler, 4)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Size, 5)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Enum_table, 6)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *Field) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// FinFactorDouble struct implement
type FinFactorDouble struct {
	Date   int32   `json:"date"`
	Report int32   `json:"report"`
	Value  float64 `json:"value"`
}

func (st *FinFactorDouble) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *FinFactorDouble) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadInt32(&st.Date, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Report, 1, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Value, 2, false)
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
func (st *FinFactorDouble) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require FinFactorDouble, but not exist. tag %d", tag)
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
func (st *FinFactorDouble) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteInt32(st.Date, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Report, 1)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Value, 2)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *FinFactorDouble) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// FinFactorString struct implement
type FinFactorString struct {
	Date   int32  `json:"date"`
	Report int32  `json:"report"`
	Value  string `json:"value"`
}

func (st *FinFactorString) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *FinFactorString) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadInt32(&st.Date, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Report, 1, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Value, 2, false)
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
func (st *FinFactorString) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require FinFactorString, but not exist. tag %d", tag)
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
func (st *FinFactorString) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteInt32(st.Date, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Report, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Value, 2)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *FinFactorString) WriteBlock(buf *codec.Buffer, tag byte) error {
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