package secrets

import (
	"github.com/goatcms/goatcli/cliapp/common/cio"
	"github.com/goatcms/goatcli/cliapp/common/config"
	"github.com/goatcms/goatcli/cliapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/varutil/plainmap"
)

// Secrets provide project secrets data
type Secrets struct {
	deps struct {
		FS     filesystem.Filespace `filespace:"root"`
		Input  app.Input            `dependency:"InputService"`
		Output app.Output           `dependency:"OutputService"`
	}
}

// Factory create new repositories instance
func Factory(dp dependency.Provider) (interface{}, error) {
	var err error
	instance := &Secrets{}
	if err = dp.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return services.SecretsService(instance), nil
}

// ReadDefFromFS read secrets definitions from filespace
func (p *Secrets) ReadDefFromFS(fs filesystem.Filespace) (secrets []*config.Property, err error) {
	var json []byte
	if !fs.IsFile(SecretsDefPath) {
		return make([]*config.Property, 0), nil
	}
	if json, err = fs.ReadFile(SecretsDefPath); err != nil {
		return nil, err
	}
	if secrets, err = config.NewProperties(json); err != nil {
		return nil, err
	}
	return secrets, nil
}

// ReadDataFromFS read secrets data from filespace
func (p *Secrets) ReadDataFromFS(fs filesystem.Filespace) (data map[string]string, err error) {
	var json []byte
	if !fs.IsFile(SecretsDataPath) {
		return make(map[string]string, 0), nil
	}
	if json, err = fs.ReadFile(SecretsDataPath); err != nil {
		return nil, err
	}
	if data, err = plainmap.JSONToPlainStringMap(json); err != nil {
		return nil, err
	}
	return data, nil
}

// FillData read lost secrets data to curent data map
func (p *Secrets) FillData(def []*config.Property, data map[string]string, defaultData map[string]string) (isChanged bool, err error) {
	return cio.ReadProperties("", p.deps.Input, p.deps.Output, def, data, defaultData)
}

// WriteDataToFS write secrets data to fs file
func (p *Secrets) WriteDataToFS(fs filesystem.Filespace, data map[string]string) (err error) {
	var json string
	if json, err = plainmap.PlainStringMapToFormattedJSON(data); err != nil {
		return err
	}
	return fs.WriteFile(SecretsDataPath, []byte(json), 0766)
}
