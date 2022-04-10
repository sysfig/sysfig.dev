# NestJS

NestJS is a framework for building Node.js server-side applications. NestJS is an abstraction built on top of frameworks like Express and Fastify. NestJS uses decorators.

## Concepts

modules can include:

- controllers - handling requests and sending responses
- service/providers

## Getting Started

```console
$ npm i -g @nestjs/cli
$ nest new <project-name>
```

This will install the NestJS CLI, create a new directory at `<project-name>`, add boilerplate code inside of it, and install any dependencies. After running those commands, you can then go into the directory and run `nest start` and the application will be available on `http://localhost:3000/`.

If you navigate to `http://localhost:3000/`, you'll see `Hello World!` printed.

The files that makes this happen is in the `src/` directory:

- `app.controller.ts` - a _controller_ containing a single _route_
- `app.controller.spec.ts` - unit tests for the controller
- `app.module.ts` - root _module_ of the application
- `app.service.ts` - a _service_ with a single _method_
- `main.ts` - entry point of the application. Here is where we run `NestFactory.create` to create a new Nest application instance. You can change the entrypoint using the `entryFile` property in `nest-cli.json`

Each module should be in its own directory.

### Decorators

decorators associate classes with ?metadata?


- [RxJS](https://rxjs.dev/) - Reactive Extensions Library for JavaScript
- [`reflect-metadata`](https://github.com/rbuckton/reflect-metadata)