const fastify = require('fastify')({ logger: {
  level: 'debug'
} })
const fs = require('fs')
const util = require('util')
const { pipeline } = require('stream')
const pump = util.promisify(pipeline)
const path = require('path');
fastify.register(require('@fastify/multipart'))

fastify.register(require('@fastify/static'), {
  root: path.join(__dirname, 'uploads'),
  prefix: '/public/',
})

fastify.get('/', async (req, reply) => {
  reply.type('text/html').send(fs.readFileSync('index.html'))
})

fastify.post("/upload", async (req, reply) => {
  const data = await req.file()
  fastify.log.info("data.filename ====>", data)
  await pump(data.file, fs.createWriteStream('uploads/' + data.filename))
  reply.send()
});

fastify.listen({ port: 3000 })