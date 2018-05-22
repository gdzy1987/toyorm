/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package toyorm

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func BenchmarkStandardInsert(b *testing.B) {
	brick := TestDB.Model(&TestBenchmarkTable{})
	createTableUnit(brick)(b)
	b.StartTimer()
	// get insert exec
	result, err := brick.Insert(&TestBenchmarkTable{
		Key:   "key",
		Value: "value",
	})
	if err != nil {
		b.Error(err)
		b.FailNow()
	}
	if result.Err() != nil {
		b.Error(err)
		b.FailNow()
	}
	// insert sql action must be 1
	assert.Equal(b, len(result.ActionFlow), 1)
	exec := result.ActionFlow[0].(ExecAction)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		result, err := TestDB.db.Exec(exec.Exec.Query(), exec.Exec.Args()...)
		result.LastInsertId()
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkInsert(b *testing.B) {
	brick := TestDB.Model(&TestBenchmarkTable{})
	createTableUnit(brick)(b)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		result, err := brick.Insert(&TestBenchmarkTable{
			Key:   "key",
			Value: "value",
		})
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
		if result.Err() != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkFind(b *testing.B) {
	brick := TestDB.Model(&TestBenchmarkTable{})
	createTableUnit(brick)(b)
	// fill some data
	for i := 0; i < 1; i++ {
		result, err := brick.Insert(&TestBenchmarkTable{
			Key:   "key" + strconv.Itoa(i),
			Value: "value",
		})
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
		if result.Err() != nil {
			b.Error(err)
			b.FailNow()
		}
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		var data []TestBenchmarkTable
		result, err := brick.Find(&data)
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
		if result.Err() != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkUpdate(b *testing.B) {
	brick := TestDB.Model(&TestBenchmarkTable{})
	createTableUnit(brick)(b)
	// fill some data
	for i := 0; i < 3; i++ {
		result, err := brick.Insert(&TestBenchmarkTable{
			Key:   "key" + strconv.Itoa(i),
			Value: "value",
		})
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
		if result.Err() != nil {
			b.Error(err)
			b.FailNow()
		}
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		result, err := brick.Update(&TestBenchmarkTable{
			Value: "value" + strconv.Itoa(n),
		})
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
		if result.Err() != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}
