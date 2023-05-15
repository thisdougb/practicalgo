# Workflows

### GitHub Setup

Both the _develop_ and _main_ branches should be protected branches (configure in Settings on GitHub).
This means that you cannot push directly to them, and you must create a PR to merge your changes.

The status checks need to have run at least once before you can setup the branch protection rules.

### Pull Request

This runs tests when a PR is created, typically from your working branch to develop or main.

The *dev* tag is included when running tests.
All test files shouuld have the following annotation to ensure they are included in the test run:

```
//go:build dev
```

### Release

When you push a new tag this will build and push a new release to GitHub packages, using goreleaser.

```
$ git tag -a v1.0.0 -m "my first version"
$ git push origin v1.0.0
```