---
code: null
id: MDU6SXNzdWU4MDMyMjE5NDk=
databaseId: 803221949
number: 6
url: https://github.com/octolab/testit/issues/6
title: "add flag traceback"
issueType: null
labels:
  - "help wanted"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2021-02-08T06:01:26Z
updatedAt: 2021-02-08T06:02:24Z
lastEditedAt: 2021-02-08T06:02:03Z
closedAt: null
---

# add flag traceback

it configures `GOTRACEBACK` env variable for `go test` and allows autocompletion.

motivation: part of autocompletion, shows all available values.

--traceback: [none, **single**, all, system, crash]

See https://golang.org/pkg/runtime/
