package calcapi

import (
	fmt "fmt"
	"time"
)

func (svc *svc) EchoStream(srv Service_EchoStreamServer) error {
	for {
		req, err := srv.Recv()
		if err != nil {
			return err
		}

		err = srv.Send(&EchoStream_Output{Msg: req.Msg})
		if err != nil {
			return err
		}
	}
}

func (svc *svc) TestReadStream(in *TestReadStream_Input, stream Service_TestReadStreamServer) error {
	start := time.Now()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		err := stream.Send(&TestReadStream_Output{Msg: fmt.Sprintf("Hello! %v", time.Since(start))})
		if err != nil {
			return err
		}
	}
	return nil
}
