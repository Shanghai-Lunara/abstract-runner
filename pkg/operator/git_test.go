package operator

import (
	"reflect"
	"testing"

	"github.com/Shanghai-Lunara/abstract-runner/proto"
)

func Test_git_Clone(t *testing.T) {
	type fields struct {
		c proto.Git
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "Test_git_Clone1",
			fields: fields{c: proto.Git{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &git{
				c: tt.fields.c,
			}
			g.Clone()
		})
	}
}

func Test_git_ExecuteWithArgs(t *testing.T) {
	type fields struct {
		c proto.Git
	}
	type args struct {
		cmd  string
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &git{
				c: tt.fields.c,
			}
			gotRes, err := g.ExecuteWithArgs(tt.args.cmd, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("git.ExecuteWithArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("git.ExecuteWithArgs() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
