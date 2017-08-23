// node web-server.js 1337

const http = require('http');

const server = http.createServer(function (req, res) {
  if(req.method == "GET") {
    res.writeHead(200, {'Content-Type': 'text/html'});
    res.write('Halo Bali!');
    res.end();
  }
});

server.listen(process.argv[2]);
