# huescene

For use with Philips Hue lights.

Define and orchestrate lighting scenes from the commandline.

## Installing it

Once you have [Go](https://golang.org/) installed:

```shell
$ go get https://github.com/bruce/huescene
```

## Configuring it

Huescene configuration are defined using [YAML](https://yaml.org/).

Here's a simple example:

```yaml
scenes:
- name: working
  lights:
  - name: Office
    color: "#fab444"
    brightness: 255
  - name: Desk
    color: "#fff"
    brightness: 200
```

More examples can be found in the [`examples/`](examples) directory.

### Configuration schema

The configuration should have a `scenes` key containing a list of scene configurations.

Scene configurations have a `name` and, optionally, the following fields that act as defaults for the associated lights:

- `power` - `true` or `false`
- `color` - A hexidecimal color string, e.g., `#fab444`
- `brightness` - `0`–`255`

A scene's lights are defines as a list under its `lights` key.

Light configurations have a `name` and can override the `power`, `color`, and `brightness` settings that are set as defaults in their scenes.

## Running it

All commands must be run from within the same network as your Philips Hue Bridge.

### Authentication

TODO

### Setting a scene

To set a specific scene, simply provide its name as an argument. For instance, if our config file in `path/to/config.yml` had a `work` scene defined, this would activate it:

```shell
$ huescene -c path/to/config.yml work
```

Note that a `--key` option must be provided for authentication reasons unless your configuration provides it (in a toplevel `key` field).

You can also pipe the configuration into huescene:

```shell
$ cat path/to/config.yml | huescene work
```

## License

See [`LICENSE`](LICENSE).
