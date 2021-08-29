package metrics

type Metrics interface {
	IncCreate()
	IncUpdate()
	IncDelete()
}
