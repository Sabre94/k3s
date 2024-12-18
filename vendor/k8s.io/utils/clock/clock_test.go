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
 * File Created: 2024/12/18
 * Author: suyao (suyao2023@iscas.ac.cn)
 */

package clock

import (
	"testing"
)

func TestRealClockNow(t *testing.T) {
	// 创建一个 RealClock 实例
	realClock := RealClock{}

	// 调用 Now() 方法获取当前时间
	now := realClock.Now()

	// 获取当前时间的 Unix 时间戳（秒和纳秒）
	nowUnix := now.Unix()
	nowUnixNano := now.UnixNano()

	// 验证 now 是否大于零，确保它获取到了一个有效的时间
	if nowUnix <= 0 {
		t.Errorf("Now() returned an invalid time: %v", now)
	}

	// 验证是否返回了精确的时间
	if nowUnixNano <= 0 {
		t.Errorf("Now() returned an invalid time with nanoseconds: %v", now)
	}

	// 打印结果，便于查看当前时间
	t.Logf("Now() returned the current time: %v (Unix: %d, Nano: %d)", now, nowUnix, nowUnixNano)
}
