---
id: 6
database_id: 803221949
node_id: MDU6SXNzdWU4MDMyMjE5NDk=
status: open
title: "add flag traceback"
labels: ["help wanted"]
url: https://github.com/octolab/testit/issues/6
created_at: 2021-02-08T06:01:26Z
updated_at: 2021-02-08T06:02:24Z
---

# add flag traceback

it configures `GOTRACEBACK` env variable for `go test` and allows autocompletion.

motivation: part of autocompletion, shows all available values.

--traceback: [none, **single**, all, system, crash]

See https://golang.org/pkg/runtime/
