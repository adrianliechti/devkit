package kafka

import (
	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/container"
)

var (
	_ catalog.Manager       = &Manager{}
	_ catalog.Decorator     = &Manager{}
	_ catalog.ShellProvider = &Manager{}
)

type Manager struct {
}

func (m *Manager) Name() string {
	return "kafka"
}

func (m *Manager) Category() catalog.Category {
	return catalog.MessagingCategory
}

func (m *Manager) DisplayName() string {
	return "Kafka"
}

func (m *Manager) Description() string {
	return "Apache Kafka is a distributed event store and stream-processing platform."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (container.Container, error) {
	image := "confluentinc/cp-kafka:7.1.0"

	return container.Container{
		Image: image,

		PlatformContext: &container.PlatformContext{
			Platform: "linux/amd64",
		},

		Env: map[string]string{
			"KAFKA_NODE_ID":   "1",
			"KAFKA_BROKER_ID": "1",

			"KAFKA_LISTENERS":                      "PLAINTEXT://0.0.0.0:9092,CONTROLLER://0.0.0.0:9093",
			"KAFKA_LISTENER_SECURITY_PROTOCOL_MAP": "CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT",

			"KAFKA_PROCESS_ROLES": "broker,controller",

			"KAFKA_CONTROLLER_QUORUM_VOTERS":  "1@localhost:9093",
			"KAFKA_CONTROLLER_LISTENER_NAMES": "CONTROLLER",

			"KAFKA_INTER_BROKER_LISTENER_NAME": "PLAINTEXT",

			"KAFKA_ADVERTISED_LISTENERS": "PLAINTEXT://localhost:9092",

			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR":         "1",
			"KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS":         "0",
			"KAFKA_TRANSACTION_STATE_LOG_MIN_ISR":            "1",
			"KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR": "1",
		},

		Args: []string{
			"/bin/bash", "-c",
			"sed -i '/KAFKA_ZOOKEEPER_CONNECT/d' /etc/confluent/docker/configure && sed -i '/cub zk-ready/d' /etc/confluent/docker/ensure && echo \"kafka-storage format --ignore-formatted -t $(kafka-storage random-uuid) -c /etc/kafka/kafka.properties\" >> /etc/confluent/docker/ensure && /etc/confluent/docker/run",
		},

		Ports: []*container.ContainerPort{
			{
				Port:     9092,
				Protocol: container.ProtocolTCP,
			},
		},
	}, nil
}

func (m *Manager) Info(instance container.Container) (map[string]string, error) {
	return map[string]string{}, nil
}

func (m *Manager) Shell(instance container.Container) (string, error) {
	return DefaultShell, nil
}
