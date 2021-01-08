package fileconfig

import (
	"fmt"
	"io/ioutil"

	"github.com/faustikle/go-ponto/internal/errorhandler"
	"github.com/go-yaml/yaml"
	"github.com/mitchellh/go-homedir"
)

const path = ".go-ponto.yaml"

type Config struct {
	CompMan     string `yaml:"compMan"`
	Fingerprint string `yaml:"fingerprint"`
	Password    string `yaml:"password"`
	Lat         string `yaml:"lat"`
	Lon         string `yaml:"lon"`
}

func Load() (Config, error) {
	dat, err := ioutil.ReadFile(getConfigPath())
	if err != nil {
		return Config{}, fmt.Errorf("Não foi possível abrir o arquivo %s, execute o comando 'init' novamente.")
	}

	var cfg Config
	err = yaml.Unmarshal(dat, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("Formato do arquivo de configuração é inválido, execute o comando 'init' novamente.")
	}

	return cfg, nil
}

func Save(cfg Config) error {
	content, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Não foi possivel criar o arquivo de configuração: %s", err)
	}

	path := getConfigPath()
	err = ioutil.WriteFile(path, content, 0644)
	if err != nil {
		return fmt.Errorf("Não foi possivel salvar o arquivo de configuração no caminho %s: %s", path, err)
	}

	fmt.Printf("Configurações salvas em %s\n", path)
	return nil
}

func getConfigPath() string {
	home, err := homedir.Dir()
	if err != nil {
		errorhandler.Handler(fmt.Errorf("Não foi possivel encontrar a Home do usuário."))
		return ""
	}

	return fmt.Sprintf("%s/%s", home, path)
}