# Shopify Backend Engineer recruitment task

This repo consists of the source code for a recruitment task for Shopify's Backend Engineer Intern position.

[Summer 2019 Developer Intern Challenge Question](https://docs.google.com/document/d/1J49NAOIoWYOumaoQCKopPfudWI_jsQWVKlXmw1f1r-4/edit)

## Thought process

My plan was to create a GraphQL API with cart functionality with well-written code.

1. At first, I wrote on a sheet of paper queries, mutations and types I’d need for such API.
2. Initially, I wanted to give Prisma with TypeScript a try (never used that), but after few hours of fighting with Docker I gave up, I had some hardware problem.
3. I chose Go for this, I have some experience with Go.
4. I started with writing domain structures (that is models/entities - purchase, product, etc.).
5. Then I proceeded to writing first working resolver, for all products.
6. After writing the rest of queries and mutation for purchasing single product I refactored some of the code to be more abstract and because of that more testable and universal.
7. I began doing extra credit, I solved cart authorization using JWT. When creating a cart, you got a JWT with cart id in claims in response which you have to submit in `Authorization` header when you are dealing with cart endpoints.

## API

List of queries and mutations can be found in [schema.graphql](https://github.com/Albert221/shopify-recruitment-backend/blob/master/schema.graphql) file which describes it best (it's the actual schema).

### Cart resolvers

`showCart`, `addToCart` and `checkoutCart` resolvers need a JWT token submitted in `Authorization` header in order to work. You can acquire this token from `createCart` query.