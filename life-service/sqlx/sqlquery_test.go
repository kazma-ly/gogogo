package sqlx

import (
	"testing"
)

func Test_makeUnderscodeToUp(t *testing.T) {
	type args struct {
		some string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{some: "type_id"}, "TypeId"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeUnderscodeToUp(tt.args.some); got != tt.want {
				t.Errorf("makeUnderscodeToUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareString(t *testing.T) {
	type args struct {
		entityStr   string
		databaseStr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{entityStr: "CreateTime", databaseStr: "create_time"}, true},
		{"test2", args{entityStr: "OriginNameFile", databaseStr: "origin_name_file"}, true},
		{"test3", args{entityStr: "Sex", databaseStr: "sex"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareString(tt.args.entityStr, tt.args.databaseStr); got != tt.want {
				t.Errorf("compareString() = %v, want %v", got, tt.want)
			}
		})
	}
}
