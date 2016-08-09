package manifest_test

import (
	"fmt"
	"net"
	"testing"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/convox/rack/manifest"
	"github.com/stretchr/testify/assert"
)

// WARNING: make sure to use spaces for the yaml indentations
func TestLoadVersion1(t *testing.T) {
	m, err := manifestFixture("v1")

	if assert.Nil(t, err) {
		assert.Equal(t, m.Version, "1")
		assert.Equal(t, len(m.Services), 1)

		if web := m.Services["web"]; assert.NotNil(t, web) {
			assert.Equal(t, web.Image, "test")
		}
	}
}

func TestLoadVersion2(t *testing.T) {
	m, err := manifestFixture("v2-number")

	if assert.Nil(t, err) {
		assert.Equal(t, m.Version, "2")
		assert.Equal(t, len(m.Services), 1)

		if web := m.Services["web"]; assert.NotNil(t, web) {
			assert.Equal(t, web.Image, "test")
		}
	}

	m, err = manifestFixture("v2-string")

	if assert.Nil(t, err) {
		assert.Equal(t, m.Version, "2")
		assert.Equal(t, len(m.Services), 1)

		if web := m.Services["web"]; assert.NotNil(t, web) {
			assert.Equal(t, web.Image, "test")
		}
	}
}

func TestLoadCommandString(t *testing.T) {
	m, err := manifestFixture("command-string")

	if assert.Nil(t, err) {
		if web := m.Services["web"]; assert.NotNil(t, web) {
			assert.Equal(t, web.Command.String, manifest.Command{String: "ls -la"}.String)
		}
	}
}

func TestLoadCommandArray(t *testing.T) {
	m, err := manifestFixture("command-array")

	if assert.Nil(t, err) {
		if web := m.Services["web"]; assert.NotNil(t, web) {
			assert.Equal(t, web.Command.Array, manifest.Command{Array: []string{"ls", "-la"}}.Array)
		}
	}
}

func TestLoadFullVersion1(t *testing.T) {
	m, err := manifestFixture("full-v1")

	if assert.Nil(t, err) {
		if web := m.Services["web"]; assert.NotNil(t, web) {
			assert.Equal(t, web.Build.Context, ".")
			assert.Equal(t, web.Command.String, manifest.Command{String: "bin/web"}.String)
			assert.Equal(t, web.Dockerfile, "Dockerfile.dev")
			assert.Equal(t, web.Entrypoint, "/sbin/init")
			assert.Equal(t, len(web.Environment), 2)
			assert.Equal(t, web.Environment["FOO"], "bar")
			assert.Equal(t, web.Environment["BAZ"], "")
			assert.Equal(t, len(web.Labels), 2)
			assert.Equal(t, web.Labels["convox.foo"], "bar")
			assert.Equal(t, web.Labels["convox.baz"], "4")
			assert.Equal(t, web.Privileged, true)

			if assert.Equal(t, len(web.Links), 1) {
				assert.Equal(t, web.Links[0], "database")
			}

			if assert.Equal(t, len(web.Ports), 2) {
				assert.True(t, web.Ports.External())
				assert.True(t, web.Ports[0].External())
				assert.Equal(t, web.Ports[0].Balancer, 80)
				assert.Equal(t, web.Ports[0].Container, 5000)
				assert.True(t, web.Ports[1].External())
				assert.Equal(t, web.Ports[1].Balancer, 443)
				assert.Equal(t, web.Ports[1].Container, 5001)
			}

			if assert.Equal(t, len(web.Volumes), 1) {
				assert.Equal(t, web.Volumes[0], "/var/db")
			}
		}

		if db := m.Services["database"]; assert.NotNil(t, db) {
			assert.Equal(t, len(db.Environment), 2)
			assert.Equal(t, db.Environment["FOO"], "bar")
			assert.Equal(t, db.Environment["BAZ"], "qux")
			assert.Equal(t, db.Image, "convox/postgres")
			assert.Equal(t, len(db.Labels), 2)
			assert.Equal(t, db.Labels["convox.aaa"], "4")
			assert.Equal(t, db.Labels["convox.ccc"], "ddd")

			if assert.Equal(t, len(db.Ports), 1) {
				assert.False(t, db.Ports.External())
				assert.False(t, db.Ports[0].External())
				assert.Equal(t, db.Ports[0].Balancer, 5432)
				assert.Equal(t, db.Ports[0].Container, 5432)
			}
		}
	}
}

