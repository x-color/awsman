package awsman

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pkg/browser"
	"github.com/xlab/treeprint"
)

type account struct {
	ID    string `json:"id,omitempty"`
	Role  string `json:"role,omitempty"`
	Alias string `json:"alias,omitempty"`
	URL   string `json:"url,omitempty"`
}

var accounts map[string]account

func Add(id, role, alias string) error {
	ac := account{
		ID:    id,
		Role:  role,
		Alias: alias,
	}
	ac.URL = generateURL(ac)
	if ac.Role == "" {
		ac.Role = "<No Role>"
	}

	file, err := dataFile()
	if err != nil {
		return err
	}
	accounts, err := load(file)
	if err != nil {
		return err
	}

	if _, ok := accounts[alias]; ok {
		return fmt.Errorf("alias(%s) already exists", alias)
	}
	accounts[alias] = ac

	return save(accounts, file)
}

func generateURL(ac account) string {
	if ac.Role == "" {
		return fmt.Sprintf("https://%s.signin.aws.amazon.com/", ac.ID)
	}
	return fmt.Sprintf("https://%s.signin.aws.amazon.com/switchrole?roleName=%s", ac.ID, ac.Role)
}

func Remove(alias string) error {
	file, err := dataFile()
	if err != nil {
		return err
	}
	accounts, err := load(file)
	if err != nil {
		return err
	}

	if _, ok := accounts[alias]; !ok {
		return fmt.Errorf("alias(%s) does not exists", alias)
	}
	delete(accounts, alias)

	return save(accounts, file)
}

func TreeView() (string, error) {
	file, err := dataFile()
	if err != nil {
		return "", err
	}
	accounts, err := load(file)
	if err != nil {
		return "", err
	}

	keys := []string{}
	for k := range accounts {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	tree := treeprint.New()
	for _, k := range keys {
		t := tree.AddBranch(accounts[k].Alias)
		t.AddMetaNode("ID  ", accounts[k].ID)
		t.AddMetaNode("Role", accounts[k].Role)
	}

	return strings.Replace(tree.String(), ".", "AWS Accounts", 1), nil
}

func Oneline() (string, error) {
	file, err := dataFile()
	if err != nil {
		return "", err
	}
	accounts, err := load(file)
	if err != nil {
		return "", err
	}

	keys := []string{}
	for k := range accounts {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	list := make([]string, len(keys)+1)
	for i, k := range keys {
		list[i] = fmt.Sprintf("%s (AccountID: %s, Role: %s)", accounts[k].Alias, accounts[k].ID, accounts[k].Role)
	}

	return strings.Join(list, "\n"), nil
}

func HTMLView() (string, error) {
	file, err := dataFile()
	if err != nil {
		return "", err
	}
	accounts, err := load(file)
	if err != nil {
		return "", err
	}

	return generateHTML(accounts)
}

func WebView() error {
	file, err := dataFile()
	if err != nil {
		return err
	}
	accounts, err := load(file)
	if err != nil {
		return err
	}

	text, err := generateHTML(accounts)
	if err != nil {
		return err
	}

	f, err := ioutil.TempFile("", "awsman-html.*.html")
	if err != nil {
		return err
	}

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}

	return browser.OpenFile(f.Name())
}

func generateHTML(accounts map[string]account) (string, error) {
	type wrapAccount struct {
		ID       string
		Accounts []account
	}

	type webData struct {
		Aliases  []account
		Accounts []wrapAccount
	}

	data := webData{
		Aliases:  make([]account, len(accounts)),
		Accounts: []wrapAccount{},
	}

	i := 0
	whereInAccounts := make(map[string]int)
	for _, v := range accounts {
		data.Aliases[i] = v
		i++

		if j, ok := whereInAccounts[v.ID]; ok {
			data.Accounts[j].Accounts = append(data.Accounts[j].Accounts, v)
		} else {
			whereInAccounts[v.ID] = len(data.Accounts)
			data.Accounts = append(data.Accounts, wrapAccount{
				ID:       v.ID,
				Accounts: []account{v},
			})
		}
	}

	sort.Slice(data.Aliases, func(i, j int) bool {
		return data.Aliases[i].Alias < data.Aliases[j].Alias
	})
	sort.Slice(data.Accounts, func(i, j int) bool {
		return data.Accounts[i].ID < data.Accounts[j].ID
	})

	for i := range data.Accounts {
		sort.Slice(data.Accounts[i].Accounts, func(j, k int) bool {
			return data.Accounts[i].Accounts[j].Role < data.Accounts[i].Accounts[k].Role
		})
	}

	temp, err := template.New("webview").Parse(templateText)
	if err != nil {
		return "", err
	}

	out := new(bytes.Buffer)
	err = temp.Execute(out, data)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func SignIn(alias string) error {
	file, err := dataFile()
	if err != nil {
		return err
	}
	accounts, err := load(file)
	if err != nil {
		return err
	}

	if _, ok := accounts[alias]; !ok {
		return fmt.Errorf("alias(%s) does not exists", alias)
	}

	return browser.OpenURL(accounts[alias].URL)
}

func SignInURL(alias string) (string, error) {
	file, err := dataFile()
	if err != nil {
		return "", err
	}
	accounts, err := load(file)
	if err != nil {
		return "", err
	}

	if _, ok := accounts[alias]; !ok {
		return "", fmt.Errorf("alias(%s) does not exists", alias)
	}

	return accounts[alias].URL, nil
}

func save(accounts map[string]account, file string) error {
	bytes, err := json.Marshal(accounts)
	if err != nil {
		return err
	}

	f, err := os.Create(file)
	if os.IsNotExist(err) {
		err = os.Mkdir(filepath.Dir(file), 0755)
		if err != nil {
			return err
		}
		f, err = os.Create(file)
	}
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(bytes)
	return err
}

func load(file string) (map[string]account, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]account{}, nil
		}
		return nil, err
	}

	accounts := make(map[string]account)
	err = json.Unmarshal(bytes, &accounts)
	return accounts, err
}

func dataFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".awsman", "accounts.json"), nil
}
