package rabbitmq

import (
	"testing"
	"time"
)

func TestConsume(t *testing.T) {
	client,err := NewRecvClient(RecvConfig{
		Config:Config{
			User:"admin",
			Password:"admin",
			Host:"10.1.11.127",
			Port:5672,
			VHost:"Affair",
		},
		ExName:"Affair.Exchange",
		QueueName:"Test",
		RouteKey:"Test.#",
	});

	if err != nil {
		t.Log(err);
		t.Errorf("TestConsume failed");
	}

	f := func(msg []byte) bool {
		str := string(msg);
		t.Log(str);
		return true;
	};

	err = client.Consume(f,true);
	if err != nil {
		t.Log(err);
		t.Errorf("TestConsume failed");
	}

	time.Sleep(20*time.Second);
}