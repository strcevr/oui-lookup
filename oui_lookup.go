package oui_lookup

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

type OuiDb struct {
	items map[string]string
}

func New(file string) *OuiDb {
	db := &OuiDb{}
	if err := db.Load(file); err != nil {
		return nil
	}
	return db
}

func (db *OuiDb) VendorLookup(s string) (string, error) {

	if vendor, found := db.items[s[:6]]; found {
		return vendor, nil
	} else {
		return "", errors.New("not found")
	}
}

func (db *OuiDb) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`^(\w+)\s{2,}(.*?)\s{2,}(.*?)$`)
	m := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		if line == "" {
			continue
		}
		sub := re.FindStringSubmatch(line)
		if len(sub) > 3 {
			m[strings.ToLower(sub[1])] = sub[3]
		}

	}

	if err := scanner.Err(); err != nil {
		return err
	}

	db.items = m

	return nil
}
