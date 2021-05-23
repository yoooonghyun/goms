package contoller

type ControllerType int
const (
   ControllerTypeRest ControllerType = 1 + iota
   ControllerTypeRpc ControllerType
)

type BaseController interface {
  handle(command string, dto *DTO) bool
}

type Controller struct {
  controller_map map[ControllerType]BaseController
}

func NewController() *Controller {
  ret := Controller{
    controller_map := make(map[ControllerType]*BaseController)
  }
  controller_map[ControllerTypeRest] = NewRestController()
  controller_map[ControllerTypeRpc] = NewRpcController()
  return ret;
}

func GetInstance() *Controller {
  if instance == nil {
    instance = NewController();
  }
  return instance
}

func (ctrl *Controller) Post() {

}
