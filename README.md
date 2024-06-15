# Notion to Presentation

[n2p.dev](https://n2p.dev) - convert Notion page to presentation.

1. Get data from Notion Web API
1. Convert to HTML in a format that can be used in reveal.js

# Stack

1. Go
1. Pocketbase
1. Templ
1. HTMX
1. Vue
1. Tailwind
1. daisyUI

# Word of caution

This project is NOT about best practices. It's about making product
and do it efficiently. I haven't been working with Vue for a long time,
and this is first time for me to use Pocketbase, Templ and HTMX.

Don't take this project as a reference for best practices.

# Project structure

1. `cmd/cli` - cli tool to convert Notion page to presentation
1. `cmd/saas` - N2P.dev SaaS server
1. `infra` - some infrastructure code
1. `libs` - libraries

# Roadmap

## User Features

1. Presentation notes
1. Public presentation description
1. Author public page
    1. Profile
        1. Picture
        1. Name
        1. Info
1. Presentations config
    1. Images customizations
    1. Background image
1. Features requested page

## UX improvements

1. Color picker
1. Add FAQ on how to create presentation
    1. Divider
    1. Supported elements
1. Main page error
1. Mobile fixes
1. Add reveal.js to `/app`
1. Make video
1. Adapt to mobile
1. Presentations pagination
1. Search in presentations
1. Island-based links
1. Theme editor
1. Dark theme
1. Toast

## Tech

1. Check cache for static files (JS, CSS)
1. Compile JS and CSS
1. Remove HTMX
1. Change CDN to local files
1. CI/CD