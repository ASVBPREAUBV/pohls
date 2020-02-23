# Pohls Media Manager

[![Build Status](https://travis-ci.com/ASVBPREAUBV/pohls.svg?branch=master)](https://travis-ci.com/ASVBPREAUBV/pohls)

CLI that converts any chaotic download folder into organised file structure:

## Goal Structure:

https://kodi.wiki/view/Naming_video_files

## Thanks to 

[parse-torrent-name](https://github.com/divijbindlish/parse-torrent-name) for Filename Parser Regex

## Project Layout

[Build with Standard Go Project Layout](https://github.com/golang-standards/project-layout)

## Development

### Concepts

A *file* is a single file on your computer with its content.

All the information about a content of a *file* should be in the *filename*.

From the *filename* pohls extracts all the media-contents information. (e.g. name, year...).
Additional info gets loaded through external apis. (e.g. [tmdb](https://www.themoviedb.org/))

The media information is saved in the *media* model.

From the *media* model we can generate the correct *filename* with the corresponding path.

### rough description

filename -> detox -> beautiful filename
beautiful filename -> extractor -> api -> media
media -> mediawriter -> filename



