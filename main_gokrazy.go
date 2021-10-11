// Copyright 2020 Matt Layher and Michael Stapelberg
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build gokrazy
// +build gokrazy

package main

// filePaths provides hardcoded /perm paths on gokrazy.
func filePaths() (cfg string, hostKey string) {
	return "/perm/consrv/consrv.toml", "/perm/consrv/host_key"
}
