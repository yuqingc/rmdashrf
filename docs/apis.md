# APIs

<a href="https://docs.microsoft.com/en-us/rest/api/storageservices/file-service-rest-api" target="_blank">Microsoft Azure APIs</a> is a good practice for API design. The API of RMDASHRF is designed by referring to it.

## Group: `/api/v1`

Caveats: 
- `default` is the current volume name or owner name although there is only a single account and a single mounted volume at present. We will support multi mounted volumes for different account in future versions

- Relative parent path (`..`) is not allowed, so there is no possibility you reach anything outside the mounted directory

### List directories and files of a specified directory

This operation returns a list of files or directories under the specified directory. It lists the contents for a single level of the directory hierarchy.

```
GET /default/<my directory path>?restype=directory&comp=list&all=<bool>&maxresults=<int>&extension=<string>
```

**Request params**

|Name|Description|
|-|-|
|restype|Required. Set it to "directory"|
|comp|Required. Set it to "list"|
|all|Optional. Lists all files when set to `true`. Default to `false`, files of which names start with dot `.` are not included|
|maxresults|Optional. Specifies the maximum number of files and/or directories to return. If the request does not specify `maxresults` or specifies a value greater than 5,000, the server will return up to 5,000 items.|
|extension|Optional. Filters the results to return only files and directories whose name has the specified extension or suffix. For example, `extension=.js`|

**Response status**

```
200 OK
```

**Response body**

JSON

### Create a file

```
PUT
```