// SiYuan - https://github.com/siyuan-note/siyuan
// Copyright (c) 2020-present, b3log.org
// SPDX-License-Identifier: AGPL-3.0-or-later

package apicontract

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestSettingThemeCompatibility(t *testing.T) {
	for _, entry := range []struct {
		body    string
		modes   []float64
		message string
	}{
		{`{}`, []float64{}, ""},
		{`{"theme":" t ","modes":[0.9,1.9,false,0],"appearanceMode":" dark "}`, []float64{0.9, 1.9}, ""},
		{`{"theme":"t","modes":[-0.9,0,0,1,2,0]}`, []float64{-0.9, 0, 0, 1}, ""},
		{`{"modes":[false,{},null]}`, []float64{}, ""},
		{`{"modes":false,"appearanceMode":false}`, nil, "Field [modes] should be of type [Array]"},
		{`{"theme":"t","modes":[false],"appearanceMode":false}`, nil, "Field [appearanceMode] should be of type [String]"},
	} {
		request, err := SetTheme.Decode(strings.NewReader(entry.body))
		if entry.message != "" {
			if err == nil || err.Error() != entry.message {
				t.Fatalf("theme error: %v", err)
			}
			continue
		}
		if err != nil || !reflect.DeepEqual(request.Modes, entry.modes) {
			t.Fatalf("theme modes: %+v %v", request, err)
		}
	}
}

func TestSettingKeymapOpaqueValues(t *testing.T) {
	request, err := SetKeymap.Decode(strings.NewReader(`{"data":{"custom":{"key":[null,false,1.5,"",{}]},"empty":null}}`))
	if err != nil {
		t.Fatal(err)
	}
	encoded, err := json.Marshal(request.Data)
	if err != nil || string(encoded) != `{"custom":{"key":[null,false,1.5,"",{}]},"empty":null}` {
		t.Fatalf("keymap payload: %s %v", encoded, err)
	}
	for _, body := range []string{`{}`, `{"data":null}`} {
		if _, err = SetKeymap.Decode(strings.NewReader(body)); err == nil {
			t.Fatalf("keymap must reject incomplete data: %s", body)
		}
	}
}
