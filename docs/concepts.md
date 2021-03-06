# Key Concepts

In order to make your interaction with `ralph-cli` easier, we need to establish
some common vocabulary - hence the following section, where key concepts behind
`ralph-cli` are explained.

## Config

`ralph-cli` has its own configuration file `~/.ralph-cli/config.toml`, in
[TOML][] format. This file is created automatically at the first run of
`ralph-cli` (e.g. `ralph-cli --help`), and it contains some default settings, as
presented below:

```no-highlight
RalphAPIURL = "change_me"
RalphAPIKey = "change_me"
ManagementUserName = "change_me"
ManagementUserPassword = "change_me"
```

* `RalphAPIURL` - this should be a string with an URL to your Ralph instance,
  with `/api` path appended, e.g. `"https://my-ralph-instance.local/api"`
* `RalphAPIKey` - your personal API key that `ralph-cli` should use for
  accessing Ralph; you can find this key by visiting Ralph's GUI (go to the
  top-right corner, click on your user name and then select "My profile" - in
  "Personal info" section you should see a field named "API Token" - that's it).
* `ManagementUserName` - user name that is used to access iDRAC/iLO, which
  `ralph-cli` exposes to scan scripts as `MANAGEMENT_USER_NAME` environment
  variable (see [Scripts Contract][self-contract])
* `ManagementUserPassword` - password for the above, exposed as
  `MANAGEMENT_USER_PASSWORD` environment variable

All of them are required, so remember to replace `change_me` strings with the
real values.

Please note that due to the presence of credentials in `config.toml`, this file
should remain readable only to its owner (it is `0600` by default, so you don't
have to do anything) - otherwise `ralph-cli` will refuse to cooperate.

Also, keep in mind that this is just an initial structure of `ralph-cli` config
file, and it is subject to change until version `1.0.0` is reached.

## Scan

Scan is one of the commands available via `ralph-cli` (well, at this moment,
this is the only command available, but we are going to add more - see section
[Ideas for Future Development][ideas]). The idea behind this is simple:

1. Access a host given as an IP address (via iDRAC, iLO, Puppet, SSH or whatever
   method you'd find useful).
2. Gather some info regarding its configuration (hardware components, software
   etc. - see [Scripts Contract][self-contract]).
3. Process it in some way (e.g. find a difference between what has been
   discovered on this host and what is stored in Ralph, and mark it for
   update/deletion).
4. Send it to Ralph.

The first two steps are handled by scan scripts (see next section), the last two
are handled exclusively by `ralph-cli`, freeing you from the extra work
associated with communication with Ralph.

## Scan scripts

Scripts are the meat of the `scan` command (see previous section). You may think
of them as "plugins", although they are called "scripts" intentionally, to
encourage modifications. Scan scripts can be written in *any* language you want,
as long as they conform to [Scripts Contract][self-contract], and you have means
for running them from your host (e.g. access to interpreter, required libraries
etc.). They should be placed in `~/.ralph-cli/scripts` directory, with an
optional manifest file (see next section).

If you write your scripts in Python, you may need so-called
["virtual environments"][virtualenv] for them - but you don't have to create
and/or activate them by hand, `ralph-cli` will do that for you, as long as you
provide a manifest file listing packages that your script requires. If your Python
script doesn't require anything, then fine, no manifest (and hence no virtualenv)
is needed.

`ralph-cli` comes with two default scripts - `idrac.py` and `ilo.py` - which
should give you an idea what scan scripts should do. You are encouraged to
experiment with them (e.g. by modifying them in-place) - don't worry if you
break them or something - when you delete one of those default scripts (or their
manifests or virtualenvs), `ralph-cli` will restore them to the default
state. Be aware, though, that this behavior doesn't apply to your own scripts
(i.e., with different names that `idrac.py` and `ilo.py`).

## Manifests

These are the files in [TOML][] format that contain some meta data needed for
launching scan scripts. At this moment, they are used exclusively for Python
scripts, and their structure will probably change a lot in the future. Here's an
example of such file, `idrac.toml`:

```no-highlight
Language = "python"
LanguageVersion = 3

[[requirement]]
name = "requests"
version = "2.10.0"
```

Its name (`idrac.toml`) states that this manifest is for `idrac.py` script, and
the fields visible above are rather self-explanatory - `idrac.py` requires
Python 3 and `requests` library in version `2.10.0` (if this field is omitted,
then the newest available version will be used). If you need to specify more
requirements, just add them as another `[[requirement]]` entry, e.g.:

```no-highlight
Language = "python"
LanguageVersion = 3

[[requirement]]
name = "requests"
version = "2.10.0"

[[requirement]]
name = "pyaml"
version = "15.8.2"
```

Manifest files are optional (yet), assuming that your Python script doesn't need
any extra packages, apart from the ones provided by the standard library.

## Scripts Contract

As mentioned in the [Going further][quickstart-further] section, scan scripts
can be written in any language you want, as long as you have means for launching
them. But in order to tame this diversity, they need some well-defined way to
communicate with `ralph-cli` binary. And here comes our Scripts Contract, which
describes exactly this.

### Input

Each script should receive its input parameters only from environment
variables. At this moment, there are three of them:

* `IP_TO_SCAN` - an IP address of a host that we want to scan
* `MANAGEMENT_USER_NAME` - user name that is used to access iDRAC/iLO
* `MANAGEMENT_USER_PASSWORD` - password for the above

`ralph-cli` executes scan scripts as sub-processes, with their own set of
environment variables, so they won't be visible in the parent shell (i.e., the
one from which `ralph-cli` is started).

