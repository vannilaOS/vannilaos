package cli

/*	License: GPLv3
	Authors:
		Mirko Brombin <brombin94@gmail.com>
		Pietro di Caprio <pietro@fabricators.ltd>
		Vanilla OS Contributors <https://github.com/vanilla-os/>
	Copyright: 2024
	Description: Apx is a wrapper around multiple package managers to install packages and run commands inside a managed container.
*/

import (
	"github.com/vanilla-os/sdk/pkg/v1/app"
	"github.com/vanilla-os/sdk/pkg/v1/cli"
)

var Apx *app.App

type RootCmd struct {
	cli.Base
	Version string

	Stacks      StacksCmd      `cmd:"stacks" help:"pr:apx.cmd.stacks"`
	Subsystems  SubsystemsCmd  `cmd:"subsystems" help:"pr:apx.cmd.subsystems"`
	PkgManagers PkgManagersCmd `cmd:"pkgmanagers" help:"pr:apx.cmd.pkgmanagers"`

	DynamicSubsystems *map[string]*SubsystemCmd `cmd:"*" help:"apx.subsystem"`
}

type SubsystemCmd struct {
	cli.Base
	Name string `json:"-"`

	Enter      SubsystemEnterCmd      `cmd:"enter" help:"pr:apx.cmd.subsystem.enter"`
	Run        SubsystemRunCmd        `cmd:"run" help:"pr:apx.cmd.subsystem.run"`
	Install    SubsystemInstallCmd    `cmd:"install" help:"pr:apx.cmd.subsystem.install"`
	Remove     SubsystemRemoveCmd     `cmd:"remove" help:"pr:apx.cmd.subsystem.remove"`
	Update     SubsystemUpdateCmd     `cmd:"update" help:"pr:apx.cmd.subsystem.update"`
	Upgrade    SubsystemUpgradeCmd    `cmd:"upgrade" help:"pr:apx.cmd.subsystem.upgrade"`
	List       SubsystemListCmd       `cmd:"list" help:"pr:apx.cmd.subsystem.list"`
	Search     SubsystemSearchCmd     `cmd:"search" help:"pr:apx.cmd.subsystem.search"`
	Show       SubsystemShowCmd       `cmd:"show" help:"pr:apx.cmd.subsystem.show"`
	Export     SubsystemExportCmd     `cmd:"export" help:"pr:apx.cmd.subsystem.export"`
	Unexport   SubsystemUnexportCmd   `cmd:"unexport" help:"pr:apx.cmd.subsystem.unexport"`
	Start      SubsystemStartCmd      `cmd:"start" help:"pr:apx.cmd.subsystem.start"`
	Stop       SubsystemStopCmd       `cmd:"stop" help:"pr:apx.cmd.subsystem.stop"`
	AutoRemove SubsystemAutoRemoveCmd `cmd:"autoremove" help:"pr:apx.cmd.subsystem.autoremove"`
	Clean      SubsystemCleanCmd      `cmd:"clean" help:"pr:apx.cmd.subsystem.clean"`
	Purge      SubsystemPurgeCmd      `cmd:"purge" help:"pr:apx.cmd.subsystem.purge"`
}

type SubsystemEnterCmd struct {
	cli.Base
	Name string `json:"-"`
}

type SubsystemRunCmd struct {
	cli.Base
	Name string   `json:"-"`
	Args []string `arg:"" optional:"" name:"command" help:"pr:apx.arg.command"`
}

type SubsystemInstallCmd struct {
	cli.Base
	Name     string   `json:"-"`
	NoExport bool     `flag:"short:n, long:no-export, name:pr:apx.cmd.subsystem.install.options.noExport"`
	Args     []string `arg:"" optional:"" name:"packages" help:"pr:apx.arg.packages"`
}

type SubsystemRemoveCmd struct {
	cli.Base
	Name string   `json:"-"`
	Args []string `arg:"" optional:"" name:"packages" help:"pr:apx.arg.packages"`
}

type SubsystemUpdateCmd struct {
	cli.Base
	Name string `json:"-"`
}

type SubsystemUpgradeCmd struct {
	cli.Base
	Name string `json:"-"`
}

type SubsystemListCmd struct {
	cli.Base
	Name string `json:"-"`
}

type SubsystemSearchCmd struct {
	cli.Base
	Name string   `json:"-"`
	Args []string `arg:"" optional:"" name:"query" help:"pr:apx.arg.query"`
}

type SubsystemShowCmd struct {
	cli.Base
	Name string   `json:"-"`
	Args []string `arg:"" optional:"" name:"package" help:"pr:apx.arg.package"`
}

type SubsystemExportCmd struct {
	cli.Base
	Name string   `json:"-"`
	Args []string `arg:"" optional:"" name:"packages" help:"pr:apx.arg.packages"`
}

type SubsystemUnexportCmd struct {
	cli.Base
	Name string   `json:"-"`
	Args []string `arg:"" optional:"" name:"packages" help:"pr:apx.arg.packages"`
}

type SubsystemStartCmd struct {
	cli.Base
	Name string `json:"-"`
}

type SubsystemStopCmd struct {
	cli.Base
	Name string `json:"-"`
}

