package privileged_pod

default deny = false

deny {
    input.kind == "Pod"
    container := input.spec.containers[_]
    container.securityContext.privileged == "true"

}