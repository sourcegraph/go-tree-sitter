package yaml

//#include "parser.h"
//TSLanguage *tree_sitter_yaml();
import "C"
import (
	"unsafe"

	sitter "github.com/sourcegraph/go-tree-sitter"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_yaml())
	return sitter.NewLanguage(ptr)
}
