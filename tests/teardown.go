package tests

import "log"

func TeardownPostgresContainer(containerMeta PostgresContainerMeta) {
	if err := containerMeta.Container.Terminate(containerMeta.Context); err != nil {
		log.Fatalf("failed to terminate container: %s", err)
	}
}

func Teardown(containerMetas []PostgresContainerMeta) {
	log.Println("tearing down test containers")
	for _, containerMeta := range containerMetas {
		TeardownPostgresContainer(containerMeta)
	}
}
