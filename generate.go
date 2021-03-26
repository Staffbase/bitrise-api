package tools

// todo: Find a better way to ignore tags or make a breaking release
//go:generate gobin -m -run github.com/go-swagger/go-swagger/cmd/swagger flatten swagger_patched.json -o swagger_patched_flat.json
//go:generate sh -c "cat swagger_patched_flat.json | jq '[., (.paths | map_values(.[] |= del(.tags?)) | {paths: .})] | add' > swagger_patched_go.json"
//go:generate gobin -m -run github.com/go-swagger/go-swagger/cmd/swagger generate client -A bitrise-api -f swagger_patched_go.json -t go --default-scheme=https --with-flatten=full
