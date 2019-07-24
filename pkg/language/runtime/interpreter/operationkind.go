package interpreter

import "github.com/dapperlabs/bamboo-node/pkg/language/runtime/errors"

//go:generate stringer -type=OperationKind

type OperationKind int

const (
	OperationKindUnary OperationKind = iota
	OperationKindBinary
	OperationKindTernary
)

func (k OperationKind) Name() string {
	switch k {
	case OperationKindUnary:
		return "unary"
	case OperationKindBinary:
		return "binary"
	case OperationKindTernary:
		return "ternary"
	}

	panic(&errors.UnreachableError{})
}
