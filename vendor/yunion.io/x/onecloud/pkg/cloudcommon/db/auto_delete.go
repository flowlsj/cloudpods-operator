// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by model-api-gen. DO NOT EDIT.

package db

import (
	"context"

	"yunion.io/x/sqlchemy"

	"yunion.io/x/onecloud/pkg/apis"
	"yunion.io/x/onecloud/pkg/mcclient"
)

// +onecloud:model-api-gen
type SAutoDeleteResourceBase struct {
	// 是否跟随资源自动删除
	// example: false
	AutoDelete bool `nullable:"false" default:"false" get:"user" update:"user" json:"auto_delete"`
}

type SAutoDeleteResourceBaseManager struct{}

func (manager *SAutoDeleteResourceBaseManager) ListItemFilter(
	ctx context.Context,
	q *sqlchemy.SQuery,
	userCred mcclient.TokenCredential,
	query apis.AutoDeleteResourceBaseListInput,
) (*sqlchemy.SQuery, error) {
	if query.AutoDelete != nil {
		q = q.Equals("auto_delete", *query.AutoDelete)
	}

	return q, nil
}

func (self *SAutoDeleteResourceBase) SetAutoDelete(model IModel, userCred mcclient.TokenCredential, autoDelete bool) error {
	_, err := Update(model, func() error {
		self.AutoDelete = autoDelete
		return nil
	})
	return err
}
