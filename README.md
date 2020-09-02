# abstract-runner
 The abstract-runner was a gitlab-runner like tool which was worked for pulling the project from the remote git server, running and debugging in a k8s pod.

## git to do
- install a git: the installation was better completed by the Docker environment
- combine the git commands inside code logic instead of depending on shell scripts
- it needs a git hook or custom trigger for triggering the workers to take operations
- `Important` git must remove the local branch which has been removed by others.

## worker to do 
- async queue
- status
- full log 
- mode (e.g. auto or manual)

## api to do
- http/tcp
- push