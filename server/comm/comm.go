package comm

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/link1st/go-stress-testing/Proto/Taurus"
	"go.k8sf.cloud/go/KsfGo/ksf"
	"math"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/bytedance/sonic"
	"github.com/golang-jwt/jwt/v4"
	"go.k8sf.cloud/go/KsfGo/ksf/protocol/codec"
	"go.k8sf.cloud/go/KsfGo/ksf/protocol/res/basef"
	"go.k8sf.cloud/go/KsfGo/ksf/protocol/res/requestf"
	"go.k8sf.cloud/go/KsfGo/ksf/util/endpoint"
	"go.k8sf.cloud/go/KsfGo/ksf/util/tools"
)

// 是否是测试模式
const IS_TEST bool = false
const maxInt32 int32 = 1<<31 - 1

// 利用map key进行索引，map value没有实际意义用Empty填充
const (
	//用户
	UserServant           = "Taurus.UserAuthServer.TaurusUserObj"
	UserCloudServant      = "Taurus.UserCloudServer.UserCloudObj"
	RouterServant         = "Taurus.RouterServer.RouterObj"
	StrategyRouterServant = "Taurus.StrategyRouterServer.RouterObj"

	UserDZServant           = "DZStgy.UserAuthServer.TaurusUserObj"
	UserCloudDZServant      = "DZStgy.UserCloudServer.UserCloudObj"
	RouterDZServant         = "DZStgy.RouterServer.RouterObj"
	StrategyRouterDZServant = "DZStgy.StrategyRouterServer.RouterObj"

	AccessServant   = "Access"
	AccessDZServant = "DZStgy.AccessServer.FlowProxyObj"
)

func GenUniqKsfObj(servant string, host string, port string) string {
	return fmt.Sprintf("%s@tcp -h %s -p %s", servant, host, port)
}

type ProgressTimer struct {
	TimeZone *time.Location
	Start    int64
}

type WatchResult struct {
	Type          string `json:"type"`
	K             string `json:"k"`
	V             string `json:"v"`
	CreateVersion int64  `json:"create_version"`
	ModVersion    int64  `json:"mod_version"`
	Version       int64  `json:"version"`
	LeaseId       int64  `json:"lease_id"`
}

func (p *ProgressTimer) Init() {
	p.TimeZone, _ = time.LoadLocation("Asia/Shanghai")
	p.Start = time.Now().UnixMilli()
}

func (p *ProgressTimer) Elapse() int64 {
	return time.Now().UnixMilli() - p.Start
}

func IsTaurusErr(ret Taurus.Result) bool {
	return ret.Code < 0
}

