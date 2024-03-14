package main

import (
	"distributed_log/cmd"
	"distributed_log/cmd/config"
	dsLog "distributed_log/internal/log"
	"fmt"
	"github.com/hashicorp/raft"
	"log/slog"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"
)

func main() {
	app := New()

	go app.Grpc.Serve()

	go func() {
		err := http.ListenAndServe(":3333", app.Router)
		if err != nil {
			slog.Error(err.Error(), err)
		}
	}()

	pprof()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGINT)
	<-c
	app.Stop()
}

var cfgLocation = "cmd/config/.config.json"
var dataDir = "tmp"

func New() *cmd.Services {
	//c := config.NewConfig(cfgLocation)
	c := config.Config{AppName: "Test"}
	slog.Info(fmt.Sprintf("Started %s service", c.AppName))

	services, err := cmd.NewServices(
		cmd.WithIndexService(c),
		cmd.WithApiService(c),
		cmd.WithGrpcService(c),
		//cmd.WithLogStore(c),
		//cmd.WithDsLog(dsConfig(0), dataDir, 0),
		//cmd.WithDsLog(dsConfig(1), dataDir, 1),
		//cmd.WithDsLog(dsConfig(2), dataDir, 2),
	)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	nodeCount := 3

	for i := 0; i < nodeCount; i++ {
		//dataDir, err := ioutil.TempDir("", "distributed-log-test")
		//defer func(dir string) {
		//	_ = os.RemoveAll(dir)
		//}(dataDir)
		port, err := GetFreePort()
		nodeId := fmt.Sprintf("%d", i)
		ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))

		c := dsLog.Config{}
		c.Raft.StreamLayer = dsLog.NewStreamLayer(ln, nil, nil)
		c.Raft.LocalID = raft.ServerID(nodeId)
		c.Raft.HeartbeatTimeout = 50 * time.Millisecond
		c.Raft.ElectionTimeout = 50 * time.Millisecond
		c.Raft.LeaderLeaseTimeout = 50 * time.Millisecond
		c.Raft.CommitTimeout = 5 * time.Millisecond
		c.Raft.BindAddr = ln.Addr().String()

		if i == 0 {
			c.Raft.Bootstrap = true
		}

		l, err := dsLog.NewDistributedLog(path.Join(dataDir, nodeId), nodeId, c)

		if i != 0 {
			err = services.Logs[0].Join(
				fmt.Sprintf("%d", i), ln.Addr().String(),
			)
		} else {
			err = l.WaitForLeader(3 * time.Second)
			slog.Info(fmt.Sprintf("Wait for leader %v", err))
		}
		services.Logs = append(services.Logs, l)
	}

	services.EventHandler.WithStoreLog(services.Logs)
	return services
}

func dsConfig(id int) dsLog.Config {
	port, err := GetFreePort()
	if err != nil {
		slog.Error(err.Error())
	}
	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))

	c := dsLog.Config{}
	c.Raft.StreamLayer = dsLog.NewStreamLayer(ln, nil, nil)
	c.Raft.LocalID = raft.ServerID(fmt.Sprintf("%d", id))
	c.Raft.HeartbeatTimeout = 50 * time.Millisecond
	c.Raft.ElectionTimeout = 50 * time.Millisecond
	c.Raft.LeaderLeaseTimeout = 50 * time.Millisecond
	c.Raft.CommitTimeout = 5 * time.Millisecond
	c.Raft.BindAddr = ln.Addr().String()

	return c
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

// https://scene-si.org/2018/08/06/basic-monitoring-of-go-apps-with-the-runtime-package/
// Alloc - currently allocated number of bytes on the heap,
// TotalAlloc - cumulative max bytes allocated on the heap (will not decrease),
// Sys - total memory obtained from the OS,
// Mallocs and Frees - number of allocations, deallocations, and live objects (mallocs - frees),
// PauseTotalNs - total GC pauses since the app has started,
// NumGC - number of completed GC cycles
func pprof() {
	//http://localhost:6060/debug/pprof/
	// go tool pprof http://localhost:6060/debug/pprof/heap then png
	go func() {
		err := http.ListenAndServe("0.0.0.0:6060", nil)
		if err != nil {
			slog.Error(err.Error(), err)
		}
	}()
}
