// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package cli_service

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var PRIMITIVE_TYPES map[TTypeId]bool
var COMPLEX_TYPES map[TTypeId]bool
var COLLECTION_TYPES map[TTypeId]bool
var TYPE_NAMES map[TTypeId]string

func init() {
	PRIMITIVE_TYPES = map[TTypeId]bool{
		0: true,
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
		6: true,
		7: true,
		8: true,
		9: true,
	}

	COMPLEX_TYPES = map[TTypeId]bool{
		10: true,
		11: true,
		12: true,
		13: true,
		14: true,
	}

	COLLECTION_TYPES = map[TTypeId]bool{
		10: true,
		11: true,
	}

	TYPE_NAMES = map[TTypeId]string{
		0:  "BOOLEAN",
		1:  "TINYINT",
		2:  "SMALLINT",
		3:  "INT",
		4:  "BIGINT",
		5:  "FLOAT",
		6:  "DOUBLE",
		7:  "STRING",
		8:  "TIMESTAMP",
		9:  "BINARY",
		10: "ARRAY",
		11: "MAP",
		12: "STRUCT",
		13: "UNIONTYPE",
	}

}
