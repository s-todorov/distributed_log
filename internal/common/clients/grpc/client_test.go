package grpc

import (
	"context"
	v4 "distributed_log/internal/common/api/protobuf/v4"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	c, err := NewClient("localhost:5002")
	if err != nil {
		fmt.Println(err)
	}
	prodStream, err := c.Client.ProduceStream(context.TODO())
	assert.NoError(t, err, "Grpc Client")
	loop := 5

	offset := make([]uint64, loop)

	for i := 0; i < loop; i++ {

		index := v4.ProduceRequest{Record: &v4.Record{Value: []byte(fmt.Sprintf("Hallo %d", i))}}
		err = prodStream.Send(&index)
		if err != nil {
			fmt.Println(err)
		}
		res, err := prodStream.Recv()
		if err != nil {
			fmt.Println(err)
		}
		offset[i] = res.Offset
		fmt.Println(res.Offset, i)
	}

	for i := 0; i < len(offset); i++ {
		consumStream, err := c.Client.ConsumeStream(context.TODO(), &v4.ConsumeRequest{Offset: offset[i]})
		if err != nil {
			fmt.Println(err)
		}
		res, err := consumStream.Recv()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(res.Record.Value))
	}

}
