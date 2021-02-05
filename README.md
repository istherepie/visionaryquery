# Visionary Query

A cli tool to query the `Visionary` database.


### Installation

More to come ...


## Getting Started

View the cli options:

```
	./VisionaryQuery

	Usage of ./VisionaryQuery:
		-actor
				Result will include actor
		-character
				Result will include character

		-config string
    			Config file location (unix path) (default: ...)

		-timecode
				Result will include timecode

```

Get actor information:

```
	./VisionaryQuery -actor

	Torkel Fredly

```

Get actor and timecode:

```
	./VisionaryQuery -actor -timecode

	Torkel Fredly,01:06:26:00

```

## Configuration

The cli assumes that the `config.yml` file is located relative to the exectuable. 

However it is possible to set a custom path to the file:

```
	./VisionaryQuery -config /etc/custom/custom.yml

```

This is useful if the config file is either in a seperate location or must have a specific name.



## License

Distributed under the MPL2.0 License. See `LICENSE` for more information.


## Contact

Steffen Park <dev@istherepie.com>

