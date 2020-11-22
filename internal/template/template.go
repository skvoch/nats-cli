package template

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/shibukawa/configdir"
)

var templates Templates

type Item struct {
	NatsServer    string
	NatsSubject   string
	NatsClusterID string
}

type Templates struct {
	Values map[string]Item
}

const (
	templateFileName = "templates.json"
)

func init() {
	templates.Values = make(map[string]Item)
}

func Add(name string, item Item) error {
	_, exist := templates.Values[name]
	if exist {
		return fmt.Errorf("failed to add new tempalte: already exist")
	}

	templates.Values[name] = item

	return nil
}

func Get(name string) (Item, error) {
	item, ok := templates.Values[name]
	if !ok {
		return Item{}, fmt.Errorf("failed to get template: doens't exist")
	}

	return item, nil
}

func Remove(name string) error {
	_, ok := templates.Values[name]
	if !ok {
		return fmt.Errorf("failed to remove tempalte: doesn't exist")
	}
	delete(templates.Values, name)

	return nil
}

func List() Templates {
	out := Templates{
		Values: make(map[string]Item),
	}

	for name, item := range templates.Values {
		out.Values[name] = item
	}

	return out
}

func Save() error {
	data, err := json.Marshal(&templates)
	if err != nil {
		return fmt.Errorf("failed to marhsal templates: %w", err)
	}

	folders := getDir().QueryFolders(configdir.Global)
	if err := folders[0].WriteFile(templateFileName, data); err != nil {
		return fmt.Errorf("failed to write templates file: %w", err)
	}

	return nil
}

func Read() error {
	dir := getDir()
	dir.LocalPath, _ = filepath.Abs(".")
	folder := dir.QueryFolderContainsFile(templateFileName)

	if folder != nil {
		data, _ := folder.ReadFile(templateFileName)
		if err := json.Unmarshal(data, &templates); err != nil {
			return fmt.Errorf("failed to unmarhsal template file: %w", err)
		}
	}

	return nil
}

func getDir() configdir.ConfigDir {
	return configdir.New("nats-cli", "nats-cli")
}
