package util

import (
  "runtime"
  "path/filepath"
)

func Dirname() string {
  _, filename, _, ok := runtime.Caller(1)
  if !ok {
    panic("getting caller functions")
  }
  return filepath.Dir(filename)
}
