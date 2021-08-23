package producer

import (
	"encoding/json"

	"github.com/Shopify/sarama"
)

type producer struct {
	prod sarama.SyncProducer
}

type event struct {
	Type string `json:"type"`
	Id   uint64 `json:"id"`
}

func New(brokers []string) (Producer, error) {
	conf := sarama.NewConfig()
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	prod, err := sarama.NewSyncProducer(brokers, conf)
	if err != nil {
		return nil, err
	}

	return &producer{
		prod: prod,
	}, nil
}

func (p *producer) Send(topic string, event Event) error {
	prodMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(event.Value()),
	}

	_, _, err := p.prod.SendMessage(prodMsg)
	return err
}

func (p *producer) Close() {
	p.prod.Close()
}

func PrepareEvent(typ EventType, id uint64) Event {
	return &event{
		Type: typ.String(),
		Id:   id,
	}
}

func (e *event) Value() string {
	if e == nil {
		return ""
	}

	msg, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(msg)
}
