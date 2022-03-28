package conf

import (
	"testing"

	"github.com/spf13/viper"
)

func TestInit(t *testing.T) {
	type args struct {
		configFileName string
		configPath     string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				configFileName: "conf.json",
				configPath:     ".",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initConfig(tt.args.configFileName, tt.args.configPath)
		})
	}
}

func TestYamlConfig(t *testing.T) {
	type args struct {
		configFile string
		configPath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			YamlConfig(tt.args.configFile, tt.args.configPath)
		})
	}
}

func TestTomlConfig(t *testing.T) {
	type args struct {
		configFile string
		configPath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TomlConfig(tt.args.configFile, tt.args.configPath)
		})
	}
}

func TestJsonConfig(t *testing.T) {
	type args struct {
		configFile string
		configPath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				configFile: "conf.json",
				configPath: ".",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JsonConfig(tt.args.configFile, tt.args.configPath)
			if viper.GetString("app") != "cc" {
				t.Errorf("JsonConfig() error = %v, want %v", viper.GetString("app"), "cc")
			}
			if viper.GetString("db.host") != "127.0.0.1" ||
				viper.GetString("db.user") != "root" ||
				viper.GetString("db.password") != "123456" ||
				viper.GetInt("db.port") != 3310 {
				t.Errorf("JsonConfig() error = %v, want %v", viper.Get("db"), nil)
			}
		})
	}
}

func TestIniConfig(t *testing.T) {
	type args struct {
		configFile string
		configPath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				configFile: "conf.ini",
				configPath: ".",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IniConfig(tt.args.configFile, tt.args.configPath)
			t.Log(viper.GetViper().AllKeys())
			if viper.GetString("app.name") != "cc" {
				t.Errorf("IniConfig() error = %v, want %v", viper.GetString("app.name"), "cc")
			}
			if viper.GetString("db.host") != "127.0.0.1" ||
				viper.GetString("db.user") != "root" ||
				viper.GetString("db.password") != "123456" ||
				viper.GetInt("db.port") != 3310 {
				t.Errorf("JsonConfig() error = %v, want %v", viper.Get("db"), nil)
			}
		})
	}
}

func TestEnvConfig(t *testing.T) {
	type args struct {
		configFile string
		configPath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				configFile: "conf.env",
				configPath: ".",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EnvConfig(tt.args.configFile, tt.args.configPath)
			t.Log(viper.GetViper().AllKeys())
			if viper.GetString("app") != "cc" {
				t.Errorf("IniConfig() error = %v, want %v", viper.GetString("app"), "cc")
			}
			if viper.GetString("db_host") != "127.0.0.1" ||
				viper.GetString("db_user") != "root" ||
				viper.GetString("db_password") != "123456" ||
				viper.GetInt("db_port") != 3310 {
				t.Errorf("JsonConfig() error = %v, want %v", viper.Get("db_*"), nil)
			}
		})
	}
}
