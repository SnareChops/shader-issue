## Example for bug report for Ebiten Shader language Kage

To run `go run github.com/SnareChops/shader-issue`

### Expected behaviour
This is the simplest possible example reproduction of the issue. The shader returns a static vec4 color with an alpha of 0.1. The background is set to blue to demonstrate the transparency of the white square. The expected result is that the white square should blend with the blue background, effectively making the white square seem semi-transparent.

## Actual behaviour
The white square renders with full alpha, no matter what alpha value is set in the shader