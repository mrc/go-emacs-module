# go-emacs-module

This doesn't work, but it's a start at getting a module written in Go that Emacs can use.

## Building

Requires `emacs-module.h`, which is created when Emacs configure is run.

Some package managers that don't keep source might remove it. macOS Homebrew removes it -- some patches ([example](https://github.com/railwaycat/homebrew-emacsmacport/pull/112)) restore it, or brew cask emacs puts it in `/Applications/Emacs.app/Contents/Resources/include/emacs-module.h`.

```
% go build -o emacsmodtest.so -buildmode=c-shared && emacs --batch -nw -Q -L $(pwd) --eval "(require 'emacsmodtest)" --eval "(princ (frob))"
emacs version: 26
hello from frob!
borf%
```

## What works

* intern a symbol, get the symbol value
* call a Go func from Emacs (just one, with no arguments)

## Todo

* Convert args for the thunk to a Go slice
* Make the thunk call a different Go func
* Figure out if all this casting is safe

## Notes

Emacs requires the symbol `plugin_is_GPL_compatible`. It doesn't need it to be anything in particular, and cgo is happy exporting functions, so I made it a function.

`fmt.Printf` is nice for testing from the command line, but would be bad at runtime.
