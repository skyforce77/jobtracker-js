# jobtracker-js

[![go report](https://goreportcard.com/badge/github.com/skyforce77/jobtracker-js)](https://goreportcard.com/report/github.com/skyforce77/jobtracker-js)

JobTracker aims to help you find your future dream job. You can use our library to scrap and export jobs from 150+ providers.

### Compile

```sh
cd jobtracker-js
go get -u github.com/gopherjs/gopherjs
go get -u github.com/IDerr/jobtracker/providers
gopherjs build -m
```

If you are using NodeJS, you'll also need xhr2
```sh
npm install xhr2
```

### Example

```js
jt = require("./jobtracker-js.js")

providers = jt.getProviders()
for (var i = 0; i < providers.length; i++) {
  var provider = providers[i];
  provider.retrieveJobs(function(job) {
    console.log(job);
  });
}
```

Or by selecting a provider

```js
jt = require("./jobtracker-js.js")

jt.newWorkday().retrieveJobs(console.log)
```

### Providers and Companies

[List available right here](https://github.com/IDerr/jobtracker#providers-and-companies)
