# dashboard
- Weather Information 
- Song Facts for currently playing song (https://kjhk.org/api/now-playing.php)
- Form for manual music logging (I will send you API route when you get here)
- Integration with Google Calendar to display upcoming football games (no calendar atm)
- Upcoming meeting feed (no calendar atm)
- Dynamic announcements (need a way for manager to update this) 

There are a few other features but this is the foundation.

There should be basic authentication (user/pass). There should be two accounts (DJ/Admin) 

Admin would just dynamically render a few more things like the form to update announcements etc.

For the design, I'm imagining a widget style dashboard kind of like https://bento.me/ each "feature" has it's own small dedicated widget. This will make it so we don't have to follow a traditional dashboard layout and if we want to add a new feature, it's as easy as adding a widget/block

If there's any DB needed, you can use vercel pg or vercel kv
Use Next.js 13 and take advantage of SSR whenever possible