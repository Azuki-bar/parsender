package setting

import (
	"github.com/adrg/xdg"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"os"
)

type Conf struct {
	SmtpConf SmtpConf `toml:"smtpConf"`
}
type SmtpConf struct {
	Auth   smtpAuth   `toml:"auth"`
	Server smtpServer `toml:"server"`
}
type smtpAuth struct {
	Identity string `toml:"identity"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
}
type smtpServer struct {
	Address string `toml:"address"`
	From    string `toml:"from"`
}

func GetToml() error {
	a := Conf{}
	def := toml.NewEncoder(os.Stdout)
	if err := def.Encode(a); err != nil {
		return err
	}
	return nil
}
func parseSmtpConf(filePath string) (*SmtpConf, error) {
	a := Conf{}
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = toml.Unmarshal(f, &a)
	if err != nil {
		return nil, err
	}
	return &a.SmtpConf, nil
}
func GetSmtpConf(filePath string) (*SmtpConf, error) {
	var authFile string
	if (len(filePath)) != 0 {
		authFile = filePath
	}
	authFile, err := xdg.SearchConfigFile("parsender/config.toml")
	if err != nil {
		return nil, err
	}
	return parseSmtpConf(authFile)
}
