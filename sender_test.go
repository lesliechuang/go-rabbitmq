package rabbitmq

import (
	"testing"
)

func TestSend(t *testing.T) {
	client,err := NewSendClient(SendConfig{
		Config:Config{
			User:"admin",
			Password:"admin",
			Host:"10.1.11.127",
			Port:5672,
			VHost:"Affair",
		},
		ExName:"Affair.Exchange",
	});

	if err != nil {
		t.Log(err);
		t.Errorf("TestSend failed");
	}

	msg := "this is a new message from test send";
	err = client.Send("Test.1",[]byte(msg));

	if err != nil {
		t.Log(err);
		t.Errorf("TestSend failed");
	}

	t.Logf("send result is %s",err);
}