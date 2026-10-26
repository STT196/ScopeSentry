# Release Guide for ScopeSentry

This project uses GitHub Actions plus GoReleaser to publish:

- GitHub release assets
- Docker image tags
- a `latest` Docker tag

The release workflow is defined in `.github/workflows/release.yml` and the GoReleaser config is in `.goreleaser.yaml`.

## Prerequisites

Before creating a release, make sure you have:

- Git installed and repository access
- GitHub CLI (`gh`) installed if you want to trigger the release manually
- Docker Hub credentials configured in GitHub repository secrets
- a valid version tag such as `v0.0.5`

## Required GitHub secrets

Set these in the repository settings under Settings -> Secrets and variables -> Actions:

- `DOCKER_IMAGE` — for example `autumn27/scopesentry`
- `DOCKERHUB_USERNAME` — Docker Hub username
- `DOCKERHUB_TOKEN` — Docker Hub password or access token

GitHub also provides `GITHUB_TOKEN` automatically for the release step.

## Release flow

### Option 1: Create a release by pushing a tag

1. Update the code and commit everything:

   ```bash
   git add .
   git commit -m "Prepare release v0.0.5"
   git push origin main
   ```

2. Create and push the version tag:

   ```bash
   git tag v0.0.5
   git push origin v0.0.5
   ```

3. The GitHub Actions workflow starts automatically because this repo listens for tags matching `v*`.

### Option 2: Trigger the workflow manually

You can also run it from GitHub actions or with GitHub CLI:

```bash
gh workflow run release.yml -f version=v0.0.5
```

This workflow will:

- build the frontend under `frontend/`
- copy the frontend build into `cmd/main/static/`
- install Go and run GoReleaser
- create GitHub release artifacts
- login to Docker Hub
- build the Docker image
- push tags like:
  - `${DOCKER_IMAGE}:0.0.5`
  - `${DOCKER_IMAGE}:latest`

## Version format

Use semantic version tags like:

```bash
v0.0.5
v1.2.3
```

The workflow strips the leading `v` internally and publishes the version as a semver tag.

## Important notes

- The project uses GoReleaser for GitHub release assets.
- The Docker image is built separately after the GoReleaser step.
- The frontend must be built before the Go binary is packaged, otherwise the static app bundle can be missing.
- If Docker image publishing fails, check the Docker Hub credentials and the `DOCKER_IMAGE` secret value.

## Quick commands

```bash
# Create a release tag
git tag v0.0.5

# Push tag to trigger GitHub workflow
git push origin v0.0.5

# Or manually trigger the workflow
gh workflow run release.yml -f version=v0.0.5
```

## After the release

Confirm the following:

- GitHub release page contains the generated binary archives
- Docker Hub has the new version tag
- `latest` tag exists and points to the newest image

If you want, I can also add a short section to the main README linking to this new release guide.
