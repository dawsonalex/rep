# Rep

![example workflow](https://github.com/dawsonalex/rep/workflows/Build/badge.svg)

A logger and stats tool for exercises. Track the number of reps, weight, and RPE of exercises and
see how you've progressed over time. 

## Building

Use `make release` to build the application, which outputs the executable to the `bin` dir.

Use `make clean` to clean up built executables.

## GitHub Workflow

The GitHub workflow defined in `base.yml` attempts to do some common things in a simple way. Currently, it does the
following steps under a single job called `Build`:

    - Set up Go environment.
    - Run Go Tests (ignoring if there are none).
    - Runs `make release` to create all binaries.
    - Bump the version based on the commit message.
        - Use `#major`, `#minor`, or `#patch` in your commit message to bump the version and create a new release.
        - Leaving out the above tags will not create a new tag or release version.
    - Generate release logs from the commits between this tag and the last.
    - Create a GitHub release and upload the content of `bin`.