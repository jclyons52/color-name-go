# color-name-go

Go port of the npm package [`color-name`](https://www.npmjs.com/package/color-name)
(1.1.4), verified against the real JS via a JS-oracle parity gate.

## API

`Lookup(name string) (RGB, bool)` — returns the `[r, g, b]` values for a CSS
color name and whether the name exists (JS `colorName[name]` is `undefined`
for unknown names).

`Colors` — the full `map[string]RGB` (148 entries).

## Parity

`go test ./... -run TestParity` runs the Go port against the real
`color-name` under node and requires **0 mismatches**.
