package bot

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shahinrahimi/teleprompt/store"
)

type Bot struct {
	l                 *log.Logger
	s                 store.Storage
	api               *tgbotapi.BotAPI
	routers           map[string]*Router
	globalMiddlewares []Middleware
}

type Router struct {
	middlewares []Middleware
	handlers    map[string]Handler
}

type Handler func(*tgbotapi.Update)
type BotHandler func(*tgbotapi.Update) error
type Middleware func(Handler) Handler

func NewBot(l *log.Logger, s store.Storage, apiKey string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		l.Printf("error creating a new bot api: %v", err)
		return nil, err
	}
	return &Bot{
		l:                 l,
		s:                 s,
		api:               api,
		routers:           make(map[string]*Router),
		globalMiddlewares: []Middleware{},
	}, nil
}

func (b *Bot) Use(m Middleware) {
	b.globalMiddlewares = append(b.globalMiddlewares, m)
}

func (b *Bot) NewRouter(routerName string) *Router {
	router := &Router{
		handlers: make(map[string]Handler),
	}
	b.routers[routerName] = router
	return router
}

func (r *Router) Use(m Middleware) {
	r.middlewares = append(r.middlewares, m)
}

func (r *Router) Handle(command string, handler Handler) {
	r.handlers[command] = handler
}

func (b *Bot) Start(ctx context.Context) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	us := b.api.GetUpdatesChan(u)
	go b.receiveUpdates(ctx, us)
}

func (b *Bot) receiveUpdates(ctx context.Context, us tgbotapi.UpdatesChannel) {
	for {
		select {
		case <-ctx.Done():
			return
		case u := <-us:
			b.handleUpdate(u)
		}
	}
}

func (b *Bot) handleUpdate(u tgbotapi.Update) {

	command := u.Message.Command()
	for _, router := range b.routers {
		if handler, exists := router.handlers[command]; exists {
			// Start with the actual handler
			finalHandler := handler

			// Apply route-specific middlewares in reverse order
			for i := len(router.middlewares) - 1; i >= 0; i-- {
				finalHandler = router.middlewares[i](finalHandler)
			}
			// Apply global middlewares in reverse order
			for i := len(b.globalMiddlewares) - 1; i >= 0; i-- {
				finalHandler = b.globalMiddlewares[i](finalHandler)
			}
			// Execute the final composed handler
			finalHandler(&u)
			return
		}
	}
	// Handle unknown command
	b.l.Printf("Unknown command: %s", command)

}

func (t *Bot) SendMessage(userID int64, msgStr string) error {
	msg := tgbotapi.NewMessage(userID, msgStr)
	_, err := t.api.Send(msg)
	return err
}

func (b *Bot) Shutdown() {
	b.l.Println("Bot shutting down...")
	b.api.StopReceivingUpdates()
}
