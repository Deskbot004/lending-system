# TYPO3 Projects Dashboard

The typo3 projects dashboard is a web application which shows the multiple informations about given [Gitlab](https://gitlab.itx.de) Projects.

## Requirements

To successfully run the dashboard from serverside three environmental variables should be set:
- **"PRIVATETOKEN"** : (Required) Token required for each get request, as such needed to validate read rights
- **"UPDATETIMER"** : (Optional) Timer for the db so update by itself should be in a format like "5m" for 5 minutes
- **"RESETTIMER"** : (Optional) Timer for the db to reset and redeploy with previous projects by itself should be in a format like "5m" for 5 minutes

## Features
- **Dashboard**: Landing page with general overview of the number of Projects/Packages and the current typo3 news feed.
- **Project Overview**: Projects with [Conluence](https://itx-online.atlassian.net/wiki/home) linkings, instances, notes and more.
- Instances of each Project with links to the frontend, backend and deployment.
- **Package Versions**: Packages of each Project containing version informations constraints and the current update satus in project.
- Each Package has linkings to its [Packagist](https://packagist.org/) and [Typo3](https://extensions.typo3.org/) if available.
- **Commit Tracker**: A Tracker for a project that compares the Production and Test instance for commits containing a Ticket
- Each Commit has its date, the author and a link to the [TRON Ticket](https://tron.itx.de).
- **Deployments**: The most recent 10 Deployments of every Project listed, sorted by the deploy date.
- Each Deployment shows its connected merge requests with author and date information.

---

## General Usage
After deploying simply use the "Add Project" Button to add a project into the database.
As such the **correct Project id** needs to be input from [Gitlab](https://gitlab.itx.de).
After waiting for the database to update, the basic functionality (Packages/Deployments) is given.

Below a short list of already tested and noted projects:
| GitlabID | Projectname |
| ------ | ------ |
| 198 | Bwegt |
| 125| Telegärtner|
| 153| BoshEvamo|
| 143| Bildungslandkarte|
| 95| Müller|

## Further Usage
A Project can be deleted or further configured from "Project overview".
Here the User can create multiple instances, input various links(confluence/Frontend/Backend/Deployment) and add further information for the project.
To enable the commit tracker for a project a few requirements must be given:
- At least two instances for the project
- Exactly one instance classified Production
- Exactly one instance classified Test
- The correct branchnames of two different branches in Production and Test

If all this is input the whole functionality is given

---

## Credit
The first Version project was created by Marvin Deichsel with supervision from Benjamin Jasper in context of a month long internship at it.x (07.2023/08.2023).
> I want to thank it.x for giving me the opportunity of an internship and Benjamin Jasper for the great support!
> The time i had here was very insightful and fun, as the whole team of it.x was simply great ^^

