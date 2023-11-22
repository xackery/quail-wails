package app

import (
	"context"

	"github.com/xackery/quail-wails/slog"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

// New creates a new app
func New() *App {
	return &App{}
}

// OnStartup should only be called during app initialization
func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx
}

// OnDomReady should only be called during ap pstart, after dom is ready
func (a *App) OnDomReady(ctx context.Context) {
	runtime.WindowShow(ctx)
}

// OnShutdown is invoked when the program exits
func (a *App) OnShutdown(ctx context.Context) {
	slog.Dump("quail-view.log")
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return "test"
}
