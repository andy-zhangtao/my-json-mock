package config

import (
	"reflect"
	"testing"
)

func TestParseParams(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantRp  RunParams
		wantErr bool
	}{
		{
			name: "normal test",
			args: args{data: []byte(`
[configure]
	debug = true
	port = 18000
[mysql]
	addr = "127.0.0.1:3600"
	user = "admin"
	passwd = "321456"
	db = "my-json-mock"
`)},
			wantRp: RunParams{
				Configure: Configure{
					Debug: true,
					Port:  18000,
				},
				Mysql: Mysql{
					Dsn:    "127.0.0.1:3600",
					User:   "admin",
					Passwd: "321456",
					Db:     "my-json-mock",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRp, err := ParseParams(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRp, tt.wantRp) {
				t.Errorf("ParseParams() gotRp = %v, want %v", gotRp, tt.wantRp)
			}
		})
	}
}
