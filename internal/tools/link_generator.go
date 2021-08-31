package tools

import (
	"math/rand"
	"sync"
	"time"
)

type Link struct {
	token   string
	genTime time.Time
}

type Generator struct {
	mutex       sync.Mutex
	n           int32
	letterRunes []rune
	activeTime  time.Duration
	links       map[int64]Link
}

func NewGenerator() *Generator {
	return &Generator{
		n:           50,
		activeTime:  time.Hour,
		links:       make(map[int64]Link),
		letterRunes: []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"),
	}
}

func (g *Generator) GenerateLink() string {
	b := make([]rune, g.n)
	for i := range b {
		b[i] = g.letterRunes[rand.Intn(len(g.letterRunes))]
	}

	return string(b)
}

func (g *Generator) ClearInactiveLink() {
	for {
		time.Sleep(10 * time.Minute)

		g.mutex.Lock()
		now := time.Now()

		for email, link := range g.links {
			if link.genTime.Add(g.activeTime).Before(now) {
				delete(g.links, email)
			}
		}
		g.mutex.Unlock()
	}
}

func (g *Generator) GetUserIDByToken(token string) int64 {
	for ID, link := range g.links {
		if link.token == token {
			return ID
		}
	}
	return -1
}

func (g *Generator) AddLink(ID int64, token string) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.links[ID] = Link{
		token:   token,
		genTime: time.Now(),
	}
}

func (g *Generator) DeleteLink(ID int64) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if _, ok := g.links[ID]; !ok {
		return
	}
	delete(g.links, ID)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
