package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

type GreetResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) (resp *GreetResponse) {
	resp = &GreetResponse{}
	if len(name) < 3 {
		resp.Error = fmt.Errorf("name must be at least 3 characters")
		return
	}
	resp.Message = fmt.Sprintf("Hello %s, It's show time!", name)
	return
}

type FooResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

// OnFoo is invoked by javascript when foo occurs
func (a *App) OnFoo(arg string) *FooResponse {
	msg, err := a.foo(arg)
	if err != nil {
		return &FooResponse{Error: err.Error()}
	}
	return &FooResponse{Message: msg}
}

func (a *App) foo(arg string) (string, error) {
	if len(arg) < 3 {
		return "", fmt.Errorf("arg must be at least 3 characters")
	}
	return fmt.Sprintf("Hello %s, It's show time!", arg), nil
}

func (a *App) rawr() {

}

// domReady is called after front-end resources have been loaded
func (a *App) domReady(ctx context.Context) {
	fmt.Println("showing window")
	runtime.WindowShow(ctx)
}
