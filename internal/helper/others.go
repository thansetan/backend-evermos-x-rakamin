package helper

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _       = runtime.Caller(0)
	ProjectDirectory = filepath.Join(filepath.Dir(b), "../..")
)
