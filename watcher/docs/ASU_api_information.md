## ASU Information

ASU likes to organize classes into terms, which all have unique codes.
Term codes are not *super* trivial to get, but are not too hard either.

Known Term Codes:
* Summer 2023: "2234"
* Fall   2023: "2237"

These could change, but it seems like there is not much changing going on there.
The second piece of information that is required to check a classes availability,
is the class number. These can be easily scraped off ASU's website when adding
a class to the *watchlist*.

Once we have a term code and class number, we can search the class via this
endpoint:

```
"https://eadvs-cscc-catalog-api.apps.asu.edu/catalog-microservices/api/v1/search/...
classes?&refine=Y&advanced=true&classNbr=70473&searchType=all&term=2237"
```
Refine, advanced, and searchType are all required, and those values needn't change.
However building a URL builder for such a query isn't a bad idea.

