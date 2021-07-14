# fvv - Fusozay Var Var

A CLI tool for quick text template rendering

Fusozay Var Var means "have fun"
It is a reference to something I see a lot

Fusozay Var Var is a text pre-processor application for quickly rendering out text templates.
I often write outfit and character descriptions that reuses a lot of elements.
This allows me to DRY up my descriptions and still quickly get results.

Template requirements:
- The template must be valid golang templates.
- The template must not require any variables to be passed in.
- The template must not be a named definition.

## Usage

All errors go to stderr, so it can work as a preprocessor or be piped to pbcopy.

1. Set up a golang template
1. run `fvv < <template file>`
