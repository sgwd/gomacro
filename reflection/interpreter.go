// this file was generated by gomacro command: import "github.com/cosmos72/gomacro/interpreter"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package reflection

import (
	. "reflect"

	. "github.com/cosmos72/gomacro/imports"
	"github.com/cosmos72/gomacro/interpreter"
)

// reflection: allow interpreted code to import "github.com/cosmos72/gomacro/interpreter"
func init() {
	Packages["github.com/cosmos72/gomacro/interpreter"] = Package{
		Binds: map[string]Value{
			"DefaultImporter":      ValueOf(interpreter.DefaultImporter),
			"New":                  ValueOf(interpreter.New),
			"NewEnv":               ValueOf(interpreter.NewEnv),
			"NewInterpreterCommon": ValueOf(interpreter.NewInterpreterCommon),
			"Nil":                        ValueOf(&interpreter.Nil).Elem(),
			"None":                       ValueOf(&interpreter.None).Elem(),
			"OptDebugCallStack":          ValueOf(interpreter.OptDebugCallStack),
			"OptDebugMacroExpand":        ValueOf(interpreter.OptDebugMacroExpand),
			"OptDebugPanicRecover":       ValueOf(interpreter.OptDebugPanicRecover),
			"OptDebugQuasiquote":         ValueOf(interpreter.OptDebugQuasiquote),
			"OptShowAfterEval":           ValueOf(interpreter.OptShowAfterEval),
			"OptShowAfterMacroExpansion": ValueOf(interpreter.OptShowAfterMacroExpansion),
			"OptShowAfterParse":          ValueOf(interpreter.OptShowAfterParse),
			"OptShowEvalDuration":        ValueOf(interpreter.OptShowEvalDuration),
			"OptShowPrompt":              ValueOf(interpreter.OptShowPrompt),
			"OptTrapPanic":               ValueOf(interpreter.OptTrapPanic),
			"Read":                       ValueOf(interpreter.Read),
			"ReadMultiline":              ValueOf(interpreter.ReadMultiline),
			"Unknown":                    ValueOf(&interpreter.Unknown).Elem(),
		},
		Types: map[string]Type{
			"Builtin":           TypeOf((*interpreter.Builtin)(nil)).Elem(),
			"CallFrame":         TypeOf((*interpreter.CallFrame)(nil)).Elem(),
			"CallStack":         TypeOf((*interpreter.CallStack)(nil)).Elem(),
			"Cmd":               TypeOf((*interpreter.Cmd)(nil)).Elem(),
			"Env":               TypeOf((*interpreter.Env)(nil)).Elem(),
			"Error_builtin":     TypeOf((*interpreter.Error_builtin)(nil)).Elem(),
			"Function":          TypeOf((*interpreter.Function)(nil)).Elem(),
			"Importer":          TypeOf((*interpreter.Importer)(nil)).Elem(),
			"Inspector":         TypeOf((*interpreter.Inspector)(nil)).Elem(),
			"InterpreterCommon": TypeOf((*interpreter.InterpreterCommon)(nil)).Elem(),
			"Macro":             TypeOf((*interpreter.Macro)(nil)).Elem(),
			"Options":           TypeOf((*interpreter.Options)(nil)).Elem(),
			"PackageRef":        TypeOf((*interpreter.PackageRef)(nil)).Elem(),
		},
		Proxies: map[string]Type{}}
}
