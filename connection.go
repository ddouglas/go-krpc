package krpc

import (
	"fmt"
	"net/url"

	"github.com/ddouglas/go-krpc/internal/pb"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

// Connection contains all inforamtion required by websocket to establish connection
type Connection struct {
	URI url.URL

	Upgrader websocket.Upgrader
	Conn     *websocket.Conn
}

// NewConnection sets informations required by the connection and also creates said. If the connection cannot be established returns error and struct member Conn is unset
func NewConnection(clientName string, addr, port string) (conn Connection, e error) {
	conn = Connection{}
	conn.Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	q := url.Values{}
	q.Add("name", clientName)

	conn.URI = url.URL{
		Scheme:   "ws",
		Host:     fmt.Sprintf("%s:%s", addr, port),
		RawQuery: q.Encode(),
	}

	conn.Conn, _, e = websocket.DefaultDialer.Dial(conn.URI.String(), nil)
	return
}

// NewDefaultConnection simply calls InitializeAPI with the default port of the krpc mod
// func NewDefaultConnection() (conn Connection, e error) {
// 	conn, e = NewConnection("go-krpc", "50000")
// 	return
// }

// Close closes the current connection
func (conn *Connection) Close() error {
	if conn.Conn == nil {
		return fmt.Errorf("Connection is nil")
	}
	err := conn.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return fmt.Errorf("failed to write connection message: %w", err)
	}

	return nil
}

func (conn *Connection) sendMessage(r *pb.Request) ([]byte, error) {
	req, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	if conn.Conn == nil {
		return nil, fmt.Errorf("Connection is nil, check if any server is running on given port")
	}

	err = conn.Conn.WriteMessage(websocket.BinaryMessage, req)
	if err != nil {
		return nil, fmt.Errorf("failed to write message to socket: %w", err)
	}
	_, msg, err := conn.Conn.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("failed to write message: %w", err)
	}

	return msg, nil
}

// REQUEST AND ARGMUENT CREATION

func createRequest(service string, procedure string, arguments []*pb.Argument) *pb.Request {
	pc := &pb.ProcedureCall{
		Service:   service,
		Procedure: procedure,
		Arguments: arguments,
	}
	return &pb.Request{
		Calls: []*pb.ProcedureCall{pc},
	}

}

func createArguments(argsIn [][]byte) []*pb.Argument {

	var args = make([]*pb.Argument, 0, len(argsIn))
	for pos, val := range argsIn {
		args = append(args, &pb.Argument{
			Position: uint32(pos),
			Value:    val,
		})
	}
	return args
}
