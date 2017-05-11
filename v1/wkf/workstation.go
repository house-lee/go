package wkf

import (
    "github.com/house-lee/SoarGO/dsa"
    "github.com/house-lee/SoarGO/v1/req"
)

type workStation struct {
    ID         string
    WorkersNum uint32
    HandleTask TaskHandler
    InputQueue dsa.IQueue
    OutputQueue dsa.IQueue
}


func (ws *workStation) mainLoop() {
    for ; ;  {
        request := ws.InputQueue.Dequeue().(req.IRequest)
        result, err := ws.HandleTask(request.CurrentJob())
        if err != nil {
            request.SetRequestResult(/*TODO: deal with header and body*/)
        }
        //TODO: request: SetCurrentJob
        ws.OutputQueue.Enqueue(result)
    }
}