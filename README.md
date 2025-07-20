# learn-cicd-starter (Notely)

This repo contains the starter code for the "Notely" application for the "Learn CICD" course on [Boot.dev](https://boot.dev).

## Local Development

Make sure you're on Go version 1.22+.

Create a `.env` file in the root of the project with the following contents:

```bash
PORT="8080"
```

Run the server:

```bash
go build -o notely && ./notely
```

_This starts the server in non-database mode._ It will serve a simple webpage at `http://localhost:8080`.

You do _not_ need to set up a database or any interactivity on the webpage yet. Instructions for that will come later in the course!

twalker's version of Boot.dev's Notely app.

![test status](https://github.com/twalker/learn-cicd-starter/actions/workflows/ci.yml/badge.svg)

## Notes

Docker commands

```bash
docker build -t timwalker2k/notely:latest .
docker run -e PORT=8080 -p 8080:8080 timwalker2k/notely:latest
```

```bash
gcloud auth configure-docker us-central1-docker.pkg.dev
# gcloud builds submit --tag REGION-docker.pkg.dev/PROJECT_ID/REPOSITORY/IMAGE:TAG .
gcloud builds submit --tag us-central1-docker.pkg.dev/bootscicd/notely-ar-repo/notely:latest .



gcloud artifacts repositories list
```

```bash
https://test-983191307534.us-central1.run.app
https://notely-983191307534.us-central1.run.app
bootdev config base_url https://notely-983191307534.us-central1.run.app


turso auth login
turso db list
turso db shell notely-db

libsql://notely-db-twalker.aws-us-east-2.turso.io
```
