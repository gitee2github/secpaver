/*
 * Copyright (c) Huawei Technologies Co., Ltd. 2020-2021. All rights reserved.
 * secPaver is licensed under the Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *     http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
 * PURPOSE.
 * See the Mulan PSL v2 for more details.
 */

package project

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"io/ioutil"
	"path/filepath"
	"secpaver/api/proto/project"
	"secpaver/common/client"
	"secpaver/common/errdefs"
	"secpaver/common/global"
	"secpaver/common/utils"
)

func newExportCommand() *cli.Command {
	return &cli.Command{
		Name:        "export",
		Usage:       "Export project files",
		UsageText:   "pav project export [-f] <PROJECT> <PATH>",
		Description: "Export files in the project",
		Action:      exportAction,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:   "force, f",
				Usage:  "Force export",
				Hidden: false,
			},
		},
	}
}

func exportCheck(ctx *cli.Context) error {
	if err := utils.CheckCommandArgs(ctx, 2, utils.CheckExactArgs); err != nil {
		return err
	}

	dirExists, err := utils.DirExist(ctx.Args()[1])
	if err != nil {
		return err
	}
	if !dirExists {
		return errdefs.NewFileNotFoundError(ctx.Args()[1])
	}

	return nil
}

func exportAction(ctx *cli.Context) error {
	if err := exportCheck(ctx); err != nil {
		return err
	}

	c, err := client.NewClientFromContext(ctx)
	if err != nil {
		return err
	}
	defer c.Close()

	svc := project.NewProjectMgrClient(c.Connection())

	req := &project.Req{
		Name: ctx.Args().First(),
	}

	prjFile, err := svc.ExportProject(context.Background(), req)
	if err != nil {
		return err
	}

	filePath := filepath.Join(ctx.Args()[1], prjFile.GetFile().GetFilename())
	exists, err := utils.PathExist(filePath)
	if err != nil {
		return errors.Wrap(err, "fail to search output path")
	}

	if exists {
		if !ctx.Bool("force") {
			return errors.New("project file already exists")
		}
	}

	if err := ioutil.WriteFile(filePath, prjFile.GetFile().GetData(), global.DefaultFilePerm); err != nil {
		return errors.Wrap(err, "fail to write file")
	}

	fmt.Printf("Finish exporting: %s\n", filePath)

	return nil
}
