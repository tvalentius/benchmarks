// node index.js 1337

const express = require('express');
const app = express();


app.get('/', function(req, res) {
  res.send('Halo Bali!');
});

app.listen( process.argv[2] || 1337 , function() {
  console.log('server listen on',process.argv[2] || 1337);
})
