package ssa

func SomeFunc(instr interface{}, v int) {
    switch instr.(type) {
    case Value:
      instr.(interface {
        setNum(int)
      }).setNum(v)
    }
}

