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
// 客户端
package backend

import (
	pb "github.com/chenquan/hit/internal/remotecache"
)

// Discovery 服务发现
type Discovery interface {
	// 拉取所有节点
	PullAllNodes() ([]string, error)
	// 拉取指定前戳节点
	PullNodes(prefix string) ([]string, error)
	// 关闭
	Close()
	// 获取节点数据
	GetNodes() map[string]string
}
type NodePicker interface {
	PickNode(key string) (node Nodor, ok bool)
}

type NodeGetter interface {
	Get(in *pb.GetRequest, out *pb.GetResponse) error
}
type NodeSetter interface {
	Set(in *pb.SetRequest, out *pb.SetResponse) error
}
type NodeDeler interface {
	Del(in *pb.DelRequest, out *pb.DelResponse) error
}

// 节点
type Nodor interface {
	NodeGetter
	NodeSetter
	NodeDeler
	Url() string
}
