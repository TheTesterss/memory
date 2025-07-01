package types;

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

type Variable struct {
	Name  string // The variable's name.
	T     string // The variable's type (can change it once the variable is declared at one type).
	Value string // The variable's value (matches with the given type).
}

// The function has for goal to return a structured version of every found functions.
// If you use a $if[1==1] it should return { name: "$if", full_version: "$if[1==1]", args: [1==1], opened: true, closed: true, line: 0, subFunction: nil }
//
// name is how the function is called.
// full_version includes the function name + the given arguments (with brackets).
// args is a list of every arguments following the argument structure.
// opened and closed are there to show if its a variable or not (both true/false).
// line is the line where the function has been called.
// subItems is a list of every Items found in the argument.
type Item struct {
	Name        string // The name gave to the item.
	Listed_args string // A string of every arguments (before getting retrieved by SplitArgs function).
	Opened      bool   // Is the function opened? (does not count for $end).
	Closed      bool   // Is the function closed? (does not count for $end).
	Line        int32  // At which line has the item been found.
	SubFunction *Item  // If Item is a class, what's the method used to accompaign the current item.
}