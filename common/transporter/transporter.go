package transporter

import "context"

type Transporter interface {
	Name() string
	Start() error
	Stop(ctx context.Context)
}
