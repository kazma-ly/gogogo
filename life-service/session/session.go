package session

import (
	"life-service/tool"
	"log"
	"net/http"
	"sync"
	"time"
)

type (
	// Session 的基本操作
	Session interface {
		Set(key, value interface{}) error // set session value
		Get(key interface{}) interface{}  // get session value
		Delete(key interface{}) error     // delete session value
		SessionID() string                // back current sessionID
	}

	// Manager 全局的session管理器
	Manager struct {
		cookieName string     // private cookieName
		lock       sync.Mutex // protects session
		//provider    Provider
		maxlifetime int64 // 秒
		info        *SessionInfo
	}

	// SessionInfo session里的信息
	SessionInfo struct {
		sid          string                      // session id
		timeAccessed time.Time                   // 最后访问的时间
		value        map[interface{}]interface{} // session里的值 key是session的键
	}
)

var (
	loadOnce            sync.Once                      // 初始化一次session GC
	lock                sync.Mutex                     // 新建session的锁
	globalSessionManage = make(map[string]*Manager, 0) // globalSessionManage 全局session管理器
)

const (
	SESSION_NAME = "GSESSIONID" // SESSION_NAME session的名字
)

// GetSessionManage 获得session， 或有会新建一个
func GetSessionManage(req *http.Request, w http.ResponseWriter) *Manager {
	lock.Lock()
	defer lock.Unlock()

	gSessionID, _ := req.Cookie(SESSION_NAME)
	if gSessionID == nil || globalSessionManage[gSessionID.Value] == nil {
		gID := tool.MakeRandomNum()
		// 添加cookie 下次可以或者这个session
		http.SetCookie(w, &http.Cookie{
			Name:     SESSION_NAME,
			Value:    string(gID),
			Path:     "/",
			Expires:  time.Now().Add(time.Minute * 30),
			HttpOnly: true,
		})
		return newSessionManager(gID, 60*30)
	}
	return globalSessionManage[gSessionID.Value]
}

// ***********
// 操作Session
// ***********

// GetSessionManager 初始化一个session管理器
func newSessionManager(cookieName string, maxlifetime int64) *Manager {
	v := make(map[interface{}]interface{}, 0)
	m := &Manager{
		cookieName:  cookieName,
		maxlifetime: maxlifetime,
		info: &SessionInfo{
			sid:          cookieName,
			timeAccessed: time.Now(),
			value:        v,
		},
	}
	globalSessionManage[cookieName] = m
	return m
}

// Set set session value
func (m *Manager) Set(key, value interface{}) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.info.value[key] = value
	m.updateTime()
	return nil
}

// Get get session value
func (m *Manager) Get(key interface{}) interface{} {
	m.updateTime()
	if v, ok := m.info.value[key]; ok {
		return v
	}
	return nil
}

// Delete delete session value
func (m *Manager) Delete(key interface{}) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.updateTime()
	delete(m.info.value, key)
	return nil
}

// SessionID back current sessionID
func (m *Manager) SessionID() string {
	m.updateTime()
	return m.cookieName
}

// 更新操作时间
func (m *Manager) updateTime() {
	m.info.timeAccessed = time.Now()
}

//  ***********
// Session管理器操作
//  ***********

// DestroySession 删除一个session
func (m *Manager) DestroySession(sid string) {
	delete(globalSessionManage, sid)
}

// ProcessSession 自动删除过期Session
func ProcessSession() {
	loadOnce.Do(makeSessionGC)
}

func makeSessionGC() {
	log.Println("设置检查session")
	go func() {
		for {
			for _, m := range globalSessionManage {
				if (time.Now().Unix() - m.info.timeAccessed.Unix()) > m.maxlifetime { // session 过期
					m.DestroySession(m.cookieName)
				}
			}
			time.Sleep(time.Second * 10)
		}
	}()
}
