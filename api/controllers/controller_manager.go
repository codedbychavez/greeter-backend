package controllers

type ControllerManager struct {
	HelloController HelloControllerManager
}

func NewControllerManager() ControllerManager {
	helloController := HelloController{}

	return ControllerManager{
		HelloController: helloController,
	}
}