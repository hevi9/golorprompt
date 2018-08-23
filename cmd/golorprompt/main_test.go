package main

import (
	"reflect"
	"testing"

	_ "github.com/hevi9/golorprompt/seg/cwd"
	_ "github.com/hevi9/golorprompt/seg/disk"
	_ "github.com/hevi9/golorprompt/seg/envvar"
	_ "github.com/hevi9/golorprompt/seg/exitcode"
	_ "github.com/hevi9/golorprompt/seg/hostname"
	_ "github.com/hevi9/golorprompt/seg/ifile"
	_ "github.com/hevi9/golorprompt/seg/jobs"
	_ "github.com/hevi9/golorprompt/seg/level"
	_ "github.com/hevi9/golorprompt/seg/load"
	_ "github.com/hevi9/golorprompt/seg/mem"
	_ "github.com/hevi9/golorprompt/seg/stub"
	_ "github.com/hevi9/golorprompt/seg/text"
	_ "github.com/hevi9/golorprompt/seg/time"
	_ "github.com/hevi9/golorprompt/seg/user"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_readAndResolveConfigFile(t *testing.T) {
	type args struct {
		config string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readAndResolveConfigFile(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readAndResolveConfigFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resolveConfigFile(t *testing.T) {
	type args struct {
		config string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolveConfigFile(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("resolveConfigFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("resolveConfigFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
