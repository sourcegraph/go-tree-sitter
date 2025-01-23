package ocaml

//#include "parser.h"
//TSLanguage *tree_sitter_ocaml();
import "C"
import (
	"unsafe"

	sitter "github.com/sourcegraph/go-tree-sitter"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_ocaml())
	return sitter.NewLanguage(ptr)
}
