# svninfo
extract subversion (svn) information from a working copy and prints them on stdout

#how to
`go install github.com/genez/svninfo`  
  
  

- **MUST** be executed with cwd set on SVN working copy
- current supported arguments:
    - `svninfo revision` prints the commit revision number (e.g. `26654`)
    - `svninfo timestamp` prints the commit timestamp (e.g. `2017-01-10T10:34:01.027421Z`)
    
#Usage Example in `go build`

Just define two (exported) variables in a package of your choice. They will be set at compile time with SVN revision and timestamp using `ldflags` of `go build` (or `go install`) command:
```
go build -ldflags="-X main.Commit_Revision=`svninfo revision` -X main.Commit_TimeStamp=`svninfo timestamp`"
```
