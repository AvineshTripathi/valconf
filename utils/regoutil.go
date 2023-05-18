package utils

var RegoUtils = `
package example

pod_privileged[msg] {
	
	input.pod[i].pod[j].privileged == "true"
    msg = sprintf("Pod name is %v", [input.pod[i].pod[j].name])
}

pod_limits[msg] {

	input.pod[i].pod[j].limits == "false"
    msg = sprintf("Pod %v has no set limits", [input.pod[i].pod[j].name])
}

deploy_replicas[msg] {
	input.deploy[i].replica > 2
    msg = sprintf("Deployement %v exceeds the maximum number of replicas", [input.deploy[i].name])
}

deploy_strategy[msg] {
	input.deploy[i].strategy != "RollingUpdate"
    msg = sprintf("Deployement %v strategy is not recommmended", [input.deploy[i].name])
}

deploy_labels[msg] {
	input.deploy[i].labels == null
    msg = sprintf("Deployement %v requires a Label", [input.deploy[i].name])
}

service_ports[msg] {
	allowed_ports := {3030, 120, 3036}
	input.svc[i].ports != allowed_ports[_]
    msg = sprintf("Service %s has ports which is not secured", [input.svc[i].name])
}

roles_verbs[msg] {
    input.role[i].rules[j].verbs[_] == "update"
    msg = sprintf("An update permission is given to %v", [input.role[i].name])
}

svcAcc_namespace[msg] {
	input.svc_account[i].namespace == "default"
    msg = sprintf("Service account %v is with default namespace", [input.svc_account[i].name])
}
`

// secured_path := ["/etc","/var/run/docker.sock", "/var/run/secrets/kubernetes.io/serviceaccount", "/proc"]
// host_mounts[msg] if {
// 	some i
// 	some path in secured_path
// 	input.volumes[_].path == path
// 	msg = sprintf("Pod is %v and is having a hostPath", [input.volumes[i].name])
// }

// input.input.request.object.spec.containers[_].image == "example.com/image"
// input.input[_].privileged == "true"