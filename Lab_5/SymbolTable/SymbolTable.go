package SymbolTable

import (
	"strconv"
)

// Name is the identifier of the variable
// Type of the variable (int, string, etc.)
// Kind of the variable (STATIC, FIELD, ARG, VAR)
// Index is the running index of the variable in the current scope
type SymbolRow struct {
	Name  string
	Type  string
	Kind  string
	Index int
}

type SymbolTable struct {
	ClassScope      map[string]SymbolRow
	SubroutineScope map[string]SymbolRow
}

// Creates a new empty symbol tables, ClassScope and SubroutineScope
func (X *SymbolTable) Constructor() {
	X.ClassScope = make(map[string]SymbolRow)
	X.SubroutineScope = make(map[string]SymbolRow)
}

// Starts a new subroutine scope (i.e., resets the subroutine's symbol table)
func (X *SymbolTable) StartSubroutine() {
	X.SubroutineScope = make(map[string]SymbolRow)
}

// Defines a new identifier of a given name, type, and kind and assigns it
// a running index and puts it in the current scope
func (X *SymbolTable) Define(name string, typ string, kind string) {
	if kind == "STATIC" || kind == "FIELD" {
		X.ClassScope[name] = SymbolRow{name, kind, typ, X.VarCount(kind)}
	} else {
		X.SubroutineScope[name] = SymbolRow{name, kind, typ, X.VarCount(kind)}
	}
}

// Returns the number of variables of the given kind already defined in the current scope
// Kind can be STATIC, FIELD, ARG, VAR.
func (X *SymbolTable) VarCount(kind string) int {
	count := 0
	if kind == "STATIC" || kind == "FIELD" {
		for _, v := range X.ClassScope {
			if v.Kind == kind {
				count++
			}
		}
	} else {
		for _, v := range X.SubroutineScope {
			if v.Kind == kind {
				count++
			}
		}
	}
	return count
}

// Returns the kind of the named identifier in the current scope.
// Kind can be STATIC, FIELD, ARG, VAR.
// If the identifier is unknown in the current scope, returns NONE.
func (X *SymbolTable) KindOf(name string) string {
	if _, ok := X.SubroutineScope[name]; ok { // If the identifier is in the subroutine scope (if the key value pair exists)
		return X.SubroutineScope[name].Kind
	} else if _, ok := X.ClassScope[name]; ok {
		return X.ClassScope[name].Kind
	}
	return "NONE"
}

// Returns the type of the named identifier in the current scope.
// Type is the type of the variable (int, string, etc.)
func (X *SymbolTable) TypeOf(name string) string {
	if _, ok := X.SubroutineScope[name]; ok {
		return X.SubroutineScope[name].Type
	} else {
		return X.ClassScope[name].Type
	}
}

// Returns the index assigned to the named identifier.
func (X *SymbolTable) IndexOf(name string) int {
	if _, ok := X.SubroutineScope[name]; ok {
		return X.SubroutineScope[name].Index
	} else {
		return X.ClassScope[name].Index
	}
}

// Returns test XML string for testing the symbol table rows
func (X *SymbolTable) IdentifierToXML(name string, isDeclaration bool) string {
	return "<identifier>\n" +
		"<name>" + name + "</name>\n" +
		"<isDeclaration>" + strconv.FormatBool(isDeclaration) + "</isDeclaration>\n " +
		"<type>" + X.TypeOf(name) + "</type>\n" +
		"<kind>" + X.KindOf(name) + "</kind>\n" +
		"<index>" + strconv.Itoa(X.IndexOf(name)) + "</index>\n" +
		"</identifier>\n"
}
