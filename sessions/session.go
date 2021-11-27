package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxlifetime int64
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) (Session, error)
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

// var gS *session.Manager

var provides = make(map[string]Provider)

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}

	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}

func NewManager(providerName, cookieName string, maxlifetime int64) (*Manager, error) {
	provider, ok := provides[providerName]
	if !ok {
		return nil, fmt.Errorf("session:unknown provide %q (forgotten import?)", providerName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

// func Init() {
// 	globalSessions, _ = NewManager("memory", "gosessionid", 3600)
// 	go globalSessions.GC()
// }

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

// func (manager *Manager) count(w http.ResponseWriter, r *http.Request) {
// 	sess := globalSessions.SessionStart(w, r)
// 	createtime := sess.Get("createtime")
// 	if createtime == nil {
// 		sess.Set("createtime", time.Now().Unix())
// 	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
// 		globalSessions.SessionDestroy(w, r)
// 		sess = globalSessions.SessionStart(w, r)
// 	}
//
// 	ct := sess.Get("countnum")
// 	if ct == nil {
// 		sess.Set("countnum", 1)
// 	} else {
// 		sess.Set("countnum", (ct.(int) + 1))
// 	}
// 	t, _ := template.ParseFiles("count.gtpl")
// 	w.Header().Set("Content-Type", "text/html")
// 	t.Execute(w, sess.Get("countnum"))
// }

func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}

func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxlifetime)
	time.AfterFunc(time.Duration(manager.maxlifetime), func() { manager.GC() })
}
