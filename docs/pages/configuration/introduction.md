---
title: Introduction
sidebar: documentation
permalink: documentation/configuration/introduction.html
author: Alexey Igrychev <alexey.igrychev@flant.com>, Timofey Kirillov <timofey.kirillov@flant.com>
---

## What is werf config?

To use werf, you should configure an application accordingly. The configuration involves:

1. Defining project meta-information (such as a project name), which will affect build, deploy, and other commands.
2. Defining images to be built.

werf uses the YAML configuration file `werf.yaml` placed in the root folder of your application. The configuration data itself is a collection of configuration sections -- parts of YAML file separated by [three hyphens](http://yaml.org/spec/1.2/spec.html#id2800132):

```yaml
CONFIG_SECTION
---
CONFIG_SECTION
---
CONFIG_SECTION
```

Each configuration section has a type. There are currently 3 types of sections:

1. Config section to describe project meta information. We will refer to it as a *meta config section*.
2. Config section to describe instructions for building images. We will refer to it as an *image config section*.
3. Config section to describe instructions for building artifacts. We will refer to it as an *artifact config section*.

More types can be added in the future.

### Meta config section

```yaml
project: PROJECT_NAME
configVersion: CONFIG_VERSION
OTHER_FIELDS
---
```

The section with the key `project: PROJECT_NAME` and `configVersion: CONFIG_VERSION` is the meta config section. This section is obligatory. There should be only one meta config section in any `werf.yaml` configuration.

#### Project name

The `project` key defines the unique project name of your application. Project name affects build cache image names, Kubernetes Namespace, the Helm Release name, and other derived names (see [deploy to Kubernetes]({{ site.baseurl }}/documentation/reference/deploy_process/deploy_into_kubernetes.html)) for a detailed description. It is the only required field in meta configuration.

The project name should be unique within the group of projects which share build hosts and are deployed into the same Kubernetes cluster (i.e., unique across all groups within the single gitlab).

The project name can consist of up to 50 characters; only lowercase alphabet letters, numbers, and dashes are allowed.

**WARNING**. You should never change the name of the project once it has been set up (unless you know what you are doing).

Changing the name of the project leads to several issues:
1. Invalidation of a build cache. In this case, you will have to build new images and manually clean up old images from the localhost and the Docker registry.
2. Creation of the completely new Helm Release. Suppose that you have deployed an application successfully and later changed the name of the project. In this case, if you will try to deploy it again, the new instance of the same application will be created.

werf cannot cope with the change of the project name automatically. You have to resolve the issues described above manually.

#### Config version

The `configVersion` defines the format of a `werf.yaml` file. For now, it must be set to `1`.

### Image config section

Each image config section contains instructions for building one individual docker image. You can define several image config sections in the single `werf.yaml` configuration file to build multiple images.

An image config section starts with the key `image: IMAGE_NAME`. Here, `image` defines the short name of the docker image to be built. This name must be unique in a single `werf.yaml` config.

```yaml
image: IMAGE_NAME_1
OTHER_FIELDS
---
image: IMAGE_NAME_2
OTHER_FIELDS
---
...
---
image: IMAGE_NAME_N
OTHER_FIELDS
```

### Artifact config section

Artifact config section also defines instructions to build one independent artifact docker image. Arifact is a secondary image aimed to isolate a build process and build tools resources (environments, software, data, see [artifacts article for the details]({{ site.baseurl }}/documentation/configuration/stapel_artifact.html)). There may be multiple artifact config sections for multiple artifact config sections defined in the same `werf.yaml` config.

Config section with the key `artifact: IMAGE_NAME` is the artifact config section. `artifact` defines short name of the artifact to be referred to from another config sections. This name must be unique in a single `werf.yaml` config.

### Minimal configuration example

```yaml
project: my-project
configVersion: 1
```

## Organizing configuration

Part of the configuration can be moved in ***separate template files*** and then included into __werf.yaml__. _Template files_ should reside in the ***.werf*** directory with **.tmpl** extension (any nesting is supported).

> **Tip:** templates can be generated or downloaded before running werf. You can use them, for example, for sharing common logic between projects

werf parses all files in one environment, thus described [define](#include) of one _template file_ becomes available in other files, including _werf.yaml_.

<div class="details active">
<a href="javascript:void(0)" class="details__summary">werf.yaml</a>
<div class="details__content" markdown="1">

{% raw %}
```yaml
{{ $_ := set . "RubyVersion" "2.3.4" }}
{{ $_ := set . "BaseImage" "alpine" }}

project: my-project
configVersion: 1
---

image: rails
from: {{ .BaseImage }}
ansible:
  beforeInstall:
  {{- include "(component) mysql client" . }}
  {{- include "(component) ruby" . }}
  install:
  {{- include "(component) Gemfile dependencies" . }}
```
{% endraw %}

</div>
</div>

<div class="details">
<a href="javascript:void(0)" class="details__summary">.werf/ansible/components.tmpl</a>
<div class="details__content" markdown="1">

{% raw %}
```yaml
{{- define "(component) Gemfile dependencies" }}
  - file:
      path: /root/.ssh
      state: directory
      owner: root
      group: root
      recurse: yes
  - name: "Setup ssh known_hosts used in Gemfile"
    shell: |
      set -e
      ssh-keyscan github.com >> /root/.ssh/known_hosts
      ssh-keyscan mygitlab.myorg.com >> /root/.ssh/known_hosts
    args:
      executable: /bin/bash
  - name: "Install Gemfile dependencies with bundler"
    shell: |
      set -e
      source /etc/profile.d/rvm.sh
      cd /app
      bundle install --without development test --path vendor/bundle
    args:
      executable: /bin/bash
{{- end }}

{{- define "(component) mysql client" }}
  - name: "Install mysql client"
    apt:
      name: "{{`{{ item }}`}}"
      update_cache: yes
    with_items:
    - libmysqlclient-dev
    - mysql-client
    - g++
{{- end }}

{{- define "(component) ruby" }}
  - command: gpg --keyserver hkp://keys.gnupg.net --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3
  - get_url:
      url: https://raw.githubusercontent.com/rvm/rvm/master/binscripts/rvm-installer
      dest: /tmp/rvm-installer
  - name: "Install rvm"
    command: bash -e /tmp/rvm-installer
  - name: "Install ruby {{ .RubyVersion }}"
    raw: bash -lec {{`{{ item | quote }}`}}
    with_items:
    - rvm install {{ .RubyVersion }}
    - rvm use --default {{ .RubyVersion }}
    - gem install bundler --no-ri --no-rdoc
    - rvm cleanup all
{{- end }}
```
{% endraw %}

</div>
</div>

> If there are templates with the same name werf will use template defined in _werf.yaml_ or the latest described in _templates files_

If need to use the whole _template file_, use template file path relative to _.werf_ directory as a template name in [include](#include) function.

<div class="details active">
<a href="javascript:void(0)" class="details__summary">werf.yaml</a>
<div class="details__content" markdown="1">

{% raw %}
```yaml
project: my-project
configVersion: 1
---
image: app
from: java:8-jdk-alpine
shell:
  beforeInstall:
  - mkdir /app
  - adduser -Dh /home/gordon gordon
import:
- artifact: appserver
  add: '/usr/src/atsea/target/AtSea-0.0.1-SNAPSHOT.jar'
  to: '/app/AtSea-0.0.1-SNAPSHOT.jar'
  after: install
- artifact: storefront
  add: /usr/src/atsea/app/react-app/build
  to: /static
  after: install
docker:
  ENTRYPOINT: ["java", "-jar", "/app/AtSea-0.0.1-SNAPSHOT.jar"]
  CMD: ["--spring.profiles.active=postgres"]
---
{{ include "artifact/appserver.tmpl" . }}
---
{{ include "artifact/storefront.tmpl" . }}
```
{% endraw %}

</div>
</div>

<div class="details">
<a href="javascript:void(0)" class="details__summary">.werf/artifact/appserver.tmpl</a>
<div class="details__content" markdown="1">

{% raw %}
```yaml
artifact: appserver
from: maven:latest
git:
- add: '/app'
  to: '/usr/src/atsea'
shell:
  install:
  - cd /usr/src/atsea
  - mvn -B -f pom.xml -s /usr/share/maven/ref/settings-docker.xml dependency:resolve
  - mvn -B -s /usr/share/maven/ref/settings-docker.xml package -DskipTests
```
{% endraw %}

</div>
</div>

<div class="details">
<a href="javascript:void(0)" class="details__summary">.werf/artifact/storefront.tmpl</a>
<div class="details__content" markdown="1">

{% raw %}
```yaml
artifact: storefront
from: node:latest
git:
- add: /app/react-app
  to: /usr/src/atsea/app/react-app
shell:
  install:
  - cd /usr/src/atsea/app/react-app
  - npm install
  - npm run build
```
{% endraw %}

</div>
</div>

## Processing of config

The following steps could describe the processing of a YAML configuration file:
1. Reading `werf.yaml` and extra templates from `.werf` directory.
2. Executing Go templates.
3. Saving dump into `.werf.render.yaml` (this file remains after the command execution and will be removed automatically with GC procedure).
4. Splitting rendered YAML file into separate config sections (part of YAML stream separated by three hyphens, https://yaml.org/spec/1.2/spec.html#id2800132).
5. Validating each config section:
  * Validating YAML syntax (you could read YAML reference [here](http://yaml.org/refcard.html)).
  * Validating werf syntax.
6. Generating a set of images.

### Go templates

Go templates are available within YAML configuration. The following functions are supported:

* [Built-in Go template functions](https://golang.org/pkg/text/template/#hdr-Functions) and other language features. E.g. using common variable:<a id="go-templates" href="#go-templates" class="anchorjs-link " aria-label="Anchor link for: go templates" data-anchorjs-icon=""></a>

  {% raw %}
  ```yaml
  {{ $base_image := "golang:1.11-alpine" }}

  project: my-project
  configVersion: 1
  ---

  image: gogomonia
  from: {{ $base_image }}
  ---
  image: saml-authenticator
  from: {{ $base_image }}
  ```
  {% endraw %}

* [Sprig functions](http://masterminds.github.io/sprig/). E.g. using environment variable:<a id="sprig-functions" href="#sprig-functions" class="anchorjs-link " aria-label="Anchor link for: sprig functions" data-anchorjs-icon=""></a>

  {% raw %}
  ```yaml
  project: my-project
  configVersion: 1
  ---

  {{ $_ := env "SPECIFIC_ENV_HERE" | set . "GitBranch" }}

  image: ~
  from: alpine
  git:
  - url: https://github.com/company/project1.git
    branch: {{ .GitBranch }}
    add: /
    to: /app/project1
  - url: https://github.com/company/project2.git
    branch: {{ .GitBranch }}
    add: /
    to: /app/project2
  ```
  {% endraw %}

* `include` function with `define` for reusing configs:<a id="include" href="#include" class="anchorjs-link " aria-label="Anchor link for: include" data-anchorjs-icon=""></a>

  {% raw %}
  ```yaml
  project: my-project
  configVersion: 1
  ---

  image: app1
  from: alpine
  ansible:
    beforeInstall:
    {{- include "(component) ruby" . }}
  ---
  image: app2
  from: alpine
  ansible:
    beforeInstall:
    {{- include "(component) ruby" . }}

  {{- define "(component) ruby" }}
    - command: gpg --keyserver hkp://keys.gnupg.net --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3
    - get_url:
        url: https://raw.githubusercontent.com/rvm/rvm/master/binscripts/rvm-installer
        dest: /tmp/rvm-installer
    - name: "Install rvm"
      command: bash -e /tmp/rvm-installer
    - name: "Install ruby 2.3.4"
      raw: bash -lec {{`{{ item | quote }}`}}
      with_items:
      - rvm install 2.3.4
      - rvm use --default 2.3.4
      - gem install bundler --no-ri --no-rdoc
      - rvm cleanup all
  {{- end }}
  ```
  {% endraw %}

* `.Files.Get` function for getting project file content:<a id="files-get" href="#files-get" class="anchorjs-link " aria-label="Anchor link for: .Files.Get" data-anchorjs-icon=""></a>

  <div class="tabs">
    <a href="javascript:void(0)" class="tabs__btn active" onclick="openTab(event, 'tabs__btn', 'tabs__content', 'ansible')">Ansible</a>
    <a href="javascript:void(0)" class="tabs__btn" onclick="openTab(event, 'tabs__btn', 'tabs__content', 'shell')">Shell</a>
  </div>
  
  <div id="ansible" class="tabs__content active" markdown="1">
  {% raw %}
  ```yaml
  project: my-project
  configVersion: 1
  ---

  image: app
  from: alpine
  ansible:
    setup:
    - name: "Setup /etc/nginx/nginx.conf"
      copy:
        content: |
  {{ .Files.Get ".werf/nginx.conf" | indent 8 }}
        dest: /etc/nginx/nginx.conf
  ```
  {% endraw %}
  </div>
  
  <div id="shell" class="tabs__content" markdown="1">
  {% raw %}
  ```yaml
  project: my-project
  configVersion: 1
  ---
  
  image: app
  from: alpine
  shell:
    setup:
    - |
      head -c -1 <<EOF | tee file > /etc/nginx/nginx.conf
  {{ .Files.Get ".werf/nginx.conf" | indent 4 }}
      EOF
  ```
  {% endraw %}
  </div>
  
