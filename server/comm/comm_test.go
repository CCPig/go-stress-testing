// Package comm.comm_test.go create by chaochen at 2024/6/21 下午3:50:00
package comm

import "testing"

func TestDESEncrypt(t *testing.T) {
	type args struct {
		raw       string
		pwd       string
		encryMode DESEncryptMode
		padMode   DESPadMode
		outMode   DESOutPutMode
		iv        string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				raw:       "abc123!",
				pwd:       "KsTaurus",
				encryMode: DESEncryptModeECB,
				padMode:   DESPadModePKCS5,
				outMode:   DESOutPutModeBase64,
				iv:        "Tomorrow",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DESEncrypt(tt.args.raw, tt.args.pwd, tt.args.encryMode, tt.args.padMode, tt.args.outMode, tt.args.iv)
			if (err != nil) != tt.wantErr {
				t.Errorf("DESEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DESEncrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDESUserPasswd(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name    string
		args    args
		wantDes string
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				raw: "abc123!",
			},
			wantDes: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDes, _ := DESUserPasswd(tt.args.raw)
			if gotDes != tt.wantDes {
				t.Errorf("DESUserPasswd() gotDes = %v, want %v", gotDes, tt.wantDes)
			}
		})
	}
}
