---
id: 4
database_id: 802659254
node_id: MDU6SXNzdWU4MDI2NTkyNTQ=
status: closed
title: "fix broken build related to TestGoTestStackTrace"
labels: ["help wanted"]
url: https://github.com/octolab/testit/issues/4
created_at: 2021-02-06T09:52:18Z
updated_at: 2021-02-08T05:50:39Z
---

# fix broken build related to TestGoTestStackTrace

```
        	            	Diff:
290        	            	--- Expected
291        	            	+++ Actual
292        	            	@@ -8,9 +8,9 @@
293        	            	 
294        	            	-1: running [Created by testing.(*T).Run @ /usr/local/Cellar/go/1.15.6/libexec/src/testing/testing.go:1168]
295        	            	-    testing  /usr/local/Cellar/go/1.15.6/libexec/src/testing/testing.go:1072                             tRunner.func1.1(*T(#1), func(#2))
296        	            	-    testing  /usr/local/Cellar/go/1.15.6/libexec/src/testing/testing.go:1075                             tRunner.func1(*T(#3))
297        	            	-             /usr/local/Cellar/go/1.15.6/libexec/src/runtime/panic.go:969                                panic(interface{}(#1))
298        	            	+1: running [Created by testing.(*T).Run @ /home/travis/.gimme/versions/go1.15.8.linux.amd64/src/testing/testing.go:1168]
299        	            	+    testing  /home/travis/.gimme/versions/go1.15.8.linux.amd64/src/testing/testing.go:1072               tRunner.func1.1(*T(#1), func(#2))
300        	            	+    testing  /home/travis/.gimme/versions/go1.15.8.linux.amd64/src/testing/testing.go:1075               tRunner.func1(*T(#3))
301        	            	+             /home/travis/.gimme/versions/go1.15.8.linux.amd64/src/runtime/panic.go:969                  panic(interface{}(#1))
302        	            	     panicked /Users/ksamigullin/Development/public/testit/internal/testdata/panicked/panicked.go:19      TheBad.Divide(...)
303        	            	-    panicked /Users/ksamigullin/Development/public/testit/internal/testdata/panicked/panicked_test.go:13 TestTheBad_Divide(*T(#3))
304        	            	-    testing  /usr/local/Cellar/go/1.15.6/libexec/src/testing/testing.go:1123                             tRunner(*T(#3), func(0x114df50))
305        	            	+    panicked /Users/ksamigullin/Development/public/testit/internal/testdata/panicked/panicked_test.go:13 TestTheBad_Divide(#3)
306        	            	+    testing  /home/travis/.gimme/versions/go1.15.8.linux.amd64/src/testing/testing.go:1123               tRunner(*T(#3), func(0x114df50))
```
the problem is related to a different working directory and GOROOT
