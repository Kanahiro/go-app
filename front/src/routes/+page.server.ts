import type { PageServerLoad } from './$types';
import { Client, fetchExchange, gql } from '@urql/core';

const client = new Client({
	url: 'http://front:5173/api/graphql',
	exchanges: [fetchExchange]
});

export const load: PageServerLoad = async ({ depends, setHeaders }) => {
	const result = await client
		.query(
			gql`
				query {
					todoList {
						id
						text
						done
					}
				}
			`,
			{}
		)
		.toPromise();

	depends('app:todolist');
	return { todoList: result.data.todoList };
};
