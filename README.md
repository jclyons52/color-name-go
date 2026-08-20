# color-name-go

Go port of the npm package [`color-name`](https://www.npmjs.com/package/color-name)
(1.1.4), verified against the real JS via a JS-oracle parity gate.

The vendored original (for the parity oracle) lives in `original/`.

## API

`Names` — the full `map[string]RGB` mirroring `module.exports` (148 entries,
lowercase CSS color names → `[r, g, b]`).

`Get(name string) (RGB, bool)` — returns the `[r, g, b]` values for a CSS
color name and whether the name exists (JS `colorName[name]` is `undefined`
for unknown names).

```go
v, ok := colorname.Get("rebeccapurple") // v = RGB{102, 51, 153}, ok = true
```

## Parity

`go test ./...` runs the Go port against the real `color-name` under node and
requires **0 mismatches** across all 148 names (count, keys, and every RGB
value).
