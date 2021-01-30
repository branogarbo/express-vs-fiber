const express = require('express');
const server = express()

server.get('/', (req,res) => {
   res.send('THIS A TEST')
})

server.listen(3000, console.log('listening on port 3000'))