package ast

type Statement interface {
	Element
	isStatement()
}

// ReturnStatement

type ReturnStatement struct {
	Expression    Expression
	StartPosition Position
	EndPosition   Position
}

func (ReturnStatement) isStatement() {}

func (s ReturnStatement) Accept(visitor Visitor) Repr {
	return visitor.VisitReturnStatement(s)
}

// IfStatement

type IfStatement struct {
	Test          Expression
	Then          Block
	Else          Block
	StartPosition Position
	EndPosition   Position
}

func (IfStatement) isStatement() {}

func (s IfStatement) Accept(visitor Visitor) Repr {
	return visitor.VisitIfStatement(s)
}

// WhileStatement

type WhileStatement struct {
	Test  Expression
	Block Block
}

func (WhileStatement) isStatement() {}

func (s WhileStatement) Accept(visitor Visitor) Repr {
	return visitor.VisitWhileStatement(s)
}

// Assignment

type Assignment struct {
	Target Expression
	Value  Expression
}

func (Assignment) isStatement() {}

func (s Assignment) Accept(visitor Visitor) Repr {
	return visitor.VisitAssignment(s)
}

// ExpressionStatement

type ExpressionStatement struct {
	Expression Expression
}

func (ExpressionStatement) isStatement() {}

func (s ExpressionStatement) Accept(visitor Visitor) Repr {
	return visitor.VisitExpressionStatement(s)
}
