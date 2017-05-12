package work

import (
	"github.com/house-lee/SoarGO/coordinator"
	"github.com/house-lee/SoarGO/dsa"
)

type workStation struct {
	id        string
	WorkerNum int

	HandleTask TaskHandler
	Stop       chan bool

	InputQueue  dsa.IQueue
	OutputQueue dsa.IQueue

	coordinator coordinator.ICoordinator

	private iWorkStationPrivate
}

func (ws *workStation) ID() string {
	return ws.id
}

func (ws *workStation) TaskQueue() dsa.IQueue {
	return ws.InputQueue
}

func (ws *workStation) SetCoordinator(c coordinator.ICoordinator) {
	ws.coordinator = c
}
func (ws *workStation) SetTaskHandler(handler TaskHandler) {
	ws.HandleTask = handler
}

func (ws *workStation) SetWorkerNum(num int) {
	ws.WorkerNum = num
}

func (ws *workStation) Start() {
	for i := 0; i != ws.WorkerNum; i++ {
		go ws.workerLoop()
	}
}

func (ws *workStation) Restart() {
	close(ws.Stop)
	ws.Start()
}

type iWorkStationPrivate interface {
	workerLoop()
    stop()
    start()
}

type wsPrivate struct {
    //TODO
}

func (ws *workStation) workerLoop() {
	for {
		select {
		case <-ws.Stop:
			return
		default:
			request := ws.InputQueue.Dequeue().(IRequest)
			ws.coordinator.UpdateRequestProgress(request.ID(), ws.id, request.Job())
			res, err := ws.HandleTask(request.Job())
			if err != nil {
				ws.coordinator.SaveFailedRequest(request.ID(), err.Code(), err.Message())
				continue
			}
			ws.OutputQueue.Enqueue(res)
		}
	}
}
