// Copyright © 2014 Steve Francia <spf@spf13.com>.
// Copyright 2013 tsuru authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package afero provides types and methods for interacting with the filesystem,
// as an abstraction layer.

// Afero also provides a few implementations that are mostly interoperable. One that
// uses the operating system filesystem, one that uses memory to store files
// (cross platform) and an interface that should be implemented if you want to
// provide your own filesystem.

package afero

import (
	"errors"
	"io"
	"os"
	"time"
)

type Afero struct {
	Fs
}

// File 表示文件系统中的一个文件.
type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	io.WriterAt

	Name() string
	Readdir(count int) ([]os.FileInfo, error)
	Readdirnames(n int) ([]string, error)
	Stat() (os.FileInfo, error)
	Sync() error
	Truncate(size int64) error
	WriteString(s string) (ret int, err error)
}

// Fs 是一个文件系统接口.
//
// 所有模拟或真实文件系统都应该实现此接口
type Fs interface {

	// Create 在文件系统创建一个文件，返回文件和error
	Create(name string) (File, error)

	// Mkdir 在文件系统创建一个文件夹，返回error
	Mkdir(name string, perm os.FileMode) error

	// MkdirAll 创建一个目录地址，所有的父目录都尚不存在
	MkdirAll(path string, perm os.FileMode) error

	// Open 打开一个文件
	Open(name string) (File, error)

	// OpenFile 以给定的flag和mode打开文件.
	OpenFile(name string, flag int, perm os.FileMode) (File, error)

	// Remove 根据文件名称删除一个文件
	Remove(name string) error

	// RemoveAll 删除一个目录以及其中包含的子目录，如果不存在不会发生错误.
	RemoveAll(path string) error

	// Rename 重命名一个文件
	Rename(oldname, newname string) error

	// Stat 返回给定文件的描述信息
	Stat(name string) (os.FileInfo, error)

	// Name 文件系统的名称
	Name() string

	//Chmod 更改一个文件的mod.
	Chmod(name string, mode os.FileMode) error

	//Chtimes 更改给定文件的访问和修改时间
	Chtimes(name string, atime time.Time, mtime time.Time) error
}

var (
	ErrFileClosed        = errors.New("File is closed")
	ErrOutOfRange        = errors.New("Out of range")
	ErrTooLarge          = errors.New("Too large")
	ErrFileNotFound      = os.ErrNotExist
	ErrFileExists        = os.ErrExist
	ErrDestinationExists = os.ErrExist
)
