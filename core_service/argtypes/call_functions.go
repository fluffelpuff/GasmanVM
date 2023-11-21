package argtypes

type CallSharedFunctionArgs struct {
	FunctionId     string
	FunctionIdType string
	GroupName      string
	FunctionName   string
	Timeout        uint64
	Args           []interface{}
}

type CallSharedFunctionReturn struct {
	Result interface{}
}
