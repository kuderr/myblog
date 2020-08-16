# frontend

## Build Setup

```bash
# install dependencies
$ npm install

# serve with hot reload at localhost:8080
$ npm run dev

# build for production and launch server
$ npm run build
$ npm run start # to test if it works
$ pm2 start npm --name "blog" -- start # supervisor for nodejs server
$ pm2 monit
$ pm2 kill

# generate static project
$ npm run generate
```

For detailed explanation on how things work, check out [Nuxt.js docs](https://nuxtjs.org).
