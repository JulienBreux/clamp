# 🗜 Clamp

[![Actions Status](https://github.com/JulienBreux/clamp/workflows/Build%20and%20test%20Go/badge.svg)](https://github.com/JulienBreux/clamp/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/JulienBreux/clamp)](https://goreportcard.com/report/github.com/JulienBreux/clamp)
[![codebeat badge](https://codebeat.co/badges/2730cae5-0cf6-43f7-a04d-2163b5902013)](https://codebeat.co/projects/github-com-julienbreux-clamp-master)
[![GitHub tag](https://img.shields.io/github/tag/JulienBreux/clamp.svg)](Tag)
[![Go version](https://img.shields.io/github/go-mod/go-version/JulienBreux/clamp)](https://golang.org/dl/#stable)

Clamp is a useful tool to help to replace environment variables in any file using go template syntax.

---

## 🔧 Installation

Clamp is available on Linux, OSX and Windows platforms.

* Binaries for Mac OS, Linux and Windows are available as tarballs in the [release](https://github.com/JulienBreux/clamp/releases) page.

* Via Homebrew (Mac OS) or LinuxBrew (Linux)

   ```shell
   brew tap JulienBreux/clamp
   brew install clamp
   ```

* Building from source
   Clamp was built using go 1.12 or above. In order to build Clamp from source you must:
   1. Clone this repository
   2. Add the following command in your go.mod file

      ```text
      replace (
        github.com/JulienBreux/clamp => CLONED_GIT_REPOSITORY
      )
      ```

   3. Build and run the executable

        ```shell
        go run main.go
        ```

   4. Use it

        ```shell
        ./clamp
        ```

---

## 🌎 Availability

* [NPM](https://www.npmjs.com/package/@julienbreux/clamp)

---

## 📘 Help

### How to use from pipe

```bash
echo "{{ .USER }}" | clamp
# JulienBreux
```

### How to use from file

```bash
echo "{{ .HOME }}" > home.txt
clamp home.txt
# /Users/julienbreux
```

---

## 📮 Contact Info

1. **Email**:   julien.breux@gmail.com
2. **GitHub**:  [@JulienBreux](https://github.com/JulienBreux)
3. **Twitter**: [@JulienBreux](https://twitter.com/JulienBreux)

---

## 👮‍♂️ Security info

### GPG Signature

You can download Julien Breux's public key to verify the signature.

```shell
gpg --keyserver hkps://hkps.pool.sks-keyservers.net --recv-keys 0BD023FA
```
