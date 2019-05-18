package main

// Provide calls `provide`
func (env *EmacsEnv) Provide(feature string) {
	feat := env.Intern(feature)
	prov := env.Intern("provide")
	args := []EmacsValue{feat}
	env.Funcall(prov, args)
}

// SymbolValue calls `symbol-value`
func (env *EmacsEnv) SymbolValue(sym EmacsValue) EmacsValue {
	symbolValue := env.Intern("symbol-value")
	return env.Funcall(symbolValue, []EmacsValue{sym})
}

// FSet calls `fset`
func (env *EmacsEnv) FSet(name string, sfun EmacsValue) {
	fset := env.Intern("fset")
	sym := env.Intern(name)
	env.Funcall(fset, []EmacsValue{sym, sfun})
}
