/*
 *
 * Copyright (C)   ISCAS
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 * File Created: 2024/12/17
 * Author: suyao (suyao2023@iscas.ac.cn)
 */

package v1

import (
	"testing"
	"time"
)

// TestNow 测试 Now 方法和 MarshalQueryParameter 方法
func TestNow(t *testing.T) {
	// 获取当前时间
	got := Now()

	// 获取现在的时间并格式化为查询参数格式（RFC3339）
	want := time.Now().Format(time.RFC3339)

	// 使用 MarshalQueryParameter 获取 got 的查询参数格式
	gotQueryParam, _ := got.MarshalQueryParameter()

	// 比较两者是否相等
	if gotQueryParam != want {
		t.Errorf("Now().MarshalQueryParameter() = %v; want %v", gotQueryParam, want)
	}
}
