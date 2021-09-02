package container

import (
	transporter2 "apus-sample/common/transporter"
	"apus-sample/container/rest"
	"apus-sample/internal/appconf"
	"context"
	"sync"
)

type Container struct {
	transporters []transporter2.Transporter
}

func (c Container)Start() error {
	for _, tsp := range c.transporters {
		err := tsp.Start()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c Container)Stop(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(len(c.transporters))
	for _, tsp := range c.transporters {
		go func(tsp transporter2.Transporter) {
			tsp.Stop(ctx)
			wg.Done()
		}(tsp)
	}
	wg.Wait()
}

func New() (*Container, error) {
	conf := appconf.Transporter()
	restTsp, err := transporter2.RestTransporter(conf.Rest.Host, conf.Rest.Port, rest.SetupRoute)
	if err != nil {
		return nil, err
	}
	return &Container{transporters: []transporter2.Transporter{
		restTsp,
	}},nil
}
