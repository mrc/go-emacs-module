package main

//#import <stdlib.h>
//#import "emacs-module.h"
//
//emacs_env* bridge_get_environment(struct emacs_runtime *ert) {
//  return ert->get_environment(ert);
//}
//emacs_value bridge_make_function(
//  emacs_env *env,
//  ptrdiff_t min_arity,
//  ptrdiff_t max_arity,
//  emacs_value (*function) (
//    emacs_env *env,
//    ptrdiff_t nargs,
//    emacs_value args[],
//    void *),
//  const char *documentation,
//  void *data) {
//  return env->make_function(env, min_arity, max_arity, function, documentation, data);
//}
//
//emacs_value bridge_funcall(emacs_env *env, emacs_value function, ptrdiff_t nargs, emacs_value args[]) {
//  return env->funcall(env, function, nargs, args);
//}
//
//emacs_value bridge_intern(emacs_env *env, const char *symbol_name) {
//  return env->intern(env, symbol_name);
//}
//
//intmax_t bridge_extract_integer(emacs_env *env, emacs_value value) {
//  return env->extract_integer(env, value);
//}
//
//extern emacs_value FThunk(emacs_env *env, ptrdiff_t nargs, emacs_value* args, void* data);
//
//emacs_value Fcallthunk(emacs_env *env, ptrdiff_t nargs, emacs_value* args, void* data) {
//  return FThunk(env, nargs, args, data);
//}
import "C"
import "unsafe"

type EmacsRuntime C.struct_emacs_runtime
type EmacsEnv C.emacs_env
type EmacsValue C.emacs_value

type GoEmacsFunction = func(env *EmacsEnv, args []EmacsValue) EmacsValue

func (env *EmacsEnv) RegisterFunction(name, doc string, minArity, maxArity int, function GoEmacsFunction) {
	fun := env.MakeFunction(minArity, maxArity, C.Fcallthunk, doc, nil)
	env.FSet(name, fun)
}

func ERTGetEnvironment(ert *EmacsRuntime) *EmacsEnv {
	ertC := (*C.struct_emacs_runtime)(ert)
	envC := C.bridge_get_environment(ertC)
	env := (*EmacsEnv)(unsafe.Pointer(envC))
	return env
}

type EmacsFunc = unsafe.Pointer

func (env *EmacsEnv) MakeFunction(minArity, maxArity int,
	function EmacsFunc, doc string, data unsafe.Pointer) EmacsValue {
	envC := (*C.emacs_env)(unsafe.Pointer(env))
	minA, maxA := C.ptrdiff_t(minArity), C.ptrdiff_t(maxArity)
	docC := C.CString(doc)
	defer C.free(unsafe.Pointer(docC))
	thunk := (*[0]byte)(function)
	retC := C.bridge_make_function(envC, minA, maxA, thunk, docC, data)
	ret := EmacsValue(unsafe.Pointer(retC))
	return ret
}

func (env *EmacsEnv) Funcall(function EmacsValue, args []EmacsValue) EmacsValue {
	envC := (*C.emacs_env)(unsafe.Pointer(env))
	funcC := C.emacs_value(unsafe.Pointer(function))
	nargs := C.ptrdiff_t(len(args))
	var argsC *C.emacs_value
	if nargs > 0 {
		argsC = (*C.emacs_value)(&args[0])
	}
	retC := C.bridge_funcall(envC, funcC, nargs, argsC)
	ret := EmacsValue(unsafe.Pointer(retC))
	return ret
}

func (env *EmacsEnv) Intern(symbol string) EmacsValue {
	envC := (*C.emacs_env)(unsafe.Pointer(env))
	symbolC := C.CString(symbol)
	defer C.free(unsafe.Pointer(symbolC))
	valC := C.bridge_intern(envC, symbolC)
	val := EmacsValue(unsafe.Pointer(valC))
	return val
}

func (env *EmacsEnv) ExtractInteger(val EmacsValue) int {
	envC := (*C.emacs_env)(unsafe.Pointer(env))
	valC := C.emacs_value(val)
	retC := C.bridge_extract_integer(envC, valC)
	ret := int(retC)
	return ret
}
