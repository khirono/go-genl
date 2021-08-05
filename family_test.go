package genl

import (
	"encoding/json"
	"log"
	"syscall"
	"testing"

	"github.com/khirono/go-nl"
)

func TestGetFamily(t *testing.T) {
	conn, err := nl.Open(syscall.NETLINK_GENERIC)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	mux, err := nl.NewMux()
	if err != nil {
		t.Fatal(err)
	}
	defer mux.Close()
	go mux.Serve()

	c := nl.NewClient(conn, mux)

	f, err := GetFamily(c, "gtp5g")
	if err != nil {
		t.Fatal(err)
	}

	j, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("family: %s\n", j)
}

func TestGetFamilyAll(t *testing.T) {
	conn, err := nl.Open(syscall.NETLINK_GENERIC)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	mux, err := nl.NewMux()
	if err != nil {
		t.Fatal(err)
	}
	defer mux.Close()
	go mux.Serve()

	c := nl.NewClient(conn, mux)

	fs, err := GetFamilyAll(c)
	if err != nil {
		t.Fatal(err)
	}

	j, err := json.MarshalIndent(fs, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("family: %s\n", j)
}
