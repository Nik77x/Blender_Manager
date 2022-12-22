package backend

import (
	"context"
	"fmt"
	"github.com/Nik77x/Blender_Manager/backend/Data"
	"github.com/Nik77x/Blender_Manager/backend/Data/Config"
	"github.com/cavaliergopher/grab/v3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"time"
)

// App struct
type App struct {
	ctx        context.Context
	BlendInfos []Data.BlendInfo
	config     Config.CfgManager
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	cfgManager := Config.NewManager()

	log.Printf(cfgManager.GetConfig().LibraryPath)

	a.UpdateVersions()

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

	req, e := grab.NewRequest(a.config.GetConfig().LibraryPath, bi.DownloadLink)

	if e != nil {
		log.Fatalf("Failed to download blender %v", e)
	}

	resp := dlClient.Do(req)

	updateProgress(&bi, resp, &a.ctx)

}

func updateProgress(info *Data.BlendInfo, res *grab.Response, ctx *context.Context) {

	// TODO: make the progress update rate a setting
	ticker := time.NewTicker(200 * time.Millisecond)

	runtime.EventsOn(*ctx, fmt.Sprintf("%v-cancel", info.CommitHash), func(optionalData ...interface{}) {
		err := res.Cancel()
		if err != nil {
			log.Println(err)
		}
	})

	go func() {

		for {
			select {

			case <-ticker.C:
				{
					// send progress event
					runtime.EventsEmit(*ctx, fmt.Sprintf("%v_dl-progress", info.CommitHash),
						res.Progress())
					// TODO: pass ETA res.ETA().Format(time.)

					log.Printf("Progress = %v", res.Progress()*100)

				}
			case <-res.Done:
				{
					// send done event
					runtime.EventsEmit(*ctx, fmt.Sprintf("%v_dl-progress", info.CommitHash), 1)

					ticker.Stop()
					return
				}

			}
		}

	}()

}
