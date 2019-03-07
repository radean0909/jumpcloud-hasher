package database

import (
	"reflect"
	"testing"
)

func TestConnect(t *testing.T) {
	type args struct {
		dbName string
	}

	tests := []struct {
		name    string
		args    args
		want    *DB
		wantErr bool
	}{
		{
			name: "Invalid DB Name",
			args: args{
				dbName: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Connect(tt.args.dbName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
