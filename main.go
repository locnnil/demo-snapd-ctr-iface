package main

import (
	"fmt"
	"os"

	"github.com/snapcore/snapd/client"
	"github.com/snapcore/snapd/dirs"
)

var clientConfig = client.Config{
	// snapctl should not try to read $HOME/.snap/auth.json, this will
	// result in apparmor denials and configure task failures
	// (LP: #1660941)
	DisableAuth: true,

	// we need the less privileged snap socket in snapctl
	Socket: dirs.SnapSocket,
}

func main() {
	cli := client.New(&clientConfig)
	cookie := os.Getenv("SNAP_CONTEXT")

	fmt.Println("cookie: ", cookie)
	fmt.Println("cli: ", cli)

	patchValues := make(map[string]interface{})

	patchValues = map[string]interface{}{
		"system": map[string]interface{}{
			"kernel": map[string]interface{}{
				"dangerous-cmdline-append": "foo=bar",
			},
		},
	}

	// not sure if this actually is being effective on passing the SNAP_CONTEXT
	cli.RunSnapctl(&client.SnapCtlOptions{
		ContextID: cookie,
		Args:      nil,
	}, nil)

	// NOTE: This currently gives me the "access denied" error
	out, err := cli.SetConf("system", patchValues)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("out: ", out)
}
