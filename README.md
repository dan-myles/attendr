# GO Class Scraper
A simple implementation of a web scraper that will find available seats
at any university that exposes a public api. While I have already written this in rust,
I thought it would be good to get an idea of how well it would perform in GO.

Here is a mock chart of the final implementation. While I do not plan to open source
the final implementation, this can give you a rough idea of what it would be.
```mermaid
flowchart TD
A[User creates an account] --> B[User adds a class to watchlist]
B -- HTTP POST --> C{Backend Recieves Request}
C -- Go Routine --> D[Added to URL map]
C -- Go Routine --> F[Added to long term DB storage]
F --> Z
D --> X[Grab Snapshot of URL MAP]
X --> G[Scrape class search]
G --> E(Class has open seats?)
E -- no --> I(Has it been 24HR since last update?)
I -- no --> K[Do nothing and sleep for SLEEP_TIME]
K -- repeat --> X
I -- yes --> L[Sync with database]
L --> Z
E -- yes --> H[Remove from URL map]
H --> J[Send text/email]
J -- Remove from DB of URLS --> Z{Database}
```

Obviously this chart does not account for all cases such as
the class code already being in the database and it already
being scraped. However this should provide a good high-level overview.

