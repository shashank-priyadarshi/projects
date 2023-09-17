# Portfolio application ToDos

## PHASE-1: IN PROGRESS

    - Dev setup for portfolio codebase: DONE
    - Merge actions and charts: DONE
    - Remove todos: DONE
    - Revamp home page, make it static: DONE
    - Remove todos and biodata endpoints: DONE
    - Disable GitHub activity graph: DONE
    - Add [Upwork](https://www.upwork.com/freelancers/~01812e02099fc1c28c) and [Fiverr](https://www.fiverr.com/s/yBzv5A): DONE
    - Remove obsolete CSS: IN PROGRESS
        app.component.sass
        login.component.sass
        actions.component.sass
        admin.component.sass
        games.component.sass
        list.component.sass
        time.component.sass
        articles.component.sass
        about.component.sass
        home.component.sass
        projects.component.sass
        policy.component.sass
        menu.component.sass
    - Fully responsive portfolio page: TODO
    - Cleanup dependencies in Angular application: TODO
    - Charts open as pop-ups on hover instead of in page rendering: TODO
    - Setup script for configuring local setup, and running frontend and backend services: TODO
    - Configure development and production env Dockerfiles for portfolio: TODO
      profiles: TODO

## PHASE-2: TODO

    - Folder structure and expandable folder structure from GitHub repository
    - In memory caching service to cache folder structure, auto refreshes every hour, endpoint for manual refresh
    - File names and render files hosted on GitHub when file clicked upon
    - Make MyFile repo private, server fetches pdf using GitHub token and stream back to client
    - If file is pdf, pdf rendering in book format

## PHASE-3: TODO

    - Temporary tokens
    - Access levels for temporary tokens
    - Duration for temporary tokens

## PHASE-4: TODO

    - If frequent API access request, block IP for 2 months
    - If IP blocked more than once within 6 months, block IP forever
    - Allow IP unblock request through email, request service of admin dashboard and email updates for request status

## PHASE-5: TODO

    - Allow marking files as accessible/inaccessible based on token access level: general, private
    - Access levels for tokens for each feature on the whole dashboard
    - Caching service caches files based on number of requests

## PHASE-6: TODO

    - Allow book access for users from particular location
    - Allow location addition using admin dashboard
    - Enable light/dark mode toggline: TODO
