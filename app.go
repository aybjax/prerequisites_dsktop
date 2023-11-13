package main

import (
	"changeme/repo"
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const AppVersion = "0.0.1"

type Conn struct {
	Uri  string `json:"uri"`
	Port string `json:"port"`
	User string `json:"user"`
	Pswd string `json:"pswd"`
}

// App struct
type App struct {
	ctx    context.Context
	driver neo4j.DriverWithContext
	repo   repo.Repo
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

func (b *App) shutdown(ctx context.Context) {}

func (b *App) HelpMenu(cd *menu.CallbackData) {
	runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   "Prereq Test App",
		Message: fmt.Sprintf("App version is %s", AppVersion),
	})
}

func (b *App) ConnectNeo(conn Conn) error {
	runtime.LogDebugf(b.ctx, "%#v\n", conn)
	driver, err := neo4j.NewDriverWithContext(
		fmt.Sprintf("bolt://%s:%s", conn.Uri, conn.Port),
		neo4j.BasicAuth(conn.User, conn.Pswd, ""))
	if err != nil {
		runtime.LogError(b.ctx, err.Error())
		return err
	}
	err = driver.VerifyConnectivity(context.Background())
	if err != nil {
		runtime.LogError(b.ctx, err.Error())
		return err
	}

	b.driver = driver
	b.repo = repo.NewRepo(driver)

	return nil
}

func (b *App) GetEdgeList() (map[string][]repo.Edge, error) {
	return b.repo.GetEdgeList()
}

func (b *App) GetLookUp() (map[string]string, error) {
	return b.repo.GetLookUp()
}

func (b *App) IgnoreMe(_ repo.Edge) {}
