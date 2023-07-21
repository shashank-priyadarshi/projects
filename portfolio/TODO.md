1. PHASE-1:
- Merge actions and charts, charts open in new tab instead of in page rendering
- Folder structure and expandable folder structure from GitHub repository
- In memory caching service to cache folder structure, auto refreshes every hour, endpoint for manual refresh
- File names and render files hosted on GitHub when file clicked upon
- Make MyFile repo private, server fetches pdf using GitHub token and stream back to client
- If file is pdf, pdf rendering in book format
2. PHASE-2:
- Temporary tokens
- Access levels for temporary tokens
- Duration for temporary tokens
3. PHASE-3:
- If frequent API access request, block IP for 2 months
- If IP blocked more than once within 6 months, block IP forever
- Allow IP unblock request through email, request service of admin dashboard and email updates for request status
4. PHASE-4:
- Allow marking files as accessible/inaccessible based on token access level: general, private
- Access levels for tokens for each feature on the whole dashboard
- Caching service caches files based on number of requests
5. PHASE-5:
- Allow book access for users from particular location
- Allow location addition using admin dashboard
