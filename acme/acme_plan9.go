package acme

import "9fans.net/go/plan9/client"

// mountAcme sets the default Fsys for Plan 9.
// On Plan 9 acme's filesystem is already in the namespace at /mnt/acme;
// no dial is needed and the connection never breaks.
func mountAcme() {
	defaultFsys = &Fsys{fs: &client.Fsys{Mtpt: "/mnt/acme"}}
}

// Mount returns the default Fsys on Plan 9.
// It always succeeds: acme is always mounted in the namespace.
func Mount() (*Fsys, error) {
	defaultOnce.Do(mountAcme)
	return defaultFsys, nil
}
