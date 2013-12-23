github-repo-gen
===============

_Version 0.1_

__github-repo-gen__ is a toy to generate a Markdown index document of GitHub user's starred repositories. 

# Install
```
$ go get -u github.com/google/go-github/github
$ go get -u github.com/skyjia/github-repo-gen

```

Exuecte it:

```
$ github-repo-gen -u USERNAME
```

Which takes the following flags:

- __-u:__ This the GitHub username of which starred repositories you want to generate. (Required)

Output to a Markdown document:
```
$ github-repo-gen -u USERNAME > /path/to/your/document.md
```