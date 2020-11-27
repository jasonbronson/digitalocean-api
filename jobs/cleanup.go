package jobs

import "log"

type Cleanup struct {
	//logger *utils.Log
}

func (i *Cleanup) Run() {
	log.Println("Running cleanup job")
	//ctx := context.Background()

}
