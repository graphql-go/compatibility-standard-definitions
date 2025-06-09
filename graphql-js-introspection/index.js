const {
  graphql,
  GraphQLSchema,
  GraphQLObjectType,
  GraphQLString,
  introspectionQuery,
} = require("graphql");

const schema = new GraphQLSchema({
  query: new GraphQLObjectType({
    name: "RootQueryType",
    fields: {
      echo: {
        type: GraphQLString,
        resolve() {
          return "ok";
        },
      },
    },
  }),
});

graphql(schema, introspectionQuery).then((result) => {
  console.log(JSON.stringify(result.data, null, 4));
});
