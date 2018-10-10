# APIs

<a href="https://docs.microsoft.com/en-us/rest/api/storageservices/file-service-rest-api" target="_blank">Microsoft Azure APIs</a> is a good practice for API design. The API of RMDASHRF is designed by learning from it.

Table of Contents
=================

<!--ts-->
   * [APIs](#apis)
   * [Table of Contents](#table-of-contents)
      * [Group: /api/v1](#group-apiv1)
         * [List directories and files of a specified directory](#list-directories-and-files-of-a-specified-directory)
         * [Create a file](#create-a-file)
         * [Create a directory](#create-a-directory)
         * [Remove a file or directory](#remove-a-file-or-directory)
         * [Move or rename a file or directory](#move-or-rename-a-file-or-directory)
         * [Copy a file or directory](#copy-a-file-or-directory)
         * [Download a file](#download-a-file)
         * [Download a directory as a zip file](#download-a-directory-as-a-zip-file)

<!-- Added by: matt, at: 2018-09-29T01:21+08:00 -->

<!--te-->

## Group: `/api/v1`

Caveats: 
- `default` is the current volume name or owner name although there is only a single account and a single mounted volume at present. We will support multi mounted volumes for different accounts in future versions

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

application/json

```json
{
    "metadata": {
        "total": 4
    },
    "items": [
        {
            "name": "run.sh",
            "size": 4096,
            "mode": "-rwxr-xr-x",
            "modTime": "2018-09-24 00:31:44.047714198 +0800 CST",
            "isDir": false
        }
    ]
}
```

### Create a file

Fails if the file is already exists or the path to the directory does not exist

```
PUT /default/<my directory path>/<my file>
```

**Response status**

```
201 Created
```

**Response body**

```
NONE
```

### Create a directory

```
PUT /default/<my directory path>/<new directory>?action=create&restype=directory&parents=<bool>
```

**Request params**

|Name|Description|
|-|-|
|action|Required. Set it to "create"|
|restype|Required. Set it to "directory"|
|parents|Optional. Default to `false`. Make parents directories if needed when set to `true`. Similar to `mkdir -p`, which means this only works when creating a *directory*|

**Response status**

```
201 Created
```

**Response body**

```
NONE
```

### Remove a file or directory

```
DELETE /default/<my directory or file path>?recursive=<bool>
```

**Request params**

|Name|Description|
|-|-|
|recursive|Optional. Default to `false`. Delete directories and their children recursively when set to `true`|

**Response status**

```
204 No Content
```

**Response body**

```
NONE
```

### Move or rename a file or directory

It is similar to the `mv` command. The difference is that the new path cannot be a path which has already existed

```
PATCH /default/<old path>?action=rename&to=<new path>
```

**Request params**

|Name|Description|
|-|-|
|action|Required. Set it to "rename"|
|to|New path|

**Response status**

```
204 No Content
```

**Response body**

```
NONE
```

### Copy a file or directory

```
PUT /default/<new path>?action=copy&from=<old path>&restype=<directory>
```

`<new path>` does not refer to the directory into which you copy your file/directory. It is the path of the copied file/directory itself

**Request params**

|Name|Description|
|-|-|
|action|Required. Set it to "copy"|
|from|Required. Old path|
|restype|Optional. Set it to "directory" to copy a directory, or copy a file|

**Response status**

```
201 Created
```

**Response body**

```
NONE
```

### Download a file

```
GET /default/<path to a file>
```

### Download a directory as a `zip` file

```
GET /default/<path to a directory>?restype=directory
```

### Upload a file

```
POST /default/<path to a directory where you want to put your file>
```

The request will be rejected either when the directory does not exist or when the file already exists

**Request content type**

form-data

**Request body**

|Name|Description|
|-|-|
|file|Required. File|

**Response status**

```
200 OK
```