// Copyright 2020 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tags

import (
	"reflect"
	"testing"

	"github.com/MasterYang7/xorm/caches"
	"github.com/MasterYang7/xorm/dialects"
	"github.com/MasterYang7/xorm/names"
	"github.com/stretchr/testify/assert"
)

type ParseTableName1 struct{}

type ParseTableName2 struct{}

func (p ParseTableName2) TableName() string {
	return "p_parseTableName"
}

func TestParseTableName(t *testing.T) {
	parser := NewParser(
		"xorm",
		dialects.QueryDialect("mysql"),
		names.SnakeMapper{},
		names.SnakeMapper{},
		caches.NewManager(),
	)
	table, err := parser.Parse(reflect.ValueOf(new(ParseTableName1)))
	assert.NoError(t, err)
	assert.EqualValues(t, "parse_table_name1", table.Name)

	table, err = parser.Parse(reflect.ValueOf(new(ParseTableName2)))
	assert.NoError(t, err)
	assert.EqualValues(t, "p_parseTableName", table.Name)

	table, err = parser.Parse(reflect.ValueOf(ParseTableName2{}))
	assert.NoError(t, err)
	assert.EqualValues(t, "p_parseTableName", table.Name)
}
