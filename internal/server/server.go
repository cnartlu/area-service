package server

import (
	"errors"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/cnartlu/area-service/component/app"
	"github.com/cnartlu/area-service/component/log"
	"github.com/cnartlu/area-service/internal/server/cron"
	"github.com/cnartlu/area-service/internal/server/grpc"
	"github.com/cnartlu/area-service/internal/server/http"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport"
	"go.uber.org/zap"
)

type Server struct {
	app    *app.App
	logger *log.Logger
	gs     *grpc.Server
	hs     *http.Server
	cn     *cron.Server
	server *kratos.App
}

func (s *Server) Run(sig string) error {
	if sig != "" {
		switch sig {
		case "quit":
			process, err := s.app.GetProcess().GetPidFile().GetProcess()
			if err != nil {
				return err
			}
			return process.Signal(os.Interrupt)
		case "stop":
			process, err := s.app.GetProcess().GetPidFile().GetProcess()
			if err != nil {
				return err
			}
			return process.Kill()
		case "reload":
			process, err := s.app.GetProcess().GetPidFile().GetProcess()
			if err != nil {
				return err
			}
			return process.Signal(syscall.SIGHUP)
		default:
			buf := strings.Builder{}
			buf.WriteString("nginx")
			buf.WriteString(": invalid option: \"-")
			buf.WriteString("s ")
			buf.WriteString(sig)
			buf.WriteString("\"")
			return errors.New(buf.String())
		}
	}
	if err := s.app.GetProcess().WritePidFile(); err != nil {
		return err
	}
	// 监听信号，重启应用

	if os.Getppid() > 0 {
		var err error
		var wg = make(chan bool, 1)
		go func() {
			err = s.server.Run()
			wg <- true
		}()
		time.Sleep(time.Millisecond * 500)
		if err == nil {
			// 子进程启动完成，通知父进程结束进程
			err = s.app.GetProcess().GetProcess().Signal(syscall.SIGINT)
		} else {
			// 子进程启动失败，还原重启配置
		}
		<-wg
		return err
	}
	return s.server.Run()
}

func (s *Server) Restart() error {
	// 1、pid修复

	// 2、启动监听事件，检查子程序完成重启通知

	// 3、启动子进程，并传入响应的监听文件
	cmd := exec.Command(os.Args[0], os.Args[1:]...)
	cmd.Stdin = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}

	return s.Stop()
}

func (s *Server) Stop() error {
	if err := s.app.GetProcess().RemovePidFile(); err != nil {
		s.logger.Error("remove pid file failed", zap.Error(err))
	}
	return s.server.Stop()
}

func NewServer(
	app *app.App,
	logger *log.Logger,
	gs *grpc.Server,
	hs *http.Server,
	cn *cron.Server,
) *Server {
	var servers []transport.Server
	if cn != nil {
		servers = append(servers, cn)
	}
	if gs != nil {
		servers = append(servers, gs)
	}
	if hs != nil {
		servers = append(servers, hs)
	}
	options := []kratos.Option{
		kratos.ID(app.GetName()),
		kratos.Name(app.GetName()),
		kratos.Metadata(map[string]string{}),
		// kratos.Logger(log.NewKratosLogger(logger)),
		kratos.Server(servers...),
	}

	server := kratos.New(options...)
	return &Server{
		app:    app,
		logger: logger,
		server: server,
	}
}
