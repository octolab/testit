---
code: null
id: MDU6SXNzdWU3OTYyODgyMzg=
databaseId: 796288238
number: 2
url: https://github.com/octolab/testit/issues/2
title: "testit go quick"
issueType: null
labels:
  - "help wanted"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2021-01-28T19:31:51Z
updatedAt: 2021-02-04T09:04:18Z
lastEditedAt: null
closedAt: 2021-01-28T21:29:38Z
---

# testit go quick

equivalent for

```bash
$ go test -run ^Fake$$ -tags $(GOTAGS) ./... | { grep -v 'no tests to run' || true; }
```
