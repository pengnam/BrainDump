# GIT

## Git hooks
Hooks are programs you can place in a hooks directory to trigger actions at certain points in git’s execution.
$GIT_DIR/hooks/\* (or `git config core.hooksPath`/\*)

Before Git invokes a hook, it changes its working directory to either $GIT_DIR in a bare repository or the root of the working tree in a non-bare repository. An exception are hooks triggered during a push (pre-receive, update, post-receive, post-update, push-to-checkout) which are always executed in $GIT_DIR.

### Common hooks
pre-commit: invoked before git commit

pre-push: invoked before push (can stop push)

## Git rebase
git rebase: reapply all commits from current branch to tip of another branch.

Note:
- These commits are new commits, not identical to previous commits.
- **Golden rule**: Never rebase a branch shared by others

- If x number of people were working on that branch, and they have their own commits, *each person will have to create a new merge commit after the rebase*.
