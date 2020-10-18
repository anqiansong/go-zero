package mock

type (
	Visitor struct {
		invoker *InvokeMatcher
		value   *Value
	}
)

func NewVisitor(invoker *InvokeMatcher) {

}
