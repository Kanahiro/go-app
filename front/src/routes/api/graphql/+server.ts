import { createSchema, createYoga } from 'graphql-yoga';
import type { RequestEvent, RequestHandler } from '@sveltejs/kit';

const yogaApp = createYoga<RequestEvent>({
	schema: createSchema({
		typeDefs: `
			type Query {
				todoList: [Todo!]!
			}
			type Mutation {
				postTodo (text: String!): Todo!
				updateTodo (id: ID!, text: String!, done: Boolean!): Todo!
				deleteTodo (id: ID!): Boolean!
			}
			type Todo {
				id: ID!
				text: String!
				done: Boolean!
			}
		`,
		resolvers: {
			Query: {
				todoList: async () => {
					const res = await fetch('http://app:3000/api/todo/');
					const val = await res.json();
					return val;
				}
			},
			Mutation: {
				postTodo: async (parent: unknown, args: { text: string }) => {
					const res = await fetch('http://app:3000/api/todo/', {
						method: 'POST',
						headers: {
							'Content-Type': 'application/json'
						},
						body: JSON.stringify({ text: args.text, done: false })
					});
					const val = await res.json();
					return val;
				},
				updateTodo: async (parent: unknown, args: { id: string; text: string; done: boolean }) => {
					const res = await fetch(`http://app:3000/api/todo/${args.id}`, {
						method: 'PUT',
						headers: {
							'Content-Type': 'application/json'
						},
						body: JSON.stringify({ text: args.text, done: args.done })
					});
					const val = await res.json();
					return val;
				},
				deleteTodo: async (parent: unknown, args: { id: string }) => {
					const res = await fetch(`http://app:3000/api/todo/${args.id}`, {
						method: 'DELETE'
					});
					const val = (await res.json()) as boolean;
					return val;
				}
			}
		}
	}),
	// Needed to be defined explicitly because our endpoint lives at a different path other than `/graphql`
	graphqlEndpoint: '/api/graphql',

	// Needed to let Yoga use sveltekit's Response object
	fetchAPI: { Response }
}) satisfies RequestHandler;

export { yogaApp as GET, yogaApp as POST };