type SubsystemAutoRemoveCmd struct {
	cli.Base
	Name string `json:"-"`
}

type SubsystemCleanCmd struct {
	cli.Base
	Name string `json:"-"`
}

type SubsystemPurgeCmd struct {
	cli.Base
	Name string `json:"-"`
}

type StacksCmd struct {
	cli.Base

	List StacksListCmd `cmd:"list" help:"pr:apx.cmd.stacks.list"`
	Show StacksShowCmd `cmd:"show" help:"pr:apx.cmd.stacks.show"`
	New  StacksNewCmd  `cmd:"new" help:"pr:apx.cmd.stacks.new"`
}

type StacksListCmd struct {
	cli.Base
	Json bool `flag:"short:j, long:json, help:pr:apx.cmd.stacks.list.options.json"`
}

type StacksShowCmd struct {
	cli.Base
	Args []string `arg:"" optional:"" name:"stack" help:"pr:apx.arg.stack"`
}

type StacksNewCmd struct {
	cli.Base
	Name      string `flag:"short:n, long:name, help:pr:apx.cmd.stacks.new.options.name"`
	BaseImage string `flag:"short:b, long:base, help:pr:apx.cmd.stacks.new.options.baseImage"`
	NoPrompt  bool   `flag:"short:q, long:no-prompt, help:pr:apx.cmd.stacks.new.options.noPrompt"`
}

type SubsystemsCmd struct {
	cli.Base

	List StacksListCmd `cmd:"list" help:"pr:apx.cmd.subsystems.list"`
	New  SubsystemsNewCmd `cmd:"new" help:"pr:apx.cmd.subsystems.new"`
}

type SubsystemsListCmd struct {
	cli.Base
	Json bool `flag:"short:j, long:json, help:pr:apx.cmd.subsystems.list.options.json"`
}

type SubsystemsNewCmd struct {
	cli.Base
	Name  string `flag:"short:n, long:name, help:pr:apx.cmd.subsystems.new.options.name"`
	Stack string `flag:"short:s, long:stack, help:pr:apx.cmd.subsystems.new.options.stack"`
}

type PkgManagersCmd struct {
	cli.Base

	List   PkgManagersListCmd   `cmd:"list" help:"pr:apx.cmd.pkgmanagers.list"`
	Show   PkgManagersShowCmd   `cmd:"show" help:"pr:apx.cmd.pkgmanagers.show"`
	New    PkgManagersNewCmd    `cmd:"new" help:"pr:apx.cmd.pkgmanagers.new"`
	Export PkgManagersExportCmd `cmd:"export" help:"pr:apx.cmd.pkgmanagers.export"`
	Import PkgManagersImportCmd `cmd:"import" help:"pr:apx.cmd.pkgmanagers.import"`
}

type PkgManagersListCmd struct {
	cli.Base
	Json bool `flag:"short:j, long:json, help:pr:apx.cmd.pkgmanagers.list.options.json"`
}

type PkgManagersShowCmd struct {
	cli.Base
	Args []string `arg:"" optional:"" name:"pkgmanager" help:"pr:apx.arg.pkgmanager"`
}

type PkgManagersNewCmd struct {
	cli.Base
	Name           string `flag:"short:n, long:name, help:pr:apx.cmd.pkgmanagers.new.options.name"`
	Install        string `flag:"long:install, help:pr:apx.cmd.pkgmanagers.new.options.install"`
	Remove         string `flag:"long:remove, help:pr:apx.cmd.pkgmanagers.new.options.remove"`
	Search         string `flag:"long:search, help:pr:apx.cmd.pkgmanagers.new.options.search"`
	Show           string `flag:"long:show, help:pr:apx.cmd.pkgmanagers.new.options.show"`
	List           string `flag:"long:list, help:pr:apx.cmd.pkgmanagers.new.options.list"`
	Update         string `flag:"long:update, help:pr:apx.cmd.pkgmanagers.new.options.update"`
	Upgrade        string `flag:"long:upgrade, help:pr:apx.cmd.pkgmanagers.new.options.upgrade"`
	Clean          string `flag:"long:clean, help:pr:apx.cmd.pkgmanagers.new.options.clean"`
	AutoRemove     string `flag:"long:autoremove, help:pr:apx.cmd.pkgmanagers.new.options.autoremove"`
	Purge          string `flag:"long:purge, help:pr:apx.cmd.pkgmanagers.new.options.purge"`
	NeedSudo       bool   `flag:"long:need-sudo, help:pr:apx.cmd.pkgmanagers.new.options.needSudo"`
	NoPrompt       bool   `flag:"short:q, long:no-prompt, help:pr:apx.cmd.pkgmanagers.new.options.noPrompt"`
}

type PkgManagersExportCmd struct {
	cli.Base
	Name   string `flag:"short:n, long:name, help:pr:apx.cmd.pkgmanagers.export.options.name"`
	Output string `flag:"short:o, long:output, help:pr:apx.cmd.pkgmanagers.export.options.output"`
}

type PkgManagersImportCmd struct {
	cli.Base
	Input string `flag:"short:i, long:input, help:pr:apx.cmd.pkgmanagers.import.options.input"`
}