func TestLoadFullVersion2(t *testing.T) {
	m, err := manifestFixture("full-v2")

	if assert.Nil(t, err) {
		if web := m.Services["web"]; assert.NotNil(t, web) {
			assert.Equal(t, web.Build.Context, ".")
			assert.Equal(t, web.Command.String, manifest.Command{String: "bin/web"}.String)
			assert.Equal(t, web.Dockerfile, "Dockerfile.dev")
			assert.Equal(t, web.Entrypoint, "/sbin/init")
			assert.Equal(t, len(web.Environment), 2)
			assert.Equal(t, web.Environment["FOO"], "bar")
			assert.Equal(t, web.Environment["BAZ"], "")
			assert.Equal(t, len(web.Labels), 2)
			assert.Equal(t, web.Labels["convox.foo"], "bar")
			assert.Equal(t, web.Labels["convox.baz"], "4")
			assert.Equal(t, web.Privileged, true)

			if assert.Equal(t, len(web.Links), 1) {
				assert.Equal(t, web.Links[0], "database")
			}

			if assert.Equal(t, len(web.Ports), 2) {
				assert.True(t, web.Ports.External())
				assert.True(t, web.Ports[0].External())
				assert.Equal(t, web.Ports[0].Balancer, 80)
				assert.Equal(t, web.Ports[0].Container, 5000)
				assert.True(t, web.Ports[0].External())
				assert.Equal(t, web.Ports[1].Balancer, 443)
				assert.Equal(t, web.Ports[1].Container, 5001)
			}

			if assert.Equal(t, len(web.Volumes), 1) {
				assert.Equal(t, web.Volumes[0], "/var/db")
			}
		}

		if db := m.Services["database"]; assert.NotNil(t, db) {
			assert.Equal(t, len(db.Environment), 2)
			assert.Equal(t, db.Environment["FOO"], "bar")
			assert.Equal(t, db.Environment["BAZ"], "qux")
			assert.Equal(t, db.Image, "convox/postgres")
			assert.Equal(t, len(db.Labels), 2)
			assert.Equal(t, db.Labels["convox.aaa"], "4")
			assert.Equal(t, db.Labels["convox.ccc"], "ddd")

			if assert.Equal(t, len(db.Ports), 1) {
				assert.False(t, db.Ports.External())
				assert.False(t, db.Ports[0].External())
				assert.Equal(t, db.Ports[0].Balancer, 5432)
				assert.Equal(t, db.Ports[0].Container, 5432)
			}
		}
	}
}

func TestLoadGarbage(t *testing.T) {
	m, err := manifest.Load([]byte("\t\003//783bfkl1f"))

	if assert.Nil(t, m) && assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), "could not parse manifest: yaml: control characters are not allowed")
	}
}

func TestLoadIdleTimeoutUnset(t *testing.T) {
	m, err := manifestFixture("idle-timeout-unset")

	if assert.Nil(t, err) {
		if assert.Equal(t, 1, len(m.Balancers())) {
			b := m.Balancers()[0]
			if val, err := b.IdleTimeout(); assert.Nil(t, err) {
				assert.Equal(t, val, "3600")
			}
		}
	}
}

func TestLoadIdleTimeoutSet(t *testing.T) {
	m, err := manifestFixture("idle-timeout-set")

	if assert.Nil(t, err) {
		if assert.Equal(t, 1, len(m.Balancers())) {
			b := m.Balancers()[0]
			if val, err := b.IdleTimeout(); assert.Nil(t, err) {
				assert.Equal(t, val, "99")
			}
		}
	}
}

func TestLoadBadVersion1(t *testing.T) {
	m, err := manifestFixture("bad-v1")

	if assert.Nil(t, m) && assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), "error loading manifest: yaml: unmarshal errors:\n  line 3: cannot unmarshal !!map into []string")
	}
}

func TestLoadBadVersion2(t *testing.T) {
	m, err := manifestFixture("bad-v2")

	if assert.Nil(t, m) && assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), "error loading manifest: yaml: unmarshal errors:\n  line 5: cannot unmarshal !!map into []string")
	}
}

func TestLoadNonexistentFile(t *testing.T) {
	m, err := manifest.LoadFile("/foo/bar/hope/this/doesnt/exist")

	if assert.Nil(t, m) && assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), "open /foo/bar/hope/this/doesnt/exist: no such file or directory")
	}
}

func TestLoadUnknownVersion(t *testing.T) {
	m, err := manifestFixture("unknown-version")

	if assert.Nil(t, m) && assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), "unknown manifest version: 3")
	}
}

func TestExternalPorts(t *testing.T) {
	m, err := manifestFixture("full-v1")

	if assert.Nil(t, err) {
		ports := m.ExternalPorts()

		if assert.Equal(t, len(ports), 2) {
			assert.Equal(t, ports[0], 80)
			assert.Equal(t, ports[1], 443)
		}
	}
}

