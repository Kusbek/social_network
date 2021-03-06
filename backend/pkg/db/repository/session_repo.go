package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/lithammer/shortuuid"

	"git.01.alem.school/Kusbek/social-network/backend/entity"
)

type session struct {
	user       *entity.User
	expireTime time.Time
}

//SessionRepository ...
type SessionRepository struct {
	expiration time.Duration
	sessions   map[string]*session
	mu         *sync.Mutex
}

//NewSessionRepository ...
func NewSessionRepository() *SessionRepository {
	c := &SessionRepository{
		expiration: 10,
		sessions:   make(map[string]*session),
		mu:         &sync.Mutex{},
	}
	go c.monitor()
	return c
}

//Create ...
func (c *SessionRepository) Create(u *entity.User) string {
	uuid := shortuuid.New()
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sessions[uuid] = &session{user: u, expireTime: time.Now().Add(c.expiration * time.Hour)}
	return uuid
}

//Get ...
func (c *SessionRepository) Get(uuid string) (*entity.User, error) {
	session, ok := c.sessions[uuid]
	if !ok {
		return nil, errors.New("Unauthorized")
	}
	return session.user, nil
}

//Delete ...
func (c *SessionRepository) Delete(uuid string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.sessions, uuid)
}

func (c *SessionRepository) monitor() {
	for {
		// fmt.Println("cookie monitoring!!!")
		time.Sleep(10 * time.Second)
		for key, value := range c.sessions {

			if value.expireTime.Before(time.Now()) {
				c.mu.Lock()
				delete(c.sessions, key)
				c.mu.Unlock()
			}
		}
	}
}
