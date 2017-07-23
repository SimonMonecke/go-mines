package highscore

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"

	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/smonecke/go-mines/game"
)

type Entries []Entry

func (entries Entries) Len() int {
	return len(entries)
}

func (entries Entries) Less(i, j int) bool {
	if entries[i].DurationInSeconds == entries[j].DurationInSeconds {
		return entries[i].Date > entries[j].Date
	}
	return entries[i].DurationInSeconds < entries[j].DurationInSeconds
}

func (entries Entries) Swap(i, j int) {
	entries[i], entries[j] = entries[j], entries[i]
}

type Entry struct {
	Date              string
	DurationInSeconds int
}

type List struct {
	Easy   Entries `json:",omitempty"`
	Normal Entries `json:",omitempty"`
	Hard   Entries `json:",omitempty"`
}

func LoadList() (List, error) {
	highscoreListPath := path.Join(getHomeDir(), ".go-mines", "highscores.json")
	if _, err := os.Stat(highscoreListPath); os.IsNotExist(err) {
		return List{}, nil
	}
	highscoresFile, err := ioutil.ReadFile(highscoreListPath)
	if err != nil {
		return List{}, fmt.Errorf("cannot read highscore file: %s", err)
	}

	var list List
	err = json.Unmarshal(highscoresFile, &list)
	if err != nil {
		return List{}, fmt.Errorf("cannot parse highscores file [%s]: %s", highscoreListPath, err)
	}
	return list, nil
}

func getHomeDir() string {
	homedir, err := homedir.Dir()
	if err != nil {
		log.Fatal("Cannot determine homedir")
	}
	return homedir
}

func AddEntry(mode game.Mode, date string, durationInSeconds int) error {
	list, err := LoadList()
	if err != nil {
		return err
	}
	err = createGoMinesDirectory()
	if err != nil {
		return err
	}

	modifiedList := addToList(list, mode, date, durationInSeconds)

	modifiedListAsJSON := toJSON(modifiedList)
	err = saveList(modifiedListAsJSON)
	if err != nil {
		return err
	}
	return nil
}

func addToList(list List, mode game.Mode, date string, durationInSeconds int) List {
	var sublistToModify *Entries
	switch mode {
	case game.Easy:
		sublistToModify = &list.Easy
	case game.Normal:
		sublistToModify = &list.Normal
	case game.Hard:
		sublistToModify = &list.Hard
	default:
		log.Fatal("unknown mode")
	}
	*sublistToModify = append(*sublistToModify, Entry{Date: date, DurationInSeconds: durationInSeconds})
	sort.Sort(*sublistToModify)
	if len(*sublistToModify) > 5 {
		*sublistToModify = (*sublistToModify)[:5]
	}
	return list
}

func toJSON(list List) []byte {
	listAsJSON, err := json.MarshalIndent(list, "", "    ")
	if err != nil {
		log.Fatalf("cannot marshal highscores: %s", err)
	}
	return listAsJSON
}

func createGoMinesDirectory() error {
	err := os.MkdirAll(path.Join(getHomeDir(), ".go-mines"), 0775)
	if err != nil {
		return fmt.Errorf("cannot create directory ~/.go-mines to save the highscore list: %s", err)
	}
	return nil
}

func saveList(listAsJson []byte) error {
	err := ioutil.WriteFile(path.Join(getHomeDir(), ".go-mines", "highscores.json"), listAsJson, 0644)
	if err != nil {
		return fmt.Errorf("cannot write highscores file: %s", err)
	}
	return nil
}
