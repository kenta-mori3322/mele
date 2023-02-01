package types

// DONTCOVER

const (
	QueryParams     = "params"
	QueryExecutions = "executions"
	QueryExecution  = "execution"
)

type QueryExecutionParams struct {
	ExecutionID uint64
}

func NewQueryExecutionParams(executionID uint64) QueryExecutionParams {
	return QueryExecutionParams{
		ExecutionID: executionID,
	}
}

