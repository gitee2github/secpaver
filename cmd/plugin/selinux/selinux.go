/*
 * Copyright (c) Huawei Technologies Co., Ltd. 2020-2021. All rights reserved.
 * secPaver is licensed under the Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *     http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
 * PURPOSE.
 * See the Mulan PSL v2 for more details.
 */

/*
Package main is the entry of selinux engine plugin
*/
package main

import (
	"gitee.com/openeuler/secpaver/engine"
	"gitee.com/openeuler/secpaver/engine/selinux"
	"sync"
)

var e *selinux.Engine
var once sync.Once

// GetEngine is used for engine plugin loading
func GetEngine() engine.Engine {
	once.Do(func() {
		e = selinux.NewEngine()
	})

	return e
}
