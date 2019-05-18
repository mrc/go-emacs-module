package main

//#import <stdlib.h>
//#import "emacs-module.h"
import "C"
import "fmt"
import "unsafe"

// Emacs checks for this symbol's presence when loading a module.
//export plugin_is_GPL_compatible
func plugin_is_GPL_compatible() {}

//export Answer
func Answer() C.int {
	return 42
}

//export FThunk
func FThunk(env *EmacsEnv, nargs C.ptrdiff_t, argsv *C.emacs_value, data unsafe.Pointer) EmacsValue {
	var args []EmacsValue
	return Frob(env, args)
}

func Frob(env *EmacsEnv, args []EmacsValue) EmacsValue {
	fmt.Printf("hello from frob!\n")
	return env.Intern("borf")
}

//export emacs_module_init
func emacs_module_init(ert *EmacsRuntime) C.int {

	env := ERTGetEnvironment(ert)
	env.Provide("emacsmodtest")

	// extract an int from a symbol
	versym := env.Intern("emacs-major-version")
	verval := env.SymbolValue(versym)
	ver := env.ExtractInteger(verval)
	fmt.Printf("emacs version: %d\n", ver)

	env.RegisterFunction("frob", "Frob something", 0, 0, Frob)

	return 0
}

func main() {}
