package Config

import (
	"github.com/kirsle/configdir"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
)

type Config struct {
	LibraryPath string `json:"library_path"`
}

type CfgManager struct {
	config     Config
	vi         *viper.Viper
	configPath string
}

func NewManager() CfgManager {

	configName := "BVM-CfgManager"

	c := CfgManager{}

	c.configPath = configdir.LocalConfig("Blender_Manager")

	err := configdir.MakePath(c.configPath)

	if err != nil {
		log.Fatalf("Failed to create config directory!\nconfig path = %q", c.configPath)
	}

	c.setupViper(&configName)

	c.checkConfigFile(&configName)

	return c
}

func (c *CfgManager) GetConfig() *Config {

	return &c.config

}

func (c *CfgManager) LoadFromDisk() Config {
	return Config{
		LibraryPath: c.vi.GetString("LibraryPath"),
	}
}

func (c *CfgManager) SaveToDisk() {

	err := viper.WriteConfig()
	if err != nil {
		log.Println("Failed to write config")
	}

}

func (c *CfgManager) setupViper(configName *string) {
	c.vi = viper.New()
	c.vi.AddConfigPath(c.configPath)
	c.vi.SetDefault("LibraryPath", c.configPath)
	c.vi.SetConfigName(*configName)
	c.vi.SetConfigType("yaml")
}

func (c *CfgManager) checkConfigFile(configName *string) {
	log.Println("Checking config file...")
	filePath := path.Join(c.configPath, *configName+".yaml")

	log.Printf("Config file path: %q", filePath)
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		log.Printf("Config does not exist, creating a new one...")
		err := c.vi.WriteConfigAs(path.Join(c.configPath, *configName+".yaml"))

		if err != nil {
			log.Fatalf("Failed to write config file!")
		}

	}
}
