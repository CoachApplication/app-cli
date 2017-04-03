# COACH: CLI

This is a CLI implementation of the Coach API, using primarily the 
local handler.

## Radi

This effort is heavily based on the Wunderkraut radi-cli tool, with
large parts of code copied.  Most of the other coach elements are
full re-writes.

https://github.com/wunderkraut/radi-cli

## CLI

The CLI is a command line tool that implements the CoachAPI by
wrapping Operations produced from a handler-local builder in 
the urfave/cli library.
Operations that are marked with external usage are interpreted
as cli commands, and optional/required Properties are made available
as command line flags, where their Type() can be interpreted.

