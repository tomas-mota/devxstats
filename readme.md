# Devxstats
⚠️ Warning: Active Development ⚠️
 
Service that fetches and stores developer experience related metrics from different sources and makes them available through a rest api

# TODO:
- define a separate model for each client communicating to a source
- decide on database
- restructure storage package to mirror sources structure
- document public methods
- handle pagination
- make different syncs run in paralel (e.g. PullRequests and repos)
- different entities can be synced in different frequencies, such as repos vs commits
- put constants such as database name somewhere central
- NEXT: figure out pk issues

## Bitbucket todos
- go-scm does not support listing projects, add that capability

# Supported data sources

## Git

### Github
Configuration

| Environment Variable | Description                                                                          | Required |
| -------------------- | ------------------------------------------------------------------------------------ | -------- |
| GITHUB_URL           | URL pointing to the api (default https://api.github.com)                             | No       |
| GITHUB_TOKEN         | Token used for authentication (can be created at https://github.com/settings/tokens) | Yes      |

### Bitbucket
Configuration

| Environment Variable | Description                                                 | Required |
| -------------------- | ----------------------------------------------------------- | -------- |
| BITBUCKET_URL        | URL pointing to the api (default https://api.bitbucket.org) | No       |
| BITBUCKET_TOKEN      | Token used for authentication                               | Yes      |

## CD

### Octopus
Configuration

| Environment Variable | Description                   | Required |
| -------------------- | ----------------------------- | -------- |
| OCTOPUS_URL          | URL pointing to the api       | Yes      |
| OCTOPUS_TOKEN        | Token used for authentication | Yes      |