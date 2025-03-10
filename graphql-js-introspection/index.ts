import {
  graphql,
  getIntrospectionQuery,
  GraphQLSchema,
  GraphQLObjectType,
  GraphQLString,
} from "graphql";

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

const source = getIntrospectionQuery({
  descriptions: true,
  specifiedByUrl: true,
  directiveIsRepeatable: true,
  schemaDescription: true,
  inputValueDeprecation: true,
  oneOf: true,
});

graphql({ schema, source }).then((result) => {
  console.log(JSON.stringify(result.data, null, 4));
});
