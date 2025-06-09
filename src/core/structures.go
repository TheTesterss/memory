package core;

type Function struct {
	Name        string				// The name gave to the function.
	Args        []Arg				// The Listed_args value, split and organized with precised types.
	ArgsCount   float32				// The amount of args counted and necessary for the function to run.
	SubFunction *Function           // If Item is a class, what's the method used to accompaign the current item.
	Execute     func(*Function) any // The function to execute when the function is ready to be executed.
	ReturnT     string              // "any" | "nil" | "int" | "str" | "bool". Basically, the types.
}

type Arg struct {
	Name  string // The argument's name (gave at the function declaration & completed by the user).
	T     string // "any" | "nil" | "int" | "str" | "bool". Basically, the types.
	Value string // The value of this function.
}