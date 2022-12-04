package krpc

import (
	"github.com/ddouglas/go-krpc/internal/pb"
	"google.golang.org/protobuf/proto"
)

// GetStatus calls remote procedure `GetStatus` and returns the response in form of the proto buffer message
func (conn *Connection) GetStatus() (res pb.Status, e error) {
	pc := pb.ProcedureCall{
		Service:   "KRPC",
		Procedure: "GetStatus",
	}
	pr := pb.Request{
		Calls: []*pb.ProcedureCall{&pc},
	}

	p, e := conn.sendMessage(&pr)
	if e != nil {
		return
	}

	r := pb.Response{}
	e = proto.Unmarshal(p, &r)
	if e != nil {
		return
	}
	e = proto.Unmarshal(r.GetResults()[0].GetValue(), &res)

	return
}

// GetServices calls KRPC remote procedure `GetServices` and returns Service Structure
func (conn *Connection) GetServices() (res pb.Services, e error) {
	pc := pb.ProcedureCall{
		Service:   "KRPC",
		Procedure: "GetServices",
	}
	pr := pb.Request{
		Calls: []*pb.ProcedureCall{&pc},
	}

	p, e := conn.sendMessage(&pr)
	if e != nil {
		return
	}

	r := &pb.Response{}
	e = proto.Unmarshal(p, r)
	if e != nil {
		return
	}
	e = proto.Unmarshal(r.GetResults()[0].GetValue(), &res)
	return
}
