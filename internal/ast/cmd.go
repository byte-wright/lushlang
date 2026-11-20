package ast

import "github.com/byte-wright/lush/internal/common"

// Cmd runs a command, Args[0] is its name. The options with a type are its results, in the
// order stdout, stderr, exit.
type Cmd struct {
	Args []*CmdArg
	// Stdout and Stderr are nil for the script's own stream
	Stdout *Target
	Stderr *Target
	Exit   ExitMode
	// Stdin is nil for the script's own stdin
	Stdin *Input
	// Dir is nil for the script's working directory
	Dir Expression
	Env []*CmdEnv
	// Body is the block after the call, it reads the pipe. It's nil without one.
	Body     *Block
	Use      CmdUse
	Try      bool
	Location *common.Location
	// calculated
	ReturnTypes []common.Type
	CanFail     bool
}

type CmdArg struct {
	Value    Expression
	Spread   bool
	Location *common.Location
}

type TargetKind int

const (
	CollectString TargetKind = iota
	CollectLines
	PipeTarget
	FileTarget
	AppendTarget
	DiscardTarget
	// StdoutTarget is only for stderr, it goes where stdout goes
	StdoutTarget
)

type Target struct {
	Kind TargetKind
	// Path is set for files
	Path     Expression
	Location *common.Location
}

func (t *Target) Print() string {
	switch t.Kind {
	case CollectString:
		return "string"
	case CollectLines:
		return "[]string"
	case PipeTarget:
		return "pipe"
	case DiscardTarget:
		return "_"
	case StdoutTarget:
		return "stdout"
	case AppendTarget:
		return "appendFile"
	}

	return "file"
}

type ExitMode int

const (
	// ExitFails is the mode without exit:, a nonzero exit is an error
	ExitFails ExitMode = iota
	ExitInt
	ExitBool
	ExitIgnore
)

func (e ExitMode) Print() string {
	switch e {
	case ExitInt:
		return "int"
	case ExitBool:
		return "bool"
	case ExitIgnore:
		return "_"
	}

	return ""
}

type InputKind int

const (
	InputString InputKind = iota
	InputFile
	// InputCmd reads the pipe of another command
	InputCmd
)

type Input struct {
	Kind InputKind
	// Value is the string, the path or the *Cmd
	Value Expression
}

func (i *Input) IsFile() bool {
	return i.Kind == InputFile
}

type CmdEnv struct {
	Name     string
	Value    Expression
	Location *common.Location
}

// CmdUse says who reads the pipe of a cmd
type CmdUse int

const (
	// UseValue is a cmd whose pipe, if it has one, is read by its Body
	UseValue CmdUse = iota
	// UseSource is the stdin or an argument of another cmd, which reads its stdout pipe
	UseSource
	// UseLoop is the value of a for in, the loop reads its pipe
	UseLoop
)

// Pipe is the stream that is a pipe, "stdout", "stderr" or "" without one. With stderr:
// stdout, both go into the pipe of stdout.
func (c *Cmd) Pipe() string {
	switch {
	case c.Stdout != nil && c.Stdout.Kind == PipeTarget:
		return "stdout"
	case c.Stderr != nil && c.Stderr.Kind == PipeTarget:
		return "stderr"
	}

	return ""
}

func (c *Cmd) Types() []common.Type {
	return c.ReturnTypes
}
