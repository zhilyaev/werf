---
title: Releases
permalink: releases.html
sidebar: documentation
layout: default
---

{%- asset releases.css %}

{%- assign releases = site.data.releases.releases %}

<div class="page__container page_releases">

<div class="releases__block-title">
    Release channels
</div>

<!-- Releases description -->
<div class="releases__info">
    Each werf release progresses through all release channels, starting with Alpha → Beta → Early-Access → Stable → Rock-Solid. You can think of each release on a lower channel as a release-candidate for the higher one. Once a release is considered bug-free, it is promoted to the next channel.
</div>

{%- assign groups = site.data.releases_history.history | map: "group" | uniq %}
{%- assign channels_sorted = site.data.channels_info.channels | sort: "stability" %}
{%- assign channels_sorted_reverse = site.data.channels_info.channels | sort: "stability" | reverse  %}

<div class="releases__menu">
{%- for channel in channels_sorted_reverse %}
{%- assign channel_latest_versions = site.data.releases_history.latest | where: "name",  channel.name | first| map: "versions" | first| default: nil %}
    <div class="releases__menu-item">
        <div class="releases__menu-item-header">            
            <div class="releases__menu-item-title">
                {{ channel.title }}
            </div>
            <div class="releases__menu-item-versions">
            {%- for version in channel_latest_versions %}
            {%- if version != nil  %}
            {%- assign version_info = site.data.releases.releases | where: "tag_name", version | first %}
                <a href="{{ version_info.html_url }}" class="releases__btn">
                {{ version }}
                </a>
            {%- endif %}
            {%- endfor %}
            </div>
        </div>        
        <div class="releases__menu-item-description">
            {{ channel.description[page.lang] }}
        </div>
    </div>
{%- endfor %}
</div>

<div class="releases__block-title">Releases: 
    {%- for group in groups %}
    <a href="javascript:void(0)" class="tabs__btn tabs__group__btn{% if group == groups[0] %} active{% endif %}" onclick="openTab(event, 'tabs__group__btn', 'tabs__group__content', 'group-{{group}}')">{{group}}</a>
    {%- endfor %}
</div>

{%- for group in groups %}

<div id="group-{{group}}" class="releases tabs__group__content{% if group == groups[0] %} active{% endif %}">
    <div class="tabs">
    {%- for channel in channels_sorted_reverse %}
    <a href="javascript:void(0)" class="tabs__btn tabs__{{group}}__channel__btn{% if channel == channels_sorted_reverse[1] %} active{% endif %}" onclick="openTab(event, 'tabs__{{group}}__channel__btn', 'tabs__{{group}}__channel__content', 'id-{{group}}-{{channel.name}}')">{{channel.title}}</a>
    {%- endfor %}
    <a href="javascript:void(0)" class="tabs__btn tabs__{{group}}__channel__btn" onclick="openTab(event, 'tabs__{{group}}__channel__btn', 'tabs__{{group}}__channel__content', 'id-{{group}}-all')">All channels</a>
    </div>

    {%- for channel in channels_sorted_reverse %}
    <div id="id-{{group}}-{{ channel.name }}" class="tabs__content tabs__{{group}}__channel__content{% if channel == channels_sorted_reverse[1] %} active{% endif %}">
    <div class="releases__info">
        <p>{{ channel.tooltip[page.lang] }}</p>
        <p class="releases__info-text">{{ channel.description[page.lang] }}</p>
    </div>

    {%- assign group_history = site.data.releases_history.history | reverse | where: "group", group %}
    {%- assign channel_history = group_history | where: "name", channel.name %}

    {%- if channel_history.size > 0 %}
        {%- for channel_action in channel_history %}
           {%- assign release = site.data.releases.releases | where: "tag_name", channel_action.version | first %}
            <div class="releases__title">
                <a href="{{ release.html_url }}">
                    {{ release.tag_name }}
                </a>
            </div>
            <div class="releases__body">
                {{ release.body | markdownify }}
            </div>
        {%- endfor %}
    {%- else %}
        <div class="releases__info releases__info_notification">
            <p>There are no versions on the channel yet, but they will appear soon.</p>
        </div>
    {%- endif %}

    </div>
    {%- endfor %}

    <div id="id-{{group}}-all" class="tabs__content tabs__{{group}}__channel__content">
        <div class="releases__info">
            <p>This is a list of all of the releases (Alpha, Beta, Early-Access, Stable and Rock-Solid) combined in chronological order.</p>
        </div>
    {%- for release_data in group_history %}
            {%- assign release = site.data.releases.releases | where: "tag_name", release_data.version | first %}
            <div class="releases__title">
                <a href="{{ release.html_url }}">
                    {{ release.tag_name }}
                </a>
            </div>
            <div class="releases__body">
                {{ release.body | markdownify }}
            </div>
    {%- endfor %}
    </div>
</div>
{%- endfor %}
