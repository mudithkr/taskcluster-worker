// This file is AUTO-GENERATED by github.com/taskcluster/taskcluster-worker/codegen/schema-gen - do not edit!

package docker

func ConfigSchema() string {
	return "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"additionalProperties\":false,\"description\":\"Config applicable to docker engine\\n\",\"properties\":{\"rootVolume\":{\"description\":\"Root Volume blah blah\\n\",\"title\":\"Root Volume\",\"type\":\"string\"}},\"required\":[\"rootVolume\"],\"title\":\"Config\",\"type\":\"object\"}"
}

func NewConfig() *Config {
	return new(Config)
}