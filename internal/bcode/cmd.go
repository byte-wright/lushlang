package bcode

// Cmd starts Procs, runs Body while they run and waits until they have ended. Then it sets
// the Code, Stopped and collected values of each Proc. The errors are checked after it.
//
// A Proc that is the source of another one comes before it, the last one is the command of
// the statement. When Body is left early, by a return, a failure or a break of an outer
// loop, the backend stops the procs and waits for them.
type Cmd struct {
	Procs []*Proc
	// Pipe is the pipe of the last proc that Body reads with ReadLines, nil without one
	Pipe *VarValue
	// Body is nil without a pipe
	Body *Block
}

type Proc struct {
	Name Value
	Args []*ProcArg
	// Dir is nil for the script's working directory
	Dir Value
	Env []*ProcEnv
	// Stdin, Stdout and Stderr are nil for the script's own streams
	Stdin  *ProcInput
	Stdout *ProcOutput
	Stderr *ProcOutput
	Code   *VarValue
	// Stopped is true when nobody read the proc's pipe to its end and it was stopped, or
	// it ended because its reader went away. It's nil for a proc that is never stopped.
	Stopped *VarValue
}

type ProcArg struct {
	// Value is a string, int or bool, or an array with Spread
	Value  Value
	Spread bool
	// Source is set instead of Value for the pipe of another proc, the argument is a path
	// to read it
	Source *Proc
}

type ProcEnv struct {
	Name  string
	Value Value
}

type InputKind int

const (
	InputString InputKind = iota
	InputFile
	InputSource
)

type ProcInput struct {
	Kind InputKind
	// Value is the string or the path
	Value  Value
	Source *Proc
}

type OutputKind int

const (
	OutCollect OutputKind = iota
	OutPipe
	OutFile
	OutAppend
	OutDiscard
	// OutStdout is for stderr, it goes into the same target as stdout
	OutStdout
)

type ProcOutput struct {
	Kind OutputKind
	// Path is the file of OutFile and OutAppend
	Path Value
	// Var is the collected output of OutCollect, a string or []string
	Var *VarValue
	// Text is the collected output as a string for the error message, nil when not needed.
	// It can be Var.
	Text *VarValue
}

func (o *ProcOutput) Name() string {
	switch o.Kind {
	case OutCollect:
		return "collect"
	case OutPipe:
		return "pipe"
	case OutFile:
		return "file"
	case OutAppend:
		return "append"
	case OutDiscard:
		return "discard"
	}

	return "stdout"
}

// ReadLines runs Block for each line of Pipe, Line is the line without its newline. The
// last line counts also without a newline.
type ReadLines struct {
	Pipe  *VarValue
	Line  *VarValue
	Block *Block
}
