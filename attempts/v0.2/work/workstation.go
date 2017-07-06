package work

import (
	"github.com/house-lee/SoarGO/coordinator"
	"github.com/house-lee/SoarGO/dsa"
)

type workStation struct {
	id        string
	workerNum int

	handleTask TaskHandler
    stop chan bool

	inputQueue  dsa.IQueue
	outputQueue dsa.IQueue

	coordinator coordinator.ICoordinator

	iWorkStationPrivate
}

func (ws *workStation) ID() string {
	return ws.id
}

func (ws *workStation) TaskQueue() dsa.IQueue {
	return ws.inputQueue
}

func (ws *workStation) SetCoordinator(c coordinator.ICoordinator) {
	ws.coordinator = c
}
func (ws *workStation) SetTaskHandler(handler TaskHandler) {
	ws.handleTask = handler
}

func (ws *workStation) SetWorkerNum(num int) {
	ws.workerNum = num
}

func (ws *workStation) Start() {
	ws.start(ws)
}

func (ws *workStation) Restart() {
    ws.restart(ws)
}

type iWorkStationPrivate interface {
	mainLoop(caller *workStation)
    restart(caller *workStation)
    start(caller *workStation)
}

type wsPrivate struct {
}

func (*wsPrivate) mainLoop(caller *workStation) {
	for {
		select {
		case <-caller.stop:
			return
		default:
			request := caller.inputQueue.Dequeue().(IRequest)
			caller.coordinator.UpdateRequestProgress(request.ID(), caller.id, request.Job())
			res, err := caller.handleTask(request.Job())
			if err != nil {
				caller.coordinator.SaveFailedRequest(request.ID(), err.Code(), err.Message())
				request.Respond(err)
				continue
			}
            request.UpdateJob(res)
			caller.outputQueue.Enqueue(request)
		}
	}
}

func (ws *wsPrivate) start(caller *workStation)  {
    for i:= 0; i != caller.workerNum; i++ {
        ws.mainLoop(caller)
    }
}

func (ws *wsPrivate) restart(caller *workStation)  {
    close(caller.stop)
    ws.start(caller)
}