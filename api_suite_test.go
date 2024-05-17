// Copyright (C) 2015-Present Pivotal Software, Inc. All rights reserved.

// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package brokerapi_test

import (
	"embed"
	"io"
	"path"
	"testing"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Suite")
}

//go:embed fixtures/*.json
var fixtures embed.FS

func fixture(name string) string {
	GinkgoHelper()

	fh := must(fixtures.Open(path.Join("fixtures", name)))
	defer fh.Close()
	return string(must(io.ReadAll(fh)))
}

func uniqueID() string {
	return uuid.NewString()
}

func uniqueInstanceID() string {
	return uniqueID()
}

func uniqueBindingID() string {
	return uniqueID()
}
