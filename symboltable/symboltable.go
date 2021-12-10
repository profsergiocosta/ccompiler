package symboltable

type Symbol struct {
	Name  string
	Type  string
	Index int
}

type SymbolTable struct {
	Variables  map[string]Symbol
	StartIndex int
}

func NewSymbolTable() *SymbolTable {

	return &SymbolTable{Variables: make(map[string]Symbol), StartIndex: -4}
}

func (s *SymbolTable) Define(name string) {
	symbol := Symbol{Name: name, Type: "int", Index: s.StartIndex}
	s.StartIndex = s.StartIndex - 4
	s.Variables[name] = symbol
}

func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	sym, hasSub := s.Variables[name]
	return sym, hasSub
}
