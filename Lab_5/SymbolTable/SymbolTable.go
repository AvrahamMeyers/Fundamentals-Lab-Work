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
