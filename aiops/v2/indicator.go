package v2

import (
	"bytes"
	"encoding/json"
	"text/template"
)

type IndicatorItem struct {
	AlertEnabled bool              `json:"alertEnabled"`
	Annotations  map[string]string `json:"annotations"`
	DisplayName  string            `json:"displayName"`
	Kind         string            `json:"kind"`
	Name         string            `json:"name"`
	Query        string            `json:"query"`
	Summary      string            `json:"summary"`
	Type         string            `json:"type"`
	Unit         string            `json:"unit"`
	Legend       string            `json:"legend"`
	Variables    []string          `json:"variables"`
}

func (i *IndicatorItem) Clone() *IndicatorItem {
	result := &IndicatorItem{}
	data, _ := json.Marshal(i)
	_ = json.Unmarshal(data, result)
	return result
}

func (i *IndicatorItem) TemplatedQuery(data interface{}) (string, error) {
	t, err := template.New("query").Parse(i.Query)
	if err != nil {
		return "", err
	}
	expr := bytes.NewBufferString("")
	if err = t.Execute(expr, data); err != nil {
		return "", err
	}
	return expr.String(), nil
}