### Output

Each script should print discovered data on its `stdout`, in the form of a
stringified JSON, pretty-printed for your convenience in the example presented
below:

```no-highlight
{
    "model_name": "Dell PowerEdge R620",
    "firmware_version": "1.1.1",
    "bios_version": "2.2.2",
    "processors": [
        {
            "model_name": "Intel(R) Xeon(R) CPU E5-2650 v2 @ 2.60GHz",
            "speed": 2600, // in Mhz
            "cores": 8
        }
    ],
    "fibre_channel_cards": [
        {
            "firmware_version": "1.1.1",
            "model_name": "Saturn-X: LightPulse Fibre Channel Host Adapter",
            "speed": "4 Gbit", // 1, 2, 4, 8, 16, 32 and "unknown speed" values are valid
            "wwn": "aabbccddeeff0011"
        }
    ],
    "ethernets": [
        {
            "mac": "AA:AA:AA:AA:AA:AA",
            "model_name": "Intel(R) Ethernet 10G 4P X520/I350 rNDC",
            "speed": "10 Gbps", // in Mbps/Gbps (e.g. "10 Mbps", "10 Gbps")
            "firmware_version": "1.1.1",
        }
    ],
    "disks": [
        {
            "model_name": "ATA Samsung SSD 840",
            "size": 476, // in GiB
            "serial_number": "S1AXNSAD8000000",
            "slot": "1",
            "firmware_version": "1.1.1"
        },
    ],
    "memory": [
        {
            "model_name": "Samsung DDR3 DIMM",
            "speed": 1600, // in MHz
            "size": 16384 // in MiB
        },
    ]
}
```

As you can see, this structure is quite flat (and we will do our best to keep it
that way), consisting mostly of lists of dicts.

Keep in mind though, that this is just an initial version of contract, which are
subject to heavy changes until `ralph-cli` will reach `1.0.0` version.

If you have any thoughts on this (or if you need to add something here), please
let us know by opening a new issue on [our GitHub profile][issues].


[self-contract]: concepts.md#scripts-contract
[quickstart-further]: quickstart.md#going-further
[ideas]: development.md#ideas-for-future-development

[TOML]: https://github.com/toml-lang/toml
[virtualenv]: https://packaging.python.org/en/latest/installing/#creating-and-using-virtual-environments
[issues]: https://github.com/allegro/ralph-cli/issues
