# GIT

## Git hooks
Hooks are programs you can place in a hooks directory to trigger actions at certain points in git’s execution.
$GIT_DIR/hooks/\* (or `git config core.hooksPath`/\*)

Before Git invokes a hook, it changes its working directory to either $GIT_DIR in a bare repository or the root of the working tree in a non-bare repository. An exception are hooks triggered during a push (pre-receive, update, post-receive, post-update, push-to-checkout) which are always executed in $GIT_DIR.

### Common hooks
pre-commit: invoked before git commit

pre-push: invoked before push (can stop push)
