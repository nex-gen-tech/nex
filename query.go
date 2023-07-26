package nex

// import (
// 	"github.com/google/uuid"
// )

// type query struct {
// 	Name       string  `json:"name"`
// 	Value      string  `json:"value"`
// 	Default    string  `json:"default"`
// 	IsProvided bool    `json:"isProvided"`
// 	IsRequired bool    `json:"isRequired"`
// 	Context    Context `json:"context"`
// }

// func newQuery(context Context) *query {
// 	return &query{
// 		Context: context,
// 	}
// }

// func (q *query) GetQuery(key string) *query {
// 	q.Name = key
// 	value := q.Context.QueryGet(key)
// 	if value != "" {
// 		q.Value = value
// 		q.IsProvided = true
// 	} else {

// 		q.Value = q.Default
// 		q.IsProvided = false
// 	}
// 	return q
// }

// func (q *query) WithDefault(value string) *query {
// 	q.Default = value
// 	return q
// }

// func (q *query) WithRequired() *query {
// 	q.IsRequired = true
// 	return q
// }

// type EmptyQueryError struct {
// 	Name string
// }

// func (e *EmptyQueryError) Error() string {
// 	return "The query '" + e.Name + "' is required but was not provided."
// }

// func (q *query) AsString() (string, error) {
// 	if q.IsRequired && !q.IsProvided && q.Default == "" {
// 		return "", &EmptyQueryError{Name: q.Name}
// 	}
// 	return q.Value, nil
// }

// func (q *query) AsInt() (int64, error) {
// 	if q.IsRequired && !q.IsProvided {
// 		return 0, &EmptyQueryError{Name: q.Name}
// 	}
// 	return q.Context.QueryGetInt(q.Name)
// }

// func (q *query) AsBool() (bool, error) {
// 	if q.IsRequired && !q.IsProvided && q.Default == "" {
// 		return false, &EmptyQueryError{Name: q.Name}
// 	}
// 	return q.Context.QueryGetBool(q.Name)
// }

// func (q *query) AsUUID() (uuid.UUID, error) {
// 	if q.IsRequired && !q.IsProvided && q.Default == "" {
// 		return uuid.Nil, &EmptyQueryError{Name: q.Name}
// 	}
// 	return q.Context.QueryGetUuid(q.Name)
// }
