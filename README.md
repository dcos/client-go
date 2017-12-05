# WIP: DC/OS Api for Go

dcos-go is a Go client library accessing the DC/OS api to enable automation of
DC/OS tasks.

The main purpose if this project is to support the development of a DC/OS terraform provider. But it might also be beneficial for other Project.

## Configuration

dcos-go is compatible with [dcos-cli](https://github.com/dcos/dcos-cli). It can
parse the configuration from a chain of common configuration options.

See: [config package](./dcos/config)

`config.NewConfigFromChain` tries to read `~/.dcos/dcos.toml`, the
`~/.dcos/clusters/<attached cluster>/dcos.toml` or tries to find whatever is in
`$DCOS_CLUSTER`

## Commands / Services
dcos-go will support every command which `dcos-cli` supports.

- Already available:
- Work in progress:
  - job
  - marathon
  - node
  - package
  - service
  - task

## Plugins
TBD. Basically it would be a good idea to provide an interface which can be used by DC/OS packages to hook into this
