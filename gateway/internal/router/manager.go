package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
	"strings"
	"time"
)

type Manager struct {
	handle *gin.Engine
}

func New() *Manager {
	return &Manager {
		handle: gin.Default(),
	}
}

// 管理员 处理所有的路由
func (m *Manager) Run() error {
	err := m.load()
	if err != nil {
		return err
	}
	// 
	return nil
}

// 
func (m *Manager) load() (err error) {
	err = m.loadPlugin()
	if err != nil {
		return
	}
	return m.loadRoute()
}

// 加载中间件
func (m *Manager) loadPlugin() error {
	if err := m.loadCORS(); err != nil {
		return err
	}
	return nil
}

// 使用 cors 中间件，设置允许来源以及允许的方法
func (m *Manager) loadCORS() error {
	// iris cors
	//crs := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"*"},
	//	AllowedMethods:   []string{"POST", "GET", "OPTIONS", "DELETE"},
	//	MaxAge:           3600,
	//	AllowedHeaders:   []string{"*"},
	//	AllowCredentials: true,
	//})
	//m.handle.UseRouter(crs)

	// gin cors
	crs := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 24 * time.Hour,
	})
	m.handle.Use(crs)
	return nil
}

func (m *Manager) loadRoute() error {
	t := reflect.TypeOf(m)
	for i := 0; i < t.NumMethod(); i++ {
		f := t.Method(i)
		if strings.HasPrefix(f.Name, "Route") &&
			f.Type.NumOut() == 0 &&
			f.Type.NumIn() == 1 {
			log.Println("[GATEWAY] LOAD ROUTE:", f.Name)
			f.Func.Call([]reflect.Value{
				reflect.ValueOf(m),
			})
		}
	}
	return nil
}
