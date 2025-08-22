Title: Git Guidelines
Type: Doc

# Description:

Defines Git branching strategy, commit guidelines, and PR process for Eventrix project.

# What problem we are solving:

Ensures that all contributors follow a consistent, clean Git workflow → improving code quality and project maintainability.

# Contents:

- Branching model → feature branches, hotfixes, releases
- Commit message guidelines
- PR process

# Actions:

## Branching model

- `main` → stable, production-ready → protected branch
- `develop` (optional, if needed for big feature waves)
- `feature/*` → feature development
- `bugfix/*` → bug fixes
- `hotfix/*` → urgent fixes on main
- `release/*` → release preparation branches

## Commit message guidelines

Format:
<type>: <short description>

[optional longer body]

BREAKING CHANGE: [optional breaking change description]


Allowed types:

- feat → new feature
- fix → bug fix
- docs → documentation only changes
- refactor → code refactoring
- test → adding or correcting tests
- chore → build process or auxiliary tool changes

Example:

feat: implement KafkaBus publish


## PR process

- PRs must be linked to Issues → use `Closes #ISSUE_NUMBER`
- PRs must pass CI (tests + lint)
- PR title must follow commit message guidelines
- PR must be reviewed by at least 1 reviewer
- Squash and merge → preferred strategy

---

## Summary

Following this Git model ensures:

- Clean history  
- Easier release management  
- Easier collaboration  
- Professional OSS quality

---


