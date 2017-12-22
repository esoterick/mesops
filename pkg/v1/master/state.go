// MIT License
//
// Copyright (c) [2017-2018] [Demitri Swan]
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package master

type GetStateResponse struct {
	Type     *string `json:"type"`
	GetState *struct {
		GetAgents *struct {
			Agents []*struct {
				Active    *bool `json:"active"`
				AgentInfo *struct {
					Hostname *string `json:"hostname"`
					ID       *struct {
						Value *string `json:"value"`
					} `json:"id"`
					Port      *int `json:"port"`
					Resources []*struct {
						Name   *string `json:"name"`
						Role   *string `json:"role"`
						Scalar *struct {
							Value *float64 `json:"value"`
						} `json:"scalar,omitempty"`
						Type   *string `json:"type"`
						Ranges *struct {
							Range []*struct {
								Begin *int `json:"begin"`
								End   *int `json:"end"`
							} `json:"range"`
						} `json:"ranges,omitempty"`
					} `json:"resources"`
				} `json:"agent_info"`
				Pid            *string `json:"pid"`
				RegisteredTime *struct {
					Nanoseconds *int64 `json:"nanoseconds"`
				} `json:"registered_time"`
				TotalResources []*struct {
					Name   *string `json:"name"`
					Role   *string `json:"role"`
					Scalar *struct {
						Value *float64 `json:"value"`
					} `json:"scalar,omitempty"`
					Type   *string `json:"type"`
					Ranges *struct {
						Range []*struct {
							Begin *int `json:"begin"`
							End   *int `json:"end"`
						} `json:"range"`
					} `json:"ranges,omitempty"`
				} `json:"total_resources"`
				Version *string `json:"version"`
			} `json:"agents"`
		} `json:"get_agents"`
		GetExecutors *struct {
			Executors []*struct {
				AgentID *struct {
					Value *string `json:"value"`
				} `json:"agent_id"`
				ExecutorInfo *struct {
					Command *struct {
						Shell *bool   `json:"shell"`
						Value *string `json:"value"`
					} `json:"command"`
					ExecutorID *struct {
						Value *string `json:"value"`
					} `json:"executor_id"`
					FrameworkID *struct {
						Value *string `json:"value"`
					} `json:"framework_id"`
				} `json:"executor_info"`
			} `json:"executors"`
		} `json:"get_executors"`
		GetFrameworks *struct {
			Frameworks []*struct {
				Active        *bool `json:"active"`
				Connected     *bool `json:"connected"`
				FrameworkInfo *struct {
					Checkpoint      *bool    `json:"checkpoint"`
					FailoverTimeout *float64 `json:"failover_timeout"`
					Hostname        *string  `json:"hostname"`
					ID              *struct {
						Value *string `json:"value"`
					} `json:"id"`
					Name      *string `json:"name"`
					Principal *string `json:"principal"`
					Role      *string `json:"role"`
					User      *string `json:"user"`
				} `json:"framework_info"`
				RegisteredTime *struct {
					Nanoseconds *int64 `json:"nanoseconds"`
				} `json:"registered_time"`
				ReregisteredTime *struct {
					Nanoseconds *int64 `json:"nanoseconds"`
				} `json:"reregistered_time"`
			} `json:"frameworks"`
		} `json:"get_frameworks"`
		GetTasks *struct {
			CompletedTasks []*struct {
				AgentID *struct {
					Value *string `json:"value"`
				} `json:"agent_id"`
				ExecutorID *struct {
					Value *string `json:"value"`
				} `json:"executor_id"`
				FrameworkID *struct {
					Value *string `json:"value"`
				} `json:"framework_id"`
				Name      *string `json:"name"`
				Resources []*struct {
					Name   *string `json:"name"`
					Role   *string `json:"role"`
					Scalar *struct {
						Value *float64 `json:"value"`
					} `json:"scalar,omitempty"`
					Type   *string `json:"type"`
					Ranges *struct {
						Range []*struct {
							Begin *int `json:"begin"`
							End   *int `json:"end"`
						} `json:"range"`
					} `json:"ranges,omitempty"`
				} `json:"resources"`
				State             *string `json:"state"`
				StatusUpdateState *string `json:"status_update_state"`
				StatusUpdateUUID  *string `json:"status_update_uuid"`
				Statuses          []*struct {
					AgentID *struct {
						Value *string `json:"value"`
					} `json:"agent_id"`
					ContainerStatus *struct {
						NetworkInfos []*struct {
							IPAddresses []*struct {
								IPAddress *string `json:"ip_address"`
							} `json:"ip_addresses"`
						} `json:"network_infos"`
					} `json:"container_status"`
					ExecutorID *struct {
						Value *string `json:"value"`
					} `json:"executor_id"`
					Source *string `json:"source"`
					State  *string `json:"state"`
					TaskID *struct {
						Value *string `json:"value"`
					} `json:"task_id"`
					Timestamp *float64 `json:"timestamp"`
					UUID      *string  `json:"uuid"`
				} `json:"statuses"`
				TaskID *struct {
					Value *string `json:"value"`
				} `json:"task_id"`
			} `json:"completed_tasks"`
		} `json:"get_tasks"`
	} `json:"get_state"`
}