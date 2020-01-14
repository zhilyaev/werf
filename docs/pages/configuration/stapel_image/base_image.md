---
title: Base image
sidebar: documentation
permalink: documentation/configuration/stapel_image/base_image.html
author: Alexey Igrychev <alexey.igrychev@flant.com>
summary: |
  <a class="google-drawings" href="https://docs.google.com/drawings/d/e/2PACX-1vReDSY8s7mMtxuxwDTwtPLFYjEXePaoIB-XbEZcunJGNEHrLbrb9aFxyOoj_WeQe0XKQVhq7RWnG3Eq/pub?w=2031&amp;h=144" data-featherlight="image">
      <img src="https://docs.google.com/drawings/d/e/2PACX-1vReDSY8s7mMtxuxwDTwtPLFYjEXePaoIB-XbEZcunJGNEHrLbrb9aFxyOoj_WeQe0XKQVhq7RWnG3Eq/pub?w=1016&amp;h=72">
  </a>

  <div class="language-yaml highlighter-rouge"><div class="highlight"><pre class="highlight"><code><span class="na">from</span><span class="pi">:</span> <span class="s">&lt;image[:&lt;tag&gt;]&gt;</span>
  <span class="na">fromLatest</span><span class="pi">:</span> <span class="s">&lt;bool&gt;</span>
  <span class="na">fromCacheVersion</span><span class="pi">:</span> <span class="s">&lt;arbitrary string&gt;</span>
  <span class="na">fromImage</span><span class="pi">:</span> <span class="s">&lt;image name&gt;</span>
  <span class="na">fromImageArtifact</span><span class="pi">:</span> <span class="s">&lt;artifact name&gt;</span>
  </code></pre></div>
  </div>
---

Here's an example of the simplest `werf.yaml` file. It describes an _image_ named `example` that is based on the _base image_ named `alpine`:

```yaml
project: my-project
configVersion: 1
---
image: example
from: alpine
```

_Base image_ can be declared by `from`, `fromImage` or `fromImageArtifact` directives.

## from, fromLatest

The `from` directive defines the name and tag of a _base image_. If absent, the tag defaults to `latest`.

```yaml
from: <image>[:<tag>]
```

By default, an assembly process does not depend on the actual digest of the _base image_ in the repository. It only depends on the value of the _from_ directive.
Therefore, changing the _base image_ in the local storage or the Docker registry will not affect the build process if the _from_ stage already exists in _stages storage_.

If you prefer to build an image with the current _base image_ only, then you should use the _fromLatest_ directive.
The _fromLatest_ directive connects an assembly process with a digest of a _base image_, getting it from the repository in the process.
```yaml
fromLatest: true
```

> Pay attention, werf uses actual _base image_ digest as extra _from_ stage dependency if _fromLatest_ is true.
Therefore, using this directive implies not reproducible signatures:
after changing _base image_ in repository, all previously built stages, also like related images, become not usable.
The problem might occur:
- between jobs of one pipeline (e.g. build and deploy) or
- when you re-run the previous job (e.g. deploy)

## fromImage and fromImageArtifact

Besides using docker image from a repository, the _base image_ can refer to _image_ or [_artifact_]({{ site.baseurl }}/documentation/configuration/stapel_artifact.html), that is described in the same `werf.yaml`.

```yaml
fromImage: <image name>
fromImageArtifact: <artifact name>
```

If a _base image_ is specific to a particular application,
it is reasonable to store its description with _images_ and _artifacts_ which are used it as opposed to storing the _base image_ in a Docker registry.

Also, this method can be useful if the stages of _stage conveyor_ are not enough for building the image. You can design your _stage conveyor_.

<a class="google-drawings" href="https://docs.google.com/drawings/d/e/2PACX-1vTmQBPjB6p_LUpwiae09d_Jp0JoS6koTTbCwKXfBBAYne9KCOx2CvcM6DuD9pnopdeHF--LPpxJJFhB/pub?w=1629&amp;h=1435" data-featherlight="image">
<img src="https://docs.google.com/drawings/d/e/2PACX-1vTmQBPjB6p_LUpwiae09d_Jp0JoS6koTTbCwKXfBBAYne9KCOx2CvcM6DuD9pnopdeHF--LPpxJJFhB/pub?w=850&amp;h=673">
</a>

## fromCacheVersion

The `fromCacheVersion` directive allows to manage image reassembly.

```yaml
fromCacheVersion: <arbitrary string>
```
