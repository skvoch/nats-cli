package natsc

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/stan.go"
)

type Client struct {
	conn stan.Conn
	sub  stan.Subscription
}

func Connect(addr, clusterID, clientID string) (*Client, error) {
	conn, err := stan.Connect(clusterID, clientID, stan.NatsURL(addr))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to nats: %w", err)
	}

	return &Client{
		conn: conn,
	}, nil
}

func (c *Client) Subscribe(subject string, delta time.Duration) (chan json.RawMessage, error) {
	out := make(chan json.RawMessage, 1)

	sub, err := c.conn.Subscribe(subject, func(msg *stan.Msg) {
		out <- msg.Data
	}, stan.StartAtTimeDelta(delta))

	if err != nil {
		return nil, fmt.Errorf("failed to subscrbie to subject %s: %w", subject, err)
	}
	c.sub = sub

	return out, nil
}
