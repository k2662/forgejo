// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package integration

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"code.gitea.io/gitea/modules/setting"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/routers"

	"github.com/stretchr/testify/assert"
)

func TestNodeinfo(t *testing.T) {
	setting.Federation.Enabled = true
	c = routers.NormalRoutes(context.TODO())
	defer func() {
		setting.Federation.Enabled = false
		c = routers.NormalRoutes(context.TODO())
	}()

	onGiteaRun(t, func(*testing.T, *url.URL) {
		req := NewRequestf(t, "GET", "/api/v1/nodeinfo")
		resp := MakeRequest(t, req, http.StatusOK)
		var nodeinfo api.NodeInfo
		DecodeJSON(t, resp, &nodeinfo)
		assert.True(t, nodeinfo.OpenRegistrations)
		assert.Equal(t, "forgejo", nodeinfo.Software.Name)
		assert.Equal(t, 24, nodeinfo.Usage.Users.Total)
		assert.Equal(t, 17, nodeinfo.Usage.LocalPosts)
		assert.Equal(t, 2, nodeinfo.Usage.LocalComments)
	})
}
