// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// +build e2e atlas,search

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"text/template"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestSearch(t *testing.T) {
	clusterName, err := deployCluster()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	defer func() {
		if e := deleteCluster(clusterName); e != nil {
			t.Errorf("error deleting test cluster: %v", e)
		}
	}()

	cliPath, err := e2e.Bin()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	n, err := e2e.RandInt(1000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	indexName := fmt.Sprintf("index-%v", n)
	collectionName := fmt.Sprintf("collection-%v", n)
	var indexID string

	t.Run("Create via file", func(t *testing.T) {
		fileName := fmt.Sprintf("create_index_search_test-%v.json", n)

		file, err := os.Create(fileName)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer func() {
			if e := os.Remove(fileName); e != nil {
				t.Errorf("error deleting file '%v': %v", fileName, e)
			}
		}()

		tpl := template.Must(template.New("").Parse(`
{
	"collectionName": "{{ .collectionName }}",
	"database": "test",
	"name": "{{ .indexName }}",
	"mappings": {
		"dynamic": true
	}
}`))
		err = tpl.Execute(file, map[string]string{
			"collectionName": collectionName,
			"indexName":      indexName,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			searchEntity,
			indexEntity,
			"create",
			"--clusterName="+clusterName,
			"--file",
			fileName,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		var index mongodbatlas.SearchIndex
		if err := json.Unmarshal(resp, &index); assert.NoError(t, err) {
			assert.Equal(t, index.Name, indexName)
			indexID = index.IndexID
		}
	})

	t.Run("list", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			searchEntity,
			indexEntity,
			"list",
			"--clusterName="+clusterName,
			"--db=test",
			"--collection="+collectionName,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var indexes []mongodbatlas.SearchIndex
		if err := json.Unmarshal(resp, &indexes); assert.NoError(t, err) {
			assert.NotEmpty(t, indexes)
		}
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			searchEntity,
			indexEntity,
			"describe",
			indexID,
			"--clusterName="+clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var index mongodbatlas.SearchIndex
		if err := json.Unmarshal(resp, &index); assert.NoError(t, err) {
			assert.Equal(t, indexID, index.IndexID)
		}
	})

	t.Run("Update via file", func(t *testing.T) {
		analyzer := "lucene.simple"
		fileName := fmt.Sprintf("update_index_search_test-%v.json", n)

		file, err := os.Create(fileName)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer func() {
			if e := os.Remove(fileName); e != nil {
				t.Errorf("error deleting file '%v': %v", fileName, e)
			}
		}()

		tpl := template.Must(template.New("").Parse(`
{
	"collectionName": "{{ .collectionName }}",
	"database": "test",
	"name": "{{ .indexName }}",
	"analyzer": "{{ .analyzer }}",
	"mappings": {
		"dynamic": true
	}
}`))
		err = tpl.Execute(file, map[string]string{
			"collectionName": collectionName,
			"indexName":      indexName,
			"analyzer":       analyzer,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			searchEntity,
			indexEntity,
			"update",
			indexID,
			"--clusterName="+clusterName,
			"--file",
			fileName,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		var index mongodbatlas.SearchIndex
		if err := json.Unmarshal(resp, &index); assert.NoError(t, err) {
			a := assert.New(t)
			a.Equal(indexID, index.IndexID)
			a.Equal(analyzer, index.Analyzer)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			searchEntity,
			indexEntity,
			"delete",
			indexID,
			"--clusterName="+clusterName,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		expected := fmt.Sprintf("Index '%s' deleted\n", indexID)
		assert.Equal(t, string(resp), expected)
	})
}
