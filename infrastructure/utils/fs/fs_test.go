package fs

import (
	"os"
	"testing"
)

func TestCreateDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "create success",
			args: args{path: "/mnt/hgfs/Program/Go/src/CleanCode/clean_code/infrastructure/pkg/fs/f1"},
			wantErr: false,
		},
		{
			name: "create exists file success",
			args: args{path: "/mnt/hgfs/Program/Go/src/CleanCode/clean_code/infrastructure/pkg/fs"},
			wantErr: false,
		},
		{
			name: "create exists file but not a dir path fail",
			args: args{path: "/mnt/hgfs/Program/Go/src/CleanCode/clean_code/infrastructure/pkg/fs/fs.go"},
			wantErr: true,
		},
		{
			name: "create not exists not a dir path success",
			args: args{path: "/mnt/hgfs/Program/Go/src/CleanCode/clean_code/infrastructure/pkg/fs/fs1.go"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateDir(tt.args.path, os.ModePerm); (err != nil) != tt.wantErr {
				t.Errorf("CreateDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateDirs(t *testing.T) {
	type args struct {
		path []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "create dirs success",
			args: args{path: []string{"./f1", "./f2", "./f3"}},
			wantErr: false,
		},
		{
			name: "create dirs(exists path already have) success",
			args: args{[]string{"./f1", "./f2","./f3"}},
			wantErr: false,
		},
		{
			name: "create dirs(exists file already have) fail",
			args: args{[]string{"./f1", "./f2","./f3", "./fs.go"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateDirs(os.ModePerm, tt.args.path...); (err != nil) != tt.wantErr {
				t.Errorf("CreateDirs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "is dir",
			args: args{path: "/mnt/hgfs/Program/Go/src/na_cai"},
			want: true,
		},
		{
			name: "is not dir",
			args: args{path: "/mnt/hgfs/Program/Go/src/na_cai/README.md"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDir(tt.args.path); got != tt.want {
				t.Errorf("IsDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "is not file",
			args: args{path: "/mnt/hgfs/Program/Go/src/na_cai"},
			want: false,
		},
		{
			name: "is file",
			args: args{path: "/mnt/hgfs/Program/Go/src/na_cai/README.md"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFile(tt.args.path); got != tt.want {
				t.Errorf("IsFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPathExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "dir path exists",
			args: args{path: "/mnt/hgfs/Program/Go/src/na_cai"},
			want: true,
		},
		{
			name: "file path exists",
			args: args{path: "/mnt/hgfs/Program/Go/src/na_cai/README.md"},
			want: true,
		},
		{
			name: "dir path not exists",
			args: args{path: "/mnt/hgfs/Program/Go/src/na_cai/docs1"},
			want: false,
		},
		{
			name: "file path not exists",
			args: args{path: "/mnt/hgfs/Program/Go/src/na_cai/README1.md"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PathExists(tt.args.path); got != tt.want {
				t.Errorf("PathExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "remove exists dir success",
			args: args{path: "./f1"},
			wantErr: false,
		},
		{
			name: "remove not exists dir success",
			args: args{path: "./f11"},
			wantErr: false,
		},
		{
			name: "remove exists file success",
			args: args{path: "file1.txt"},
			wantErr: false,
		},
		{
			name: "remove not exists file success",
			args: args{path: "file2.txt"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Remove(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBatchRemove(t *testing.T) {
	type args struct {
		path []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "remove exists dirs success",
			args: args{path: []string{"./f1", "./f4", "./f3"}},
			wantErr: false,
		},
		{
			name: "remove not exists dirs success",
			args: args{path: []string{"./f11", "./f14", "./f13"}},
			wantErr: false,
		},
		{
			name: "remove exists files success",
			args: args{path: []string{"./f1.txt", "./f34.txt", "./f3.txt"}},
			wantErr: false,
		},
		{
			name: "remove not exists files success",
			args: args{path: []string{"./f11.txt", "./f341.txt", "./f31.txt"}},
			wantErr: false,
		},
		{
			name: "remove file and dir success",
			args: args{path: []string{"./f2", "./f2.txt"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := BatchRemove(tt.args.path...); (err != nil) != tt.wantErr {
				t.Errorf("BatchRemove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	type args struct {
		srcPath string
		dstPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "cp ./f1.txt ./f2.txt success",
			args:    args{
				srcPath: "./f1.txt",
				dstPath: "./f2.txt",
			},
			wantErr: false,
		},
		{
			name:    "cp ./f1 ./f2 success",
			args:    args{
				srcPath: "./f1",
				dstPath: "./f2",
			},
			wantErr: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CopyRecursively(tt.args.srcPath, tt.args.dstPath); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCopyFile(t *testing.T) {
	type args struct {
		srcPath string
		dstPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "cp fs.go fs.go.bak",
			args:    args{
				srcPath: "fs.go",
				dstPath: "fs.go.bak",
			},
			wantErr: false,
		},
		{
			name:    "cp fs.go.bak f1/fs.go. f1 not exists",
			args:    args{
				srcPath: "fs.go.bak",
				dstPath: "f1/fs.go",
			},
			wantErr: true,
		},
		{
			name:    "cp fs.go.bak f2/fs.go. f2 exists",
			args:    args{
				srcPath: "fs.go.bak",
				dstPath: "f2/fs.go",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CopyFile(tt.args.srcPath, tt.args.dstPath); (err != nil) != tt.wantErr {
				t.Errorf("CopyFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}