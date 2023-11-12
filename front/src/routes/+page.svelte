<script lang="ts">
	import { invalidate, invalidateAll } from '$app/navigation';
	import { Client, cacheExchange, fetchExchange, gql, mutationStore } from '@urql/svelte';

	const client = new Client({
		url: 'http://localhost:5173/api/graphql',
		exchanges: [cacheExchange, fetchExchange]
	});

	type Todo = {
		id: string;
		text: string;
		done: boolean;
	};

	let newTodo: Omit<Todo, 'id'> = {
		text: '',
		done: false
	};

	export let data: { todoList: Todo[] };

	const updateTodo = (id: string, text: string, done: boolean) => {
		mutationStore({
			client,
			query: gql`
        mutation {
          updateTodo (id: "${id}", text: "${text}", done: ${done ? 'true' : 'false'}) {
            id
            text
            done
          }
        }
      `
		});
		invalidate('app:todolist');
	};

	const deleteTodo = async (id: string) => {
		mutationStore({
			client,
			query: gql`
        mutation {
          deleteTodo (id: "${id}")
        }
      `
		});
		invalidate('app:todolist');
	};

	const createTodo = async (newTodo: Omit<Todo, 'id'>) => {
		mutationStore({
			client,
			query: gql`
        mutation {
          postTodo (text: "${newTodo.text}") {
            id
            text
            done
          }
        }
      `
		});
		invalidate('app:todolist');
		newTodo.text = '';
	};
</script>

<ul>
	<!-- show todolist-->
	{#each data.todoList as todo}
		<li>
			<span>{todo.id}</span>
			<input
				type="text"
				class={`${todo.done ? 'line-through' : ''}`}
				value={todo.text}
				on:change={(val) => updateTodo(todo.id, val.target.value, todo.done)}
			/>
			<input
				type="checkbox"
				name={todo.id}
				bind:checked={todo.done}
				on:change={(val) => {
					updateTodo(todo.id, todo.text, val.target.checked);
				}}
			/>
			<button on:click={() => deleteTodo(todo.id)}> delete </button>
		</li>
	{/each}
	<li>
		<span>new</span>
		<input type="text" bind:value={newTodo.text} />
		<button on:click={() => createTodo(newTodo)}>create</button>
	</li>
</ul>