func TestPortConflictsWithoutConflict(t *testing.T) {
	m, err := manifestFixture("port-conflicts")

	if assert.Nil(t, err) {
		pc, err := m.PortConflicts()

		if assert.Nil(t, err) {
			assert.Equal(t, len(pc), 0)
		}
	}
}

func TestPortConflictsWithConflict(t *testing.T) {
	m, err := manifestFixture("port-conflicts")

	if assert.Nil(t, err) {
		l, err := net.Listen("tcp", "127.0.0.1:30544")

		defer l.Close()

		ch := make(chan error)

		go func() {
			for {
				_, err := l.Accept()
				ch <- err
			}
		}()

		if assert.Nil(t, err) {
			pc, err := m.PortConflicts()

			if assert.Nil(t, err) && assert.Equal(t, len(pc), 1) {
				assert.Equal(t, pc[0], 30544)
			}
		}

		select {
		case <-time.After(200 * time.Millisecond):
			assert.Fail(t, "nothing connected to the server")
		case <-ch:
		}
	}
}

func TestManifestNetworks(t *testing.T) {
	m, err := manifestFixture("networks")
	if assert.Nil(t, err) {
		for _, s := range m.Services {
			assert.Equal(t, s.Networks, manifest.Networks{
				"foo": manifest.InternalNetwork{
					"external": manifest.ExternalNetwork{
						Name: "foo",
					},
				},
			})
		}
	}
}

func TestShift(t *testing.T) {
	m, err := manifestFixture("shift")

	if assert.Nil(t, err) {
		m.Shift(5000)

		web := m.Services["web"]

		if assert.NotNil(t, web) && assert.Equal(t, len(web.Ports), 2) {
			assert.Equal(t, web.Ports[0].Balancer, 5000)
			assert.Equal(t, web.Ports[0].Container, 5000)
			assert.Equal(t, web.Ports[1].Balancer, 11000)
			assert.Equal(t, web.Ports[1].Container, 7000)
		}
	}
}

func TestManifestMarshalYaml(t *testing.T) {

	strCmd := manifest.Command{
		String: "bin/web",
	}

	arrayCmd := manifest.Command{
		Array: []string{"sh", "-c", "bin/web"},
	}

	m := manifest.Manifest{
		Version: "1",
		Services: map[string]manifest.Service{
			"food": manifest.Service{
				Name: "food",
				Build: manifest.Build{
					Context:    ".",
					Dockerfile: "Dockerfile",
				},
				Command: strCmd,
				Ports: manifest.Ports{
					manifest.Port{
						Public:    true,
						Balancer:  10,
						Container: 10,
					},
				},
			},
		},
	}

	byts, err := yaml.Marshal(m)
	if err != nil {
		t.Error(err.Error())
	}

	m2, err := manifest.Load(byts)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, m2.Version, "2")
	assert.Equal(t, m2.Services["food"].Name, "food")
	assert.Equal(t, m2.Services["food"].Command.String, strCmd.String)

	// Test an array Command
	food := m.Services["food"]
	food.Command = arrayCmd
	m.Services["food"] = food

	byts, err = yaml.Marshal(m)
	if err != nil {
		t.Error(err.Error())
	}

	m2, err = manifest.Load(byts)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, m2.Version, "2")
	assert.Equal(t, m2.Services["food"].Name, "food")
	assert.Equal(t, m2.Services["food"].Command.Array, arrayCmd.Array)
}

func TestManifestValidate(t *testing.T) {
	_, cerr := manifestFixture("invalid-cron")
	if assert.NotNil(t, cerr) {
		assert.Equal(t, cerr.Error(), "Cron task my_job is not valid (cron names can contain only alphanumeric characters and dashes and must be between 4 and 30 characters)")
	}

	_, lerr := manifestFixture("invalid-link")
	if assert.NotNil(t, lerr) {
		assert.Equal(t, lerr.Error(), "web links to service: database2 which does not exist")
	}

	_, lperr := manifestFixture("invalid-link-no-ports")
	if assert.NotNil(t, lperr) {
		assert.Equal(t, lperr.Error(), "web links to service: database which does not expose any ports")
	}

	_, herr := manifestFixture("invalid-health-timeout")
	if assert.NotNil(t, herr) {
		assert.Equal(t, herr.Error(), "convox.health.timeout is invalid for web, must be a number between 0 and 60")
	}
}

func manifestFixture(name string) (*manifest.Manifest, error) {
	return manifest.LoadFile(fmt.Sprintf("fixtures/%s.yml", name))
}
