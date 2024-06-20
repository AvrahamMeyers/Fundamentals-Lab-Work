package compilationengine

import (
	"github.com/AvrahamMeyers/Fundamentals-Lab-Work/Lab_4/Tokenizer"
)

type comp struct {
	token Tokenizer.Tokenizer
}

func (X *comp) Constructor() {
	/*Creates a new compilation
	engine with the given input and
	output. The next routine called
	must be compileClass()..*/

}

func (X *comp) CompileClass() {
	//Compiles a complete class.
}

func (X *comp) CompileClassVarDec() {
	/*Compiles a static declaration or
	a field declaration.*/
}

func (X *comp) CompileSubroutine() {
	//Compiles a complete method,
	//function, or constructor.
}

func (X *comp) CompileParameterList() {
	/*Compiles a (possibly empty)
	parameter list, not including the
	enclosing ‘‘()’’..*/
}

func (X *comp) CompileVarDec() {
	/* Compiles a var declaration..*/
}

func (X *comp) CompileStatements() {
	/* Compiles a sequence of state
	ments, not including the
	 enclosing ‘‘{}’’..*/
}

func (X *comp) CompileDo() {
	/* Compiles a do statement..*/
}

func (X *comp) CompileLet() {
	/*  Compiles a let statement..*/
}
func (X *comp) CompileWhile() {
	/* Compiles a while statement..*/
}
func (X *comp) CompileReturn() {
	/*  Compiles a return statement.*/
}
func (X *comp) CompileIf() {
	/*  Compiles an if statement, pos
	sibly with a trailing else clause.T.*/
}
func (X *comp) CompileExpression() {
	/* Compiles an expression..*/
}
func (X *comp) CompileTerm() {
	/*Compiles a term. This routine is
	 faced with a slight difficulty
	 when trying to decide between
	 some of the alternative parsing
	 rules. Specifically, if the current
	 token is an identifier, the routine
	 must distinguish between a
	 variable, an array entry, and a
	 subroutine call. A single look
	ahead token, which may be one
	 of ‘‘[’’, ‘‘(’’, or ‘‘.’’ suffices to dis
	tinguish between the three possi
	bilities. Any other token is not
	 part of this term and should not
	 be advanced over.*/
}
func (X *comp) CompileExpressionList() {
	/*Compiles a (possibly empty)
	comma-separated list of
	expressions.*/
}
