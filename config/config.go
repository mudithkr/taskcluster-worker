// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package config

import "github.com/taskcluster/taskcluster-worker/runtime"

type (
	// The global configuration settings used by taskcluster-worker, not particular
	// to any engine or plugin.
	Config struct {

		// The number of tasks that this worker supports running in parallel.
		Capacity int `json:"capacity"`

		// The set of credentials that should be used by the worker when
		// authenticating against taskcluster endpoints.
		Credentials struct {

			// The security-sensitive access token for the client. Do not share or expose!
			//
			// Syntax:     ^[a-zA-Z0-9_-]{22,66}$
			AccessToken string `json:"accessToken"`

			// The certificate for the client, if using temporary credentials and/or using authorized scopes.
			Certificate string `json:"certificate"`

			// The Client ID for the client. Not very helpful, am I?
			//
			// Syntax:     ^[A-Za-z0-9@/:._-]+$
			ClientID string `json:"clientId"`
		} `json:"credentials"`

		// The amount of time to wait between task polling iterations
		PollingInterval int `json:"pollingInterval,omitempty"`

		// The provisioner (if any) that is responsible for spawning instances of this worker. Typically `aws-provisioner-v1`.
		ProvisionerID string `json:"provisionerId"`

		// Configuration relating to the polling of the TaskCluster Queue.
		QueueService struct {

			// The number of seconds prior to the task claim expiry that tasks should be reclaimed.
			ExpirationOffset int `json:"expirationOffset"`
		} `json:"queueService"`

		// Properties related to the TaskCluster Platform
		Taskcluster struct {

			// Properties defining interaction with the TaskCluster queue
			Queue struct {

				// Base URL for TaskCluster queue
				URL string `json:"url,omitempty"`
			} `json:"queue,omitempty"`
		} `json:"taskcluster,omitempty"`

		// A logical group for this worker to belong to, such as an AWS region.
		WorkerGroup string `json:"workerGroup"`

		// A unique name that can be used to identify which worker instance this is (such as AWS instance id).
		WorkerID string `json:"workerId"`

		// Type of worker pool the worker belongs to.
		WorkerType string `json:"workerType,omitempty"`
	}
)

var ConfigSchema = func() runtime.CompositeSchema {
	schema, err := runtime.NewCompositeSchema(
		"config",
		`
		{
		  "$schema": "http://json-schema.org/draft-04/schema#",
		  "additionalProperties": false,
		  "description": "The global configuration settings used by taskcluster-worker, not particular\nto any engine or plugin.",
		  "properties": {
		    "capacity": {
		      "description": "The number of tasks that this worker supports running in parallel.",
		      "title": "Capacity",
		      "type": "integer"
		    },
		    "credentials": {
		      "description": "The set of credentials that should be used by the worker when\nauthenticating against taskcluster endpoints.",
		      "properties": {
		        "accessToken": {
		          "description": "The security-sensitive access token for the client. Do not share or expose!",
		          "pattern": "^[a-zA-Z0-9_-]{22,66}$",
		          "title": "AccessToken",
		          "type": "string"
		        },
		        "certificate": {
		          "description": "The certificate for the client, if using temporary credentials and/or using authorized scopes.",
		          "title": "Certificate",
		          "type": "string"
		        },
		        "clientId": {
		          "description": "The Client ID for the client. Not very helpful, am I?",
		          "pattern": "^[A-Za-z0-9@/:._-]+$",
		          "title": "ClientId",
		          "type": "string"
		        }
		      },
		      "required": [
		        "clientId",
		        "accessToken",
		        "certificate"
		      ],
		      "title": "Credentials",
		      "type": "object"
		    },
		    "pollingInterval": {
		      "description": "The amount of time to wait between task polling iterations",
		      "title": "PollingInterval",
		      "type": "integer"
		    },
		    "provisionerId": {
		      "description": "The provisioner (if any) that is responsible for spawning instances of this worker. Typically `+"`"+`aws-provisioner-v1`+"`"+`.",
		      "title": "ProvisionerID",
		      "type": "string"
		    },
		    "queueService": {
		      "description": "Configuration relating to the polling of the TaskCluster Queue.",
		      "properties": {
		        "expirationOffset": {
		          "description": "The number of seconds prior to the task claim expiry that tasks should be reclaimed.",
		          "title": "ExpirationOffset",
		          "type": "integer"
		        }
		      },
		      "required": [
		        "expirationOffset"
		      ],
		      "title": "QueueService",
		      "type": "object"
		    },
		    "taskcluster": {
		      "description": "Properties related to the TaskCluster Platform",
		      "properties": {
		        "queue": {
		          "description": "Properties defining interaction with the TaskCluster queue",
		          "properties": {
		            "url": {
		              "description": "Base URL for TaskCluster queue",
		              "title": "BaseUrl",
		              "type": "string"
		            }
		          },
		          "title": "TaskclusterQueueProperties",
		          "type": "object"
		        }
		      },
		      "title": "TaskclusterProperties",
		      "type": "object"
		    },
		    "workerGroup": {
		      "description": "A logical group for this worker to belong to, such as an AWS region.",
		      "title": "WorkerGroup",
		      "type": "string"
		    },
		    "workerId": {
		      "description": "A unique name that can be used to identify which worker instance this is (such as AWS instance id).",
		      "title": "WorkerID",
		      "type": "string"
		    },
		    "workerType": {
		      "description": "Type of worker pool the worker belongs to.",
		      "title": "WorkerType",
		      "type": "string"
		    }
		  },
		  "required": [
		    "credentials",
		    "provisionerId",
		    "workerGroup",
		    "workerId",
		    "capacity",
		    "queueService"
		  ],
		  "title": "Config",
		  "type": "object"
		}
		`,
		true,
		func() interface{} {
			return &Config{}
		},
	)
	if err != nil {
		panic(err)
	}
	return schema
}()
