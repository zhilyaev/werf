---
title: Dockerfile Image
sidebar: documentation
permalink: documentation/configuration/dockerfile_image.html
author: Alexey Igrychev <alexey.igrychev@flant.com>
summary: |
  <a class="google-drawings" href="https://docs.google.com/drawings/d/e/2PACX-1vRrzxht-PmC-4NKq95DtLS9E7JrvtuHy0JpMKdylzlZtEZ5m7bJwEMJ6rXTLevFosWZXmi9t3rDVaPB/pub?w=2031&amp;h=144" data-featherlight="image">
    <img src="https://docs.google.com/drawings/d/e/2PACX-1vRrzxht-PmC-4NKq95DtLS9E7JrvtuHy0JpMKdylzlZtEZ5m7bJwEMJ6rXTLevFosWZXmi9t3rDVaPB/pub?w=1016&amp;h=72">
  </a>

  <div class="language-yaml highlighter-rouge"><div class="highlight"><pre class="highlight"><code><span class="na">image</span><span class="pi">:</span> <span class="s">&lt;image name... || ~&gt;</span>
  <span class="na">dockerfile</span><span class="pi">:</span> <span class="s">&lt;relative path&gt;</span>
  <span class="na">context</span><span class="pi">:</span> <span class="s">&lt;relative path&gt;</span>
  <span class="na">target</span><span class="pi">:</span> <span class="s">&lt;docker stage name&gt;</span>
  <span class="na">args</span><span class="pi">:</span>
    <span class="s">&lt;build arg name&gt;</span><span class="pi">:</span> <span class="s">&lt;value&gt;</span>
  <span class="na">addHost</span><span class="pi">:</span>
  <span class="pi">-</span> <span class="s">&lt;host:ip&gt;</span>
  </code></pre></div></div>
---

Building an image using the existing Dockerfile is the easiest way to start using werf in an existing project.
Presented below is an elementary `werf.yaml` configuration file. It describes an image named `example` that relates to a project called `Dockerfile`:

```yaml
project: my-project
configVersion: 1
---
image: example
dockerfile: Dockerfile
```

Here is how to specify several images from one Dockerfile:

```yaml
image: backend
dockerfile: Dockerfile
target: backend
---
image: frontend
dockerfile: Dockerfile
target: frontend
```

The same as above but with different Dockerfiles:

```yaml
image: backend
dockerfile: dockerfiles/DockerfileBackend
---
image: frontend
dockerfile: dockerfiles/DockerfileFrontend
```

## Naming

{% include image_configuration/naming.md %}

## Dockerfile directives

Similarly to Docker, werf builds an image using the Dockerfile and context.

- `dockerfile` **(required)**: set the Dockerfile path relative to the project directory.
- `context`: set the build context PATH inside the project directory (the root of a project - `.` - is selected by default).
- `target`: link specific Dockerfile stage (by default, the most recent one is used, see the `docker build` \-\-target option).
- `args`: set build-time variables (see the `docker build` \-\-build-arg option).
- `addHost`: add a custom host-to-IP mapping (host:ip) (see the `docker build` \-\-add-host option).
