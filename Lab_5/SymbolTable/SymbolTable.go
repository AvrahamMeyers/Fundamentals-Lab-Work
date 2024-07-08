package SymbolTable

type SymbolRow struct {
	Identifier string
	Kind       string
	Type       string
	Index      int
}

type SymbolTable struct {
	ClassScope      map[string]SymbolRow
	SubroutineScope map[string]SymbolRow
}

func (X *SymbolTable) Constructor() {
	X.ClassScope = make(map[string]SymbolRow)
	X.SubroutineScope = make(map[string]SymbolRow)
}

func (X *SymbolTable) StartSubroutine() {
	X.SubroutineScope = make(map[string]SymbolRow)
}

func (X *SymbolTable) Define(name string, typ string, kind string) {
	if kind == "static" || kind == "field" {
		X.ClassScope[name] = SymbolRow{name, kind, typ, X.VarCount(kind)}
	} else {
		X.SubroutineScope[name] = SymbolRow{name, kind, typ, X.VarCount(kind)}
	}
}

func (X *SymbolTable) VarCount(kind string) int {
	count := 0
	if kind == "static" || kind == "field" {
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
