package core;

type Function struct {
	Name        string
	Args        []Arg
	ArgsCount   float32
	SubFunction *Function
	Execute     func(*Function) any
	ReturnT     string // "any" | "nil" | "int" | "str" | "bool". Basically, the types.
}

type Arg struct {
	Name  string
	T     string // "any" | "nil" | "int" | "str" | "bool". Basically, the types.
	Value string
}