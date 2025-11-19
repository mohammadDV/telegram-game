package testhelper

import (
	"os"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/sirupsen/logrus"
)

func Integration() bool {
	return os.Getenv("TEST_INTEGRATION") == "true"
}

func StartDockerPool() *dockertest.Pool {
	pool, err := dockertest.NewPool("")
	if err != nil {
		logrus.WithError(err).Fatalln("Could not construct pool")
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		logrus.WithError(err).Fatalln("Could not connect to Docker")
	}

	return pool
}

func StartDockerInstance(pool *dockertest.Pool, image , tag string, env ...string) *dockertest.Resource {

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: image,
		Tag: tag,
		Env: env,
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	}) 

	if err != nil {
		logrus.WithError(err).Fatalln("Could not start resource")
	}

	if err := resource.Expire(120); err != nil {
		logrus.WithError(err).Fatalln("Could not expire resource")
	}

	return resource

}