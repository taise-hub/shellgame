package model

type ExecutionResult struct {
	DataType 	 string //WARN: 変更の可能性大
	Command      string
	StdOut 		 []byte
	StdErr   	 []byte
	// Owner    	 bool
}