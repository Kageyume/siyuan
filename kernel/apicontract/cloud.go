// SiYuan - https://github.com/siyuan-note/siyuan
// Copyright (c) 2020-present, b3log.org
// SPDX-License-Identifier: AGPL-3.0-or-later

package apicontract

type CloudReminderRequest struct {
	ID      string `json:"id"`
	Timed   string `json:"timed"`
	Content string `json:"content"`
}

type CloudBackup struct {
	Size    int64  `json:"size"`
	HSize   string `json:"hSize"`
	Updated string `json:"updated"`
	SaveDir string `json:"saveDir"`
}

type CloudSync struct {
	Size      int64  `json:"size"`
	HSize     string `json:"hSize"`
	Updated   string `json:"updated"`
	CloudName string `json:"cloudName"`
	SaveDir   string `json:"saveDir"`
}
