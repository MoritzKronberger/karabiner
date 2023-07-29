# Karabiner-Elements Complex Modifications Generator

Generate custom JSON config for [Carabiner-Elements Complex Modifications](https://karabiner-elements.pqrs.org/docs/json/).

Inspired by [Max Stoiber's TypeScript Carabiner-Elements config generator](https://github.com/mxstbr/karabiner).

## Generate config

Modify [config.go](./config.go) to match your desired config.

Run the generator:

```bash
go run .
```

Paste the JSON config into `~/.config/karabiner/assets/complex_modifications`.

Activate the Complex Modifications using Carabiner-Elements.
