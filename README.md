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
- name: work
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

Huescene _currently_ automatically attempts to create a username (API key) using your Hue Bridge if a key isn't already provided.

At the moment, the workflow is a but clumsy :sob: and looks like this:

1. You try one of the commands below, but you don't provide a key. It fails with a noisy `"link button not pressed"` error.
2. You press the Hue Bridge link button and (within 30 seconds, reattempt the command). Now (instead of the command you really want to run) you get the API key output. A somewhat qualified success!
3. You optionally store the API key in your configuration file (under the `key` toplevel field) so you don't have to provide it as `--key` in future commands.
4. You run the commands you want (see below) passing the key in your configuration file or as `--key`.

:construction_worker_man: In the future, [there will be](https://github.com/bruce/huescene/issues/1) a separate flag to trigger authentication more explicitly. You know, like you'd expect!

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

### Listing lights

Obviously to configure lights, you need to know what they're named. Use the `--list-lights` flag.

If you're storing your API key in your config file (`key` field), you can provide the configuration as usual with `-c` or `--config`,

```shell
$ huescene -c path/to/config.yml --list-lights
```

You can also just pass the key directory:

```shell
$ huescene --key YOUR_API_KEY --list-lights
```

## License

See [`LICENSE`](LICENSE).
