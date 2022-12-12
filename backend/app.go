package backend

import (
	"context"
	"fmt"
	"github.com/Nik77x/Blender_Manager/backend/Data"
	"github.com/cavaliergopher/grab/v3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"time"
)

// App struct
type App struct {
	ctx        context.Context
	BlendInfos []Data.BlendInfo
	settings   Data.Settings
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	a.settings.LibraryPath = "H:\\AllTheStuff\\Projects\\Programming\\Go\\dl_Test"

	a.UpdateVersions()

	// TODO save and load settings + UI to change them

}

func (a *App) UpdateVersions() {

	a.BlendInfos = ScrapeVersions()

}

func (a *App) GetBlendInfos() []Data.BlendInfo {

	return a.BlendInfos

}

func (a *App) DownloadBlender(bi Data.BlendInfo) {

	log.Printf("Downloading %q", bi.VersionTitle)

	dlClient := grab.NewClient()

	req, e := grab.NewRequest(a.settings.LibraryPath, bi.DownloadLink)

	if e != nil {
		log.Fatalf("Failed to download blender %v", e)
	}

	resp := dlClient.Do(req)

	updateProgress(&bi, resp, &a.ctx)

}

func updateProgress(info *Data.BlendInfo, res *grab.Response, ctx *context.Context) {

	ticker := time.NewTicker(1 * time.Second)

	go func() {

		for {
			select {

			case <-ticker.C:
				{

					runtime.EventsEmit(*ctx, fmt.Sprintf("%v_dl-progress", info.CommitHash), res.Progress())
					log.Printf("Progress = %v", res.Progress()*100)

				}
			case <-res.Done:
				{
					ticker.Stop()
					return
				}

			}
		}

	}()

}
