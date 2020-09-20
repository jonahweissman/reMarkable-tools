package hypercard_whiteboard

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	nats "github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type natsClient struct {
	*nats.Conn
}

func (srv *Server) setupNats(ctx context.Context, host, user, pass string) (err error) {
	log := NewLogFromCtx(ctx)
	start := time.Now()

	uri := fmt.Sprintf("nats://%s:%s@%s:4222", user, pass, host)
	log.Info("connecting to NATS", zap.String("uri", uri))
	var nc *nats.Conn
	if nc, err = nats.Connect(uri); err != nil {
		log.Error("", zap.Error(err))
		return
	}
	srv.nc = &natsClient{nc}

	log.Info("connected",
		zap.String("to", host),
		zap.Duration("in", time.Since(start)),
	)
	return
}

func (nc *natsClient) publish(ctx context.Context, event *Event) (err error) {
	log := NewLogFromCtx(ctx)
	rk := event.rk()
	var payload []byte
	{
		log.Debug("encoding", zap.Reflect("event", event))
		start := time.Now()
		if payload, err = proto.Marshal(event); err != nil {
			log.Error("", zap.Error(err))
			return
		}
		log.Debug("encoded",
			zap.Int("bytes", len(payload)),
			zap.Duration("in", time.Since(start)),
		)
	}

	log.Debug("publishing", zap.String("rk", rk))
	start := time.Now()
	if err = nc.Publish(rk, payload); err != nil {
		log.Error("", zap.Error(err))
		return
	}
	log.Debug("published",
		zap.String("rk", rk),
		zap.Duration("in", time.Since(start)),
	)
	return
}
