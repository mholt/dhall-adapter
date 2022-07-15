Dhall Config Adapter for Caddy
==============================

This Caddy plugin lets you configure Caddy with the [Dhall configuration language](https://dhall-lang.org/).

The result of evaluating your Dhall document must be a valid [Caddy JSON config](https://caddyserver.com/docs/json/).

Config adapter name: `dhall`

## Installation

You can select this plugin on Caddy's [download page](https://caddyserver.com/download).

Or, you can plug it in by building Caddy yourself with [xcaddy](https://github.com/caddyserver/xcaddy):

```
$ xcaddy build --with github.com/mholt/dhall-adapter
```

## Usage

Using this config adapter is [the same as any other](https://caddyserver.com/docs/config-adapters#using-config-adapters):

```
$ caddy run --config caddy.dhall --adapter dhall
```

The key is to specify `--adapter dhall` on the command line or, if using Caddy's config API, `Content-Type: application/dhall`.
