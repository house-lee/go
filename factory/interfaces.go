package factory

import "github.com/house-lee/SoarGO/coordinator"

type IWorkflow interface {

}

type IWorkStation interface {
    Init(c coordinator.ICoordinator )
}

type IJob interface {

}