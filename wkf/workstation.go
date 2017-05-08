package wkf

type WorkStation struct {
	ID         string
	WorkersNum uint32
	Worker     IWorker
}
