package main

import (
  ssa1 "github.com/golangci/test1/ssa"
  ssa2 "github.com/golangci/test2/ssa"
)

func main() {
  v1 := &ssa1.Call{}
  ssa1.SomeFunc(v1, 1)

  v2 := &ssa2.Call{}
  ssa2.DoesNotPanic(v2, 2)
  ssa2.Panics(v2, 2)
}
