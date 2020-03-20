// Copyright 2015 Light Code Labs, LLC
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

package caddyhttp

import (
	// plug in the server
	_ "github.com/Ranger-X/caddy/caddyhttp/httpserver"

	// plug in the standard directives
	_ "github.com/Ranger-X/caddy/caddyhttp/basicauth"
	_ "github.com/Ranger-X/caddy/caddyhttp/bind"
	_ "github.com/Ranger-X/caddy/caddyhttp/browse"
	_ "github.com/Ranger-X/caddy/caddyhttp/errors"
	_ "github.com/Ranger-X/caddy/caddyhttp/expvar"
	_ "github.com/Ranger-X/caddy/caddyhttp/extensions"
	_ "github.com/Ranger-X/caddy/caddyhttp/fastcgi"
	_ "github.com/Ranger-X/caddy/caddyhttp/gzip"
	_ "github.com/Ranger-X/caddy/caddyhttp/header"
	_ "github.com/Ranger-X/caddy/caddyhttp/index"
	_ "github.com/Ranger-X/caddy/caddyhttp/internalsrv"
	_ "github.com/Ranger-X/caddy/caddyhttp/limits"
	_ "github.com/Ranger-X/caddy/caddyhttp/log"
	_ "github.com/Ranger-X/caddy/caddyhttp/markdown"
	_ "github.com/Ranger-X/caddy/caddyhttp/mime"
	_ "github.com/Ranger-X/caddy/caddyhttp/pprof"
	_ "github.com/Ranger-X/caddy/caddyhttp/proxy"
	_ "github.com/Ranger-X/caddy/caddyhttp/push"
	_ "github.com/Ranger-X/caddy/caddyhttp/redirect"
	_ "github.com/Ranger-X/caddy/caddyhttp/requestid"
	_ "github.com/Ranger-X/caddy/caddyhttp/rewrite"
	_ "github.com/Ranger-X/caddy/caddyhttp/root"
	_ "github.com/Ranger-X/caddy/caddyhttp/status"
	_ "github.com/Ranger-X/caddy/caddyhttp/templates"
	_ "github.com/Ranger-X/caddy/caddyhttp/timeouts"
	_ "github.com/Ranger-X/caddy/caddyhttp/websocket"
	_ "github.com/Ranger-X/caddy/onevent"
)
