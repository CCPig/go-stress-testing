// Package Taurus comment
// This file was generated by ksf2go 1.3.21
// Generated from TaurusUserCloud.ksf
package Taurus

import (
	"fmt"

	"go.k8sf.cloud/go/KsfGo/ksf/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// BasketInfo struct implement
type BasketInfo struct {
	User_id     string `json:"user_id"`
	Basket_id   int64  `json:"basket_id"`
	Basket_name string `json:"basket_name"`
	Insts_count int32  `json:"insts_count"`
	Remark      string `json:"remark"`
}

func (st *BasketInfo) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *BasketInfo) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.User_id, 0, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.Basket_id, 1, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Basket_name, 2, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Insts_count, 3, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Remark, 4, false)
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
func (st *BasketInfo) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require BasketInfo, but not exist. tag %d", tag)
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
func (st *BasketInfo) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.User_id, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.Basket_id, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Basket_name, 2)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Insts_count, 3)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Remark, 4)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *BasketInfo) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// BasketInstrument struct implement
type BasketInstrument struct {
	Basket_instrument_id  int64   `json:"basket_instrument_id"`
	Basket_id             int64   `json:"basket_id"`
	Instrument_id         string  `json:"instrument_id"`
	Instrument_type       string  `json:"instrument_type"`
	Exchange_id           string  `json:"exchange_id"`
	Instrument_name       string  `json:"instrument_name"`
	Order_side            string  `json:"order_side"`
	Amount                int64   `json:"amount"`
	Weight                float64 `json:"weight"`
	Remark                string  `json:"remark"`
	Price_type            string  `json:"price_type"`
	Order_price           string  `json:"order_price"`
	Estimate_order_amount float64 `json:"estimate_order_amount"`
	Per_amount            int64   `json:"per_amount"`
}

func (st *BasketInstrument) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *BasketInstrument) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadInt64(&st.Basket_instrument_id, 0, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.Basket_id, 1, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Instrument_id, 2, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Instrument_type, 3, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Exchange_id, 4, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Instrument_name, 5, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Order_side, 6, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.Amount, 7, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Weight, 8, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Remark, 9, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Price_type, 10, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Order_price, 11, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Estimate_order_amount, 12, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.Per_amount, 13, false)
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
func (st *BasketInstrument) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require BasketInstrument, but not exist. tag %d", tag)
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
func (st *BasketInstrument) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteInt64(st.Basket_instrument_id, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.Basket_id, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Instrument_id, 2)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Instrument_type, 3)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Exchange_id, 4)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Instrument_name, 5)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Order_side, 6)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.Amount, 7)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Weight, 8)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Remark, 9)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Price_type, 10)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Order_price, 11)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Estimate_order_amount, 12)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.Per_amount, 13)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *BasketInstrument) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// AccountGroup struct implement
type AccountGroup struct {
	Account_group_id   int64         `json:"account_group_id"`
	User_id            string        `json:"user_id"`
	Account_group_name string        `json:"account_group_name"`
	Account_type       string        `json:"account_type"`
	Accounts           []UserAccPerm `json:"accounts"`
}

func (st *AccountGroup) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *AccountGroup) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadInt64(&st.Account_group_id, 0, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.User_id, 1, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Account_group_name, 2, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Account_type, 3, true)
	if err != nil {
		return err
	}

	_, ty, err = readBuf.SkipToNoCheck(4, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = readBuf.ReadInt32(&length, 0, true)
		if err != nil {
			return err
		}

		st.Accounts = make([]UserAccPerm, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = st.Accounts[i0].ReadBlock(readBuf, 0, false)
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

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *AccountGroup) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require AccountGroup, but not exist. tag %d", tag)
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
func (st *AccountGroup) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteInt64(st.Account_group_id, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.User_id, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Account_group_name, 2)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Account_type, 3)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 4)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Accounts)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Accounts {

		err = v.WriteBlock(buf, 0)
		if err != nil {
			return err
		}

	}

	return err
}

// WriteBlock encode struct
func (st *AccountGroup) WriteBlock(buf *codec.Buffer, tag byte) error {
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
