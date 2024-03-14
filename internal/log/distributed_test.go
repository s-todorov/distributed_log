package log_test

import (
	v4 "distributed_log/internal/common/api/protobuf/v4"
	dsLog "distributed_log/internal/log"
	"fmt"
	"io/ioutil"
	"log"
	"log/slog"
	"net"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/raft"
	"github.com/stretchr/testify/require"
)

func TestMultipleNodes(t *testing.T) {
	var logs []*dsLog.DistributedLog
	nodeCount := 3

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	for i := 0; i < nodeCount; i++ {
		dataDir, err := ioutil.TempDir("", "distributed-log-test")
		require.NoError(t, err)
		//defer func(dir string) {
		//	_ = os.RemoveAll(dir)
		//}(dataDir)
		port, err := GetFreePort()
		require.NoError(t, err)
		ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
		require.NoError(t, err)

		config := dsLog.Config{}
		config.Raft.StreamLayer = dsLog.NewStreamLayer(ln, nil, nil)
		config.Raft.LocalID = raft.ServerID(fmt.Sprintf("%d", i))
		config.Raft.HeartbeatTimeout = 50 * time.Millisecond
		config.Raft.ElectionTimeout = 50 * time.Millisecond
		config.Raft.LeaderLeaseTimeout = 50 * time.Millisecond
		config.Raft.CommitTimeout = 5 * time.Millisecond
		config.Raft.BindAddr = ln.Addr().String()

		if i == 0 {
			config.Raft.Bootstrap = true
		}

		l, err := dsLog.NewDistributedLog(dataDir, strconv.Itoa(i), config)
		require.NoError(t, err)

		if i != 0 {
			err = logs[0].Join(
				fmt.Sprintf("%d", i), ln.Addr().String(),
			)
			require.NoError(t, err)
		} else {
			err = l.WaitForLeader(3 * time.Second)
			slog.Info(fmt.Sprintf("Wait for leader %v", err))
			require.NoError(t, err)
		}
		logs = append(logs, l)
	}

	records := []*v4.Record{
		{Value: []byte("first")},
		{Value: []byte("second")},
	}

	var leaderNode *dsLog.DistributedLog

	for _, node := range logs {
		if node.IsLeader() {
			leaderNode = node
			break
		}
	}

	for _, record := range records {

		off, err := leaderNode.Append(record)
		require.NoError(t, err)

		require.Eventually(t, func() bool {
			for j := 0; j < nodeCount; j++ {
				got, err := logs[j].Read(off)
				if err != nil {
					return false
				}
				record.Offset = off
				fmt.Println(string(got.Value))
				if !reflect.DeepEqual(got.Value, record.Value) {
					return false
				}
			}
			return true
		}, 500*time.Millisecond, 50*time.Millisecond)
	}
}

func GetFreePort() (port int, err error) {
	var a *net.TCPAddr
	if a, err = net.ResolveTCPAddr("tcp", "localhost:0"); err == nil {
		var l *net.TCPListener
		if l, err = net.ListenTCP("tcp", a); err == nil {
			defer l.Close()
			return l.Addr().(*net.TCPAddr).Port, nil
		}
	}
	return
}
