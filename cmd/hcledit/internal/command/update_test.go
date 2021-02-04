package command

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestRunUpdate(t *testing.T) {
	filename := tempFile(t, `
resource "google_container_node_pool" "nodes1" {
  node_config {
    preemptible  = false
    machine_type = "e2-medium"
  }
}
`)

	args := []string{
		"resource.google_container_node_pool.*.node_config.machine_type",
		"e2-highmem-2",
		filename,
	}
	if err := runUpdate(args); err != nil {
		t.Fatal(err)
	}

	got, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	want := `
resource "google_container_node_pool" "nodes1" {
  node_config {
    preemptible  = false
    machine_type = "e2-highmem-2"
  }
}
`
	if string(got) != want {
		t.Fatalf("\ngot  %s\nwant %s", got, want)
	}
}
