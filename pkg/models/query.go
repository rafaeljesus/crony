package models

type Query struct {
	Status     string
	Expression string
	Limit      int
	Offset     int
}

func NewQuery(status, expression string, limit, offset int) *Query {
	return &Query{status, expression, limit, offset}
}

func (q *Query) IsEmpty() bool {
	return q.Status == "" &&
		q.Expression == ""
}

func (q *Query) GetLimit() int {
	if q.Limit == 0 {
		return 10
	}
	return q.Limit
}
