# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.10.3] - Unreleased

## [0.10.2] - 2019-10-22
### Added
- appengine: Add an option to skip Realm Management checks, where possible
- Add possibility to use a token instead of the private key

## [0.10.1] - 2019-10-18
### Added
- Add commands to generate Device IDs and authentication JWTs
- Add gobuild.sh script
- Add shell completion generator command
- Add CI
- Add appengine command and subcommands

### Fixed
- Fix keypair generation
- Fix Datacenter Replication checks in realm creation command
- Fix a bug that prevented realm key to be set from the command line

### Changed
- utils gen-jwt accepts the private key through -k rather than through -p, just like all other commands

## [0.10.0] - 2019-09-20
### Added
- Initial release
