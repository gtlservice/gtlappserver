package ctrl

import "github.com/gtlservice/gtlappserver/base"
import "github.com/gtlservice/gutils/logger"
import "github.com/gtlservice/gzkwrapper"

type Controller struct {
	Args   *base.ServiceArgs
	Worker *gzkwrapper.Worker
}

func NewController(worker *gzkwrapper.Worker, args *base.ServiceArgs) *Controller {

	return &Controller{
		Args:   args,
		Worker: worker,
	}
}

func (c *Controller) Initialize() error {

	logger.INFO("[#ctrl#] controller initializeing......")
	if err := c.Worker.Open(); err != nil {
		return err
	}

	if err := c.Worker.Signin(c.Args.Service); err != nil {
		return err
	}
	return nil
}

func (c *Controller) UnInitialize() error {

	if err := c.Worker.Signout(); err != nil {
		return err
	}

	if err := c.Worker.Close(); err != nil {
		return err
	}
	return nil
}
