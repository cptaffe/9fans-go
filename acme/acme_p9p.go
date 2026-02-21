// +build !plan9

package acme

import "9fans.net/go/plan9/client"

// mountAcme is called once by defaultOnce to set up the default Fsys.
func mountAcme() {
	fs, err := client.MountService("acme")
	if err != nil {
		defaultErr = err
		return
	}
	defaultFsys = &Fsys{fs: fs}
}

// Mount opens a fresh connection to acme and returns it as an Fsys.
// Each call dials independently; retry with your own backoff on failure.
// Use the package-level functions for simple tools that do not need
// reconnect semantics.
func Mount() (*Fsys, error) {
	fs, err := client.MountService("acme")
	if err != nil {
		return nil, err
	}
	return &Fsys{fs: fs}, nil
}
