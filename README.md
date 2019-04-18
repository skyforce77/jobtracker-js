# jobtracker-js

Jobtracker contains a lot of different sources to find jobs directly on companies websites.

### Compile

```sh
cd jobtracker-js
go get -u github.com/gopherjs/gopherjs
go get -u github.com/IDerr/jobtracker/providers
gopherjs build
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