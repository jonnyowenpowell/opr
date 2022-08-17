# features

# todo

- Move prDescription into a proper package
- Integrate JIRA into prDescription
- Make setup run on first run (if config file empty)
- Integrate with git and github for full PR workflow
- Tests!

## JIRA plan

- have jira list stories check for invalid setup
- jira list stories
    - offer most recently used sprint, if set, and pick sprint
    - picking sprint sets most recently used
- integrate jira list stories into PR template


## layout

- cmd (1-1 map with commands) consider bundling weirdos into one cmd
- config
    - setup (register and run setup flows)
    - display (define theme for surveys)
- template (register and render templates)
    - sources (register funcs to source data from in templates)
- jira
- mongo
- github?


# project structure

- cmd
  - setup (auto-run on no config file) [one file?]
  - pr (default command) [one file?]
  - tool [toolName] (eg. objectid) [folder?]

- integration
  - git
  - github
  - jira

- tool: random useful things [separate project?]
  - mongo: generate object id, drop test dbs
  - password: generate password
  - fake: generate phones, names, addrs, etc. 

- pr: generate pr description
  - show available functions
  - accept new pr templates
  - run a template (ask Qs when functions hit)

- config: rw config file
  - secret: rw keychain

- persist: rw arbitrary data [diskv]
  - enforce namespacing to avoid collisions 

- display: hold constants for theme + survey

# data

- git
  - current branch
  - local branches
  - remote branches
  - current branch commit messages
- github
  - API access - setup
    - repo - default to last used
      - open PRs
- jira
  - API access - setup 
    - board - default to last used
      - my open stories
      - all open stories
        - list tasks
        