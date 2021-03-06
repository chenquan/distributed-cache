/*
 *    Copyright 2020 Chen Quan
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package main

import (
	"fmt"
	"github.com/chenquan/hit/internal/register"
	"github.com/chenquan/hit/internal/server"
	"net/http"
	"testing"
)

type Test interface {
	Name() string
	SetName(string)
}
type Atest struct {
	value string
}

func (t *Atest) Name() string {
	return t.value
}
func (t *Atest) SetName(value string) {
	t.value = value
}

func Test1(t *testing.T) {
	config := handleConfig("hit.toml")
	addr := fmt.Sprintf("%s://%s:%s", config.Protocol, config.NodeAddr, config.Port)
	_ = register.New(config).RegisterNode(config.NodeName, addr)
	httpPool := server.NewHTTPPool()
	_ = http.ListenAndServe(":"+config.Port, httpPool)

}
