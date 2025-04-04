package integration

import (
	"context"
	"os"
	"path/filepath"

	tc "github.com/testcontainers/testcontainers-go/modules/compose"
)

func StartApplication() tc.ComposeStack {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	for range 3 {
		wd = filepath.Dir(wd)
	}
	dockerComposeFilePath := filepath.Join(wd, "development/docker-compose.yaml")
	compose, err := tc.NewDockerCompose(dockerComposeFilePath)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	err = compose.Up(ctx, tc.Wait(true))
	if err != nil {
		panic(err)
	}
	return compose
}

func StopApplication(compose tc.ComposeStack) {
	ctx := context.Background()
	err := compose.Down(ctx, tc.RemoveOrphans(true), tc.RemoveImagesLocal, tc.RemoveVolumes(true))
	if err != nil {
		panic(err)
	}
}