func Req2Byte(rsp *requestf.ResponsePacket) []byte {
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

func Rsp2Byte(rsp *requestf.ResponsePacket) []byte {
	if rsp.IVersion == basef.KUPVERSION {
		return Req2Byte(rsp)
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

type UUID struct {
	msgID int32
}

/*
*
fake_iRequestId -> (raw_iRequestId, sdk_uuid)
*/
func (id *UUID) GenRequestId() (requestId int32) {
	// 尽力防止溢出
	atomic.CompareAndSwapInt32(&id.msgID, maxInt32, 1)
	for {
		// 0比较特殊,用于表示 server 端推送消息给 client 端进行主动 close()
		// 溢出后回转成负数
		if v := atomic.AddInt32(&id.msgID, 1); v != 0 {
			return v
		}
	}
}

func GenResponsePacket(reqPack *requestf.RequestPacket, rspPack *requestf.ResponsePacket) {
	rspPack.IVersion = reqPack.IVersion
	rspPack.CPacketType = reqPack.CPacketType
	rspPack.IRequestId = reqPack.IRequestId
	rspPack.IMessageType = reqPack.IMessageType
	//rspPack.IRet = reqPack.IRet
	//rspPack.Context = reqPack.Context
}

func GenNotifyRsppacket(notify *Taurus.NotifyEvent, rspPack *requestf.ResponsePacket) {
	rspPack.IVersion = basef.KSFVERSION
	rspPack.CPacketType = basef.KSFNORMAL
	rspPack.IRequestId = 0
	rspPack.IMessageType = 0
	rspPack.SResultDesc = "NotifyEvent"
	buff := codec.NewBuffer()
	err := notify.WriteBlock(buff, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	rspPack.SBuffer = tools.ByteToInt8(buff.ToBytes())
}

func GenParentOrderRsppacket(notify *Taurus.ParentOrder, rspPack *requestf.ResponsePacket) {
	rspPack.IVersion = basef.KSFVERSION
	rspPack.CPacketType = basef.KSFNORMAL
	rspPack.IRequestId = 0
	rspPack.IMessageType = 0
	rspPack.SResultDesc = "OnUpdateParentOrder"
	buff := codec.NewBuffer()
	err := notify.WriteBlock(buff, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	rspPack.SBuffer = tools.ByteToInt8(buff.ToBytes())
}

func GenChildOrderRsppacket(notify *Taurus.ChildOrder, rspPack *requestf.ResponsePacket) {
	rspPack.IVersion = basef.KSFVERSION
	rspPack.CPacketType = basef.KSFNORMAL
	rspPack.IRequestId = 0
	rspPack.IMessageType = 0
	rspPack.SResultDesc = "OnUpdateChildOrder"
	buff := codec.NewBuffer()
	err := notify.WriteBlock(buff, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	rspPack.SBuffer = tools.ByteToInt8(buff.ToBytes())
}

func GenTaskStatusRsppacket(notify *Taurus.TaskStatus, rspPack *requestf.ResponsePacket) {
	rspPack.IVersion = basef.KSFVERSION
	rspPack.CPacketType = basef.KSFNORMAL
	rspPack.IRequestId = 0
	rspPack.IMessageType = 0
	rspPack.SResultDesc = "OnUpdateTaskStatus"
	buff := codec.NewBuffer()
	err := notify.WriteBlock(buff, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	rspPack.SBuffer = tools.ByteToInt8(buff.ToBytes())
}

type TaurusClaims struct {
	jwt.RegisteredClaims
	UserID string
	Passwd string
}

type DESPadMode string

const (
	DESPadModePKCS5    DESPadMode = "pkcs5"
	DESPadModePKCS7    DESPadMode = "pkcs7"
	DESPadModeZERO     DESPadMode = "zero"
	DESPadModeISO10126 DESPadMode = "iso10126"
	DESPadModeANSIX923 DESPadMode = "ansix923"
	DESPadModeNONE     DESPadMode = "none"
)

type DESEncryptMode string

const (
	DESEncryptModeECB DESEncryptMode = "ECB"
	DESEncryptModeCBC DESEncryptMode = "CBC"
	DESEncryptModeCTR DESEncryptMode = "CTR"
	DESEncryptModeOFB DESEncryptMode = "OFB"
	DESEncryptModeCFB DESEncryptMode = "CFB"
)

type DESOutPutMode string

const (
	DESOutPutModeBase64 DESOutPutMode = "base64"
	DESOutPutModeHEX    DESOutPutMode = "hex"
)

func GenRefreshToken(userID string, duration int64, salt []byte, passwd string) (string, int64, error) {
	now := time.Now()
	claims := &TaurusClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:  "UserAuthServer",
			Subject: "RefreshToken",
			ExpiresAt: &jwt.NumericDate{
				Time: time.Unix(now.Unix()+duration, 0),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Unix(now.Unix()-60*3, 0),
			},
		},
		UserID: userID,
		Passwd: passwd,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(salt)
	expireTime := time.Unix(now.Unix()+duration, 0).UnixMilli()

	return signed, expireTime, err
}

func GenQuickToken(userID string, duration int64, salt []byte) (string, int64, error) {
	now := time.Now()
	claims := &TaurusClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:  "UserAuthServer",
			Subject: "QuickToken",
			ExpiresAt: &jwt.NumericDate{
				Time: time.Unix(now.Unix()+duration, 0),
			},
			NotBefore: &jwt.NumericDate{
				Time: time.Time{},
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Unix(now.Unix()-60*3, 0),
			},
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(salt)
	expireTime := time.Unix(now.Unix()+duration, 0).UnixMilli()

	return signed, expireTime, err
}

func ParseRefreshToken(token string, salt []byte) (*TaurusClaims, Taurus.Result) {
	jToken, err := jwt.ParseWithClaims(token, &TaurusClaims{}, func(t *jwt.Token) (interface{}, error) {
		return salt, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); !ok {
			return nil, ES(Taurus.ENO_PARAM_INVALID, "token: "+token)
		} else {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ES(Taurus.ENO_PARAM_INVALID, token+" is nota token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ES(Taurus.ENO_REFRESH_TOKEN_ERR)
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ES(Taurus.ENO_PARAM_INVALID, "token "+token+" not active yet")
			} else {
				return nil, ES(Taurus.ENO_PARAM_INVALID, "couldn't handle refresh token: "+token+", error: "+fmt.Sprintf("%+v", ve))
			}
		}
	}

	if !jToken.Valid {
		return nil, ES(Taurus.ENO_PARAM_INVALID, "token: "+token+" is invalid")
	}

	if claims, ok := jToken.Claims.(*TaurusClaims); ok {
		return claims, Taurus.Result{}
	}

	return nil, ES(Taurus.ENO_PARAM_INVALID, token+" is not a token")
}

func ParseQuickToken(token string, salt []byte) (*TaurusClaims, Taurus.Result) {
	jToken, err := jwt.ParseWithClaims(token, &TaurusClaims{}, func(t *jwt.Token) (interface{}, error) {
		return salt, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); !ok {
			return nil, ES(Taurus.ENO_PARAM_INVALID, "token: "+token)
		} else {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ES(Taurus.ENO_PARAM_INVALID, token+" is nota token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ES(Taurus.ENO_QUICK_TOKEN_ERR)
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ES(Taurus.ENO_PARAM_INVALID, "token "+token+" not active yet")
			} else {
				return nil, ES(Taurus.ENO_PARAM_INVALID, "couldn't handle quick token: "+token+", error: "+fmt.Sprintf("%+v", ve))
			}
		}
	}

	if !jToken.Valid {
		return nil, ES(Taurus.ENO_PARAM_INVALID, "token: "+token+" is invalid")
	}

	if claims, ok := jToken.Claims.(*TaurusClaims); ok {
		return claims, Taurus.Result{}
	}

	return nil, ES(Taurus.ENO_PARAM_INVALID, token+" is not a token")
}

// GetTodayZeroTimestamp 当天0点的时间戳
func GetTodayZeroTimestamp() int64 {
	now := time.Now()
	ms := now.Unix() - int64(now.Hour())*3600 - int64(now.Minute())*60 - int64(now.Second())
	return ms
}

// GetSecondDaySpecifyTime 获取第二天的指定时间点
func GetSecondDaySpecifyTime(hour int64) time.Time {
	today := GetTodayZeroTimestamp()
	return time.Unix(today+hour*3600, 0)
}

// DESEncrypt DES加密算法, raw: 待加密内容, pwd: 加密密钥, encryMode: 加密模式, padMode: 填充模式, outMode: 输出模式, iv: 偏移量
func DESEncrypt(raw string, pwd string, encryMode DESEncryptMode, padMode DESPadMode, outMode DESOutPutMode, iv string) (string, error) {
	block, err := des.NewCipher([]byte(pwd))
	if err != nil {
		return "", err
	}

	var paddingData []byte
	if padMode == DESPadModePKCS5 {
		paddingData = desPaddingPKCS5([]byte(raw), block.BlockSize())
	} else if padMode == DESPadModeNONE {
		paddingData = []byte(raw)
	} else {
		return "", errors.New("unsupported des padding mode: " + string(padMode))
	}

	var out []byte
	if encryMode == DESEncryptModeCBC {
		out = make([]byte, len(paddingData))
		mode := cipher.NewCBCEncrypter(block, []byte(iv))
		mode.CryptBlocks(out, paddingData)
	} else if encryMode == DESEncryptModeECB {
		if len(paddingData)%block.BlockSize() != 0 {
			return "", errors.New("need a multiple of the blocksize")
		}

		dst := make([]byte, block.BlockSize())
		for len(paddingData) > 0 {
			block.Encrypt(dst, paddingData[:block.BlockSize()])
			paddingData = paddingData[block.BlockSize():]
			out = append(out, dst...)
		}
	} else {
		return "", errors.New("unsupported encrypt mode: " + string(encryMode))
	}

	if outMode == DESOutPutModeBase64 {
		return base64.StdEncoding.EncodeToString(out), nil
	} else if outMode == DESOutPutModeHEX {
		return fmt.Sprintf("%X", out), nil
	} else {
		return "", errors.New("unsupported out put mode: " + string(outMode))
	}
}

func DESUserPasswd(raw string) (des string, err error) {
	return DESEncrypt(raw, "KsTaurus", DESEncryptModeECB, DESPadModePKCS5, DESOutPutModeBase64, "Tomorrow")
}

// DESDecrypt DES解密算法
func DESDecrypt(raw string, pwd string, encryMode DESEncryptMode, padMode DESPadMode, outMode DESOutPutMode, iv string) (string, error) {
	var deStr []byte
	var err error

	if outMode == DESOutPutModeBase64 {
		deStr, err = base64.StdEncoding.DecodeString(raw)
		if err != nil {
			return "", err
		}
	} else if outMode == DESOutPutModeHEX {
		deStr, err = hex.DecodeString(raw)
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("unsupported out put mode: " + string(outMode))
	}

	block, err := des.NewCipher([]byte(pwd))
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()

	// 解密出来的内容属于尚未进行逆补齐的文本
	var notUnPaddingtext []byte
	if encryMode == DESEncryptModeCBC {
		notUnPaddingtext = make([]byte, len(deStr))
		mode := cipher.NewCBCDecrypter(block, []byte(iv))
		mode.CryptBlocks(notUnPaddingtext, deStr)
	} else if encryMode == DESEncryptModeECB {
		if len(deStr)%blockSize != 0 {
			return "", errors.New("not full blocks")
		}

		for len(deStr) > 0 {
			dst := make([]byte, blockSize)
			block.Decrypt(dst, deStr[:blockSize])
			deStr = deStr[blockSize:]
			notUnPaddingtext = append(notUnPaddingtext, dst...)
		}
	} else {
		return "", errors.New("unsupported encrypt mode: " + string(encryMode))
	}

	// 逆补齐
	unPaddingtext := []byte{}
	if padMode == DESPadModePKCS5 {
		unPaddingtext = desUnPaddingPKCS5(notUnPaddingtext)
	} else if padMode == DESPadModeNONE {
		unPaddingtext = notUnPaddingtext
	} else {
		return "", errors.New("unsupported des padding mode: " + string(padMode))
	}

	return string(unPaddingtext), nil
}

// desPaddingPKCS5 DES加密算法中的pkcs5补齐算法
func desPaddingPKCS5(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// desUnPaddingPKCS5 pkcs5逆补齐
func desUnPaddingPKCS5(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// GetDiffDays 获取两个时间戳中间跨天的天数
func GetDiffDays(ts1 int64, ts2 int64) int {
	t1 := time.Unix(ts1, 0)
	t2 := time.Unix(ts2, 0)

	ts1 = ts1 - int64(t1.Hour())*3600 - int64(t1.Minute())*60 - int64(t1.Second())
	ts2 = ts2 - int64(t2.Hour())*3600 - int64(t2.Minute())*60 - int64(t2.Second())

	return int(math.Abs(float64(ts2-ts1))/86400 + 1)
}

func PrettyJsonStr(raw []byte) (pretty []byte, err error) {
	var data map[string]interface{}
	err = json.Unmarshal(raw, &data)
	if err != nil {
		return nil, err
	}

	// Marshal map as pretty JSON string
	pretty, err = json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, err
	}
	return pretty, nil
}

func DeleteSlice(list []interface{}, elem interface{}) []interface{} {
	index := 0
	for _, v := range list {
		if v != elem {
			list[index] = v
			index++
		}
	}
	return list[:index]
}

func IntInnerItems(nums1 []int, nums2 []int) []int {
	i, j, k := 0, 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}

func IntIsSubSet(nums1 []int, nums2 []int) bool {
	raw_len := len(nums1)
	if raw_len > len(nums2) {
		return false
	} else {
		tmp := IntInnerItems(nums1, nums2)
		if len(tmp) != raw_len {
			return false
		} else {
			return true
		}
	}
}

func StrInnerItems(nums1 []string, nums2 []string) []string {
	i, j, k := 0, 0, 0
	sort.Strings(nums1)
	sort.Strings(nums2)
	for i < len(nums1) && j < len(nums2) {
		if strings.Compare(nums1[i], nums2[j]) > 0 {
			j++
		} else if strings.Compare(nums1[i], nums2[j]) < 0 {
			i++
		} else {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}

func StrIsSubSet(strs1 []string, strs2 []string) bool {
	raw_len := len(strs1)
	if raw_len > len(strs2) {
		return false
	} else {
		tmp := StrInnerItems(strs1, strs2)
		if len(tmp) != raw_len {
			return false
		} else {
			return true
		}
	}
}

func RemoveDuplicates[T any](source []T) []T {
	unique := make(map[string]bool)
	var result []T

	for _, arr := range source {
		// 使用reflect.DeepEqual进行数组比较，将数组转换为字符串作为map键
		key := fmt.Sprintf("%v", arr)
		if !unique[key] {
			unique[key] = true
			result = append(result, arr)
		}
	}
	return result
}

func UniqeObj(obj string) (uniqObj string, err error) {
	defer func() {
		if r := recover(); r != nil {
			if er, ok := r.(error); ok {
				err = er
			}
			fmt.Println(err)
		}
	}()
	vecStr := strings.Split(obj, "@")
	if len(vecStr) != 2 {
		return "", errors.New("invalid obj")
	}
	ep := endpoint.Parse(vecStr[1])
	return vecStr[0] + "@" + ep.UniqString(), nil
}

// TransAccountStatus 获取账户状态
func TransAccountStatus(loginStatus string, connectionStatus string) string {
	if loginStatus == "on" {
		if connectionStatus == "on" {
			return Taurus.USER_ACCOUNT_STATUS_LOGIN_CONNECTED
		} else if connectionStatus == "off" {
			return Taurus.USER_ACCOUNT_STATUS_LOGIN_DISCONNECTED
		}
	} else if loginStatus == "off" {
		if connectionStatus == "on" {
			return Taurus.USER_ACCOUNT_STATUS_LOGOUT_CONNECTED
		} else if connectionStatus == "off" {
			return Taurus.USER_ACCOUNT_STATUS_LOGOUT_DISCONNECTED
		}
	}

	return ""
}

func MD5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func DeferFunc() {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			fmt.Printf("%+v", err)
		} else {
			fmt.Printf("%#v", r)
		}
		fmt.Printf("%s", string(debug.Stack()))
	}
}

func CurDateInt() int {
	// 获取当前日期和时间
	currentTime := time.Now()
	// 获取当天日期的年、月、日
	year := currentTime.Year()
	month := int(currentTime.Month())
	day := currentTime.Day()
	// 将日期表示为整数，格式为YYYYMMDD
	return year*10000 + month*100 + day
}

func CurDateStr() string {
	// 获取当前日期和时间
	return time.Now().Format("20060102")
}

func SliceByPage[T any](total []T, pagination Taurus.Pagination) (data []T, err error) {
	totalSize := len(total)
	if int(pagination.Offset) >= totalSize {
		return data, fmt.Errorf("无分页内容：%+v|%d", pagination, totalSize)
	}

	realSize := int(pagination.PageSize)
	if pagination.Offset+pagination.PageSize >= int64(totalSize) {
		realSize = totalSize - int(pagination.Offset)
	}
	copy(total[pagination.Offset:realSize], data)
	return
}

func SliceByPage2[T any](total []T, pagination Taurus.Pagination) (data []T, err error) {
	if pagination.Offset < 1 {
		return nil, errors.New("invalid offset")
	}
	totalSize := len(total)
	offset := (pagination.Offset - 1) * pagination.PageSize
	if int(offset) >= totalSize {
		return []T{}, nil
	}

	endIndex := int(offset) + int(pagination.PageSize)
	if offset+pagination.PageSize >= int64(totalSize) {
		endIndex = totalSize
	}
	data = append(data, total[offset:endIndex]...)
	return
}

func GenOptStationMark(clientInfo map[string]string) (map[string]string, error) {
	result := map[string]string{"OPT_STATION_CS": ""}
	if clientInfo == nil {
		return nil, errors.New("clientInfo is nil")
	}

	value := ""

	if val, exist := clientInfo["TYPE"]; !exist {
		return nil, errors.New("TYPE is not exist")
	} else {
		value += val + ";"
	}

	if val, exist := clientInfo["IIP"]; !exist {
		return nil, errors.New("IIP is not exist")
	} else {
		value += fmt.Sprintf("IIP=%s;", val)
	}

	if val, exist := clientInfo["IPORT"]; !exist {
		return nil, errors.New("IPORT is not exist")
	} else {
		value += fmt.Sprintf("IPORT=%s;", val)
	}

	if val, exist := clientInfo["LIP"]; !exist {
		return nil, errors.New("LIP is not exist")
	} else {
		value += fmt.Sprintf("LIP=%s;", val)
	}

	if val, exist := clientInfo["MAC"]; !exist {
		return nil, errors.New("MAC is not exist")
	} else {
		value += fmt.Sprintf("MAC=%s;", val)
	}

	if val, exist := clientInfo["HD"]; !exist {
		return nil, errors.New("HD is not exist")
	} else {
		value += fmt.Sprintf("HD=%s;", val)
	}

	if val, exist := clientInfo["PCN"]; !exist {
		return nil, errors.New("PCN is not exist")
	} else {
		value += fmt.Sprintf("PCN=%s;", val)
	}

	if val, exist := clientInfo["CPU"]; !exist {
		return nil, errors.New("CPU is not exist")
	} else {
		value += fmt.Sprintf("CPU=%s;", val)
	}

	if val, exist := clientInfo["PI"]; !exist {
		return nil, errors.New("PI is not exist")
	} else {
		value += fmt.Sprintf("PI=%s;", val)
	}

	if val, exist := clientInfo["VOL"]; !exist {
		return nil, errors.New("VOL is not exist")
	} else {
		value += fmt.Sprintf("VOL=%s;", val)
	}

	if val, exist := clientInfo["PROVIDER"]; !exist {
		return nil, errors.New("PROVIDER is not exist")
	} else {
		value += fmt.Sprintf("@%s;", val)
	}

	if val, exist := clientInfo["VERSION"]; !exist {
		return nil, errors.New("VERSION is not exist")
	} else {
		value += val
	}

	result["OPT_STATION_CS"] = value

	return result, nil
}

func HashFn(str string) uint32 {
	var value uint32
	for _, c := range str {
		value += uint32(c)
		value += value << 10
		value ^= value >> 6
	}
	value += value << 3
	value ^= value >> 11
	value += value << 15
	if value == 0 {
		return 1
	}
	return value
}

func StrIsInteger(s string) (bool, int) {
	i, err := strconv.Atoi(s)
	return err == nil, i
}

// 仅针对支持json序列化的结构体
func DeepCopy[T any](src *T, dest *T) error {
	buff, err := sonic.Marshal(src)
	if err != nil {
		return err
	}
	return sonic.Unmarshal(buff, dest)
}

// 将slice拆分成多个slice
func SliceSplit[T any](src []T, num int, dest *[][]T) {
	totalSize := len(src)
	for i := 0; i < totalSize+num; {
		if i+num >= totalSize {
			var tmp []T
			item := src[i:]
			err := DeepCopy(&item, &tmp)
			if err != nil {
				ksf.TLOG.Error(err)
				return
			}
			*dest = append(*dest, tmp)
			return
		} else {
			var tmp []T
			item := src[i : i+num]
			err := DeepCopy(&item, &tmp)
			if err != nil {
				ksf.TLOG.Error(err)
				return
			}
			*dest = append(*dest, tmp)
			i += num
		}
	}
}

func Yesterday(date int) int {
	t, err := time.Parse("20060102", strconv.Itoa(date))
	if err != nil {
		ksf.TLOG.Error(err)
		return 0
	}
	yesterday := t.Add(-24 * time.Hour)
	return yesterday.Year()*10000 + int(yesterday.Month())*100 + yesterday.Day()
}
