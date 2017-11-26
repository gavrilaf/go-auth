package backend

import (
	"sync"
	"time"

	"github.com/gavrilaf/spawn/pkg/cache"
	"github.com/gavrilaf/spawn/pkg/dbx"
	"github.com/gavrilaf/spawn/pkg/env"
	log "github.com/sirupsen/logrus"
)

type ServerState int32

const (
	StateCreated ServerState = iota
	StateLoading
	StateOk
	StateError
)

type Server struct {
	db    dbx.Database
	cache cache.Cache
	state ServerState
	wg    *sync.WaitGroup
}

func CreateServer(en *env.Environment) *Server {
	log.Infof("Starting backend with environment: %v", en.GetName())

	db, err := dbx.Connect(en)
	if err != nil {
		log.Errorf("Can not connect to database: %v", err)
		return nil
	}
	log.Infof("Db connection, ok")

	cache, err := cache.Connect(en)
	if err != nil {
		log.Errorf("Can not connect to cache: %v", err)
		return nil
	}
	log.Infof("Cache connection, ok")

	return &Server{db: db, cache: cache, state: StateCreated, wg: &sync.WaitGroup{}}
}

func (srv *Server) StartServer() {
	log.Infof("Server started...")
	srv.updateServerState(StateLoading)

	srv.wg.Add(1)

	go srv.LoadAuthCache()

	timeout := waitTimeout(srv.wg, 10*time.Second)

	if srv.state == StateLoading {
		if timeout {
			log.Errorf("Server loading timeout")
			srv.updateServerState(StateError)
		} else {
			srv.updateServerState(StateOk)
		}
	}

	log.Infof("Server initializing finished with state %d", srv.state)
}

func (srv *Server) GetServerState() ServerState {
	return srv.state
}

func (srv *Server) updateServerState(newState ServerState) {
	srv.state = newState
}