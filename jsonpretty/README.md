Implement a JSON pretty printer.

## Instructions

* Use `json.NewDecoder` to decode JSON value into Golang value.

```golang
var data interface{}

decoder := json.NewDecoder(reader)
err = decoder.Decode(&data)
log.Printf("decoded data: %v", data)
```

In the above example, the JSON decoder MUST convert a JSON value to a Golang value that satisfies the type of `data`, which is the type `interface{}`. Since any Golang value satisfies "interface{}", the above code can convert any JSON value to any Golang value.

* Use [json.NewEncoder](https://golang.org/pkg/encoding/json/#NewEncoder) to encode a Golang value into JSON value.
* [json.Encoder.SetIndent](https://golang.org/pkg/encoding/json/#Encoder.SetIndent)
  * Configure the JSON pretty printer.

Most of the time you'd actually want the decoder to convert to a more specific value. For example, we might want to convert a user JSON to a User struct.

The JSON value:

```
{
  "name": "howard",
  "email": "hayeah@gmail.com"
}
```

The User struct:

```golang
type User {
  Name string
  Email string
}
```

The JSON decoding:

```golang
var data User

decoder := json.NewDecoder(reader)
err = decoder.Decode(&data)
log.Printf("decoded data: %v", data)
```

For more details see: https://blog.golang.org/json-and-go

## Examples

```
$ go run jsonpretty.go octocat.json
{
  "avatar_url": "https://avatars0.githubusercontent.com/u/583231?v=3",
  "bio": null,
  "blog": "http://www.github.com/blog",
  "company": "GitHub",
  "created_at": "2011-01-25T18:44:36Z",
  "email": "octocat@github.com",
  "events_url": "https://api.github.com/users/octocat/events{/privacy}",
  "followers": 1779,
  "followers_url": "https://api.github.com/users/octocat/followers",
  "following": 6,
  "following_url": "https://api.github.com/users/octocat/following{/other_user}",
  "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
  "gravatar_id": "",
  "hireable": null,
  "html_url": "https://github.com/octocat",
  "id": 583231,
  "location": "San Francisco",
  "login": "octocat",
  "name": "The Octocat",
  "organizations_url": "https://api.github.com/users/octocat/orgs",
  "public_gists": 8,
  "public_repos": 7,
  "received_events_url": "https://api.github.com/users/octocat/received_events",
  "repos_url": "https://api.github.com/users/octocat/repos",
  "site_admin": false,
  "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
  "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
  "type": "User",
  "updated_at": "2017-04-03T23:20:54Z",
  "url": "https://api.github.com/users/octocat"
}
```

```
$ go run jsonpretty.go octocat-repos.json
[
  {
    "archive_url": "https://api.github.com/repos/octocat/git-consortium/{archive_format}{/ref}",
    "assignees_url": "https://api.github.com/repos/octocat/git-consortium/assignees{/user}",
    "blobs_url": "https://api.github.com/repos/octocat/git-consortium/git/blobs{/sha}",
    "branches_url": "https://api.github.com/repos/octocat/git-consortium/branches{/branch}",
    "clone_url": "https://github.com/octocat/git-consortium.git",
    "collaborators_url": "https://api.github.com/repos/octocat/git-consortium/collaborators{/collaborator}",
...
```






